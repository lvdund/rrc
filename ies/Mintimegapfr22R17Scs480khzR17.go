package ies

import "rrc/utils"

// MinTimeGapFR2-2-r17-scs-480kHz-r17 ::= ENUMERATED
type Mintimegapfr22R17Scs480khzR17 struct {
	Value utils.ENUMERATED
}

const (
	Mintimegapfr22R17Scs480khzR17EnumeratedNothing = iota
	Mintimegapfr22R17Scs480khzR17EnumeratedSl8
	Mintimegapfr22R17Scs480khzR17EnumeratedSl96
)
