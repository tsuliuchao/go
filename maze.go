package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main(){

	maze := readMaze("D:/go_path/src/maze/maze.in") //读maze.in文件
	for _,row := range maze {

		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}
	fmt.Println("------------------------------")

	steps :=walk(maze,point{0,0},point{len(maze)-1,len(maze[0])-1})

	//打印走的路径
	for _,row := range steps{
		for _,val := range row{
			fmt.Printf("%3d ",val)
		}
		fmt.Println()
	}
}
type point struct {
	x,y int
}
//走的路径，指下一个方向,上，左，下，右 逆时针
var dirs = []point{{-1,0},{0,-1},{1,0},{0,1}}


func (p point)add(r point) point {
	return point{p.x+r.x,p.y+r.y}
}

//获取点point在grid位置的值
func (p point)at(grid [][]int) (int,bool){
	if p.x < 0 || p.x >= len(grid){
		return 0,false
	}
	if p.y < 0 || p.y >=len(grid[p.x]){
		return 0,false
	}
	return grid[p.x][p.y],true
}
func walk(maze [][]int,start,end point) [][] int{
	steps := make([][]int,len(maze))
	queue := []point{start}
	for i := range maze{
		steps[i] = make([]int,len(maze[i]) )
	}
	for len(queue)>0{
		current := queue[0]
		queue = queue[1:]

		if current == end{
			continue
		}
		for _,dir :=range dirs{
			//maze at next is 0
			//and steps at next is 0
			//and next != start
			next := current.add(dir)
			if next == start{
				continue
			}
			//遇到为1的墙，跳过
			val,err := next.at(maze)
			if !err || val==1{
				continue
			}
			//不等于0表示已经走过，跳过
			val,err = next.at(steps)
			if !err || val !=0 {
				continue
			}
			//走过路径+1
			current_step,err := current.at(steps)
			steps[next.x][next.y] = current_step+1
			queue = append(queue,next)
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



