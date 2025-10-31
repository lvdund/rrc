package ies

import "rrc/utils"

// PUSCH-ConfigCommon-pusch-ConfigBasic-hoppingMode ::= ENUMERATED
type PuschConfigcommonPuschConfigbasicHoppingmode struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigcommonPuschConfigbasicHoppingmodeEnumeratedNothing = iota
	PuschConfigcommonPuschConfigbasicHoppingmodeEnumeratedIntersubframe
	PuschConfigcommonPuschConfigbasicHoppingmodeEnumeratedIntraandintersubframe
)
