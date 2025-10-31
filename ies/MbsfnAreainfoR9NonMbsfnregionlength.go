package ies

import "rrc/utils"

// MBSFN-AreaInfo-r9-non-MBSFNregionLength ::= ENUMERATED
type MbsfnAreainfoR9NonMbsfnregionlength struct {
	Value utils.ENUMERATED
}

const (
	MbsfnAreainfoR9NonMbsfnregionlengthEnumeratedNothing = iota
	MbsfnAreainfoR9NonMbsfnregionlengthEnumeratedS1
	MbsfnAreainfoR9NonMbsfnregionlengthEnumeratedS2
)
