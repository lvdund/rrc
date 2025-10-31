package ies

import "rrc/utils"

// UplinkPowerControlDedicatedSCell-r10 ::= SEQUENCE
type UplinkpowercontroldedicatedscellR10 struct {
	P0UePuschR10                utils.INTEGER `lb:0,ub:7`
	DeltamcsEnabledR10          UplinkpowercontroldedicatedscellR10DeltamcsEnabledR10
	AccumulationenabledR10      utils.BOOLEAN
	PsrsOffsetR10               utils.INTEGER  `lb:0,ub:15`
	PsrsOffsetapR10             *utils.INTEGER `lb:0,ub:15`
	FiltercoefficientR10        Filtercoefficient
	PathlossreferencelinkingR10 UplinkpowercontroldedicatedscellR10PathlossreferencelinkingR10
}
