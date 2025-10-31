package ies

// SpatialRelations ::= SEQUENCE
type Spatialrelations struct {
	Maxnumberconfiguredspatialrelations  SpatialrelationsMaxnumberconfiguredspatialrelations
	Maxnumberactivespatialrelations      SpatialrelationsMaxnumberactivespatialrelations
	Additionalactivespatialrelationpucch *SpatialrelationsAdditionalactivespatialrelationpucch
	MaxnumberdlRsQclTyped                SpatialrelationsMaxnumberdlRsQclTyped
}
