package main

import (
    "errors"
    "fmt"
)

var notfound=errors.New("not found")
func err1(id int)(error){
   if id==1{
	return nil
   }else if id==2{
	return fmt.Errorf("error is %w",notfound)
   }else{
     return errors.New("value not found")
   }
}
func main(){
	err:=err1(1)
    if errors.Is(err,notfound){
       fmt.Println("the value not found")
	}else if err!=nil{
      fmt.Println("error is",err)
	}else{
		fmt.Println("no error")
	}
}