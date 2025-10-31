package ies

import "rrc/utils"

// BandNR-mpr-PowerBoost-FR2-r16 ::= ENUMERATED
type BandnrMprPowerboostFr2R16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrMprPowerboostFr2R16EnumeratedNothing = iota
	BandnrMprPowerboostFr2R16EnumeratedSupported
)
