package ies

import "rrc/utils"

// SystemInformationBlockType1-v1310-IEs-bandwidthReducedAccessRelatedInfo-r13 ::= SEQUENCE
type Systeminformationblocktype1V1310IesBandwidthreducedaccessrelatedinfoR13 struct {
	SiWindowlengthBrR13                 Systeminformationblocktype1V1310IesBandwidthreducedaccessrelatedinfoR13SiWindowlengthBrR13
	SiRepetitionpatternR13              Systeminformationblocktype1V1310IesBandwidthreducedaccessrelatedinfoR13SiRepetitionpatternR13
	SchedulinginfolistBrR13             *SchedulinginfolistBrR13
	FddDownlinkortddsubframebitmapbrR13 *Systeminformationblocktype1V1310IesBandwidthreducedaccessrelatedinfoR13FddDownlinkortddsubframebitmapbrR13
	FddUplinksubframebitmapbrR13        *utils.BITSTRING `lb:10,ub:10`
	StartsymbolbrR13                    utils.INTEGER    `lb:0,ub:4`
	SiHoppingconfigcommonR13            Systeminformationblocktype1V1310IesBandwidthreducedaccessrelatedinfoR13SiHoppingconfigcommonR13
	SiValiditytimeR13                   *bool
	SysteminfovaluetaglistR13           *SysteminfovaluetaglistR13
}
