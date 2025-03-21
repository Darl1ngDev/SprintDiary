// dup1 打印标准输入中多次出现的行的个数和文本
// 它从标准输入读取，然后将出现次数大于1的行打印出来
package main 
import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		counts[input.Text()]++
	}
	// 注意：忽略 input.Err() 中可能的错误
	for line, n := range counts{
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}