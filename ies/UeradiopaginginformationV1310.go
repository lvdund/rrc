package ies

// UERadioPagingInformation-v1310-IEs ::= SEQUENCE
type UeradiopaginginformationV1310 struct {
	SupportedbandlisteutraforpagingR13 *[]FreqbandindicatorR11 `lb:1,ub:maxBands`
	Noncriticalextension               *UeradiopaginginformationV1610
}
