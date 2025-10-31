package ies

// UE-TxTEG-AssociationList-r17 ::= SEQUENCE OF UE-TxTEG-Association-r17
// SIZE (1..maxNrOfTxTEGReport-r17)
type UeTxtegAssociationlistR17 struct {
	Value []UeTxtegAssociationR17 `lb:1,ub:maxNrOfTxTEGReportR17`
}
