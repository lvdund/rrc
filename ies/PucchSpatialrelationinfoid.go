package ies

import "rrc/utils"

// PUCCH-SpatialRelationInfoId ::= utils.INTEGER (1..maxNrofSpatialRelationInfos)
type PucchSpatialrelationinfoid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSpatialRelationInfos`
}
