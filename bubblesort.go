package main
import "fmt"
func main(){
	arr:=[]int{3,2,1,4,5}
	for j:=0;j<len(arr);j++{
	  for i:=0;i<len(arr)-j-1;i++{
		if arr[i]>arr[i+1]{
            arr[i],arr[i+1]=arr[i+1],arr[i]
		}}
	}
	fmt.Println("bubble sort",arr)

}