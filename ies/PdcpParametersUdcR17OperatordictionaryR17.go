package ies

import "rrc/utils"

// PDCP-Parameters-udc-r17-operatorDictionary-r17 ::= SEQUENCE
type PdcpParametersUdcR17OperatordictionaryR17 struct {
	VersionofdictionaryR17 utils.INTEGER `lb:0,ub:15`
	AssociatedplmnIdR17    PlmnIdentity
}
