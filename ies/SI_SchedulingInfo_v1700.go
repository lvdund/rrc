package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SI_SchedulingInfo_v1700 struct {
	SchedulingInfoList2_r17    []SchedulingInfo2_r17 `lb:1,ub:maxSI_Message,madatory`
	Si_RequestConfigRedCap_r17 *SI_RequestConfig     `optional`
}

func (ie *SI_SchedulingInfo_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Si_RequestConfigRedCap_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_SchedulingInfoList2_r17 := utils.NewSequence[*SchedulingInfo2_r17]([]*SchedulingInfo2_r17{}, aper.Constraint{Lb: 1, Ub: maxSI_Message}, false)
	for _, i := range ie.SchedulingInfoList2_r17 {
		tmp_SchedulingInfoList2_r17.Value = append(tmp_SchedulingInfoList2_r17.Value, &i)
	}
	if err = tmp_SchedulingInfoList2_r17.Encode(w); err != nil {
		return utils.WrapError("Encode SchedulingInfoList2_r17", err)
	}
	if ie.Si_RequestConfigRedCap_r17 != nil {
		if err = ie.Si_RequestConfigRedCap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Si_RequestConfigRedCap_r17", err)
		}
	}
	return nil
}

func (ie *SI_SchedulingInfo_v1700) Decode(r *aper.AperReader) error {
	var err error
	var Si_RequestConfigRedCap_r17Present bool
	if Si_RequestConfigRedCap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_SchedulingInfoList2_r17 := utils.NewSequence[*SchedulingInfo2_r17]([]*SchedulingInfo2_r17{}, aper.Constraint{Lb: 1, Ub: maxSI_Message}, false)
	fn_SchedulingInfoList2_r17 := func() *SchedulingInfo2_r17 {
		return new(SchedulingInfo2_r17)
	}
	if err = tmp_SchedulingInfoList2_r17.Decode(r, fn_SchedulingInfoList2_r17); err != nil {
		return utils.WrapError("Decode SchedulingInfoList2_r17", err)
	}
	ie.SchedulingInfoList2_r17 = []SchedulingInfo2_r17{}
	for _, i := range tmp_SchedulingInfoList2_r17.Value {
		ie.SchedulingInfoList2_r17 = append(ie.SchedulingInfoList2_r17, *i)
	}
	if Si_RequestConfigRedCap_r17Present {
		ie.Si_RequestConfigRedCap_r17 = new(SI_RequestConfig)
		if err = ie.Si_RequestConfigRedCap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Si_RequestConfigRedCap_r17", err)
		}
	}
	return nil
}
