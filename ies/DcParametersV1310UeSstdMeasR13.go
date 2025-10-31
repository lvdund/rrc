package ies

import "rrc/utils"

// DC-Parameters-v1310-ue-SSTD-Meas-r13 ::= ENUMERATED
type DcParametersV1310UeSstdMeasR13 struct {
	Value utils.ENUMERATED
}

const (
	DcParametersV1310UeSstdMeasR13EnumeratedNothing = iota
	DcParametersV1310UeSstdMeasR13EnumeratedSupported
)
