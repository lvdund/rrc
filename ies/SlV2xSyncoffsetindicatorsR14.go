package ies

// SL-V2X-SyncOffsetIndicators-r14 ::= SEQUENCE
type SlV2xSyncoffsetindicatorsR14 struct {
	Syncoffsetindicator1R14 SlOffsetindicatorsyncR14
	Syncoffsetindicator2R14 SlOffsetindicatorsyncR14
	Syncoffsetindicator3R14 *SlOffsetindicatorsyncR14
}
