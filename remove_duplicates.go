package main
import "fmt"
func main(){
	a:=[]int{1,2,3,1,3,5}
	d:=[]int{}
	for i:=0;i<len(a);i++{
		ch:=a[i]
		exists:=false
		for _,v:= range d{
			if v==ch{
				exists=true
				break
			}
		}
		if exists==false{
			d=append(d,ch)
		}
	}
	fmt.Print("removing duplicates",d)
}