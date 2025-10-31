package ies

// PUR-NRSRP-ChangeThreshold-NB-r16 ::= SEQUENCE
type PurNrsrpChangethresholdNbR16 struct {
	IncreasethreshR16 NrsrpChangethreshNbR16
	DecreasethreshR16 *NrsrpChangethreshNbR16
}
