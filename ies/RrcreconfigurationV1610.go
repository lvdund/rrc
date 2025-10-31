package ies

import "rrc/utils"

// RRCReconfiguration-v1610-IEs ::= SEQUENCE
type RrcreconfigurationV1610 struct {
	OtherconfigV1610                 *OtherconfigV1610
	BapConfigR16                     *utils.Setuprelease[BapConfigR16]
	IabIpAddressconfigurationlistR16 *IabIpAddressconfigurationlistR16
	ConditionalreconfigurationR16    *ConditionalreconfigurationR16
	DapsSourcereleaseR16             *utils.BOOLEAN
	T316R16                          *utils.Setuprelease[T316R16]
	NeedforgapsconfignrR16           *utils.Setuprelease[NeedforgapsconfignrR16]
	OndemandsibRequestR16            *utils.Setuprelease[OndemandsibRequestR16]
	DedicatedpossysinfodeliveryR16   *utils.OCTETSTRING
	SlConfigdedicatednrR16           *utils.Setuprelease[SlConfigdedicatednrR16]
	SlConfigdedicatedeutraInfoR16    *utils.Setuprelease[SlConfigdedicatedeutraInfoR16]
	TargetcellsmtcScgR16             *SsbMtc
	Noncriticalextension             *RrcreconfigurationV1700
}
