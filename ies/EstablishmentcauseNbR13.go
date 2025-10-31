package ies

import "rrc/utils"

// EstablishmentCause-NB-r13 ::= ENUMERATED
type EstablishmentcauseNbR13 struct {
	Value utils.ENUMERATED
}

const (
	EstablishmentcauseNbR13EnumeratedNothing = iota
	EstablishmentcauseNbR13EnumeratedMt_Access
	EstablishmentcauseNbR13EnumeratedMo_Signalling
	EstablishmentcauseNbR13EnumeratedMo_Data
	EstablishmentcauseNbR13EnumeratedMo_Exceptiondata
	EstablishmentcauseNbR13EnumeratedDelaytolerantaccess_V1330
	EstablishmentcauseNbR13EnumeratedMt_Edt_V1610
	EstablishmentcauseNbR13EnumeratedSpare2
	EstablishmentcauseNbR13EnumeratedSpare1
)
