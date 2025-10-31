package ies

import "rrc/utils"

// PCCH-Config-nAndPagingFrameOffset ::= CHOICE
const (
	PcchConfigNandpagingframeoffsetChoiceNothing = iota
	PcchConfigNandpagingframeoffsetChoiceOnet
	PcchConfigNandpagingframeoffsetChoiceHalft
	PcchConfigNandpagingframeoffsetChoiceQuartert
	PcchConfigNandpagingframeoffsetChoiceOneeightht
	PcchConfigNandpagingframeoffsetChoiceOnesixteentht
)

type PcchConfigNandpagingframeoffset struct {
	Choice        uint64
	Onet          *struct{}
	Halft         *utils.INTEGER `lb:0,ub:1`
	Quartert      *utils.INTEGER `lb:0,ub:3`
	Oneeightht    *utils.INTEGER `lb:0,ub:7`
	Onesixteentht *utils.INTEGER `lb:0,ub:15`
}
