//给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；prices = [1, 3, 2, 8, 4, 9], fee = 2
//结果：能够达到的最大利润: 该怎么买进卖出
package main

import (
	"fmt"
	"math"
)

func main(){
	prices := []float64{1, 3, 2, 8, 4, 9}
	fmt.Println(maxProfit(prices,2))
}
func maxProfit(prices []float64,fee float64) float64{
	n := len(prices)
	if n<1 {
		return 0
	}
	buy := -prices[0]
	cash :=0.0
	for i:=1;i<n;i++{
		cash = math.Max(cash,float64(buy+prices[i]-fee))
		buy = math.Max(buy,cash-prices[i])
	}
	return cash
}

