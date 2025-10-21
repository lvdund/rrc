package ies

import "rrc/utils"

// UplinkPowerControlDedicatedSCell-r10 ::= SEQUENCE
type UplinkpowercontroldedicatedscellR10 struct {
	P0UePuschR10                utils.INTEGER
	DeltamcsEnabledR10          utils.ENUMERATED
	AccumulationenabledR10      bool
	PsrsOffsetR10               utils.INTEGER
	PsrsOffsetapR10             *utils.INTEGER
	FiltercoefficientR10        Filtercoefficient
	PathlossreferencelinkingR10 utils.ENUMERATED
}
