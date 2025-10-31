package ies

import "rrc/utils"

// Other-Parameters-v1540-inDeviceCoexInd-ENDC-r15 ::= ENUMERATED
type OtherParametersV1540IndevicecoexindEndcR15 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersV1540IndevicecoexindEndcR15EnumeratedNothing = iota
	OtherParametersV1540IndevicecoexindEndcR15EnumeratedSupported
)
