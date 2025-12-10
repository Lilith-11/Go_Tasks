package main
import "fmt"
func main(){
	var a int =12
	var b int =24
	fmt.Println("before swap of a and b",a,b)
	a=a+b
	b=a-b
	a=a-b
	fmt.Println("after swap of a and b without third variable is",a,b)
}