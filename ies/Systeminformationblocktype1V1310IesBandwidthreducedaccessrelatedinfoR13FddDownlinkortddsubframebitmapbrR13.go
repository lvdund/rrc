package ies

import "rrc/utils"

// SystemInformationBlockType1-v1310-IEs-bandwidthReducedAccessRelatedInfo-r13-fdd-DownlinkOrTddSubframeBitmapBR-r13 ::= CHOICE
const (
	Systeminformationblocktype1V1310IesBandwidthreducedaccessrelatedinfoR13FddDownlinkortddsubframebitmapbrR13ChoiceNothing = iota
	Systeminformationblocktype1V1310IesBandwidthreducedaccessrelatedinfoR13FddDownlinkortddsubframebitmapbrR13ChoiceSubframepattern10R13
	Systeminformationblocktype1V1310IesBandwidthreducedaccessrelatedinfoR13FddDownlinkortddsubframebitmapbrR13ChoiceSubframepattern40R13
)

type Systeminformationblocktype1V1310IesBandwidthreducedaccessrelatedinfoR13FddDownlinkortddsubframebitmapbrR13 struct {
	Choice               uint64
	Subframepattern10R13 *utils.BITSTRING `lb:10,ub:10`
	Subframepattern40R13 *utils.BITSTRING `lb:40,ub:40`
}
