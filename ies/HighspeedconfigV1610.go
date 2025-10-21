package ies

import "rrc/utils"

// HighSpeedConfig-v1610 ::= SEQUENCE
type HighspeedconfigV1610 struct {
	Highspeedenhmeasflag2R16  *utils.ENUMERATED
	Highspeedenhdemodflag2R16 *utils.ENUMERATED
}
