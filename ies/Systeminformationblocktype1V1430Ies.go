package ies

import "rrc/utils"

// SystemInformationBlockType1-v1430-IEs ::= SEQUENCE
type Systeminformationblocktype1V1430Ies struct {
	EcalloverimsSupportR14       *utils.ENUMERATED
	TddConfigV1430               *TddConfigV1430
	CellaccessrelatedinfolistR14 *interface{}
	Noncriticalextension         *Systeminformationblocktype1V1450Ies
}
