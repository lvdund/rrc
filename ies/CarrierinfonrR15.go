package ies

// CarrierInfoNR-r15 ::= SEQUENCE
type CarrierinfonrR15 struct {
	CarrierfreqR15          ArfcnValuenrR15
	SubcarrierspacingssbR15 CarrierinfonrR15SubcarrierspacingssbR15
	SmtcR15                 *MtcSsbNrR15
}
