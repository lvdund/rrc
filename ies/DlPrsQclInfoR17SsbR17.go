package ies

import "rrc/utils"

// DL-PRS-QCL-Info-r17-ssb-r17 ::= SEQUENCE
// Extensible
type DlPrsQclInfoR17SsbR17 struct {
	SsbIndexR17 utils.INTEGER `lb:0,ub:63`
	RsTypeR17   DlPrsQclInfoR17SsbR17RsTypeR17
}
