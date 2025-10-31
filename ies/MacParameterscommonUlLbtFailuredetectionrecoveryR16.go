package ies

import "rrc/utils"

// MAC-ParametersCommon-ul-LBT-FailureDetectionRecovery-r16 ::= ENUMERATED
type MacParameterscommonUlLbtFailuredetectionrecoveryR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonUlLbtFailuredetectionrecoveryR16EnumeratedNothing = iota
	MacParameterscommonUlLbtFailuredetectionrecoveryR16EnumeratedSupported
)
