package ies

import "rrc/utils"

// CSI-RS-ConfigBeamformed-v1430 ::= SEQUENCE
type CsiRsConfigbeamformedV1430 struct {
	CsiRsConfignzpAplistR14        *interface{}
	NzpResourceconfigoriginalV1430 *CsiRsConfigNzpV1430
	CsiRsNzpActivationR14          *CsiRsConfignzpActivationR14
}
