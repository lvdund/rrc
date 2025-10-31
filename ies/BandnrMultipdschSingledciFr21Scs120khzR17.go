package ies

import "rrc/utils"

// BandNR-multiPDSCH-SingleDCI-FR2-1-SCS-120kHz-r17 ::= ENUMERATED
type BandnrMultipdschSingledciFr21Scs120khzR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrMultipdschSingledciFr21Scs120khzR17EnumeratedNothing = iota
	BandnrMultipdschSingledciFr21Scs120khzR17EnumeratedSupported
)
