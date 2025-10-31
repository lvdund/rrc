package ies

import "rrc/utils"

// BandNR-multiPUSCH-SingleDCI-FR2-1-SCS-120kHz-r17 ::= ENUMERATED
type BandnrMultipuschSingledciFr21Scs120khzR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrMultipuschSingledciFr21Scs120khzR17EnumeratedNothing = iota
	BandnrMultipuschSingledciFr21Scs120khzR17EnumeratedSupported
)
