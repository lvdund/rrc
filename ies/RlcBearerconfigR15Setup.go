package ies

// RLC-BearerConfig-r15-setup ::= SEQUENCE
type RlcBearerconfigR15Setup struct {
	RlcConfigR15                    *RlcConfigR15
	LogicalchannelidentityconfigR15 RlcBearerconfigR15SetupLogicalchannelidentityconfigR15
	LogicalchannelconfigR15         *Logicalchannelconfig
}
