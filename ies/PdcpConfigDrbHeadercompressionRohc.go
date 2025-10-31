package ies

import "rrc/utils"

// PDCP-Config-drb-headerCompression-rohc ::= SEQUENCE
type PdcpConfigDrbHeadercompressionRohc struct {
	Maxcid          utils.INTEGER `lb:0,ub:16383`
	Profiles        PdcpConfigDrbHeadercompressionRohcProfiles
	DrbContinuerohc *utils.BOOLEAN
}
