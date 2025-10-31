package ies

// MultiBandInfoList-r11 ::= SEQUENCE OF FreqBandIndicator-r11
// SIZE (1..maxMultiBands)
type MultibandinfolistR11 struct {
	Value []FreqbandindicatorR11 `lb:1,ub:maxMultiBands`
}
