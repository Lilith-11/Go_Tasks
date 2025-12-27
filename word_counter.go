package main
import ("fmt"
"strings")
func main(){
	word:="I am a go developer"
    words:=strings.ToLower(word)
	wordss:=strings.Fields(words)
	fmt.Println(word)
	fmt.Println(words)
	fmt.Println(wordss)
	frequency:=make(map[string]int)
	for _,wor := range wordss{
		frequency[wor]++
	}
	fmt.Println(frequency)
	for i,value := range frequency{
		fmt.Println(i,":",value)
	}
}