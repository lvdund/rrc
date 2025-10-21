package ies

import "rrc/utils"

// RRCConnectionSetupComplete-NB-v1610-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteNbV1610Ies struct {
	Ng5gSTmsiR16             *Ng5gSTmsiR15
	RegisteredamfR16         *RegisteredamfR15
	GummeiTypeV1610          *utils.ENUMERATED
	GuamiTypeR16             *utils.ENUMERATED
	SNssaiListR16            *interface{}
	NgUDatatransferR16       *utils.ENUMERATED
	UpCiot5gsOptimisationR16 *utils.ENUMERATED
	RlfInfoavailableR16      *utils.ENUMERATED
	AnrInfoavailableR16      *utils.ENUMERATED
	PurConfigidR16           *PurConfigidNbR16
	Noncriticalextension     *interface{}
}
