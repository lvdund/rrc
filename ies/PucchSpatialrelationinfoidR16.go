package ies

import "rrc/utils"

// PUCCH-SpatialRelationInfoId-r16 ::= utils.INTEGER (1..maxNrofSpatialRelationInfos-r16)
type PucchSpatialrelationinfoidR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSpatialRelationInfosR16`
}
