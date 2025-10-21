package ies

import "rrc/utils"

// WLAN-ChannelList-r13 ::= SEQUENCE OF WLAN-Channel-r13
// SIZE (1..maxWLAN-Channels-r13)
type WlanChannellistR13 struct {
	Value []WlanChannelR13
}
