package ies

// MRB-PDCP-ConfigBroadcast-r17-headerCompression-r17 ::= CHOICE
const (
	MrbPdcpConfigbroadcastR17HeadercompressionR17ChoiceNothing = iota
	MrbPdcpConfigbroadcastR17HeadercompressionR17ChoiceNotused
	MrbPdcpConfigbroadcastR17HeadercompressionR17ChoiceRohc
)

type MrbPdcpConfigbroadcastR17HeadercompressionR17 struct {
	Choice  uint64
	Notused *struct{}
	Rohc    *MrbPdcpConfigbroadcastR17HeadercompressionR17Rohc
}
