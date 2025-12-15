//checking and handling errors
package main
import ("fmt"
"errors")
func filename(file string)(string,error){
	if file==""{
		return "",errors.New("file cannot be empty")
	}
	return file,nil
}
func main(){
	content,err:=filename("i am lilith")
	if err!=nil{
        fmt.Println("error is",err)
		return
	}
	fmt.Println("content:",content)
}