package ies

// PDCP-Config ::= SEQUENCE
// Extensible
type PdcpConfig struct {
	Discardtimer      *PdcpConfigDiscardtimer
	RlcAm             *PdcpConfigRlcAm
	RlcUm             *PdcpConfigRlcUm
	Headercompression PdcpConfigHeadercompression
}
