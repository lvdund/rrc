package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1610-subframeResourceResvUL-r16 ::= ENUMERATED
type PhylayerparametersNbV1610SubframeresourceresvulR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbV1610SubframeresourceresvulR16EnumeratedNothing = iota
	PhylayerparametersNbV1610SubframeresourceresvulR16EnumeratedSupported
)
