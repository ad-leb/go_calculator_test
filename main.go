package main
import (
	"os"
	"io"
)



func main () {
	var e *Expr = new(Expr)
	var err error



	for {
		if _, err = io.Copy(e, os.Stdin); err != nil && err != io.EOF		{ panic(err) }
		if _, err = io.Copy(os.Stdout, e); err != nil && err != io.EOF		{ panic(err) }
		if _, err = os.Stdout.Write(Str_newline); err != nil 				{ panic(err) }
	}
}
