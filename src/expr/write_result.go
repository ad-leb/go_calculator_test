package expr
import (
	"errors"


	"calculator/src/strings"
)


func (e *Expr) write_result () error {
	var this string = strings.Expr



	switch  e.op  {
		case sum:		e.result = e.first + e.second
		case sub:		e.result = e.first - e.second
		case mul:		e.result = e.first * e.second
		case div:		e.result = e.first / e.second
	}




	if  e.mode == roman  &&  e.result <= 0 { 
		return errors.New(this + strings.Err_roman_neg) 
	} else {
		return nil
	}
}
