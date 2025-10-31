package ies

// CSI-RS-ConfigNonPrecoded-v1430 ::= SEQUENCE
type CsiRsConfignonprecodedV1430 struct {
	CsiRsConfignzpEmimoV1430          *CsiRsConfignzpEmimoV1430
	Codebookconfign1V1430             CsiRsConfignonprecodedV1430Codebookconfign1V1430
	Codebookconfign2V1430             CsiRsConfignonprecodedV1430Codebookconfign2V1430
	NzpResourceconfigtm9OriginalV1430 CsiRsConfigNzpV1430
}
