package main
import "fmt"
func palindrome(num int)int{
    a:=0
	for num!=0{
		m:=num%10
		num=num/10
		a=a*10+m
	}
	return a
}
func main(){
	var num int 
	fmt.Print("enter your number")
	fmt.Scan(&num)
	n:=num
   p:=palindrome(num)
   if n==p{
	fmt.Println("Palindrome")
   }else{
	fmt.Println("not palindrome")
   }
}