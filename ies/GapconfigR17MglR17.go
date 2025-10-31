package ies

import "rrc/utils"

// GapConfig-r17-mgl-r17 ::= ENUMERATED
type GapconfigR17MglR17 struct {
	Value utils.ENUMERATED
}

const (
	GapconfigR17MglR17EnumeratedNothing = iota
	GapconfigR17MglR17EnumeratedMs1
	GapconfigR17MglR17EnumeratedMs1dot5
	GapconfigR17MglR17EnumeratedMs2
	GapconfigR17MglR17EnumeratedMs3
	GapconfigR17MglR17EnumeratedMs3dot5
	GapconfigR17MglR17EnumeratedMs4
	GapconfigR17MglR17EnumeratedMs5
	GapconfigR17MglR17EnumeratedMs5dot5
	GapconfigR17MglR17EnumeratedMs6
	GapconfigR17MglR17EnumeratedMs10
	GapconfigR17MglR17EnumeratedMs20
)
