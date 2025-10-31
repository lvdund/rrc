package ies

// RRCReestablishmentComplete-v1610-IEs ::= SEQUENCE
type RrcreestablishmentcompleteV1610 struct {
	UeMeasurementsavailableR16 *UeMeasurementsavailableR16
	Noncriticalextension       *RrcreestablishmentcompleteV1610IesNoncriticalextension
}
