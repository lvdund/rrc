package ies

import "rrc/utils"

// FrequencyConfig-NR-r16 ::= SEQUENCE
type FrequencyconfigNrR16 struct {
	FreqbandindicatornrR16 Freqbandindicatornr
	CarriercenterfreqNrR16 ArfcnValuenr
	CarrierbandwidthNrR16  utils.INTEGER `lb:0,ub:maxNrofPhysicalResourceBlocks`
	SubcarrierspacingNrR16 Subcarrierspacing
}
