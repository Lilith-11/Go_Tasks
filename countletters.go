package main
import "fmt"
func countletters(s string)int{
	cn:=0
	for _= range s{
		
		cn=cn+1
	}
	return cn
}
func main(){
	var s string="LILITH"
	c:=countletters(s)
	fmt.Println("count letters in a string",c)
}