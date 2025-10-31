package ies

import "rrc/utils"

// UCI-OnPUSCH-scaling ::= ENUMERATED
type UciOnpuschScaling struct {
	Value utils.ENUMERATED
}

const (
	UciOnpuschScalingEnumeratedNothing = iota
	UciOnpuschScalingEnumeratedF0p5
	UciOnpuschScalingEnumeratedF0p65
	UciOnpuschScalingEnumeratedF0p8
	UciOnpuschScalingEnumeratedF1
)
