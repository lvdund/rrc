package ies

// MRB-RLC-ConfigBroadcast-r17 ::= SEQUENCE
type MrbRlcConfigbroadcastR17 struct {
	LogicalchannelidentityR17 Logicalchannelidentity
	SnFieldlengthR17          *MrbRlcConfigbroadcastR17SnFieldlengthR17
	TReassemblyR17            *TReassembly
}
