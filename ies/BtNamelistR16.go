package ies

// BT-NameList-r16 ::= SEQUENCE OF BT-Name-r16
// SIZE (1..maxBT-Name-r16)
type BtNamelistR16 struct {
	Value []BtNameR16 `lb:1,ub:maxBTNameR16`
}
