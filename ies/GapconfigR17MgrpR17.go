package ies

import "rrc/utils"

// GapConfig-r17-mgrp-r17 ::= ENUMERATED
type GapconfigR17MgrpR17 struct {
	Value utils.ENUMERATED
}

const (
	GapconfigR17MgrpR17EnumeratedNothing = iota
	GapconfigR17MgrpR17EnumeratedMs20
	GapconfigR17MgrpR17EnumeratedMs40
	GapconfigR17MgrpR17EnumeratedMs80
	GapconfigR17MgrpR17EnumeratedMs160
)
