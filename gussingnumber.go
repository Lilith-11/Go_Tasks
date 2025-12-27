package main
import ("fmt"
"time"
"math/rand")
func main(){
rand.Seed(time.Now().UnixNano())
target:=rand.Intn(100)+1
var guess int
var attempts int=0
for{
	fmt.Println("enter your guess")
	fmt.Scan(&guess)
	attempts++
    if guess<target{
		fmt.Println("too low")
	}else if guess>target {
		fmt.Println("too high")
	}else{
		fmt.Println("you are correct")
		fmt.Println("attempts",attempts)
		break
	}
}}
