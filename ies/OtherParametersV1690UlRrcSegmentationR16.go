package ies

import "rrc/utils"

// Other-Parameters-v1690-ul-RRC-Segmentation-r16 ::= ENUMERATED
type OtherParametersV1690UlRrcSegmentationR16 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersV1690UlRrcSegmentationR16EnumeratedNothing = iota
	OtherParametersV1690UlRrcSegmentationR16EnumeratedSupported
)
