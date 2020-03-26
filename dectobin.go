package main
import (
"fmt"
"strconv"
)

func main(){
var decimal int64

fmt.Println("Enter decimal number")
fmt.Scanln(decimal)

output := strconv.FormatInt(decimal, 2)
 fmt.Println("Output ", output)
	}