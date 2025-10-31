package ies

// FreqPriorityNR-r15 ::= SEQUENCE
type FreqprioritynrR15 struct {
	CarrierfreqR15                ArfcnValuenrR15
	CellreselectionpriorityR15    Cellreselectionpriority
	CellreselectionsubpriorityR15 *CellreselectionsubpriorityR13
}
