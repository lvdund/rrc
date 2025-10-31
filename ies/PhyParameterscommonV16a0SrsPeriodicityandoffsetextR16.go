package ies

import "rrc/utils"

// Phy-ParametersCommon-v16a0-srs-PeriodicityAndOffsetExt-r16 ::= ENUMERATED
type PhyParameterscommonV16a0SrsPeriodicityandoffsetextR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonV16a0SrsPeriodicityandoffsetextR16EnumeratedNothing = iota
	PhyParameterscommonV16a0SrsPeriodicityandoffsetextR16EnumeratedSupported
)
