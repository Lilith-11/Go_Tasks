//creating and returning errors
package main
import ("fmt"
        "errors")
func divide(a,b float32)(float32,error){
	if b==0{
		return 0,errors.New("value not divisible by 0")
	}
	return a/b,nil
}
func main(){
	res,err:=divide(2,0)
	if err!=nil{
		fmt.Println("error",err)
		return
	}
	fmt.Println(res)
}