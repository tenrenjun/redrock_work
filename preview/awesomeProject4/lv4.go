package main

import "fmt"
var (
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)
var coins int=0
func dispatchCoins() {
	for _, user := range users {
		for _, c := range user {
			switch c {
			case 'e', 'E':
				distribution[user]++
				coins++
			case 'i', 'I':
				distribution[user] += 2
				coins += 2
			case 'o', 'O':
				distribution[user] += 3 //分金币
				coins += 3              //分出金币之后需要从总金币数量扣除
			case 'u', 'U':
				distribution[user] += 4
				coins += 4

			}
		}

	}
	fmt.Println(distribution) //打印每人分到的金币
	fmt.Println("一共花费的硬币", coins)  //剩余的金币
}
func main() {
	dispatchCoins()
}