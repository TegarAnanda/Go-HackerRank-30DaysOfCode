package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReader(os.Stdin)
	var message string
	message, _ = reader.ReadString('\n')

	fmt.Printf("Hello, World.\n%s", message)
}

