package ies

import "rrc/utils"

// Other-Parameters-v11d0-inDeviceCoexInd-UL-CA-r11 ::= ENUMERATED
type OtherParametersV11d0IndevicecoexindUlCaR11 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersV11d0IndevicecoexindUlCaR11EnumeratedNothing = iota
	OtherParametersV11d0IndevicecoexindUlCaR11EnumeratedSupported
)
