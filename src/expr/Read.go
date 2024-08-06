package expr
import (
	"io"


	"calculator/src/ascii"
)


func (e *Expr) Read (bs []byte) (int, error) {
	var foo func([]byte, int)error
	var err error



	if  e.mode == arabic {
		foo = ascii.Itoa
	} else {
		foo = ascii.Itor
	}

	err = foo(bs, e.result)
		if err != nil 		{ return 0, err }



	return len(bs), io.EOF
}
