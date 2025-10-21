package ies

import "rrc/utils"

// CSI-RS-ConfigNonPrecoded-v1430 ::= SEQUENCE
type CsiRsConfignonprecodedV1430 struct {
	CsiRsConfignzpEmimoV1430          *CsiRsConfignzpEmimoV1430
	Codebookconfign1V1430             utils.ENUMERATED
	Codebookconfign2V1430             utils.ENUMERATED
	NzpResourceconfigtm9OriginalV1430 CsiRsConfigNzpV1430
}
