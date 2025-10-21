package ies

import "rrc/utils"

// CarrierInfoNR-r15 ::= SEQUENCE
type CarrierinfonrR15 struct {
	CarrierfreqR15          ArfcnValuenrR15
	SubcarrierspacingssbR15 utils.ENUMERATED
	SmtcR15                 *MtcSsbNrR15
}
