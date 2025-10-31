package ies

// PDCP-Config-headerCompression ::= CHOICE
// Extensible
const (
	PdcpConfigHeadercompressionChoiceNothing = iota
	PdcpConfigHeadercompressionChoiceNotused
	PdcpConfigHeadercompressionChoiceRohc
)

type PdcpConfigHeadercompression struct {
	Choice  uint64
	Notused *struct{}
	Rohc    *PdcpConfigHeadercompressionRohc
}
