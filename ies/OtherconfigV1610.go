package ies

import "rrc/utils"

// OtherConfig-v1610 ::= SEQUENCE
type OtherconfigV1610 struct {
	IdcAssistanceconfigR16                 *utils.Setuprelease[IdcAssistanceconfigR16]
	DrxPreferenceconfigR16                 *utils.Setuprelease[DrxPreferenceconfigR16]
	MaxbwPreferenceconfigR16               *utils.Setuprelease[MaxbwPreferenceconfigR16]
	MaxccPreferenceconfigR16               *utils.Setuprelease[MaxccPreferenceconfigR16]
	MaxmimoLayerpreferenceconfigR16        *utils.Setuprelease[MaxmimoLayerpreferenceconfigR16]
	MinschedulingoffsetpreferenceconfigR16 *utils.Setuprelease[MinschedulingoffsetpreferenceconfigR16]
	ReleasepreferenceconfigR16             *utils.Setuprelease[ReleasepreferenceconfigR16]
	ReferencetimepreferencereportingR16    *utils.BOOLEAN
	BtnamelistR16                          *utils.Setuprelease[BtNamelistR16]
	WlannamelistR16                        *utils.Setuprelease[WlanNamelistR16]
	SensornamelistR16                      *utils.Setuprelease[SensorNamelistR16]
	ObtaincommonlocationR16                *utils.BOOLEAN
	SlAssistanceconfignrR16                *utils.BOOLEAN
}
