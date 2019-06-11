背景

在很多场景中，我们都会进行字符串拼接操作。

最开始的时候，你可能会使用如下的操作：

package main

func main() {
    ss := []string{
        "A",
        "B",
        "C",
    }

    var str string
    for _, s := range ss {
        str += s
    }

    print(str)
}
与许多支持string类型的语言一样，golang中的string类型也是只读且不可变的。因此，这种拼接字符串的方式会导致大量的string创建、销毁和内存分配。如果你拼接的字符串比较多的话，这显然不是一个正确的姿势。

在 Golang 1.10 以前，你可以使用bytes.Buffer来优化：

package main

import (
    "bytes"
    "fmt"
)

func main() {
    ss := []string{
        "A",
        "B",
        "C",
    }

    var b bytes.Buffer
    for _, s := range ss {
        fmt.Fprint(&b, s)
    }

    print(b.String())
}
这里使用 var b bytes.Buffer 存放最终拼接好的字符串，一定程度上避免上面 str 每进行一次拼接操作就重新申请新的内存空间存放中间字符串的问题。

但这里依然有一个小问题： b.String() 会有一次 []byte -> string 类型转换。而这个操作是会进行一次内存分配和内容拷贝的。

使用 strings.Builder 进行字符串拼接

如果你现在已经在使用 golang 1.10, 那么你还有一个更好的选择：strings.Builder:

package main

import (
    "fmt"
    "strings"
)

func main() {
    ss := []string{
        "A",
        "B",
        "C",
    }

    var b strings.Builder
    for _, s := range ss {
        fmt.Fprint(&b, s)
    }

    print(b.String())
}
Golang官方将strings.Builder作为一个feature引入，想必是有两把刷子。不信跑个分？简单来了个benchmark:

package ts

import (
    "bytes"
    "fmt"
    "strings"
    "testing"
)

func BenchmarkBuffer(b *testing.B) {
    var buf bytes.Buffer
    for i := 0; i < b.N; i++ {
        fmt.Fprint(&buf, "ABC")
        _ = buf.String()
    }
}

func BenchmarkBuilder(b *testing.B) {
    var builder strings.Builder
    for i := 0; i < b.N; i++ {
        fmt.Fprint(&builder, "ABC")
        _ = builder.String()
    }
}
╰─➤  go test -bench=. -benchmem                                                                                                                         2 ↵
goos: darwin
goarch: amd64
pkg: test/ts
BenchmarkBuffer-4         300000        101086 ns/op      604155 B/op          1 allocs/op
BenchmarkBuilder-4      20000000            90.4 ns/op        21 B/op          0 allocs/op
PASS
ok      test/ts 32.308s
性能提升感人。要知道诸如C#, Java这些自带GC的语言很早就引入了string builder, Golang 在1.10才引入，时机其实不算早，但是巨大的提升终归没让人失望。下面我们看一下标准库是如何做到的。

strings.Builder 原理解析

strings.Builder的实现在文件strings/builder.go中，一共只有120行，非常精炼。关键代码摘选如下：

type Builder struct {
    addr *Builder // of receiver, to detect copies by value
    buf  []byte // 1
}

// Write appends the contents of p to b's buffer.
// Write always returns len(p), nil.
func (b *Builder) Write(p []byte) (int, error) {
    b.copyCheck()
    b.buf = append(b.buf, p...) // 2
    return len(p), nil
}

// String returns the accumulated string.
func (b *Builder) String() string {
    return *(*string)(unsafe.Pointer(&b.buf))  // 3
}

func (b *Builder) copyCheck() {
    if b.addr == nil {
        // 4
        // This hack works around a failing of Go's escape analysis
        // that was causing b to escape and be heap allocated.
        // See issue 23382.
        // TODO: once issue 7921 is fixed, this should be reverted to
        // just "b.addr = b".
        b.addr = (*Builder)(noescape(unsafe.Pointer(b)))
    } else if b.addr != b {
        panic("strings: illegal use of non-zero Builder copied by value")
    }
}
与byte.Buffer思路类似，既然 string 在构建过程中会不断的被销毁重建，那么就尽量避免这个问题，底层使用一个 buf []byte 来存放字符串的内容。
对于写操作，就是简单的将byte写入到 buf 即可。
为了解决bytes.Buffer.String()存在的[]byte -> string类型转换和内存拷贝问题，这里使用了一个unsafe.Pointer的存指针转换操作，实现了直接将buf []byte转换为 string类型，同时避免了内存充分配的问题。
如果我们自己来实现strings.Builder, 大部分情况下我们完成前3步就觉得大功告成了。但是标准库做得要更近一步。我们知道Golang的堆栈在大部分情况下是不需要开发者关注的，如果能够在栈上完成的工作逃逸到了堆上，性能就大打折扣了。因此，copyCheck 加入了一行比较hack的代码来避免buf逃逸到堆上。关于这部分内容，你可以进一步阅读Dave Cheney的关于Go’s hidden #pragmas.
就此止步？

一般Golang标准库中使用的方式都是会逐步被推广，成为某些场景下的最佳实践方式。

这里使用到的*(*string)(unsafe.Pointer(&b.buf))其实也可以在其他的场景下使用。比如：如何比较string与[]byte是否相等而不进行内存分配呢？似乎铺垫写得太明显了，大家应该都会写了，直接给出代码吧：

func unsafeEqual(a string, b []byte) bool {
    bbp := *(*string)(unsafe.Pointer(&b))
    return a == bbp
}
