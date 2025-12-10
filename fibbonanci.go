package main
import "fmt"
func fibbonaci(num int)int{
	first:=0
    second:=1
	for i:=1;i<5;i++{
       third:=first+second
	   first=second
	   second=third
	}
	return second

}
func main(){
	var num int =5
	f:=fibbonaci(num)
	fmt.Println("fibbonaci series of 5 is",f)
}