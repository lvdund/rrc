package ies

import "rrc/utils"

// AntennaInfoDedicated-v10i0-maxLayersMIMO-r10 ::= ENUMERATED
type AntennainfodedicatedV10i0MaxlayersmimoR10 struct {
	Value utils.ENUMERATED
}

const (
	AntennainfodedicatedV10i0MaxlayersmimoR10EnumeratedNothing = iota
	AntennainfodedicatedV10i0MaxlayersmimoR10EnumeratedTwolayers
	AntennainfodedicatedV10i0MaxlayersmimoR10EnumeratedFourlayers
	AntennainfodedicatedV10i0MaxlayersmimoR10EnumeratedEightlayers
)
