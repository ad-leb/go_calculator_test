package expr


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
	


type Expr struct {
	op operator
	mode writing
	first int
	second int
	result int
}
