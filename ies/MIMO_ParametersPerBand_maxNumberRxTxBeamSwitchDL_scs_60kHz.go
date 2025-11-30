package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_scs_60kHz_Enum_n4  aper.Enumerated = 0
	MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_scs_60kHz_Enum_n7  aper.Enumerated = 1
	MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_scs_60kHz_Enum_n14 aper.Enumerated = 2
)

type MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_scs_60kHz struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_scs_60kHz) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_scs_60kHz", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_scs_60kHz) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_scs_60kHz", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
