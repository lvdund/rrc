package ies

import "rrc/utils"

// PUCCH-SpatialRelationInfoId-v1610 ::= utils.INTEGER (maxNrofSpatialRelationInfos-plus-1..maxNrofSpatialRelationInfos-r16)
type PucchSpatialrelationinfoidV1610 struct {
	Value utils.INTEGER `lb:maxNrofSpatialRelationInfosPlus1,ub:maxNrofSpatialRelationInfosR16`
}
