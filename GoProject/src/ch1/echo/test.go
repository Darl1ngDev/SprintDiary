package main
import (
	"fmt"
	"os"
)

func main(){
	// var s, sep string
	fmt.Println(os.Args[0])
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}