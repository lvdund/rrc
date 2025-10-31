package ies

import "rrc/utils"

// FR2-2-AccessParamsPerBand-r17-support32-UL-HARQ-ProcessPerSCS-r17-scs-960kHz-r17 ::= ENUMERATED
type Fr22AccessparamsperbandR17Support32UlHarqProcessperscsR17Scs960khzR17 struct {
	Value utils.ENUMERATED
}

const (
	Fr22AccessparamsperbandR17Support32UlHarqProcessperscsR17Scs960khzR17EnumeratedNothing = iota
	Fr22AccessparamsperbandR17Support32UlHarqProcessperscsR17Scs960khzR17EnumeratedSupported
)
