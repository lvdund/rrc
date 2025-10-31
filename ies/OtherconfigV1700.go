package ies

import "rrc/utils"

// OtherConfig-v1700 ::= SEQUENCE
type OtherconfigV1700 struct {
	UlGapfr2PreferenceconfigR17               *utils.BOOLEAN
	MusimGapassistanceconfigR17               *utils.Setuprelease[MusimGapassistanceconfigR17]
	MusimLeaveassistanceconfigR17             *utils.Setuprelease[MusimLeaveassistanceconfigR17]
	SuccesshoConfigR17                        *utils.Setuprelease[SuccesshoConfigR17]
	MaxbwPreferenceconfigfr22R17              *utils.BOOLEAN
	MaxmimoLayerpreferenceconfigfr22R17       *utils.BOOLEAN
	MinschedulingoffsetpreferenceconfigextR17 *utils.BOOLEAN
	RlmRelaxationreportingconfigR17           *utils.Setuprelease[RlmRelaxationreportingconfigR17]
	BfdRelaxationreportingconfigR17           *utils.Setuprelease[BfdRelaxationreportingconfigR17]
	ScgDeactivationpreferenceconfigR17        *utils.Setuprelease[ScgDeactivationpreferenceconfigR17]
	RrmMeasrelaxationreportingconfigR17       *utils.Setuprelease[RrmMeasrelaxationreportingconfigR17]
	PropdelaydiffreportconfigR17              *utils.Setuprelease[PropdelaydiffreportconfigR17]
}
