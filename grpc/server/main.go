package main

import (
	"fmt"
	"errors"
	"net"
	"net/http"
	"net/rpc"
	. "grpc/protocol"
)

type Calculator struct {
}
var(
	_DATA *Calculator
	_CAN_CANCEL chan bool
)

func main(){
	runRpcServer()
}
func init(){
	_DATA = new(Calculator)
	_CAN_CANCEL = make(chan bool)
}

func runRpcServer(){
	rpc.Register(_DATA)
	rpc.HandleHTTP()
	le,err := net.Listen("tcp",":2311")
	if err!=nil{
		fmt.Errorf("cretater listen Error %v",err)
		return
	}
	go http.Serve(le,nil)
	<- _CAN_CANCEL
}


func (*Calculator) Addition(args *Param,reply *int32) error {
	*reply = args.A + args.B
	return nil
}
func (*Calculator) Subtraction(args *Param,reply *int32) error{
	*reply = args.A - args.B
	return nil
}
func (*Calculator)Multiplication(args *Param,reply *int32) error{
	*reply = args.A * args.B
	return nil
}

func(*Calculator)Division(args *Param,reply *int32)error{
	if 0 == args.B{
		return errors.New("divide by zero")
	}
	*reply = args.A/args.B
	return nil
}
