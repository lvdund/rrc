package ies

// CarrierInfoNR ::= SEQUENCE
// Extensible
type Carrierinfonr struct {
	Carrierfreq          ArfcnValuenr
	Ssbsubcarrierspacing Subcarrierspacing
	Smtc                 *SsbMtc
}
