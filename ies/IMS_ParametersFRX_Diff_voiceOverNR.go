package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	IMS_ParametersFRX_Diff_voiceOverNR_Enum_supported aper.Enumerated = 0
)

type IMS_ParametersFRX_Diff_voiceOverNR struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *IMS_ParametersFRX_Diff_voiceOverNR) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode IMS_ParametersFRX_Diff_voiceOverNR", err)
	}
	return nil
}

func (ie *IMS_ParametersFRX_Diff_voiceOverNR) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode IMS_ParametersFRX_Diff_voiceOverNR", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
