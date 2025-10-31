package ies

// SL-PreconfigV2X-Sync-r14 ::= SEQUENCE
// Extensible
type SlPreconfigv2xSyncR14 struct {
	SyncoffsetindicatorsR14 SlV2xSyncoffsetindicatorsR14
	SynctxparametersR14     P0SlR12
	SynctxthreshoocR14      RsrpRangesl3R12
	FiltercoefficientR14    Filtercoefficient
	SyncrefminhystR14       SlPreconfigv2xSyncR14SyncrefminhystR14
	SyncrefdiffhystR14      SlPreconfigv2xSyncR14SyncrefdiffhystR14
}
