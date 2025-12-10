package main
import "fmt"
func countwords(s string)int{
	count:=0
	inword:=false
	for _,ch:= range s{
		if ch!=' ' && !inword{
			inword=true
			count++
		}else if ch==' '{
			inword =false
		}

	}
	return count
}
func main(){
	var s string="I am Lilith"
	count:=countwords(s)
	fmt.Println("Number of words in a string",count)
}