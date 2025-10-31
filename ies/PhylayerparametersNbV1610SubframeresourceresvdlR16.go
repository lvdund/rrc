package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1610-subframeResourceResvDL-r16 ::= ENUMERATED
type PhylayerparametersNbV1610SubframeresourceresvdlR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbV1610SubframeresourceresvdlR16EnumeratedNothing = iota
	PhylayerparametersNbV1610SubframeresourceresvdlR16EnumeratedSupported
)
