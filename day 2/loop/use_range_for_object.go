package main
import "fmt"
//  write the syntax of range
func main() {

	myMap:=map[string]int{"A":1,"B":2,"C":3,"D":4,"E":5,"F":6,"G":7}
	for key,value:= range myMap {
		fmt.Printf("key %s: has value: %d \n",key,value)
		}
}