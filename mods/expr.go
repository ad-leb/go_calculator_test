package mods
import (
	"io"
	"errors"
)


type writing byte
const (
	arabic writing		= iota
	roman
)
type operator byte
const (
	sum operator		= iota
	sub
	mul
	div
)


	

/*
 *
	Russion (UTF-8):

Основная структура -- Expr (expression). В ней содержатся:
	1. Тип записи чисел   ('arabic', 'roman')
	2. Тип операции   ('sum', 'sub', 'mul', 'div')
	3. Первый операнд   <int>
	4. Второй операнд   <int>
	5. Результат операции   <int>
Структура реализует интерфейсы Reader и Writer, с помощью которых логика 
вычисления инкапсулируется в методах 'Read' и 'Write'.
 *
 */
type Expr struct {
	mode writing
	op operator
	first int
	second int
	result int
}

/* io.Writer	*/
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
/* io.Reader	*/
func (e *Expr) Read (bs []byte) (int, error) {
	var foo func([]byte, int)error
	var err error



	if  e.mode == arabic {
		foo = Itoa
	} else {
		foo = Itor
	}

	err = foo(bs, e.result)
		if err != nil 		{ return 0, err }



	return len(bs), io.EOF
}










/*
 *
	Russian (UTF-8):

Здесь содержатся функции-этапы основного чтения методом Write:
	1. is_empty()
		|	функция проверки строки на наполненность. В случае отсутствия в ней
		|	символов, выдаёт соответствующую панику. Введена больше для декора 
		|	-- чтобы сообщение об ошибке выпадало соответствующее.
	2. write_operator()
		|	метод считывания оператора. Помимо ошибки возвращает индекс 
		|	положения оператора в байтовом срезе. Этот индекс позже использует-
		|	ся для деления этого среза на два соответствующих двум операндам.
	3. write_mode()
		|	метод установки режима записи чисел. Тип записи определяется по 
		|	первому встречному символу среза. Например, в выражении
		|		3 * x
		|	будет установлен режим 'arabic'.
	4. write_numbers()
		|	метод считываения обоих операндов. В начале его работы проверяется 
		|	режим записи. В соответствии с ним выбирается функция перевода
		|	'Atoi' или 'Rtoi' в локальной переменной 'foo'. Далее вызывается
		|	уже сама выбранная функция.
	5. write_result()
		|	метод. По установленным в объекте переменным 'op', 'first' и 
		|	'second' выбирается операция и выполняется вычисление с 
		|	последующей записью результата в объект.
 *
 */
func is_empty (bs []byte) error {
	var this string = Str_Expr
	var lim = len(bs)


	for i:=0; i<lim; i++ {
		if  bs[i] > 0x20			{ return nil }
	}
	

	return errors.New(this + Str_err_empty_arr)
}


func (e *Expr) write_operator (bs []byte) (int, error) {
	var this string = Str_Expr
	var lim int = len(bs)
	var i int


	for i=0; i<lim; i++ {
		switch ( bs[i] ) {
			case '+':			e.op = sum; return i, nil
			case '-':			e.op = sub; return i, nil
			case '*':			e.op = mul; return i, nil
			case '/':			e.op = div; return i, nil
		}
	}
	

	return 0, errors.New(this + Str_err_no_operator)
}


func (e *Expr) write_mode (bs []byte) error {
	var this string = Str_Expr
	var lim int = len(bs)
	var i int


	/* pass whitespaces	*/
	for i=0; bs[i]<=0x20; i++ {
		if i >= lim 				{ return errors.New(this + Str_err_empty_arr) }
	}

	if '0' <= bs[i] && bs[i] <= '9' {
		e.mode = arabic
	} else {
		e.mode = roman
	}


	return nil
}


func (e *Expr) write_numbers (bs []byte, middle int) error {
	var this string = Str_Expr
	var foo func([]byte)(int, error)
	var err error



	if e.mode == arabic { 
		foo = Atoi
	} else {
		foo = Rtoi
	}


	e.first, err = foo(bs[:middle])
	if 	err != nil						{ return err }
	if  e.first < 1 || 10 < e.first		{ return errors.New(this + Str_err_wrong_number) }

	e.second, err = foo(bs[middle + 1:])
	if 	err != nil						{ return err }
	if  e.second < 1 || 10 < e.second	{ return errors.New(this + Str_err_wrong_number) }



	return nil
}


func (e *Expr) write_result () error {
	var this string = Str_Expr



	switch  e.op  {
		case sum:		e.result = e.first + e.second
		case sub:		e.result = e.first - e.second
		case mul:		e.result = e.first * e.second
		case div:		e.result = e.first / e.second
	}




	if  e.mode == roman  &&  e.result <= 0 { 
		return errors.New(this + Str_err_roman_neg) 
	} else {
		return nil
	}
}
