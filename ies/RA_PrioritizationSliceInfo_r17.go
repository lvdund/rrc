package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RA_PrioritizationSliceInfo_r17 struct {
	Nsag_ID_List_r17      []NSAG_ID_r17     `lb:1,ub:maxSliceInfo_r17,madatory`
	Ra_Prioritization_r17 RA_Prioritization `madatory`
}

func (ie *RA_PrioritizationSliceInfo_r17) Encode(w *aper.AperWriter) error {
	var err error
	tmp_Nsag_ID_List_r17 := utils.NewSequence[*NSAG_ID_r17]([]*NSAG_ID_r17{}, aper.Constraint{Lb: 1, Ub: maxSliceInfo_r17}, false)
	for _, i := range ie.Nsag_ID_List_r17 {
		tmp_Nsag_ID_List_r17.Value = append(tmp_Nsag_ID_List_r17.Value, &i)
	}
	if err = tmp_Nsag_ID_List_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Nsag_ID_List_r17", err)
	}
	if err = ie.Ra_Prioritization_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Ra_Prioritization_r17", err)
	}
	return nil
}

func (ie *RA_PrioritizationSliceInfo_r17) Decode(r *aper.AperReader) error {
	var err error
	tmp_Nsag_ID_List_r17 := utils.NewSequence[*NSAG_ID_r17]([]*NSAG_ID_r17{}, aper.Constraint{Lb: 1, Ub: maxSliceInfo_r17}, false)
	fn_Nsag_ID_List_r17 := func() *NSAG_ID_r17 {
		return new(NSAG_ID_r17)
	}
	if err = tmp_Nsag_ID_List_r17.Decode(r, fn_Nsag_ID_List_r17); err != nil {
		return utils.WrapError("Decode Nsag_ID_List_r17", err)
	}
	ie.Nsag_ID_List_r17 = []NSAG_ID_r17{}
	for _, i := range tmp_Nsag_ID_List_r17.Value {
		ie.Nsag_ID_List_r17 = append(ie.Nsag_ID_List_r17, *i)
	}
	if err = ie.Ra_Prioritization_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Ra_Prioritization_r17", err)
	}
	return nil
}
