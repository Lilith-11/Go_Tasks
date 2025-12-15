package main
import "fmt"
readfile(filename string)(string,error){
      if filename==""{
		return "",fmt.Errorf("the file has no content")
	  }
	  return filename,nil
}
func main(){
	content,err:=readfile("")
	if err!=nil{
		fmt.Println("error:",err)
		return 
	}
	fmt.Println("content:",content)

}