package ies

// SL-RLC-ChannelConfigPC5-r17 ::= SEQUENCE
// Extensible
type SlRlcChannelconfigpc5R17 struct {
	SlRlcChannelidPc5R17            SlRlcChannelidR17
	SlRlcConfigpc5R17               *SlRlcConfigpc5R16
	SlMacLogicalchannelconfigpc5R17 *SlLogicalchannelconfigpc5R16
}
