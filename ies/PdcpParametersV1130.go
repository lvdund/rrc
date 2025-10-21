package ies

import "rrc/utils"

// PDCP-Parameters-v1130 ::= SEQUENCE
type PdcpParametersV1130 struct {
	PdcpSnExtensionR11            *utils.ENUMERATED
	SupportrohccontextcontinueR11 *utils.ENUMERATED
}
