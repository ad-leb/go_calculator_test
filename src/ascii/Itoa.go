package ascii
import (
	"errors"


	"calculator/src/strings"
)



func Itoa (bs []byte, digit int) error {
	var this string = strings.Itoa
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
		if  p >= SS 									{ return errors.New(this + strings.Err_big_number) }
		stack[p] = byte((digit % BD)) | AP
		digit /= 10
	}



	/* check array(slice) for capacity */
	if  p >= lim 										{ return errors.New(this + strings.Err_lil_buffer) }

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
