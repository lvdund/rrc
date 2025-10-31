package ies

import "rrc/utils"

// Other-Parameters-r11-inDeviceCoexInd-r11 ::= ENUMERATED
type OtherParametersR11IndevicecoexindR11 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersR11IndevicecoexindR11EnumeratedNothing = iota
	OtherParametersR11IndevicecoexindR11EnumeratedSupported
)
