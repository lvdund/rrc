package ies

import "rrc/utils"

// Other-Parameters-v1360-inDeviceCoexInd-HardwareSharingInd-r13 ::= ENUMERATED
type OtherParametersV1360IndevicecoexindHardwaresharingindR13 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersV1360IndevicecoexindHardwaresharingindR13EnumeratedNothing = iota
	OtherParametersV1360IndevicecoexindHardwaresharingindR13EnumeratedSupported
)
