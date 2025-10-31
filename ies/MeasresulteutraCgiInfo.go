package ies

// MeasResultEUTRA-cgi-Info ::= SEQUENCE
type MeasresulteutraCgiInfo struct {
	Cellglobalid     Cellglobalideutra
	Trackingareacode Trackingareacode
	PlmnIdentitylist *PlmnIdentitylist2
}
