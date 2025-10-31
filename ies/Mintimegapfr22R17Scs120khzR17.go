package ies

import "rrc/utils"

// MinTimeGapFR2-2-r17-scs-120kHz-r17 ::= ENUMERATED
type Mintimegapfr22R17Scs120khzR17 struct {
	Value utils.ENUMERATED
}

const (
	Mintimegapfr22R17Scs120khzR17EnumeratedNothing = iota
	Mintimegapfr22R17Scs120khzR17EnumeratedSl2
	Mintimegapfr22R17Scs120khzR17EnumeratedSl24
)
