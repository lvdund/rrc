package ies

// CC-State-r17 ::= SEQUENCE
type CcStateR17 struct {
	DlcarrierR17 *CarrierstateR17
	UlcarrierR17 *CarrierstateR17
}
