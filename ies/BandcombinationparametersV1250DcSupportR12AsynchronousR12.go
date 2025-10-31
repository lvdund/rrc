package ies

import "rrc/utils"

// BandCombinationParameters-v1250-dc-Support-r12-asynchronous-r12 ::= ENUMERATED
type BandcombinationparametersV1250DcSupportR12AsynchronousR12 struct {
	Value utils.ENUMERATED
}

const (
	BandcombinationparametersV1250DcSupportR12AsynchronousR12EnumeratedNothing = iota
	BandcombinationparametersV1250DcSupportR12AsynchronousR12EnumeratedSupported
)
