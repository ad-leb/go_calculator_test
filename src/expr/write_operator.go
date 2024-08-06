package expr
import (
	"errors"


	"calculator/src/strings"
)



func (e *Expr) write_operator (bs []byte) (int, error) {
	var this string = strings.Expr
	var lim int = len(bs)
	var i int


	for i=0; i<lim; i++ {
		switch ( bs[i] ) {
			case '+':			e.op = sum; return i, nil
			case '-':			e.op = sub; return i, nil
			case '*':			e.op = mul; return i, nil
			case '/':			e.op = div; return i, nil
		}
	}
	

	return 0, errors.New(this + strings.Err_no_operator)
}
