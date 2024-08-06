package ascii
import (
	"errors"


	"calculator/src/strings"
)



func Rtoi (bs []byte) (int, error) {
	var this string = strings.Rtoi
	var lim int = len(bs)
	var curr byte = 0
	var next byte = 0
	var res int = 0
	var i int = 0				/* bs index (using by 'next')	*/
	


	if ( lim == 0 ) 			{ return 0, errors.New(this + strings.Err_empty_arr) } 
	for ; bs[i]<=0x20; i++ {	/* passing whitespaces	*/
		if i>=lim				{ return 0, errors.New(this + strings.Err_empty_arr) }
	}


	curr=(bs[i]&0x5f) 
	i++
	for ; curr>0x20; curr=next {
		if i >= lim	{ next = 0 }		else { next = bs[i] & 0x5f }

		switch ( curr ) {
			case 'V':		res += 5
			case 'X':		res += 10
			case 'I':	if  next == 'V' || next == 'X' { 
							res -= 1		
						} else {
							res += 1
						}
			default	:			return res, errors.New(this + strings.Err_roman);
		}
		i++
	}


	/* check the rest of slice 	*/
	for ; i<lim; i++ {
		if  bs[i]> 0x20  						{ return 0, errors.New(this + strings.Err_foreign) }
	}


	
	return res, nil
}
