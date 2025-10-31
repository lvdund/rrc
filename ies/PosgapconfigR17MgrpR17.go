package ies

import "rrc/utils"

// PosGapConfig-r17-mgrp-r17 ::= ENUMERATED
type PosgapconfigR17MgrpR17 struct {
	Value utils.ENUMERATED
}

const (
	PosgapconfigR17MgrpR17EnumeratedNothing = iota
	PosgapconfigR17MgrpR17EnumeratedMs20
	PosgapconfigR17MgrpR17EnumeratedMs40
	PosgapconfigR17MgrpR17EnumeratedMs80
	PosgapconfigR17MgrpR17EnumeratedMs160
)
