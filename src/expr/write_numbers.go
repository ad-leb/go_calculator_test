package expr
import (
	"errors"


	"calculator/src/ascii"
	"calculator/src/strings"
)



func (e *Expr) write_numbers (bs []byte, middle int) error {
	var this string = strings.Expr
	var foo func([]byte)(int, error)
	var err error



	if e.mode == arabic { 
		foo = ascii.Atoi
	} else {
		foo = ascii.Rtoi
	}


	e.first, err = foo(bs[:middle])
	if 	err != nil						{ return err }
	if  e.first < 1 || 10 < e.first		{ return errors.New(this + strings.Err_wrong_number) }

	e.second, err = foo(bs[middle + 1:])
	if 	err != nil						{ return err }
	if  e.second < 1 || 10 < e.second	{ return errors.New(this + strings.Err_wrong_number) }



	return nil
}
