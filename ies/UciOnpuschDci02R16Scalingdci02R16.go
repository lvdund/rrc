package ies

import "rrc/utils"

// UCI-OnPUSCH-DCI-0-2-r16-scalingDCI-0-2-r16 ::= ENUMERATED
type UciOnpuschDci02R16Scalingdci02R16 struct {
	Value utils.ENUMERATED
}

const (
	UciOnpuschDci02R16Scalingdci02R16EnumeratedNothing = iota
	UciOnpuschDci02R16Scalingdci02R16EnumeratedF0p5
	UciOnpuschDci02R16Scalingdci02R16EnumeratedF0p65
	UciOnpuschDci02R16Scalingdci02R16EnumeratedF0p8
	UciOnpuschDci02R16Scalingdci02R16EnumeratedF1
)
