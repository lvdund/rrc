package ies

// QuantityConfig ::= SEQUENCE
// Extensible
type Quantityconfig struct {
	QuantityconfignrList     *[]Quantityconfignr       `lb:1,ub:maxNrofQuantityConfig`
	Quantityconfigeutra      *Filterconfig             `ext`
	QuantityconfigutraFddR16 *QuantityconfigutraFddR16 `ext`
	QuantityconfigcliR16     *FilterconfigcliR16       `ext`
}
