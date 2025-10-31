package ies

// BT-NameList-r15 ::= SEQUENCE OF BT-Name-r15
// SIZE (1..maxBT-Name-r15)
type BtNamelistR15 struct {
	Value []BtNameR15 `lb:1,ub:maxBTNameR15`
}
