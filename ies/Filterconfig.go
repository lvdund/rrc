package ies

// FilterConfig ::= SEQUENCE
type Filterconfig struct {
	Filtercoefficientrsrp   Filtercoefficient
	Filtercoefficientrsrq   Filtercoefficient
	FiltercoefficientrsSinr Filtercoefficient
}
