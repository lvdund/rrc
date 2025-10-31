package ies

import "rrc/utils"

// AntennaInfoDedicatedSTTI-r15-setup-maxLayersMIMO-STTI-r15 ::= ENUMERATED
type AntennainfodedicatedsttiR15SetupMaxlayersmimoSttiR15 struct {
	Value utils.ENUMERATED
}

const (
	AntennainfodedicatedsttiR15SetupMaxlayersmimoSttiR15EnumeratedNothing = iota
	AntennainfodedicatedsttiR15SetupMaxlayersmimoSttiR15EnumeratedTwolayers
	AntennainfodedicatedsttiR15SetupMaxlayersmimoSttiR15EnumeratedFourlayers
)
