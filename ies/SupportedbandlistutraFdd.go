package ies

// SupportedBandListUTRA-FDD ::= SEQUENCE OF SupportedBandUTRA-FDD
// SIZE (1..maxBands)
type SupportedbandlistutraFdd struct {
	Value []SupportedbandutraFdd `lb:1,ub:maxBands`
}
