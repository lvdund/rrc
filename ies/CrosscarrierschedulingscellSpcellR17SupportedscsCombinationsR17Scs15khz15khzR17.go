package ies

import "rrc/utils"

// CrossCarrierSchedulingSCell-SpCell-r17-supportedSCS-Combinations-r17-scs15kHz-15kHz-r17 ::= ENUMERATED
type CrosscarrierschedulingscellSpcellR17SupportedscsCombinationsR17Scs15khz15khzR17 struct {
	Value utils.ENUMERATED
}

const (
	CrosscarrierschedulingscellSpcellR17SupportedscsCombinationsR17Scs15khz15khzR17EnumeratedNothing = iota
	CrosscarrierschedulingscellSpcellR17SupportedscsCombinationsR17Scs15khz15khzR17EnumeratedSupported
)
