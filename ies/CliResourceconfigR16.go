package ies

import "rrc/utils"

// CLI-ResourceConfig-r16 ::= SEQUENCE
type CliResourceconfigR16 struct {
	SrsResourceconfigR16  *utils.Setuprelease[SrsResourcelistconfigcliR16]
	RssiResourceconfigR16 *utils.Setuprelease[RssiResourcelistconfigcliR16]
}
