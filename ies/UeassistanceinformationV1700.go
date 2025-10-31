package ies

import "rrc/utils"

// UEAssistanceInformation-v1700-IEs ::= SEQUENCE
type UeassistanceinformationV1700 struct {
	UlGapfr2PreferenceR17               *UlGapfr2PreferenceR17
	MusimAssistanceR17                  *MusimAssistanceR17
	OverheatingassistanceR17            *OverheatingassistanceR17
	MaxbwPreferencefr22R17              *MaxbwPreferencefr22R17
	MaxmimoLayerpreferencefr22R17       *MaxmimoLayerpreferencefr22R17
	MinschedulingoffsetpreferenceextR17 *MinschedulingoffsetpreferenceextR17
	RlmMeasrelaxationstateR17           *utils.BOOLEAN
	BfdMeasrelaxationstateR17           *utils.BITSTRING `lb:1,ub:maxNrofServingCells`
	NonsdtDataindicationR17             *UeassistanceinformationV1700IesNonsdtDataindicationR17
	ScgDeactivationpreferenceR17        *UeassistanceinformationV1700IesScgDeactivationpreferenceR17
	UplinkdataR17                       *utils.BOOLEAN
	RrmMeasrelaxationfulfilmentR17      *utils.BOOLEAN
	PropagationdelaydifferenceR17       *PropagationdelaydifferenceR17
	Noncriticalextension                *UeassistanceinformationV1700IesNoncriticalextension
}
