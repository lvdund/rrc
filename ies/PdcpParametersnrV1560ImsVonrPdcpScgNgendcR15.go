package ies

import "rrc/utils"

// PDCP-ParametersNR-v1560-ims-VoNR-PDCP-SCG-NGENDC-r15 ::= ENUMERATED
type PdcpParametersnrV1560ImsVonrPdcpScgNgendcR15 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersnrV1560ImsVonrPdcpScgNgendcR15EnumeratedNothing = iota
	PdcpParametersnrV1560ImsVonrPdcpScgNgendcR15EnumeratedSupported
)
