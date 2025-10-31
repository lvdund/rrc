package ies

// SidelinkUEInformationNR-v1700-IEs ::= SEQUENCE
type SidelinkueinformationnrV1700 struct {
	SlTxresourcereqlistV1700        *SlTxresourcereqlistV1700
	SlRxdrxReportlistV1700          *SlRxdrxReportlistV1700
	SlRxinterestedgcBcDestlistR17   *SlRxinterestedgcBcDestlistR17
	SlRxinterestedfreqlistdiscR17   *SlInterestedfreqlistR16
	SlTxresourcereqlistdiscR17      *SlTxresourcereqlistdiscR17
	SlTxresourcereqlistcommrelayR17 *SlTxresourcereqlistcommrelayR17
	UeTypeR17                       *SidelinkueinformationnrV1700IesUeTypeR17
	SlSourceidentityremoteueR17     *SlSourceidentityR17
	Noncriticalextension            *SidelinkueinformationnrV1700IesNoncriticalextension
}
