package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CA_ParametersNR_v1610_crossCarrierA_CSI_trigDiffSCS_r16_Enum_higherA_CSI_SCS aper.Enumerated = 0
	CA_ParametersNR_v1610_crossCarrierA_CSI_trigDiffSCS_r16_Enum_lowerA_CSI_SCS  aper.Enumerated = 1
	CA_ParametersNR_v1610_crossCarrierA_CSI_trigDiffSCS_r16_Enum_both            aper.Enumerated = 2
)

type CA_ParametersNR_v1610_crossCarrierA_CSI_trigDiffSCS_r16 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *CA_ParametersNR_v1610_crossCarrierA_CSI_trigDiffSCS_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode CA_ParametersNR_v1610_crossCarrierA_CSI_trigDiffSCS_r16", err)
	}
	return nil
}

func (ie *CA_ParametersNR_v1610_crossCarrierA_CSI_trigDiffSCS_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode CA_ParametersNR_v1610_crossCarrierA_CSI_trigDiffSCS_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
