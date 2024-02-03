package main
import "fmt"
//  write the syntax of range
func main() {

	numbers:=[]int{1,2,3,4,5}
	for index,number:= range numbers {
		fmt.Printf("index %d: have stored: %d \n",index,number)
		}
}