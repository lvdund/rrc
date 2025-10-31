package ies

import "rrc/utils"

// PhyLayerParameters-v1250-pusch-SRS-PowerControl-SubframeSet-r12 ::= ENUMERATED
type PhylayerparametersV1250PuschSrsPowercontrolSubframesetR12 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1250PuschSrsPowercontrolSubframesetR12EnumeratedNothing = iota
	PhylayerparametersV1250PuschSrsPowercontrolSubframesetR12EnumeratedSupported
)
