package ies

import "rrc/utils"

// PUR-Parameters-r16-pur-PUSCH-NB-MaxTBS-r16 ::= ENUMERATED
type PurParametersR16PurPuschNbMaxtbsR16 struct {
	Value utils.ENUMERATED
}

const (
	PurParametersR16PurPuschNbMaxtbsR16EnumeratedNothing = iota
	PurParametersR16PurPuschNbMaxtbsR16EnumeratedSupported
)
