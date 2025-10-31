package ies

import "rrc/utils"

// PDCP-Config-drb-headerCompression-uplinkOnlyROHC ::= SEQUENCE
type PdcpConfigDrbHeadercompressionUplinkonlyrohc struct {
	Maxcid          utils.INTEGER `lb:0,ub:16383`
	Profiles        PdcpConfigDrbHeadercompressionUplinkonlyrohcProfiles
	DrbContinuerohc *utils.BOOLEAN
}
