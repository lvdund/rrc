package ies

import "rrc/utils"

// BandNR-maxUplinkDutyCycle-PC2-FR1 ::= ENUMERATED
type BandnrMaxuplinkdutycyclePc2Fr1 struct {
	Value utils.ENUMERATED
}

const (
	BandnrMaxuplinkdutycyclePc2Fr1EnumeratedNothing = iota
	BandnrMaxuplinkdutycyclePc2Fr1EnumeratedN60
	BandnrMaxuplinkdutycyclePc2Fr1EnumeratedN70
	BandnrMaxuplinkdutycyclePc2Fr1EnumeratedN80
	BandnrMaxuplinkdutycyclePc2Fr1EnumeratedN90
	BandnrMaxuplinkdutycyclePc2Fr1EnumeratedN100
)
