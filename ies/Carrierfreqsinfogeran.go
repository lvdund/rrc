package ies

// CarrierFreqsInfoGERAN ::= SEQUENCE
// Extensible
type Carrierfreqsinfogeran struct {
	Carrierfreqs Carrierfreqsgeran
	Commoninfo   *CarrierfreqsinfogeranCommoninfo
}
