package main
//from:https://www.jianshu.com/p/31757e530144
import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)
//1.simple-json
type Account struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Money float64 `json:"money"`
}
var jsonString string = `{
	"email":"chaosir@gmail.com",
	"password":"1234",
	"money":1.23
}`
//2.tag-json
//，golang会将json的数据结构和go的数据结构进行匹配。匹配的原则就是寻找tag的相同的字段，然后查找字段。查询的时候是大小写不敏感的,但是如果结构的字段是私有的，即使tag符合，也不会被解析：
type AccountTwo struct {
	Email string `json:"email"`
	Password string
	Money float64 `json:"money"`
}
type AccountPasswordPrivate struct {
	Email    string  `json:"email"`
	password string  `json:"password"`
	Money    float64 `json:"money"`
}
//3.string-tag
//使用tag string，可以把结构定义的数字类型以字串形式编码。同样在解码的时候，只有字串类型的数字，才能被正确解析，或者会报错：panic: json: invalid use of ,string struct tag, trying to unmarshal unquoted value into float64
type AccountThree struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	//Money    float64 `json:"money,string"` //出错
	Money    float64 `json:"money,float64"`
}
//4.-tag 不解析字段  Money    float64 `json:"-"`
//Convert JSON to Go struct ：https://mholt.github.io/json-to-go/https://mholt.github.io/json-to-go/
type AccountFour struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Money    float64 `json:"-"`
}
//5.动态解析
//实际开发中可能有动态数据类型，尤其一些不确定的字段，单个字段多种数据类型支持等
//前面我们使用了简单的方法Unmarshal直接解析json字串，下面我们使用更底层的方法NewDecode和Decode方法。
//Decode函数，在这个函数进行json字串的解析。然后调用json的NewDecoder方法构造一个Decode对象，最后使用这个对象的Decode方法赋值给定义好的结构对象。
type User struct {
	UserName interface{} `json:"username"`
	Password string `json:"password"`
}
var jsonDMap string = `{
   "username":[{"name":"八爷rs7@gmail.com"},{"name":"sj217@gmail.com"}],
   "password":"123"
}`
var jsonDString string = `{
   "username":"chaosir",
   "password":"123"
}`
//6.延迟解析 因为UserName字段，实际上是在使用的时候，才会用到他的具体类型，因此我们可以延迟解析。使用json.RawMessage方式，将json的字串继续以byte数组方式存在,总体而言，延迟解析和使用空接口的方式类似。需要再次调用Unmarshal方法，对json.RawMessage进行解析。原理和解析到接口的形式类似。

type UserDefer struct {
	UserName json.RawMessage `json:"username"`
	Password string `json:"password"`
	Email string
	Phone int64
}
var jsonDeferString string = `{
    "username":"18512341234@qq.com",
    "password":"123"
}`
//7.不定字段的解析：对于未知json结构的解析，不同的数据类型可以映射到接口或者使用延迟解析。有时候，会遇到json的数据字段都不一样的情况。例如需要解析下面一个json字串：
var jsonMuxString string = `{
        "things": [
            {
                "name": "Alice",
                "age": 37
            },
            {
                "city": "Ipoh",
                "country": "Malaysia"
            },
            {
                "name": "Bob",
                "age": 36
            },
            {
                "city": "Northampton",
                "country": "England"
            }
        ]
    }`
//json字串的是一个对象，其中一个key things的值是一个数组，这个数组的每一个item都未必一样，大致是两种数据结构，可以抽象为person和place。即，定义下面的结构体，
type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
type Place struct {
	City string `json:"city"`
	Country string `json:"country"`
}
//8.混合结构混合结构的思路很简单，借助golang会初始化没有匹配的json和抛弃没有匹配的json，给特定的字段赋值。比如每一个item都具有四个字段，只不过有的会匹配person的json数据，有的则是匹配place。没有匹配的字段则是零值。接下来在根据item的具体情况，分别赋值到对于的Person或Place结构。
func DecodeMux(jsonStr []byte)(persons []Person,places []Place){
	var data map[string][]map[string]interface{}
	err := json.Unmarshal(jsonStr,&data)
	if err!=nil{
		fmt.Println(err)
		return
	}
	for i:= range data["things"]{
		item:=data["things"][i]
		if item["name"] != nil{
			persons = addPerson(persons,item)
		}else{
			places = addPlace(places,item)
		}
	}
	return
}
func addPerson(persons []Person, item map[string]interface{}) []Person {
	name := item["name"].(string)
	age := item["age"].(float64)
	person := Person{name, int(age)}
	persons = append(persons, person)
	return persons
}

func addPlace(places []Place, item map[string]interface{})([]Place){
	city := item["city"].(string)
	country := item["country"].(string)
	place := Place{City:city, Country:country}
	places = append(places, place)
	return places
}
//把things的item数组解析成一个json.RawMessage，然后再定义其他结构逐步解析。上述这些例子其实在真实的开发环境下，应该尽量避免。像person或是place这样的数据，可以定义两个数组分别存储他们，这样就方便很多。不管怎么样，通过这个略傻的例子，我们也知道了如何解析json数据。
func DecodeMuxTwo(jsonStr []byte)(persons []Person,places []Place){
	var data map[string][]json.RawMessage
	err := json.Unmarshal(jsonStr,&data)
	if err!=nil{
		fmt.Println(err)
		return
	}
	for _,item:= range data["things"]{
			persons = addPersonTwo(persons,item)
			places = addPlaceTwo(places,item)
	}
	return
}
func addPersonTwo(persons []Person, item json.RawMessage) []Person {
	person :=Person{}
	if err:=json.Unmarshal(item,&person);err != nil {
		fmt.Println(err)
	}else{
		if person != *new(Person){
			persons = append(persons,person)
		}
	}
	return persons
}

func addPlaceTwo(places []Place, item json.RawMessage)([]Place){
	place := Place{}
	if err:=json.Unmarshal(item,&place);err != nil {
		fmt.Println(err)
		return places
	}
	if place != *new(Place){
		places = append(places, place)
	}
	return places
}
func DecodeDefer(r io.Reader)(udf *UserDefer,err error){
	udf = new(UserDefer)
	err = json.NewDecoder(r).Decode(udf)
	if err != nil {
		panic(err)
	}
	var email string
	if err= json.Unmarshal(udf.UserName,&email);err == nil {
		udf.Email = email
	}
	var phone int64
	if err=json.Unmarshal(udf.UserName,&phone);err==nil{
		udf.Phone = phone
	}
	return
}

func Decode(r io.Reader)(u *User,err error){
	u = new(User)
	err = json.NewDecoder(r).Decode(u)
	if err != nil {
		return
	}
	switch t:= u.UserName.(type) {
		case string:
			u.UserName = t
		case map[string]interface{}:
			u.UserName = t
	}
	return
}
func main(){
	//Unmarshal接受一个byte数组和空接口指针的参数。和sql中读取数据类似，先定义一个数据实例，然后传其指针地址。
	account := Account{}
	err := json.Unmarshal([]byte(jsonString),&account)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n",account)
	accountThree := AccountThree{}
	err = json.Unmarshal([]byte(jsonString),&accountThree)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n",accountThree)
	user,err := Decode(strings.NewReader(jsonDMap))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n",user)
	user_string,err := Decode(strings.NewReader(jsonDString))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n",user_string)
	udf_string,err := DecodeDefer(strings.NewReader(jsonDeferString))
	if err != nil {
		//panic(err)
	}
	fmt.Printf("%+v\n",udf_string)
	personA,placeB := DecodeMux([]byte(jsonMuxString))
	fmt.Printf("%+v\n",personA)
	fmt.Printf("%+v\n",placeB)
	personA,placeB = DecodeMuxTwo([]byte(jsonMuxString))
	fmt.Printf("%+v\n",personA)
	fmt.Printf("%+v\n",placeB)
}

