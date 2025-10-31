package ies

import "rrc/utils"

// CSI-SemiPersistentOnPUSCH-TriggerState-sp-CSI-MultiplexingMode-r17 ::= ENUMERATED
type CsiSemipersistentonpuschTriggerstateSpCsiMultiplexingmodeR17 struct {
	Value utils.ENUMERATED
}

const (
	CsiSemipersistentonpuschTriggerstateSpCsiMultiplexingmodeR17EnumeratedNothing = iota
	CsiSemipersistentonpuschTriggerstateSpCsiMultiplexingmodeR17EnumeratedEnabled
)
