package ies

import "rrc/utils"

// CarrierFreqsInfoListGERAN ::= SEQUENCE OF CarrierFreqsInfoGERAN
// SIZE (1..maxGNFG)
type Carrierfreqsinfolistgeran struct {
	Value utils.Sequence[Carrierfreqsinfogeran]
}
