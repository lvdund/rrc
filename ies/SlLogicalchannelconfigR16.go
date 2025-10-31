package ies

import "rrc/utils"

// SL-LogicalChannelConfig-r16 ::= SEQUENCE
// Extensible
type SlLogicalchannelconfigR16 struct {
	SlPriorityR16                          utils.INTEGER `lb:0,ub:8`
	SlPrioritisedbitrateR16                SlLogicalchannelconfigR16SlPrioritisedbitrateR16
	SlBucketsizedurationR16                SlLogicalchannelconfigR16SlBucketsizedurationR16
	SlConfiguredgranttype1allowedR16       *utils.BOOLEAN
	SlHarqFeedbackenabledR16               *SlLogicalchannelconfigR16SlHarqFeedbackenabledR16
	SlAllowedcgListR16                     *[]SlConfigindexcgR16 `lb:0,ub:maxNrofCGSl1R16`
	SlAllowedscsListR16                    *[]Subcarrierspacing  `lb:1,ub:maxSCSs`
	SlMaxpuschDurationR16                  *SlLogicalchannelconfigR16SlMaxpuschDurationR16
	SlLogicalchannelgroupR16               *utils.INTEGER `lb:0,ub:maxLCGId`
	SlSchedulingrequestidR16               *Schedulingrequestid
	SlLogicalchannelsrDelaytimerappliedR16 *utils.BOOLEAN
}
