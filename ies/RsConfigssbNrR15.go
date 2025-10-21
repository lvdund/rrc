package ies

import "rrc/utils"

// RS-ConfigSSB-NR-r15 ::= SEQUENCE
// Extensible
type RsConfigssbNrR15 struct {
	MeastimingconfigR15     MtcSsbNrR15
	SubcarrierspacingssbR15 utils.ENUMERATED
}
