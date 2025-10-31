package ies

// CellAccessRelatedInfo-EUTRA-5GC ::= SEQUENCE
type CellaccessrelatedinfoEutra5gc struct {
	PlmnIdentitylistEutra5gc PlmnIdentitylistEutra5gc
	TrackingareacodeEutra5gc Trackingareacode
	Ranac5gc                 *RanAreacode
	CellidentityEutra5gc     CellidentityEutra5gc
}
