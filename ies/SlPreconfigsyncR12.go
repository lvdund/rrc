package ies

// SL-PreconfigSync-r12 ::= SEQUENCE
// Extensible
type SlPreconfigsyncR12 struct {
	SynccpLenR12            SlCpLenR12
	Syncoffsetindicator1R12 SlOffsetindicatorsyncR12
	Syncoffsetindicator2R12 SlOffsetindicatorsyncR12
	SynctxparametersR12     P0SlR12
	SynctxthreshoocR12      RsrpRangesl3R12
	FiltercoefficientR12    Filtercoefficient
	SyncrefminhystR12       SlPreconfigsyncR12SyncrefminhystR12
	SyncrefdiffhystR12      SlPreconfigsyncR12SyncrefdiffhystR12
}
