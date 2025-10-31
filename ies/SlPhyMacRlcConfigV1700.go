package ies

// SL-PHY-MAC-RLC-Config-v1700 ::= SEQUENCE
// Extensible
type SlPhyMacRlcConfigV1700 struct {
	SlDrxConfigR17               *SlDrxConfigR17
	SlRlcChanneltoreleaselistR17 *[]SlRlcChannelidR17     `lb:1,ub:maxSLLcidR16`
	SlRlcChanneltoaddmodlistR17  *[]SlRlcChannelconfigR17 `lb:1,ub:maxSLLcidR16`
}
