package ies

import "rrc/utils"

// CSI-RS-ConfigNZP-EMIMO-v1430 ::= SEQUENCE
type CsiRsConfignzpEmimoV1430 struct {
	NzpResourceconfiglistextR14 interface{}
	CdmtypeV1430                *utils.ENUMERATED
}
