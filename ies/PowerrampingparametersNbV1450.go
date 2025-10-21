package ies

import "rrc/utils"

// PowerRampingParameters-NB-v1450 ::= SEQUENCE
type PowerrampingparametersNbV1450 struct {
	PreambleinitialreceivedtargetpowerV1450 *utils.ENUMERATED
	Powerrampingparametersce1R14            *interface{}
}
