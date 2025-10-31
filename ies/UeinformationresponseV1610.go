package ies

// UEInformationResponse-v1610-IEs ::= SEQUENCE
type UeinformationresponseV1610 struct {
	RachReportV1610          *RachReportV1610
	MeasresultlistextidleR16 *MeasresultlistextidleR16
	MeasresultlistidlenrR16  *MeasresultlistidlenrR16
	Noncriticalextension     *UeinformationresponseV1610IesNoncriticalextension
}
