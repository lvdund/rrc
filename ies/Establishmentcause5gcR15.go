package ies

import "rrc/utils"

// EstablishmentCause-5GC-r15 ::= ENUMERATED
type Establishmentcause5gcR15 struct {
	Value utils.ENUMERATED
}

const (
	Establishmentcause5gcR15Emergency          = 0
	Establishmentcause5gcR15Highpriorityaccess = 1
	Establishmentcause5gcR15MtAccess           = 2
	Establishmentcause5gcR15MoSignalling       = 3
	Establishmentcause5gcR15MoData             = 4
	Establishmentcause5gcR15MoVoicecall        = 5
	Establishmentcause5gcR15Spare2             = 6
	Establishmentcause5gcR15Spare1             = 7
)
