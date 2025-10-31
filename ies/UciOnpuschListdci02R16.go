package ies

// UCI-OnPUSCH-ListDCI-0-2-r16 ::= SEQUENCE OF UCI-OnPUSCH-DCI-0-2-r16
// SIZE (1..2)
type UciOnpuschListdci02R16 struct {
	Value []UciOnpuschDci02R16 `lb:1,ub:2`
}
