package main
import "fmt"
func main(){
	var a,b int
	a=12
	b=5
	for b!=0{
	   a,b=b,a%b
	}
	fmt.Println(a)

}