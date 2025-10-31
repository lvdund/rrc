package ies

import "rrc/utils"

// SoundingRS-UL-ConfigCommon-setup ::= SEQUENCE
type SoundingrsUlConfigcommonSetup struct {
	SrsBandwidthconfig                 SoundingrsUlConfigcommonSetupSrsBandwidthconfig
	SrsSubframeconfig                  SoundingrsUlConfigcommonSetupSrsSubframeconfig
	AcknacksrsSimultaneoustransmission utils.BOOLEAN
	SrsMaxuppts                        *bool
}
