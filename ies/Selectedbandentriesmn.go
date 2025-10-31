package ies

// SelectedBandEntriesMN ::= SEQUENCE OF BandEntryIndex
// SIZE (1..maxSimultaneousBands)
type Selectedbandentriesmn struct {
	Value []Bandentryindex `lb:1,ub:maxSimultaneousBands`
}
