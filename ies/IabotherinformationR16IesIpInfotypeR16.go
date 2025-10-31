package ies

// IABOtherInformation-r16-IEs-ip-InfoType-r16 ::= CHOICE
// Extensible
const (
	IabotherinformationR16IesIpInfotypeR16ChoiceNothing = iota
	IabotherinformationR16IesIpInfotypeR16ChoiceIabIpRequestR16
	IabotherinformationR16IesIpInfotypeR16ChoiceIabIpReportR16
)

type IabotherinformationR16IesIpInfotypeR16 struct {
	Choice          uint64
	IabIpRequestR16 *IabotherinformationR16IesIpInfotypeR16IabIpRequestR16
	IabIpReportR16  *IabotherinformationR16IesIpInfotypeR16IabIpReportR16
}
