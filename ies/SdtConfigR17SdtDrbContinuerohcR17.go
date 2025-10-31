package ies

import "rrc/utils"

// SDT-Config-r17-sdt-DRB-ContinueROHC-r17 ::= ENUMERATED
type SdtConfigR17SdtDrbContinuerohcR17 struct {
	Value utils.ENUMERATED
}

const (
	SdtConfigR17SdtDrbContinuerohcR17EnumeratedNothing = iota
	SdtConfigR17SdtDrbContinuerohcR17EnumeratedCell
	SdtConfigR17SdtDrbContinuerohcR17EnumeratedRna
)
