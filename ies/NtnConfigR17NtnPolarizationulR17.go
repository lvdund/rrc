package ies

import "rrc/utils"

// NTN-Config-r17-ntn-PolarizationUL-r17 ::= ENUMERATED
type NtnConfigR17NtnPolarizationulR17 struct {
	Value utils.ENUMERATED
}

const (
	NtnConfigR17NtnPolarizationulR17EnumeratedNothing = iota
	NtnConfigR17NtnPolarizationulR17EnumeratedRhcp
	NtnConfigR17NtnPolarizationulR17EnumeratedLhcp
	NtnConfigR17NtnPolarizationulR17EnumeratedLinear
)
