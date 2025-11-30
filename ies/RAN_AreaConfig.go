package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RAN_AreaConfig struct {
	TrackingAreaCode TrackingAreaCode `madatory`
	Ran_AreaCodeList []RAN_AreaCode   `lb:1,ub:32,optional`
}

func (ie *RAN_AreaConfig) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Ran_AreaCodeList) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.TrackingAreaCode.Encode(w); err != nil {
		return utils.WrapError("Encode TrackingAreaCode", err)
	}
	if len(ie.Ran_AreaCodeList) > 0 {
		tmp_Ran_AreaCodeList := utils.NewSequence[*RAN_AreaCode]([]*RAN_AreaCode{}, aper.Constraint{Lb: 1, Ub: 32}, false)
		for _, i := range ie.Ran_AreaCodeList {
			tmp_Ran_AreaCodeList.Value = append(tmp_Ran_AreaCodeList.Value, &i)
		}
		if err = tmp_Ran_AreaCodeList.Encode(w); err != nil {
			return utils.WrapError("Encode Ran_AreaCodeList", err)
		}
	}
	return nil
}

func (ie *RAN_AreaConfig) Decode(r *aper.AperReader) error {
	var err error
	var Ran_AreaCodeListPresent bool
	if Ran_AreaCodeListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.TrackingAreaCode.Decode(r); err != nil {
		return utils.WrapError("Decode TrackingAreaCode", err)
	}
	if Ran_AreaCodeListPresent {
		tmp_Ran_AreaCodeList := utils.NewSequence[*RAN_AreaCode]([]*RAN_AreaCode{}, aper.Constraint{Lb: 1, Ub: 32}, false)
		fn_Ran_AreaCodeList := func() *RAN_AreaCode {
			return new(RAN_AreaCode)
		}
		if err = tmp_Ran_AreaCodeList.Decode(r, fn_Ran_AreaCodeList); err != nil {
			return utils.WrapError("Decode Ran_AreaCodeList", err)
		}
		ie.Ran_AreaCodeList = []RAN_AreaCode{}
		for _, i := range tmp_Ran_AreaCodeList.Value {
			ie.Ran_AreaCodeList = append(ie.Ran_AreaCodeList, *i)
		}
	}
	return nil
}
