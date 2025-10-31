package ies

import "rrc/utils"

// Phy-ParametersCommon-monitoringDCI-SameSearchSpace-r16 ::= ENUMERATED
type PhyParameterscommonMonitoringdciSamesearchspaceR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonMonitoringdciSamesearchspaceR16EnumeratedNothing = iota
	PhyParameterscommonMonitoringdciSamesearchspaceR16EnumeratedSupported
)
