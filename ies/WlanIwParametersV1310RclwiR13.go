package ies

import "rrc/utils"

// WLAN-IW-Parameters-v1310-rclwi-r13 ::= ENUMERATED
type WlanIwParametersV1310RclwiR13 struct {
	Value utils.ENUMERATED
}

const (
	WlanIwParametersV1310RclwiR13EnumeratedNothing = iota
	WlanIwParametersV1310RclwiR13EnumeratedSupported
)
