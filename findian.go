package main
import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main(){
	Scanner := bufio.NewScanner(os.Stdin)
	var data string
	fmt.Printf("Enter a string ")
	Scanner.Scan()
	data = Scanner.Text()
	// fmt.Scanln(&data)
	data = strings.ToLower(data)
	if (data[0] == 'i' && data[len(data) - 1 ] == 'n') && strings.Contains(data, "a") {
			 fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not found!\n")
	}
}