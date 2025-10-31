package ies

import "rrc/utils"

// NTN-Config-r17-ntn-PolarizationDL-r17 ::= ENUMERATED
type NtnConfigR17NtnPolarizationdlR17 struct {
	Value utils.ENUMERATED
}

const (
	NtnConfigR17NtnPolarizationdlR17EnumeratedNothing = iota
	NtnConfigR17NtnPolarizationdlR17EnumeratedRhcp
	NtnConfigR17NtnPolarizationdlR17EnumeratedLhcp
	NtnConfigR17NtnPolarizationdlR17EnumeratedLinear
)
