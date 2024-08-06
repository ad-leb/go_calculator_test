package ascii
import (
	"errors"


	"calculator/src/strings"
)




func Atoi (bs []byte) (int, error) {
	var this string = strings.Atoi
	const AV byte = 0x0f			/* ASCII digit value	*/
	const BD int = 10				/* bit depth 			*/
	var lim int = len(bs)
	var negative bool
	var res int = 0
	var i int = 0					/* bs index			*/



	if  lim == 0  								{ return 0, errors.New(this + strings.Err_empty_arr) }
	/* passing whitespaces	*/
	for ; bs[i]<=0x20; i++ {
		if i>=lim								{ return 0, errors.New(this + strings.Err_empty_arr) }
	}



	if bs[i] == '-'		{ negative = true; i++ }
	for ; i<lim && bs[i]>0x20; i++ {
		if ( bs[i] < '0' || bs[i] > '9' ) 		{ return 0, errors.New(this + strings.Err_arabic) }

		res = (res * BD) + int(bs[i] & AV)
	}

	if negative 			{ res *= -1 }


	/* check the rest of slice 	*/
	for ; i<lim; i++ {
		if  bs[i]> 0x20  						{ return 0, errors.New(this + strings.Err_foreign) }
	}



	return res, nil
}
