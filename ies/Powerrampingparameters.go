package ies

import "rrc/utils"

// PowerRampingParameters ::= SEQUENCE
type Powerrampingparameters struct {
	Powerrampingstep                   utils.ENUMERATED
	Preambleinitialreceivedtargetpower utils.ENUMERATED
}
