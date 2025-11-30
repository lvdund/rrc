package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RB_SetGroup_r17 struct {
	ResourceAvailability_r17 []int64 `lb:1,ub:maxNrofResourceAvailabilityPerCombination_r16,e_lb:0,e_ub:7,optional`
	Rb_Sets_r17              []int64 `lb:1,ub:maxNrofRB_Sets_r17,e_lb:0,e_ub:7,optional`
}

func (ie *RB_SetGroup_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.ResourceAvailability_r17) > 0, len(ie.Rb_Sets_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.ResourceAvailability_r17) > 0 {
		tmp_ResourceAvailability_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofResourceAvailabilityPerCombination_r16}, false)
		for _, i := range ie.ResourceAvailability_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 7}, false)
			tmp_ResourceAvailability_r17.Value = append(tmp_ResourceAvailability_r17.Value, &tmp_ie)
		}
		if err = tmp_ResourceAvailability_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ResourceAvailability_r17", err)
		}
	}
	if len(ie.Rb_Sets_r17) > 0 {
		tmp_Rb_Sets_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofRB_Sets_r17}, false)
		for _, i := range ie.Rb_Sets_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 7}, false)
			tmp_Rb_Sets_r17.Value = append(tmp_Rb_Sets_r17.Value, &tmp_ie)
		}
		if err = tmp_Rb_Sets_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Rb_Sets_r17", err)
		}
	}
	return nil
}

func (ie *RB_SetGroup_r17) Decode(r *aper.AperReader) error {
	var err error
	var ResourceAvailability_r17Present bool
	if ResourceAvailability_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rb_Sets_r17Present bool
	if Rb_Sets_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ResourceAvailability_r17Present {
		tmp_ResourceAvailability_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofResourceAvailabilityPerCombination_r16}, false)
		fn_ResourceAvailability_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 7}, false)
			return &ie
		}
		if err = tmp_ResourceAvailability_r17.Decode(r, fn_ResourceAvailability_r17); err != nil {
			return utils.WrapError("Decode ResourceAvailability_r17", err)
		}
		ie.ResourceAvailability_r17 = []int64{}
		for _, i := range tmp_ResourceAvailability_r17.Value {
			ie.ResourceAvailability_r17 = append(ie.ResourceAvailability_r17, int64(i.Value))
		}
	}
	if Rb_Sets_r17Present {
		tmp_Rb_Sets_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofRB_Sets_r17}, false)
		fn_Rb_Sets_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 7}, false)
			return &ie
		}
		if err = tmp_Rb_Sets_r17.Decode(r, fn_Rb_Sets_r17); err != nil {
			return utils.WrapError("Decode Rb_Sets_r17", err)
		}
		ie.Rb_Sets_r17 = []int64{}
		for _, i := range tmp_Rb_Sets_r17.Value {
			ie.Rb_Sets_r17 = append(ie.Rb_Sets_r17, int64(i.Value))
		}
	}
	return nil
}
