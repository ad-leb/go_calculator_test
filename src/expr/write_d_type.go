package expr
import (
	"errors"


	"calculator/src/strings"
)



func (e *Expr) write_mode (bs []byte) error {
	var this string = strings.Expr
	var lim int = len(bs)
	var i int


	/* pass whitespaces	*/
	for i=0; bs[i]<=0x20; i++ {
		if i >= lim 				{ return errors.New(this + strings.Err_empty_arr) }
	}

	if '0' <= bs[i] && bs[i] <= '9' {
		e.mode = arabic
	} else {
		e.mode = roman
	}


	return nil
}
