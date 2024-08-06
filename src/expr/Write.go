package expr
import (
	"io"
)




func (e *Expr) Write (bs []byte) (int, error) {
	var lim int = len(bs)
	var middle int
	var err error


	if err = is_empty(bs);					err != nil				{ return 0, err }
	if middle, err = e.write_operator(bs); 	err != nil 				{ return 0, err }
	if err = e.write_mode(bs); 				err != nil				{ return 0, err }
	if err = e.write_numbers(bs, middle);	err != nil				{ return 0, err }
	if err = e.write_result(); 				err != nil				{ return 0, err }

	
	return lim, io.EOF;
}
