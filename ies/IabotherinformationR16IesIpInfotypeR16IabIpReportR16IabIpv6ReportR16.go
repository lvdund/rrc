package ies

// IABOtherInformation-r16-IEs-ip-InfoType-r16-iab-IP-Report-r16-iab-IPv6-Report-r16 ::= CHOICE
// Extensible
const (
	IabotherinformationR16IesIpInfotypeR16IabIpReportR16IabIpv6ReportR16ChoiceNothing = iota
	IabotherinformationR16IesIpInfotypeR16IabIpReportR16IabIpv6ReportR16ChoiceIabIpv6AddressreportR16
	IabotherinformationR16IesIpInfotypeR16IabIpReportR16IabIpv6ReportR16ChoiceIabIpv6PrefixreportR16
)

type IabotherinformationR16IesIpInfotypeR16IabIpReportR16IabIpv6ReportR16 struct {
	Choice                  uint64
	IabIpv6AddressreportR16 *IabIpAddressandtrafficR16
	IabIpv6PrefixreportR16  *IabIpPrefixandtrafficR16
}
