package main
import "fmt"
func secondlargestelement(arr1 []int)int{
	max1:=arr1[0]
	max2:=arr1[1]
	if max2>max1{
		max1=arr1[1]
		max2=arr1[0]
	}
	for i:=2;i<len(arr1);i++{
        if max1<arr1[i]{
			max2= max1
			max1=arr1[i]
		}else if arr1[i]>max2{
			max2=arr1[i]
		}
		}
	return max2
}

func main(){
	arr1:=[]int{4,3,9,1,5}
	sl:=secondlargestelement(arr1)
	fmt.Println("second largest element is",sl)
}