package ies

import "rrc/utils"

// PDCP-Config-headerCompression-rohc ::= SEQUENCE
// Extensible
type PdcpConfigHeadercompressionRohc struct {
	Maxcid   utils.INTEGER `lb:0,ub:16383`
	Profiles PdcpConfigHeadercompressionRohcProfiles
}
