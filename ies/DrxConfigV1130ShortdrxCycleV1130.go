package ies

import "rrc/utils"

// DRX-Config-v1130-shortDRX-Cycle-v1130 ::= ENUMERATED
type DrxConfigV1130ShortdrxCycleV1130 struct {
	Value utils.ENUMERATED
}

const (
	DrxConfigV1130ShortdrxCycleV1130EnumeratedNothing = iota
	DrxConfigV1130ShortdrxCycleV1130EnumeratedSf4_V1130
)
