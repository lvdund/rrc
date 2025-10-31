package ies

import "rrc/utils"

// SDT-Config-r17-sdt-SRB2-Indication-r17 ::= ENUMERATED
type SdtConfigR17SdtSrb2IndicationR17 struct {
	Value utils.ENUMERATED
}

const (
	SdtConfigR17SdtSrb2IndicationR17EnumeratedNothing = iota
	SdtConfigR17SdtSrb2IndicationR17EnumeratedAllowed
)
