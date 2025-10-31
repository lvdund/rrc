package ies

import "rrc/utils"

// RF-Parameters-v12b0-maxLayersMIMO-Indication-r12 ::= ENUMERATED
type RfParametersV12b0MaxlayersmimoIndicationR12 struct {
	Value utils.ENUMERATED
}

const (
	RfParametersV12b0MaxlayersmimoIndicationR12EnumeratedNothing = iota
	RfParametersV12b0MaxlayersmimoIndicationR12EnumeratedSupported
)
