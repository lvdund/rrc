package ies

import "rrc/utils"

// DL-PRS-QCL-Info-r17-ssb-r17-rs-Type-r17 ::= ENUMERATED
type DlPrsQclInfoR17SsbR17RsTypeR17 struct {
	Value utils.ENUMERATED
}

const (
	DlPrsQclInfoR17SsbR17RsTypeR17EnumeratedNothing = iota
	DlPrsQclInfoR17SsbR17RsTypeR17EnumeratedTypec
	DlPrsQclInfoR17SsbR17RsTypeR17EnumeratedTyped
	DlPrsQclInfoR17SsbR17RsTypeR17EnumeratedTypec_Plus_Typed
)
