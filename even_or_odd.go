package main
import "fmt"
func main(){
	var a int
	fmt.Print("enter your number:")
	fmt.Scan(&a)
	if a%2==0{
		fmt.Println("your entered number is even")
	}else{
		fmt.Println("your entered number is odd")
	}
}