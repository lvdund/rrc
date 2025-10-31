package ies

import "rrc/utils"

// FR2-2-AccessParamsPerBand-r17-support32-DL-HARQ-ProcessPerSCS-r17-scs-480kHz-r17 ::= ENUMERATED
type Fr22AccessparamsperbandR17Support32DlHarqProcessperscsR17Scs480khzR17 struct {
	Value utils.ENUMERATED
}

const (
	Fr22AccessparamsperbandR17Support32DlHarqProcessperscsR17Scs480khzR17EnumeratedNothing = iota
	Fr22AccessparamsperbandR17Support32DlHarqProcessperscsR17Scs480khzR17EnumeratedSupported
)
