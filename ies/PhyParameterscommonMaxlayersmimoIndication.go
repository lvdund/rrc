package ies

import "rrc/utils"

// Phy-ParametersCommon-maxLayersMIMO-Indication ::= ENUMERATED
type PhyParameterscommonMaxlayersmimoIndication struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonMaxlayersmimoIndicationEnumeratedNothing = iota
	PhyParameterscommonMaxlayersmimoIndicationEnumeratedSupported
)
