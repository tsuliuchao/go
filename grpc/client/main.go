package main

import "net/rpc"
import (
	. "grpc/protocol"
	"fmt"
	"time"
)

var (    _CLIENT  *rpc.Client
	_RPC_MSG chan *rpc.Call
	_CAN_CANCEL chan bool
)

func main()  {
	DialRpcServer()
	//起个协程处理异步rpc调用结果
	go loop()

	//测试同步的方式调用rpc服务
	param := Param{A:int32(10),B:int32(30)}
	reply := int32(0)
	SyncCallRpcFunc(RPC_ADDITION, &param, &reply)
	fmt.Printf("Sync Call Addition Result %d \n", reply)
	SyncCallRpcFunc(RPC_SUBTRACTION, &param, &reply)
	fmt.Printf("Sync Call Subtraction Result %d \n", reply)
	////测试异步的方式调用rpc服务
	ASyncCallRpcFunc(RPC_MULTIPLICATION, &param, &reply)
	ASyncCallRpcFunc(RPC_DIVISION, &param, &reply)
	//阻塞等待异步调用完成
	<- _CAN_CANCEL
}

func init(){
	_RPC_MSG = make(chan *rpc.Call, 1024)
	_CAN_CANCEL = make(chan bool)
}

func DialRpcServer(){
	c, e := rpc.DialHTTP("tcp", "127.0.0.1:2311")

	if e != nil {
		fmt.Errorf("Dial RPC Error %s", e.Error())
	}

	_CLIENT = c
}
//重连RPC服务器
func ReDialRpcServer() bool{
	c, e := rpc.DialHTTP("tcp", "127.0.0.1:2311")

	if e != nil {
		fmt.Printf("ReDial RPC Error %s \n", e.Error())

		return false
	}

	_CLIENT = c

	fmt.Println("ReDial Rpc Server Succ")

	return true
}
//同步rpc调用
func SyncCallRpcFunc(method string, args interface{}, reply interface{}){
	if nil == _CLIENT{
		for{//如果断线就等到重连上为止
			if ReDialRpcServer(){
				break
			}
			time.Sleep(5000 * time.Millisecond)
		}
	}

	_CLIENT.Call(method, args, reply)
}
//异步rpc调用
func ASyncCallRpcFunc(method string, args interface{}, reply interface{}){
	if nil == _CLIENT{
		for{//如果断线就等到重连上为止
			if ReDialRpcServer(){
				break
			}
			time.Sleep(5000 * time.Millisecond)
		}
	}
	// Go(serviceMethod string, args interface{}, reply interface{}, done chan *Call)的done如果填nil会构建个新的channel用于接受结果
	_CLIENT.Go(method, args, reply, _RPC_MSG)
}
//接收异步调用的返回
func loop(){
	for{
		select {
		case rpcMsg, ok := <- _RPC_MSG:
			if !ok{
				fmt.Errorf("Rpc Call Error")
			}
			rpcMsgHandler(rpcMsg)
		}
	}

	_CAN_CANCEL <- true
}
// 处理异步rpc的返回值
func rpcMsgHandler(msg * rpc.Call){
	switch msg.ServiceMethod {
	case RPC_ADDITION:
		reply := msg.Reply.(*int32)
		fmt.Printf("Addtoion Result [%d] \n", *reply)
	case RPC_SUBTRACTION:
		reply := msg.Reply.(*int32)
		fmt.Printf("Subtraction Result [%d] \n", *reply)
	case RPC_MULTIPLICATION:
		reply := msg.Reply.(*int32)
		fmt.Printf("Multiplication Result [%d] \n", *reply)
	case RPC_DIVISION:
		reply := msg.Reply.(*int32)
		fmt.Printf("Division Result [%d] \n", *reply)
	default:
		fmt.Errorf("Can Not Handler Reply [%s] \n", msg.ServiceMethod)

	}
}