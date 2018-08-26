package main

import (
	"fmt"
	"time"
	"runtime"
)

func qie(m []int, n int) [][]int {
	res := [][]int{}
	len := len(m)
	last := 0
	for i:=1; i<=len;i++ {
		if (i % n ==0) {
			last = i
			res =  append(res,m[i-n:i])
		} else if (i == len) {
			res = append(res, m[last:len])
		}
	}

	return res
}
type Fruit struct {
	price int
	quantity int
}

func (fruit *Fruit) Cost() int {
	return fruit.price * fruit.quantity
}

type SpecialFruit struct {
	Fruit  // 匿名字段 （嵌入）
	markup int  // 具名字段 （聚合）
}

/*func (specialFruit *SpecialFruit) Cost() int {
	return specialFruit.Fruit.Cost()  * specialFruit.markup
}*/

func main() {

	fmt.Println(runtime.NumCPU())
	res := map[string]string{"name": "wuyanwen", "age":"100"}
	go func() {
		fmt.Println(res)
	}()
	go func() {
		fmt.Println(res)
	}()
	go func() {
		fmt.Println(res)
	}()
	go func() {
		fmt.Println(res)
	}()
	go func() {
		fmt.Println(res)
	}()
	go func() {
		fmt.Println(res)
	}()

	go func() {res["name"] = "dddd"
		time.Sleep(time.Microsecond * 1)

		fmt.Println(res)
	}()

	go func() {
		fmt.Println(res)
	}()

	fmt.Println(res)
	time.Sleep(time.Second * 2)
}