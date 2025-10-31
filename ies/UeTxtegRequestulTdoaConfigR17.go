package ies

import "rrc/utils"

// UE-TxTEG-RequestUL-TDOA-Config-r17 ::= CHOICE
const (
	UeTxtegRequestulTdoaConfigR17ChoiceNothing = iota
	UeTxtegRequestulTdoaConfigR17ChoiceOneshotR17
	UeTxtegRequestulTdoaConfigR17ChoicePeriodicreportingR17
)

type UeTxtegRequestulTdoaConfigR17 struct {
	Choice               uint64
	OneshotR17           *struct{}
	PeriodicreportingR17 *utils.ENUMERATED
}
