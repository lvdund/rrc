package ies

import "rrc/utils"

// SupportedBandwidth-v1700 ::= CHOICE
const (
	SupportedbandwidthV1700ChoiceNothing = iota
	SupportedbandwidthV1700ChoiceFr1R17
	SupportedbandwidthV1700ChoiceFr2R17
)

type SupportedbandwidthV1700 struct {
	Choice uint64
	Fr1R17 *utils.ENUMERATED
	Fr2R17 *utils.ENUMERATED
}
