package ascii
import (
	"errors"


	"calculator/src/strings"
)



func Itor (bs []byte, digit int) error {
	var this string = strings.Itor
	var lim int = len(bs)
	var i int
	

	if digit <= 0									{ return errors.New(this + strings.Err_roman_neg) }

	for i=0; digit > 0; i++ {
		if i >= lim 								{ return errors.New(this + strings.Err_lil_buffer) }

		switch  {
		case digit >= 100:		bs[i] = 'C';					digit -= 100
		case digit >= 90:		bs[i] = 'X'; i++; bs[i] = 'C';	digit -= 90	
		case digit >= 50:		bs[i] = 'L';					digit -= 50
		case digit >= 40:		bs[i] = 'X'; i++; bs[i] = 'L';	digit -= 40	
		case digit >= 10:		bs[i] = 'X'; 					digit -= 10	
		case digit >= 9	:		bs[i] = 'I'; i++; bs[i] = 'X';	digit -= 9	
		case digit >= 5	:		bs[i] = 'V'; 					digit -= 5	
		case digit >= 4	:		bs[i] = 'I'; i++; bs[i] = 'V';	digit -= 4	
		default			:		bs[i] = 'I';					digit--
		}
	}

	/* erase the rest of array */
	for ; i < lim; i++ {
		bs[i] = 0
	}


	return nil
}
