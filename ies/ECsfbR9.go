package ies

import "rrc/utils"

// E-CSFB-r9 ::= SEQUENCE
type ECsfbR9 struct {
	Messagecontcdma20001xrttR9    *utils.OCTETSTRING
	Mobilitycdma2000HrpdR9        *utils.ENUMERATED
	Messagecontcdma2000HrpdR9     *utils.OCTETSTRING
	Redirectcarriercdma2000HrpdR9 *Carrierfreqcdma2000
}
