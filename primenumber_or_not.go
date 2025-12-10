package main
import "fmt"
func isprime(num int)bool{
	if num<2{
      return true
	}
	if num==2{
		return true	}
	for i:=2;i<num;i++{
		if num%i==0{
			return false
		}else{
           return true
		}
	}
	return true
}
func main(){
	num:=4
	isprime(num)
	if isprime(num){
		fmt.Println("number is prime")
	}else{
        fmt.Println("number is not prime")
	}
}