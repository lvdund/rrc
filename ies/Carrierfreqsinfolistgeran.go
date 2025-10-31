package ies

// CarrierFreqsInfoListGERAN ::= SEQUENCE OF CarrierFreqsInfoGERAN
// SIZE (1..maxGNFG)
type Carrierfreqsinfolistgeran struct {
	Value []Carrierfreqsinfogeran `lb:1,ub:maxGNFG`
}
