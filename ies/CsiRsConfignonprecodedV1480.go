package ies

import "rrc/utils"

// CSI-RS-ConfigNonPrecoded-v1480 ::= SEQUENCE
type CsiRsConfignonprecodedV1480 struct {
	CsiRsConfignzpEmimoV1480          *CsiRsConfignzpEmimoV1430
	Codebookconfign1V1480             *utils.ENUMERATED
	Codebookconfign2R1480             *utils.ENUMERATED
	NzpResourceconfigtm9OriginalV1480 CsiRsConfigNzpV1430
}
