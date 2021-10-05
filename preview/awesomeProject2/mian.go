package main

import "fmt"

func main(){
	var name string
	var pwd  string
	fmt.Print("请输入账号：")
	fmt.Scanln(&name)
	fmt.Print("请输入密码：")
	fmt.Scanln(&pwd)
	if name == "asfd" && pwd == "zxvc"{
		fmt.Println("登录成功")
	}else{
		fmt.Println("密码错误!")
	}
}