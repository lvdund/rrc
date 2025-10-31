package ies

// PDCP-Config-moreThanOneRLC-primaryPath ::= SEQUENCE
type PdcpConfigMorethanonerlcPrimarypath struct {
	Cellgroup      *Cellgroupid
	Logicalchannel *Logicalchannelidentity
}
