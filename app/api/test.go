package api

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

var Test = new(TestApi)

type TestApi struct {
	name string
	age  int
	sex  int
}

func init() {
	fmt.Println("this is test.go")
}

func getInfo(input *ghttp.Request) {
	var name = input.Get("name")
	fmt.Println(name)
	fmt.Println("123")

}
