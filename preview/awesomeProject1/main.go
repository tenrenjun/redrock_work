package main
import "fmt"

func main(){
	var a,b int
	var c string
	fmt.Println("input:")
	fmt.Scan(&a)
	fmt.Scan(&c)
	fmt.Scan(&b)
	if c == "+"{
		fmt.Println(a,c,b,"=",a+b)
	}else if c == "-"{
		fmt.Println(a,c,b,"=",a-b)
	}else if c == "/"{
		fmt.Println(a,c,b,"=",a/b)
	}else if c == "*"{
		fmt.Println("output:",a,c,b,"=",a*b)
	}else{
		fmt.Println("输入错误")
	}
}