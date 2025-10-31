package ies

import "rrc/utils"

// BandCombinationParameters-r13-dc-Support-r13-asynchronous-r13 ::= ENUMERATED
type BandcombinationparametersR13DcSupportR13AsynchronousR13 struct {
	Value utils.ENUMERATED
}

const (
	BandcombinationparametersR13DcSupportR13AsynchronousR13EnumeratedNothing = iota
	BandcombinationparametersR13DcSupportR13AsynchronousR13EnumeratedSupported
)
