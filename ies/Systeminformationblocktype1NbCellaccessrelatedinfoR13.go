package ies

// SystemInformationBlockType1-NB-cellAccessRelatedInfo-r13 ::= SEQUENCE
type Systeminformationblocktype1NbCellaccessrelatedinfoR13 struct {
	PlmnIdentitylistR13     PlmnIdentitylistNbR13
	TrackingareacodeR13     Trackingareacode
	CellidentityR13         Cellidentity
	CellbarredR13           Systeminformationblocktype1NbCellaccessrelatedinfoR13CellbarredR13
	IntrafreqreselectionR13 Systeminformationblocktype1NbCellaccessrelatedinfoR13IntrafreqreselectionR13
}
