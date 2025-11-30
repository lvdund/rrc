package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM_Enum_n2 aper.Enumerated = 0
	MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM_Enum_n4 aper.Enumerated = 1
	MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM_Enum_n6 aper.Enumerated = 2
	MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM_Enum_n8 aper.Enumerated = 3
)

type MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM", err)
	}
	return nil
}

func (ie *MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
