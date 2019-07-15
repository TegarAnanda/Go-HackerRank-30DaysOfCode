package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
) 

func evenOrOdd(n int, arrInput string) {
    for i:=n; i<len(arrInput); i+=2 {
        fmt.Printf("%c", arrInput[i])
    }
    fmt.Print(" ")
}

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var T int

    fmt.Scan(&T)
    arrInput := make([]string, T)
    reader := bufio.NewReader(os.Stdin)

    for i:=0; i<T; i++ {
        arrInput[i], _ = reader.ReadString('\n')
        arrInput[i] = strings.TrimSuffix(arrInput[i], "\n")
    }

    for j:= range arrInput {
        evenOrOdd(0, arrInput[j])
        evenOrOdd(1, arrInput[j])
        fmt.Printf("\n")
    }
}

