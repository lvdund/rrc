package ies

import "rrc/utils"

// SL-LatencyBoundIUC-Report-r17 ::= utils.INTEGER (3..160)
type SlLatencyboundiucReportR17 struct {
	Value utils.INTEGER `lb:0,ub:160`
}
