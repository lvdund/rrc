package ies

// RRCReconfigurationComplete-v1610-IEs ::= SEQUENCE
type RrcreconfigurationcompleteV1610 struct {
	UeMeasurementsavailableR16 *UeMeasurementsavailableR16
	NeedforgapsinfonrR16       *NeedforgapsinfonrR16
	Noncriticalextension       *RrcreconfigurationcompleteV1640
}
