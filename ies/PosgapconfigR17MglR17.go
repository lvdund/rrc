package ies

import "rrc/utils"

// PosGapConfig-r17-mgl-r17 ::= ENUMERATED
type PosgapconfigR17MglR17 struct {
	Value utils.ENUMERATED
}

const (
	PosgapconfigR17MglR17EnumeratedNothing = iota
	PosgapconfigR17MglR17EnumeratedMs1dot5
	PosgapconfigR17MglR17EnumeratedMs3
	PosgapconfigR17MglR17EnumeratedMs3dot5
	PosgapconfigR17MglR17EnumeratedMs4
	PosgapconfigR17MglR17EnumeratedMs5dot5
	PosgapconfigR17MglR17EnumeratedMs6
	PosgapconfigR17MglR17EnumeratedMs10
	PosgapconfigR17MglR17EnumeratedMs20
)
