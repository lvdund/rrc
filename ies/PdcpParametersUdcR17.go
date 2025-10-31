package ies

// PDCP-Parameters-udc-r17 ::= SEQUENCE
type PdcpParametersUdcR17 struct {
	StandarddictionaryR17  *PdcpParametersUdcR17StandarddictionaryR17
	OperatordictionaryR17  *PdcpParametersUdcR17OperatordictionaryR17
	ContinueudcR17         *PdcpParametersUdcR17ContinueudcR17
	SupportofbuffersizeR17 *PdcpParametersUdcR17SupportofbuffersizeR17
}
