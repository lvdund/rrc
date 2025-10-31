package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-CSI-numCPU-r17 ::= ENUMERATED
type MimoParametersperbandMtrpCsiNumcpuR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpCsiNumcpuR17EnumeratedNothing = iota
	MimoParametersperbandMtrpCsiNumcpuR17EnumeratedN2
	MimoParametersperbandMtrpCsiNumcpuR17EnumeratedN3
	MimoParametersperbandMtrpCsiNumcpuR17EnumeratedN4
)
