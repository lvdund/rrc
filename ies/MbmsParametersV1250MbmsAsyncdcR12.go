package ies

import "rrc/utils"

// MBMS-Parameters-v1250-mbms-AsyncDC-r12 ::= ENUMERATED
type MbmsParametersV1250MbmsAsyncdcR12 struct {
	Value utils.ENUMERATED
}

const (
	MbmsParametersV1250MbmsAsyncdcR12EnumeratedNothing = iota
	MbmsParametersV1250MbmsAsyncdcR12EnumeratedSupported
)
