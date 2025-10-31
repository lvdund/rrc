package ies

import "rrc/utils"

// SupportedBandwidth ::= CHOICE
const (
	SupportedbandwidthChoiceNothing = iota
	SupportedbandwidthChoiceFr1
	SupportedbandwidthChoiceFr2
)

type Supportedbandwidth struct {
	Choice uint64
	Fr1    *utils.ENUMERATED
	Fr2    *utils.ENUMERATED
}
