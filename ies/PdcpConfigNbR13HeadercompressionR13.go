package ies

// PDCP-Config-NB-r13-headerCompression-r13 ::= CHOICE
// Extensible
const (
	PdcpConfigNbR13HeadercompressionR13ChoiceNothing = iota
	PdcpConfigNbR13HeadercompressionR13ChoiceNotused
	PdcpConfigNbR13HeadercompressionR13ChoiceRohc
)

type PdcpConfigNbR13HeadercompressionR13 struct {
	Choice  uint64
	Notused *struct{}
	Rohc    *PdcpConfigNbR13HeadercompressionR13Rohc
}
