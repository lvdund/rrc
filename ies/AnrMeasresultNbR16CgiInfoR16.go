package ies

// ANR-MeasResult-NB-r16-cgi-Info-r16 ::= SEQUENCE
type AnrMeasresultNbR16CgiInfoR16 struct {
	CellglobalidR16     Cellglobalideutra
	TrackingareacodeR16 Trackingareacode
	PlmnIdentitylistR16 *PlmnIdentitylist2
}
