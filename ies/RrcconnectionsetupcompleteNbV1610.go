package ies

// RRCConnectionSetupComplete-NB-v1610-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteNbV1610 struct {
	Ng5gSTmsiR16             *Ng5gSTmsiR15
	RegisteredamfR16         *RegisteredamfR15
	GummeiTypeV1610          *RrcconnectionsetupcompleteNbV1610IesGummeiTypeV1610
	GuamiTypeR16             *RrcconnectionsetupcompleteNbV1610IesGuamiTypeR16
	SNssaiListR16            *[]SNssaiR15 `lb:1,ub:maxNrofSNssaiR15`
	NgUDatatransferR16       *bool
	UpCiot5gsOptimisationR16 *bool
	RlfInfoavailableR16      *bool
	AnrInfoavailableR16      *bool
	PurConfigidR16           *PurConfigidNbR16
	Noncriticalextension     *RrcconnectionsetupcompleteNbV1610IesNoncriticalextension
}
