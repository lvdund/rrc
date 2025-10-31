package ies

// SLRB-Config-r16 ::= SEQUENCE
// Extensible
type SlrbConfigR16 struct {
	SlrbPc5ConfigindexR16           SlrbPc5ConfigindexR16
	SlSdapConfigpc5R16              *SlSdapConfigpc5R16
	SlPdcpConfigpc5R16              *SlPdcpConfigpc5R16
	SlRlcConfigpc5R16               *SlRlcConfigpc5R16
	SlMacLogicalchannelconfigpc5R16 *SlLogicalchannelconfigpc5R16
}
