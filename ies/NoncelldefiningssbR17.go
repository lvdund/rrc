package ies

// NonCellDefiningSSB-r17 ::= SEQUENCE
// Extensible
type NoncelldefiningssbR17 struct {
	AbsolutefrequencyssbR17 ArfcnValuenr
	SsbPeriodicityR17       *NoncelldefiningssbR17SsbPeriodicityR17
	SsbTimeoffsetR17        *NoncelldefiningssbR17SsbTimeoffsetR17
}
