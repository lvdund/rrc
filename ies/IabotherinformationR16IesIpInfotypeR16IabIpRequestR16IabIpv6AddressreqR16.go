package ies

// IABOtherInformation-r16-IEs-ip-InfoType-r16-iab-IP-Request-r16-iab-IPv6-AddressReq-r16 ::= CHOICE
// Extensible
const (
	IabotherinformationR16IesIpInfotypeR16IabIpRequestR16IabIpv6AddressreqR16ChoiceNothing = iota
	IabotherinformationR16IesIpInfotypeR16IabIpRequestR16IabIpv6AddressreqR16ChoiceIabIpv6AddressnumreqR16
	IabotherinformationR16IesIpInfotypeR16IabIpRequestR16IabIpv6AddressreqR16ChoiceIabIpv6AddressprefixreqR16
)

type IabotherinformationR16IesIpInfotypeR16IabIpRequestR16IabIpv6AddressreqR16 struct {
	Choice                     uint64
	IabIpv6AddressnumreqR16    *IabIpAddressnumreqR16
	IabIpv6AddressprefixreqR16 *IabIpAddressprefixreqR16
}
