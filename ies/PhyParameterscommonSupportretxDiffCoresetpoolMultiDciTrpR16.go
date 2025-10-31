package ies

import "rrc/utils"

// Phy-ParametersCommon-supportRetx-Diff-CoresetPool-Multi-DCI-TRP-r16 ::= ENUMERATED
type PhyParameterscommonSupportretxDiffCoresetpoolMultiDciTrpR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonSupportretxDiffCoresetpoolMultiDciTrpR16EnumeratedNothing = iota
	PhyParameterscommonSupportretxDiffCoresetpoolMultiDciTrpR16EnumeratedNotsupported
)
