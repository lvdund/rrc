package ies

// RRCConnectionRelease-v1530-IEs ::= SEQUENCE
type RrcconnectionreleaseV1530 struct {
	DrbContinuerohcR15      *bool
	NexthopchainingcountR15 *Nexthopchainingcount
	MeasidleconfigR15       *MeasidleconfigdedicatedR15
	RrcInactiveconfigR15    *RrcInactiveconfigR15
	CnTypeR15               *RrcconnectionreleaseV1530IesCnTypeR15
	Noncriticalextension    *RrcconnectionreleaseV1540
}
