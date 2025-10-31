package ies

import "rrc/utils"

// CSI-RS-ConfigNZP-EMIMO-v1430-cdmType-v1430 ::= ENUMERATED
type CsiRsConfignzpEmimoV1430CdmtypeV1430 struct {
	Value utils.ENUMERATED
}

const (
	CsiRsConfignzpEmimoV1430CdmtypeV1430EnumeratedNothing = iota
	CsiRsConfignzpEmimoV1430CdmtypeV1430EnumeratedCdm8
)
