package main
import "fmt"
func sumdigits(n int)int{
	sum:=0
    for i:=1;i<=n;i++{
        sum=sum+i
	}
	return sum
}
func main(){
	n:=3
   sum:=sumdigits(n)
   fmt.Println("sum of the digits upto n is",sum)
} 