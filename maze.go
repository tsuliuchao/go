package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main(){

	maze := readMaze("D:/go_path/src/EEE/maze.in") //读maze.in文件
	for _,row := range maze {

		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}
	fmt.Println("--------")

	steps :=walk(maze,point{0,0},point{len(maze)-1,len(maze[0])-1})

	//打印走的路径
	for _,row := range steps{
		for _,val := range row{
			fmt.Printf("%3d ",val)
		}
		fmt.Println()
	}
}
//坐标，i是行，j是列
type point struct {
	i,j int
}
//走的路径，指下一个方向,上，左，下，右 逆时针方向走
var dirs = []point{{-1,0},{0,-1},{1,0},{0,1}}


func (p point)add(r point) point {
	return point{p.i+r.i,p.j+r.j}
}

//获取点point在grid位置的值
func (p point)at(grid [][]int) (int,bool){
	if p.i < 0 || p.i >= len(grid){
		return 0,false
	}
	if p.j < 0 || p.j >=len(grid[p.i]){
		return 0,false
	}
	return grid[p.i][p.j],true
}

func walk(maze [][]int,start,end point) [][]int{
	//steps记录游走的信息
	steps := make([][]int,len(maze))
	for i := range steps{
		steps[i] = make([]int,len(maze[i]))
	}
	//fmt.Println(steps)
	//作为队列使用，存放走得通的点,当某个点要被探索便把它出队列
	Q := []point{start}
	for len(Q) > 0{
		cur := Q[0]
		Q = Q[1:]//切片，去掉cur，依次循环下去cur都不是以前的值
		//fmt.Println(Q,cur,end)

		if cur == end{
			break
		}

		for _,dir := range dirs{
			next := cur.add(dir)
			//maze at next is 0
			//and steps at next is 0
			//and next != start
			//判断该点是否越界或者遇到墙
			val,ok := next.at(maze)
			if !ok || val == 1{
				continue
			}
			//==0相当于还没有
			val,ok = next.at(steps)
			if !ok || val != 0{
				continue
			}
			//不能往回走
			if next == start{
				continue
			}
			//当前的步骤数
			curSteps,_ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			Q = append(Q,next)
		}
	}
	return steps
}

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	fmt.Fscanf(file, "%d")
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			_,err= fmt.Fscanf(file, "%d", &maze[i][j])
			//fmt.Println(err)
		}
		//Fscanf无法对换行符做处理，方案1在maze.in空格替换所有的换行，另外使用这种方案
		fmt.Fscanf(file, "%d")
	}
	return maze
}
//以上程序结束
//======================================================================
//maze.in路径文件解析：左上角和右下角的点分别为起点和终点，0代表可以走，1代表代表墙，走不通，
//我们要做的是从起点走到终点，我们每到一个点便从上左下右四个方向探索它周围的四个点，如果是走过的点我们不要探索，计算出它的步数，用的广度优先算法。
//第一行代表grid共6行5列
//maze.in
6 5
0 1 0 0 0
0 0 0 1 0
0 1 0 1 0
1 1 1 0 0
0 1 0 0 1
0 1 0 0 0
//run-result
  0   0   4   5   6 
  1   2   3   0   7 
  2   0   4   0   8 
  0   0   0  10   9 
  0   0  12  11   0 
  0   0  13  12  13 



