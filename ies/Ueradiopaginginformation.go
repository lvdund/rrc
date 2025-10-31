package ies

// UERadioPagingInformation-IEs ::= SEQUENCE
type Ueradiopaginginformation struct {
	Supportedbandlistnrforpaging *[]Freqbandindicatornr `lb:1,ub:maxBands`
	Noncriticalextension         *UeradiopaginginformationV15e0
}
