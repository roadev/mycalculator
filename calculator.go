package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct{}

func (Calc) Operate(input string, operator string) int {
	cleanInput := strings.Split(input, operator)
	operator1 := parse(cleanInput[0])
	operator2 := parse(cleanInput[1])
	switch operator {
	case "+":
		fmt.Println(operator1 + operator2)
		return operator1 + operator2
	case "-":
		fmt.Println(operator1 - operator2)
		return operator1 - operator2
	case "*":
		fmt.Println(operator1 * operator2)
		return operator1 * operator2
	case "/":
		fmt.Println(operator1 / operator2)
		return operator1 / operator2
	default:
		fmt.Println("No está soportado")
		return 0
	}
}

func parse(input string) int {
	operator, _ := strconv.Atoi(input)
	return operator
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	input := ReadInput()
	operator := ReadInput()
	c := Calc{}
	fmt.Println(c.Operate(input, operator))
}
