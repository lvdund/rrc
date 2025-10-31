package ies

import "rrc/utils"

// GapConfig-r17-mgta-r17 ::= ENUMERATED
type GapconfigR17MgtaR17 struct {
	Value utils.ENUMERATED
}

const (
	GapconfigR17MgtaR17EnumeratedNothing = iota
	GapconfigR17MgtaR17EnumeratedMs0
	GapconfigR17MgtaR17EnumeratedMs0dot25
	GapconfigR17MgtaR17EnumeratedMs0dot5
	GapconfigR17MgtaR17EnumeratedMs0dot75
)
