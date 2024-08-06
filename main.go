package main
import (
	"os"
	"io"

	"calculator/src/expr"
	"calculator/src/strings"
)



func main () {
	var e *expr.Expr = new(expr.Expr)
	var err error



	for {
		if _, err = io.Copy(e, os.Stdin); err != nil && err != io.EOF		{ panic(err) }
		if _, err = io.Copy(os.Stdout, e); err != nil && err != io.EOF		{ panic(err) }
		if _, err = os.Stdout.Write(strings.Newline); err != nil 			{ panic(err) }
	}
}
