package ies

import "rrc/utils"

// MIMO-ParametersPerBand-multiDCI-multiTRP-Parameters-r16 ::= SEQUENCE
type MimoParametersperbandMultidciMultitrpParametersR16 struct {
	OverlappdschsfullyfreqtimeR16       *utils.INTEGER `lb:0,ub:2`
	OverlappdschsintimepartiallyfreqR16 *MimoParametersperbandMultidciMultitrpParametersR16OverlappdschsintimepartiallyfreqR16
	OutoforderoperationdlR16            *MimoParametersperbandMultidciMultitrpParametersR16OutoforderoperationdlR16
	OutoforderoperationulR16            *MimoParametersperbandMultidciMultitrpParametersR16OutoforderoperationulR16
	SeparatecrsRatematchingR16          *MimoParametersperbandMultidciMultitrpParametersR16SeparatecrsRatematchingR16
	DefaultqclPercoresetpoolindexR16    *MimoParametersperbandMultidciMultitrpParametersR16DefaultqclPercoresetpoolindexR16
	MaxnumberactivatedtciStatesR16      *MimoParametersperbandMultidciMultitrpParametersR16MaxnumberactivatedtciStatesR16
}
