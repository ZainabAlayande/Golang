package main
import "fmt"

func main() {
    //var(
      //a = 5
      //b = 10
      //c = 15
     //)

    fmt.Print("Enter a number: ")
    var input float64
	
    fmt.Scanf("%f", &input)

    output := input * 2
    fmt.Println(output)
}