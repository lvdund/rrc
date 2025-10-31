package ies

import "rrc/utils"

// EstablishmentCause ::= ENUMERATED
type Establishmentcause struct {
	Value utils.ENUMERATED
}

const (
	EstablishmentcauseEnumeratedNothing = iota
	EstablishmentcauseEnumeratedEmergency
	EstablishmentcauseEnumeratedHighpriorityaccess
	EstablishmentcauseEnumeratedMt_Access
	EstablishmentcauseEnumeratedMo_Signalling
	EstablishmentcauseEnumeratedMo_Data
	EstablishmentcauseEnumeratedDelaytolerantaccess_V1020
	EstablishmentcauseEnumeratedMo_Voicecall_V1280
	EstablishmentcauseEnumeratedSpare1
)
