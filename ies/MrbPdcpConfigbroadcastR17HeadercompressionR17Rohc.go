package ies

import "rrc/utils"

// MRB-PDCP-ConfigBroadcast-r17-headerCompression-r17-rohc ::= SEQUENCE
type MrbPdcpConfigbroadcastR17HeadercompressionR17Rohc struct {
	MaxcidR17   utils.INTEGER `lb:0,ub:16`
	ProfilesR17 MrbPdcpConfigbroadcastR17HeadercompressionR17RohcProfilesR17
}
