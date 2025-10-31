package ies

import "rrc/utils"

// MinTimeGapFR2-2-r17-scs-960kHz-r17 ::= ENUMERATED
type Mintimegapfr22R17Scs960khzR17 struct {
	Value utils.ENUMERATED
}

const (
	Mintimegapfr22R17Scs960khzR17EnumeratedNothing = iota
	Mintimegapfr22R17Scs960khzR17EnumeratedSl16
	Mintimegapfr22R17Scs960khzR17EnumeratedSl192
)
