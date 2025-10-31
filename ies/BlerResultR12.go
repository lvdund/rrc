package ies

// BLER-Result-r12 ::= SEQUENCE
type BlerResultR12 struct {
	BlerR12           BlerRangeR12
	BlocksreceivedR12 BlerResultR12BlocksreceivedR12
}
