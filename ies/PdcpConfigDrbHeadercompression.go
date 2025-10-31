package ies

// PDCP-Config-drb-headerCompression ::= CHOICE
// Extensible
const (
	PdcpConfigDrbHeadercompressionChoiceNothing = iota
	PdcpConfigDrbHeadercompressionChoiceNotused
	PdcpConfigDrbHeadercompressionChoiceRohc
	PdcpConfigDrbHeadercompressionChoiceUplinkonlyrohc
)

type PdcpConfigDrbHeadercompression struct {
	Choice         uint64
	Notused        *struct{}
	Rohc           *PdcpConfigDrbHeadercompressionRohc
	Uplinkonlyrohc *PdcpConfigDrbHeadercompressionUplinkonlyrohc
}
