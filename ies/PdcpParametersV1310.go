package ies

import "rrc/utils"

// PDCP-Parameters-v1310 ::= SEQUENCE
type PdcpParametersV1310 struct {
	PdcpSnExtension18bitsR13 *utils.ENUMERATED
}
