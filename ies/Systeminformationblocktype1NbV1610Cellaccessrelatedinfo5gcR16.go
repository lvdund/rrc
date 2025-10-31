package ies

// SystemInformationBlockType1-NB-v1610-cellAccessRelatedInfo-5GC-r16 ::= SEQUENCE
type Systeminformationblocktype1NbV1610Cellaccessrelatedinfo5gcR16 struct {
	PlmnIdentitylistR16    PlmnIdentitylist5gcNbR16
	Trackingareacode5gcR16 Trackingareacode5gcR15
	CellidentityR16        *Cellidentity
	Cellbarred5gcR16       Systeminformationblocktype1NbV1610Cellaccessrelatedinfo5gcR16Cellbarred5gcR16
}
