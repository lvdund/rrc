package ies

// RRCConnectionRelease-NB-v1530-IEs ::= SEQUENCE
type RrcconnectionreleaseNbV1530 struct {
	DrbContinuerohcR15      *bool
	NexthopchainingcountR15 *Nexthopchainingcount
	Noncriticalextension    *RrcconnectionreleaseNbV1550
}
