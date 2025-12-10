package main
import (
	"fmt"
	"strconv")
func main(){
	var s string ="123"
	num,err:=strconv.Atoi(s)
	if err!=nil{
        fmt.Println("error happens",err)
	}
    fmt.Println("converting string to integer is",num)	
}