package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_supportInter_slotTDM_r16 struct {
	SupportRepNumPDSCH_TDRA_r16 MIMO_ParametersPerBand_supportInter_slotTDM_r16_supportRepNumPDSCH_TDRA_r16 `madatory`
	MaxTBS_Size_r16             MIMO_ParametersPerBand_supportInter_slotTDM_r16_maxTBS_Size_r16             `madatory`
	MaxNumberTCI_states_r16     int64                                                                       `lb:1,ub:2,madatory`
}

func (ie *MIMO_ParametersPerBand_supportInter_slotTDM_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.SupportRepNumPDSCH_TDRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode SupportRepNumPDSCH_TDRA_r16", err)
	}
	if err = ie.MaxTBS_Size_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxTBS_Size_r16", err)
	}
	if err = w.WriteInteger(ie.MaxNumberTCI_states_r16, &aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberTCI_states_r16", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_supportInter_slotTDM_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.SupportRepNumPDSCH_TDRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode SupportRepNumPDSCH_TDRA_r16", err)
	}
	if err = ie.MaxTBS_Size_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxTBS_Size_r16", err)
	}
	var tmp_int_MaxNumberTCI_states_r16 int64
	if tmp_int_MaxNumberTCI_states_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberTCI_states_r16", err)
	}
	ie.MaxNumberTCI_states_r16 = tmp_int_MaxNumberTCI_states_r16
	return nil
}
