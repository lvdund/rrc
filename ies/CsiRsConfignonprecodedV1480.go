package ies

// CSI-RS-ConfigNonPrecoded-v1480 ::= SEQUENCE
type CsiRsConfignonprecodedV1480 struct {
	CsiRsConfignzpEmimoV1480          *CsiRsConfignzpEmimoV1430
	Codebookconfign1V1480             *CsiRsConfignonprecodedV1480Codebookconfign1V1480
	Codebookconfign2R1480             *CsiRsConfignonprecodedV1480Codebookconfign2R1480
	NzpResourceconfigtm9OriginalV1480 CsiRsConfigNzpV1430
}
