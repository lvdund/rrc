package ies

// UEInformationResponse-r9-IEs ::= SEQUENCE
type UeinformationresponseR9 struct {
	RachReportR9         *RachReportR16
	RlfReportR9          *RlfReportR9
	Noncriticalextension *UeinformationresponseV930
}
