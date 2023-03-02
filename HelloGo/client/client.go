package main

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars"

	"github.com/Tars/test/tars-protocol/TestApp"
)

func main() {
	comm := tars.NewCommunicator()
	obj := fmt.Sprintf("TestApp.HelloGo.SayHelloObj@tcp -h 172.1.1.2 -p 22857 -t 60000")
	//obj := fmt.Sprintf("TestApp.HelloGo.SayHelloObj@tcp -h 127.0.0.1 -p 10015 -t 60000")
	app := new(TestApp.SayHello)
	comm.StringToProxy(obj, app)

	//var i, r int32
	//i = 1223
	//ret, err := app.Add(i, i*2, &r)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(ret, r)
	
	input := "i am client"
	var output string
	ret, err := app.ShowGoReq(input, &output)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println(ret, output)
}

