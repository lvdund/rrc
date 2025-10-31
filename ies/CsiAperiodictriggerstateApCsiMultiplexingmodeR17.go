package ies

import "rrc/utils"

// CSI-AperiodicTriggerState-ap-CSI-MultiplexingMode-r17 ::= ENUMERATED
type CsiAperiodictriggerstateApCsiMultiplexingmodeR17 struct {
	Value utils.ENUMERATED
}

const (
	CsiAperiodictriggerstateApCsiMultiplexingmodeR17EnumeratedNothing = iota
	CsiAperiodictriggerstateApCsiMultiplexingmodeR17EnumeratedEnabled
)
