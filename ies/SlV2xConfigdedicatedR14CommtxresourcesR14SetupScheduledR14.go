package ies

import "rrc/utils"

// SL-V2X-ConfigDedicated-r14-commTxResources-r14-setup-scheduled-r14 ::= SEQUENCE
type SlV2xConfigdedicatedR14CommtxresourcesR14SetupScheduledR14 struct {
	SlVRntiR14                CRnti
	MacMainconfigR14          MacMainconfigslR12
	V2xSchedulingpoolR14      *SlCommresourcepoolv2xR14
	McsR14                    *utils.INTEGER `lb:0,ub:31`
	LogicalchgroupinfolistR14 LogicalchgroupinfolistR13
}
