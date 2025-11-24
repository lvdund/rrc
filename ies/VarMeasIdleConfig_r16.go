package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarMeasIdleConfig_r16 struct {
	MeasIdleCarrierListNR_r16    []MeasIdleCarrierNR_r16                    `lb:1,ub:maxFreqIdle_r16,optional`
	MeasIdleCarrierListEUTRA_r16 []MeasIdleCarrierEUTRA_r16                 `lb:1,ub:maxFreqIdle_r16,optional`
	MeasIdleDuration_r16         VarMeasIdleConfig_r16_measIdleDuration_r16 `madatory`
	ValidityAreaList_r16         *ValidityAreaList_r16                      `optional`
}

func (ie *VarMeasIdleConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.MeasIdleCarrierListNR_r16) > 0, len(ie.MeasIdleCarrierListEUTRA_r16) > 0, ie.ValidityAreaList_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.MeasIdleCarrierListNR_r16) > 0 {
		tmp_MeasIdleCarrierListNR_r16 := utils.NewSequence[*MeasIdleCarrierNR_r16]([]*MeasIdleCarrierNR_r16{}, uper.Constraint{Lb: 1, Ub: maxFreqIdle_r16}, false)
		for _, i := range ie.MeasIdleCarrierListNR_r16 {
			tmp_MeasIdleCarrierListNR_r16.Value = append(tmp_MeasIdleCarrierListNR_r16.Value, &i)
		}
		if err = tmp_MeasIdleCarrierListNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasIdleCarrierListNR_r16", err)
		}
	}
	if len(ie.MeasIdleCarrierListEUTRA_r16) > 0 {
		tmp_MeasIdleCarrierListEUTRA_r16 := utils.NewSequence[*MeasIdleCarrierEUTRA_r16]([]*MeasIdleCarrierEUTRA_r16{}, uper.Constraint{Lb: 1, Ub: maxFreqIdle_r16}, false)
		for _, i := range ie.MeasIdleCarrierListEUTRA_r16 {
			tmp_MeasIdleCarrierListEUTRA_r16.Value = append(tmp_MeasIdleCarrierListEUTRA_r16.Value, &i)
		}
		if err = tmp_MeasIdleCarrierListEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasIdleCarrierListEUTRA_r16", err)
		}
	}
	if err = ie.MeasIdleDuration_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MeasIdleDuration_r16", err)
	}
	if ie.ValidityAreaList_r16 != nil {
		if err = ie.ValidityAreaList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ValidityAreaList_r16", err)
		}
	}
	return nil
}

func (ie *VarMeasIdleConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var MeasIdleCarrierListNR_r16Present bool
	if MeasIdleCarrierListNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasIdleCarrierListEUTRA_r16Present bool
	if MeasIdleCarrierListEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ValidityAreaList_r16Present bool
	if ValidityAreaList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasIdleCarrierListNR_r16Present {
		tmp_MeasIdleCarrierListNR_r16 := utils.NewSequence[*MeasIdleCarrierNR_r16]([]*MeasIdleCarrierNR_r16{}, uper.Constraint{Lb: 1, Ub: maxFreqIdle_r16}, false)
		fn_MeasIdleCarrierListNR_r16 := func() *MeasIdleCarrierNR_r16 {
			return new(MeasIdleCarrierNR_r16)
		}
		if err = tmp_MeasIdleCarrierListNR_r16.Decode(r, fn_MeasIdleCarrierListNR_r16); err != nil {
			return utils.WrapError("Decode MeasIdleCarrierListNR_r16", err)
		}
		ie.MeasIdleCarrierListNR_r16 = []MeasIdleCarrierNR_r16{}
		for _, i := range tmp_MeasIdleCarrierListNR_r16.Value {
			ie.MeasIdleCarrierListNR_r16 = append(ie.MeasIdleCarrierListNR_r16, *i)
		}
	}
	if MeasIdleCarrierListEUTRA_r16Present {
		tmp_MeasIdleCarrierListEUTRA_r16 := utils.NewSequence[*MeasIdleCarrierEUTRA_r16]([]*MeasIdleCarrierEUTRA_r16{}, uper.Constraint{Lb: 1, Ub: maxFreqIdle_r16}, false)
		fn_MeasIdleCarrierListEUTRA_r16 := func() *MeasIdleCarrierEUTRA_r16 {
			return new(MeasIdleCarrierEUTRA_r16)
		}
		if err = tmp_MeasIdleCarrierListEUTRA_r16.Decode(r, fn_MeasIdleCarrierListEUTRA_r16); err != nil {
			return utils.WrapError("Decode MeasIdleCarrierListEUTRA_r16", err)
		}
		ie.MeasIdleCarrierListEUTRA_r16 = []MeasIdleCarrierEUTRA_r16{}
		for _, i := range tmp_MeasIdleCarrierListEUTRA_r16.Value {
			ie.MeasIdleCarrierListEUTRA_r16 = append(ie.MeasIdleCarrierListEUTRA_r16, *i)
		}
	}
	if err = ie.MeasIdleDuration_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MeasIdleDuration_r16", err)
	}
	if ValidityAreaList_r16Present {
		ie.ValidityAreaList_r16 = new(ValidityAreaList_r16)
		if err = ie.ValidityAreaList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ValidityAreaList_r16", err)
		}
	}
	return nil
}
