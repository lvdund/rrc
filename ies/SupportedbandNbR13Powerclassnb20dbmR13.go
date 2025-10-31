package ies

import "rrc/utils"

// SupportedBand-NB-r13-powerClassNB-20dBm-r13 ::= ENUMERATED
type SupportedbandNbR13Powerclassnb20dbmR13 struct {
	Value utils.ENUMERATED
}

const (
	SupportedbandNbR13Powerclassnb20dbmR13EnumeratedNothing = iota
	SupportedbandNbR13Powerclassnb20dbmR13EnumeratedSupported
)
