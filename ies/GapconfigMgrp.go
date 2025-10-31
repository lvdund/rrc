package ies

import "rrc/utils"

// GapConfig-mgrp ::= ENUMERATED
type GapconfigMgrp struct {
	Value utils.ENUMERATED
}

const (
	GapconfigMgrpEnumeratedNothing = iota
	GapconfigMgrpEnumeratedMs20
	GapconfigMgrpEnumeratedMs40
	GapconfigMgrpEnumeratedMs80
	GapconfigMgrpEnumeratedMs160
)
