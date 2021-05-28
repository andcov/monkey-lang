package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"regexp"

	"github.com/andcov/monkey-lang/evaluator"
	"github.com/andcov/monkey-lang/lexer"
	"github.com/andcov/monkey-lang/object"
	"github.com/andcov/monkey-lang/parser"
	"github.com/andcov/monkey-lang/repl"
)

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	args := os.Args[1:]

	if len(args) == 1 {
		runFromFile(args[0])
	} else {
		fmt.Print(MONKEY_FACE)
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}

func runFromFile(filename string) {
	isFileNameOk, err := regexp.Match(`^[\w-_\.]+$`, []byte(filename))

	if err != nil {
		log.Fatal(err)
	}

	if isFileNameOk {
		content, err := ioutil.ReadFile(filename)

		if err != nil {
			log.Fatal(err)
		}

		l := lexer.New(string(content))
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(p.Errors())
			return
		}

		env := object.NewEnvironment()
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			fmt.Print(evaluated.Inspect())
		}
	}
}

func printParserErrors(errs []string) {
	fmt.Println("Woops! We ran into some monkey business here!")
	fmt.Println(" parser errors:")
	for _, msg := range errs {
		fmt.Print("\t" + msg + "\n")
	}
}
