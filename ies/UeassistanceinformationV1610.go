package ies

import "rrc/utils"

// UEAssistanceInformation-v1610-IEs ::= SEQUENCE
type UeassistanceinformationV1610 struct {
	IdcAssistanceR16                 *IdcAssistanceR16
	DrxPreferenceR16                 *DrxPreferenceR16
	MaxbwPreferenceR16               *MaxbwPreferenceR16
	MaxccPreferenceR16               *MaxccPreferenceR16
	MaxmimoLayerpreferenceR16        *MaxmimoLayerpreferenceR16
	MinschedulingoffsetpreferenceR16 *MinschedulingoffsetpreferenceR16
	ReleasepreferenceR16             *ReleasepreferenceR16
	SlUeAssistanceinformationnrR16   *SlUeAssistanceinformationnrR16
	ReferencetimeinfopreferenceR16   *utils.BOOLEAN
	Noncriticalextension             *UeassistanceinformationV1700
}
