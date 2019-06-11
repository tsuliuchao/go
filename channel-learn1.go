1. buffered channel and non-buffered channel

buffered channel

bufferedchan := make(chan int, 1)
when adding element into bufferedchan, it will not blocked if it is not full.

for example:

复制代码
package main

import "fmt"

func main() {

messages := make(chan string, 1)

messages <- "buffered"

fmt.Println("quit!")
}
复制代码
the program won't block and the answer is:

quit!



2.channel is thread safe, and we can read data from a closed channel

按 Ctrl+C 复制代码

package main

import "fmt"

func main() {
jobs := make(chan int, 5)
done := make(chan bool)

go func() {                           ------------------------------- (Part I)
fmt.Println("enter go routinue")

for {
j, more := <-jobs
if more {
fmt.Println("received job", j)
} else {
fmt.Println("received all jobs")
done <- true
return
}
}
}()

jobs <- 1                             ------------------------------- (part II)
jobs <- 2
jobs <- 3
close(jobs)
fmt.Println("sent all jobs and closed")


<-done                                ------------------------------- (Part III)
}
按 Ctrl+C 复制代码
the result is:

sent all jobs and closed
enter go routinue
received job 1
received job 2
received job 3
received all jobs



firstly, it executes (Part II)  and stop at  <-done.

and then a child routinue executes (Part I) and done<-true.

Now the channel done is actived and continue to execute.



*If the channel jobs is not closed in main, the statement "j, more := <-jobs" in (Part I) will block.  the deadlock will happen.

*If we don't set "<-done" in main, main will not wait the child routinue to finish.



3. range iterates buffered channel seeing https://gobyexample.com/range-over-channels



按 Ctrl+C 复制代码

package main

import "fmt"

func main() {

queue := make(chan string, 2)
queue <- "one"
queue <- "two"
close(queue)

for elem := range queue {
fmt.Println(elem)
}
fmt.Println("done!")
}
按 Ctrl+C 复制代码


if we don't close the queue, it will block in  "elem := range queue".



4. a closed channel won't block;  we can't write to it, but we can read from it infinitly. the read value is always 0 value like 0/false/nil

复制代码
package main

import "fmt"

func main() {

queue := make(chan int, 2)
queue <- 10
queue <- 20
close(queue)

for i := 0; i < 10; i++ {
fmt.Println(<-queue)
}

fmt.Println("done!")
}
复制代码
the result is

10
20
0
0
0
0
0
0
0
0
done!



!#Never too late to do it now#!