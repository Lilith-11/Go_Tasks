package main
import "fmt"
func main(){
	a:="lilith"
    b:=[]byte{}
	for i:=0;i<len(a);i++{
		ch:=a[i]
		exists:=false
		for _,v:= range b{
			if v==ch{
				exists=true
				break
			}
		}
		if exists==false{
			b=append(b,ch)
		}

	}
	fmt.Println(string(b))
}