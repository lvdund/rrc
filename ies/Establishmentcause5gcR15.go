package ies

import "rrc/utils"

// EstablishmentCause-5GC-r15 ::= ENUMERATED
type Establishmentcause5gcR15 struct {
	Value utils.ENUMERATED
}

const (
	Establishmentcause5gcR15EnumeratedNothing = iota
	Establishmentcause5gcR15EnumeratedEmergency
	Establishmentcause5gcR15EnumeratedHighpriorityaccess
	Establishmentcause5gcR15EnumeratedMt_Access
	Establishmentcause5gcR15EnumeratedMo_Signalling
	Establishmentcause5gcR15EnumeratedMo_Data
	Establishmentcause5gcR15EnumeratedMo_Voicecall
	Establishmentcause5gcR15EnumeratedSpare2
	Establishmentcause5gcR15EnumeratedSpare1
)
