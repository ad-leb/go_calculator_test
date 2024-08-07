package mods
import (
	"errors"
)





/* Arabic to Integer 	*/
func Atoi (bs []byte) (int, error) {
	var this string = Str_Atoi
	const AV byte = 0x0f			/* ASCII digit value	*/
	const BD int = 10				/* bit depth 			*/
	var lim int = len(bs)
	var negative bool
	var res int = 0
	var i int = 0					/* bs index			*/



	if  lim == 0  								{ return 0, errors.New(this + Str_err_empty_arr) }
	/* passing whitespaces	*/
	for ; bs[i]<=0x20; i++ {
		if i>=lim								{ return 0, errors.New(this + Str_err_empty_arr) }
	}



	if bs[i] == '-'		{ negative = true; i++ }
	for ; i<lim && bs[i]>0x20; i++ {
		if ( bs[i] < '0' || bs[i] > '9' ) 		{ return 0, errors.New(this + Str_err_arabic) }

		res = (res * BD) + int(bs[i] & AV)
	}

	if negative 			{ res *= -1 }


	/* check the rest of slice 	*/
	for ; i<lim; i++ {
		if  bs[i]> 0x20  						{ return 0, errors.New(this + Str_err_foreign) }
	}



	return res, nil
}










/* Roman to Integer 	*/
func Rtoi (bs []byte) (int, error) {
	var this string = Str_Rtoi
	var lim int = len(bs)
	var curr byte = 0
	var next byte = 0
	var res int = 0
	var i int = 0				/* bs index (using by 'next')	*/
	


	if ( lim == 0 ) 			{ return 0, errors.New(this + Str_err_empty_arr) } 
	for ; bs[i]<=0x20; i++ {	/* passing whitespaces	*/
		if i>=lim				{ return 0, errors.New(this + Str_err_empty_arr) }
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
			default	:			return res, errors.New(this + Str_err_roman);
		}
		i++
	}


	/* check the rest of slice 	*/
	for ; i<lim; i++ {
		if  bs[i]> 0x20  						{ return 0, errors.New(this + Str_err_foreign) }
	}


	
	return res, nil
}













/* Integer to Arabic	*/
func Itoa (bs []byte, digit int) error {
	var this string = Str_Itoa
	const AP byte = 0x30			/* ASCII digit prefix 	*/
	const BD int = 10				/* bit depth 			*/
	const SS int = 16				/* stack size			*/
	var lim int = len(bs)
	var negative bool
	var stack [SS]byte 
	var p int = 0					/* stack pointer		*/
	var i int = 0					/* bs index			*/



	/* push characters to stack in reverse order */
	if digit < 0 		{ negative = true; digit *= -1 }
	if digit == 0 		{ stack[p] = '0'; p++ }

	for ; digit != 0; p++ {
		if  p >= SS 									{ return errors.New(this + Str_err_big_number) }
		stack[p] = byte((digit % BD)) | AP
		digit /= 10
	}



	/* check array(slice) for capacity */
	if  p >= lim 										{ return errors.New(this + Str_err_lil_buffer) }

	/* pop characters from stack to bsing in direct order */
	if negative 		{ bs[i] = '-'; i++ }

	for ; p >= 0; i++ { 
		bs[i] = stack[p]
		p--
	}



	/* erase the rest of array */
	for ; i < lim; i++ {
		bs[i] = 0
	}



	return nil
}













/* Integer to Roman		*/
func Itor (bs []byte, digit int) error {
	var this string = Str_Itor
	var lim int = len(bs)
	var i int
	

	if digit <= 0									{ return errors.New(this + Str_err_roman_neg) }

	for i=0; digit > 0; i++ {
		if i >= lim 								{ return errors.New(this + Str_err_lil_buffer) }

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



/* really heavy and scary. But they works good	*/
