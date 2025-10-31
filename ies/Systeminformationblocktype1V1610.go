package ies

// SystemInformationBlockType1-v1610-IEs ::= SEQUENCE
type Systeminformationblocktype1V1610 struct {
	EdrxAllowed5gcR16                *bool
	TransmissionincontrolchregionR16 *bool
	CampingallowedinceR16            *bool
	PlmnIdentitylistV1610            *PlmnIdentitylistV1610
	Noncriticalextension             *Systeminformationblocktype1V1610IesNoncriticalextension
}
