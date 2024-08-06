package expr
import (
	"errors"


	"calculator/src/strings"
)


func is_empty (bs []byte) error {
	var this string = strings.Expr
	var lim = len(bs)


	for i:=0; i<lim; i++ {
		if  bs[i] > 0x20			{ return nil }
	}
	

	return errors.New(this + strings.Err_empty_arr)
}
