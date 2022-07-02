package main

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"math"
	"os"
	"os/exec"
	"strconv"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clear()
	for {
		var in1 string
		var in2 string
		var in3 string
		var num1 float64
		var num2 float64
		var result float64
		fmt.Print("Input : ")
		fmt.Scanln(&in1, &in2, &in3)
		num1, _ = strconv.ParseFloat(in1, 64)
		num2, _ = strconv.ParseFloat(in3, 64)
		if in2 == "+" {
			result = num1 + num2
		} else if in2 == "-" {
			result = num1 - num2
		} else if in2 == "*" {
			result = num1 * num2
		} else if in2 == "/" {
			result = num1 / num2
		} else if in2 == "%" {
			result = math.Remainder(num1, num2)
		} else if in2 == "!" {
			result = 1
			for i := 1.00; i <= num1; i++ {
				result = result * i
			}
		} else if in2 == "^" {
			r := math.Pow(num1, num2)
			result = r
		} else if in1 == "quit" || in1 == "exit" || in1 == "stop" {
			break
		}

		if in1 == "Hello" || in1 == "hello" {
			fmt.Println("Hello~")
		} else if in1 == "clear" || in1 == "cls" {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		} else if in1 == "fuck" || in1 == "FUCK" {
			fmt.Println("    _\n   | |\n   | |\n __| |__\n|       |\n|_______|")
		} else if in1 == "?" || in1 == "help" {
			fmt.Println("----------\n[Form]\nnum1 string num2\n\n[Expression]\n1. + : num1 + num2\n2. - : num1 - num2\n3. * : num1 * num2\n4. / : num1 / num2\n5. % : num1 % num2\n6. ! : num1's factorial\n7. ^ : num1 ^ num2\n\n[Command]\n1. ? or help : Help\n2. clear or cls : Screen Clear\n3. exit or quit or stop: Program exit\n4. ????\n5. ?????\n----------")
		} else {
			p := message.NewPrinter(language.English)
			s := p.Sprint(result)
			fmt.Println("Result :", s)
		}
	}
}
