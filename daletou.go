package main

import (
	"math/rand"
	"fmt"
	"time"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func contain(a []int,i int) bool {
	for _,v := range a {
		if v == i {
			return true
		}
	}
	return false
}

func checkA(a []int) bool {
	if len(a)!= 5{
		fmt.Println("输入有误，请输入5位数字\n")
		return false
	}

	for _,v := range a {
		if v>35 || v<= 0{
			fmt.Println("输入数字有误:",v)
			fmt.Println()
			return false
		}
	}
	return true
}

func checkB(b []int) bool {
	if len(b)!= 2{
		fmt.Println("输入有误，请输入2位数字\n")
		return false
	}

	for _,v := range b {
		if v>12 || v<= 0{
			fmt.Println("输入数字有误:",v)
			fmt.Println()
			return false
		}
	}
	return true
}

func transformA(a []int,n int) {
	for changed := 0;changed<n;{
		for i,_ := range a {
			if changed == n {
				break
			}
			p := rand.Float64()
			if p<=0.5 {
				rn := rand.Intn(34)+1
				if !contain(a,rn){
					a[i] = rn
					changed += 1
				}
			}
		}
	}
}


func transformB(b []int,n int) {
	for changed := 0;changed<n;{
		for i,_ := range b {
			if changed == n {
				break
			}
			p := rand.Float64()
			if p<=0.5 {
				rn := rand.Intn(11)+1
				if !contain(b,rn){
					b[i] = rn
					changed += 1
				}
			}
		}
	}
}


func getNumbers(a []int,b[]int)  {
	rand.Seed(time.Now().Unix())

	//随机替换3个位置的数
	transformA(a,3)

	transformA(a,2)

	//随机替换1个位置的数
	transformB(b,1)

	transformB(b,1)

	fmt.Println(a,b)
	time.Sleep(1*time.Second)
}

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}


func main() {
	//

	cancel := make(chan int)

	go func() {
		for {
			var a  = make([]int,0,5)
			var b  = make([]int,0,5)

			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("请输入前5个数字")
			if scanner.Scan(){
				a = numbers(scanner.Text())
			}
			if !checkA(a) {
				continue
			}
			fmt.Println("请输入后两个数字")
			if scanner.Scan(){
				b = numbers(scanner.Text())
			}
			if !checkB(b) {
				continue
			}

			fmt.Println("请输入要生成的号码个数：")
			var n int
			_,err := fmt.Scanf("%d",&n)
			if err != nil {
				fmt.Println("输入有误，退出")
				cancel <- 0
			}

			for j := 0;j<n;j++{
				var ac  = make([]int,5,5)
				var bc  = make([]int,2,2)
				copy(ac,a)
				copy(bc,b)
				getNumbers(ac,bc)
			}
			fmt.Println()
		}
	}()

	<-cancel

}