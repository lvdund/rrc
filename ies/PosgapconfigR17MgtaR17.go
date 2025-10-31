package ies

import "rrc/utils"

// PosGapConfig-r17-mgta-r17 ::= ENUMERATED
type PosgapconfigR17MgtaR17 struct {
	Value utils.ENUMERATED
}

const (
	PosgapconfigR17MgtaR17EnumeratedNothing = iota
	PosgapconfigR17MgtaR17EnumeratedMs0
	PosgapconfigR17MgtaR17EnumeratedMs0dot25
	PosgapconfigR17MgtaR17EnumeratedMs0dot5
)
