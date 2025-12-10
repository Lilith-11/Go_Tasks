package main
import "fmt"
func main(){
	var F float64
	fmt.Print("enter your fahrenheit")
	fmt.Scan(&F)
	C:=(F-32)*5/9
	fmt.Println("converted fahrenheit to celsius is",C)

}