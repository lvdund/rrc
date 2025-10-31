package ies

import "rrc/utils"

// PhyLayerParameters-v1020-simultaneousPUCCH-PUSCH-r10 ::= ENUMERATED
type PhylayerparametersV1020SimultaneouspucchPuschR10 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1020SimultaneouspucchPuschR10EnumeratedNothing = iota
	PhylayerparametersV1020SimultaneouspucchPuschR10EnumeratedSupported
)
