package ies

import "rrc/utils"

// UEInformationResponse-r16-IEs ::= SEQUENCE
type UeinformationresponseR16 struct {
	MeasresultidleeutraR16   *MeasresultidleeutraR16
	MeasresultidlenrR16      *MeasresultidlenrR16
	LogmeasreportR16         *LogmeasreportR16
	ConnestfailreportR16     *ConnestfailreportR16
	RaReportlistR16          *RaReportlistR16
	RlfReportR16             *RlfReportR16
	MobilityhistoryreportR16 *MobilityhistoryreportR16
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *UeinformationresponseV1700
}
