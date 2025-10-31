package ies

import "rrc/utils"

// RLC-BearerConfig ::= SEQUENCE
// Extensible
type RlcBearerconfig struct {
	Logicalchannelidentity       Logicalchannelidentity
	Servedradiobearer            *RlcBearerconfigServedradiobearer
	Reestablishrlc               *utils.BOOLEAN
	RlcConfig                    *RlcConfig
	MacLogicalchannelconfig      *Logicalchannelconfig
	RlcConfigV1610               *RlcConfigV1610               `ext`
	RlcConfigV1700               *RlcConfigV1700               `ext`
	LogicalchannelidentityextR17 *LogicalchannelidentityextR17 `ext`
	MulticastrlcBearerconfigR17  *MulticastrlcBearerconfigR17  `ext`
	Servedradiobearersrb4R17     *SrbIdentityV1700             `ext`
}
