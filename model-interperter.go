package main

import (
	"fmt"
)

type Context struct {
	text string
}

// 抽象表达式
type IAbstractExpression interface {
	Interpret(*Context)
}

// 终结符表达式
type TerminalExpression struct {
}

func (t *TerminalExpression) Interpret(context *Context) {
	if t == nil {
		return
	}
	context.text = context.text[:len(context.text)-1]
	fmt.Println(context)
}

// 非终结符表达式
type NonterminalExpression struct {
}

func (t *NonterminalExpression) Interpret(context *Context) {
	if t == nil {
		return
	}
	context.text = context.text[:len(context.text)-1]
	fmt.Println(context)
}

func main(){
	context := Context{"abc"}

	list := []IAbstractExpression{}

	list = append(list, new(TerminalExpression))
	list = append(list, new(TerminalExpression))
	list = append(list, new(NonterminalExpression))

	for _, val := range list {
		val.Interpret(&context)
	}
}