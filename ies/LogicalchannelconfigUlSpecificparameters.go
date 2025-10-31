package ies

import "rrc/utils"

// LogicalChannelConfig-ul-SpecificParameters ::= SEQUENCE
// Extensible
type LogicalchannelconfigUlSpecificparameters struct {
	Priority                          utils.INTEGER `lb:0,ub:16`
	Prioritisedbitrate                LogicalchannelconfigUlSpecificparametersPrioritisedbitrate
	Bucketsizeduration                LogicalchannelconfigUlSpecificparametersBucketsizeduration
	Allowedservingcells               *[]Servcellindex     `lb:1,ub:maxNrofServingCells1`
	AllowedscsList                    *[]Subcarrierspacing `lb:1,ub:maxSCSs`
	MaxpuschDuration                  *LogicalchannelconfigUlSpecificparametersMaxpuschDuration
	Configuredgranttype1allowed       *utils.BOOLEAN
	Logicalchannelgroup               *utils.INTEGER `lb:0,ub:maxLCGId`
	Schedulingrequestid               *Schedulingrequestid
	LogicalchannelsrMask              utils.BOOLEAN
	LogicalchannelsrDelaytimerapplied utils.BOOLEAN
	Bitratequeryprohibittimer         *LogicalchannelconfigUlSpecificparametersBitratequeryprohibittimer
	AllowedcgListR16                  *[]ConfiguredgrantconfigindexmacR16 `lb:0,ub:maxNrofConfiguredGrantConfigMAC1R16`
	AllowedphyPriorityindexR16        *LogicalchannelconfigUlSpecificparametersAllowedphyPriorityindexR16
	LogicalchannelgroupiabExtR17      *utils.INTEGER `lb:0,ub:maxLCGIdIabR17`
	AllowedharqModeR17                *LogicalchannelconfigUlSpecificparametersAllowedharqModeR17
}
