package ies

// PDCP-Config-NB-r13 ::= SEQUENCE
// Extensible
type PdcpConfigNbR13 struct {
	DiscardtimerR13      *PdcpConfigNbR13DiscardtimerR13
	HeadercompressionR13 PdcpConfigNbR13HeadercompressionR13
}
