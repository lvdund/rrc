package ies

import "rrc/utils"

// SystemInformationBlockType1-v1610-IEs ::= SEQUENCE
type Systeminformationblocktype1V1610Ies struct {
	EdrxAllowed5gcR16                *utils.ENUMERATED
	TransmissionincontrolchregionR16 *utils.ENUMERATED
	CampingallowedinceR16            *utils.ENUMERATED
	PlmnIdentitylistV1610            *PlmnIdentitylistV1610
	Noncriticalextension             *interface{}
}
