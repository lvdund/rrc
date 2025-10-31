package ies

import "rrc/utils"

// MIMO-ParametersPerBand-multiDCI-multiTRP-Parameters-r16-defaultQCL-PerCORESETPoolIndex-r16 ::= ENUMERATED
type MimoParametersperbandMultidciMultitrpParametersR16DefaultqclPercoresetpoolindexR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMultidciMultitrpParametersR16DefaultqclPercoresetpoolindexR16EnumeratedNothing = iota
	MimoParametersperbandMultidciMultitrpParametersR16DefaultqclPercoresetpoolindexR16EnumeratedSupported
)
