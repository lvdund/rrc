package ies

import "rrc/utils"

// BandNR-trs-AdditionalBandwidth-r16 ::= ENUMERATED
type BandnrTrsAdditionalbandwidthR16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrTrsAdditionalbandwidthR16EnumeratedNothing = iota
	BandnrTrsAdditionalbandwidthR16EnumeratedTrs_Addbw_Set1
	BandnrTrsAdditionalbandwidthR16EnumeratedTrs_Addbw_Set2
)
