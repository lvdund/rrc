package ies

import "rrc/utils"

// RN-SubframeConfig-r10-rpdcch-Config-r10-pucch-Config-r10-tdd-channelSelectionMultiplexingBundling ::= SEQUENCE
type RnSubframeconfigR10RpdcchConfigR10PucchConfigR10TddChannelselectionmultiplexingbundling struct {
	N1pucchAnListR10 []utils.INTEGER `lb:1,ub:4`
}
