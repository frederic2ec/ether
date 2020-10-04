package repl

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/frederic2ec/ether/evaluator"
	"github.com/frederic2ec/ether/lexer"
	"github.com/frederic2ec/ether/object"
	"github.com/frederic2ec/ether/parser"
)

const PROMPT = "#=> "

type Options struct {
	Interactive bool
}

type REPL struct {
	user string
	args []string
	opts *Options
}

func New(user string, args []string, opts *Options) *REPL {
	object.StandardInput = os.Stdin
	object.StandardOutput = os.Stdout
	object.ExitFunction = os.Exit

	return &REPL{user, args, opts}
}

func (r *REPL) Eval(f io.Reader, out io.Writer) (env *object.Environment) {
	env = object.NewEnvironment()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading source file: %s", err)
		return
	}

	l := lexer.New(string(b))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(os.Stderr, p.Errors())
		return
	}

	evaluator.Eval(program, env)
	return
}

func (r *REPL) StartEvalLoop(in io.Reader, out io.Writer, env *object.Environment) {
	scanner := bufio.NewScanner(in)

	if env == nil {
		env = object.NewEnvironment()
	}

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		obj := evaluator.Eval(program, env)
		if obj != nil {
			if _, ok := obj.(*object.Null); !ok {
				io.WriteString(out, obj.Inspect())
				io.WriteString(out, "\n")
			}
		}
	}
}

func (r *REPL) Run() {
	object.Arguments = make([]string, len(r.args))
	copy(object.Arguments, r.args)

	if len(r.args) == 0 {
		fmt.Printf("Hello %s! This is the Ether programming language!\n", r.user)
		fmt.Printf("Feel free to type in commands\n")
		r.StartEvalLoop(os.Stdin, os.Stdout, nil)
		return
	}

	if len(r.args) > 0 {
		f, err := os.Open(r.args[0])
		if err != nil {
			log.Fatalf("could not open source file %s: %s", r.args[0], err)
		}

		// Remove program argument (zero)
		r.args = r.args[1:]
		object.Arguments = object.Arguments[1:]
		env := r.Eval(f, os.Stdout)
		if r.opts.Interactive {
			r.StartEvalLoop(os.Stdin, os.Stdout, env)
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "Woops! We ran into some issue here!\n")
		io.WriteString(out, " parser errors:\n")
		io.WriteString(out, "\t"+msg+"\n")
	}
}
