package ies

// CA-ParametersNR-v1630 ::= SEQUENCE
type CaParametersnrV1630 struct {
	SimultxSrsAntswitchinginterbandulCaR16 *SimulsrsForantennaswitchingR16
	BeammanagementtypeR16                  *CaParametersnrV1630BeammanagementtypeR16
	IntrabandfreqseparationulAggbwGapbwR16 *CaParametersnrV1630IntrabandfreqseparationulAggbwGapbwR16
	IntercaNonalignedframeBR16             *CaParametersnrV1630IntercaNonalignedframeBR16
}
