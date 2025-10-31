package ies

// UCI-OnPUSCH-ListDCI-0-1-r16 ::= SEQUENCE OF UCI-OnPUSCH
// SIZE (1..2)
type UciOnpuschListdci01R16 struct {
	Value []UciOnpusch `lb:1,ub:2`
}
