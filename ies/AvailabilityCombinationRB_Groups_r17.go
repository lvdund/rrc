package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AvailabilityCombinationRB_Groups_r17 struct {
	AvailabilityCombinationId_r17 AvailabilityCombinationId_r16 `madatory`
	Rb_SetGroups_r17              []RB_SetGroup_r17             `lb:1,ub:maxNrofRB_SetGroups_r17,optional`
	ResourceAvailability_r17      []int64                       `lb:1,ub:maxNrofResourceAvailabilityPerCombination_r16,e_lb:0,e_ub:7,optional`
}

func (ie *AvailabilityCombinationRB_Groups_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Rb_SetGroups_r17) > 0, len(ie.ResourceAvailability_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.AvailabilityCombinationId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode AvailabilityCombinationId_r17", err)
	}
	if len(ie.Rb_SetGroups_r17) > 0 {
		tmp_Rb_SetGroups_r17 := utils.NewSequence[*RB_SetGroup_r17]([]*RB_SetGroup_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofRB_SetGroups_r17}, false)
		for _, i := range ie.Rb_SetGroups_r17 {
			tmp_Rb_SetGroups_r17.Value = append(tmp_Rb_SetGroups_r17.Value, &i)
		}
		if err = tmp_Rb_SetGroups_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Rb_SetGroups_r17", err)
		}
	}
	if len(ie.ResourceAvailability_r17) > 0 {
		tmp_ResourceAvailability_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofResourceAvailabilityPerCombination_r16}, false)
		for _, i := range ie.ResourceAvailability_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 7}, false)
			tmp_ResourceAvailability_r17.Value = append(tmp_ResourceAvailability_r17.Value, &tmp_ie)
		}
		if err = tmp_ResourceAvailability_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ResourceAvailability_r17", err)
		}
	}
	return nil
}

func (ie *AvailabilityCombinationRB_Groups_r17) Decode(r *uper.UperReader) error {
	var err error
	var Rb_SetGroups_r17Present bool
	if Rb_SetGroups_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ResourceAvailability_r17Present bool
	if ResourceAvailability_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.AvailabilityCombinationId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode AvailabilityCombinationId_r17", err)
	}
	if Rb_SetGroups_r17Present {
		tmp_Rb_SetGroups_r17 := utils.NewSequence[*RB_SetGroup_r17]([]*RB_SetGroup_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofRB_SetGroups_r17}, false)
		fn_Rb_SetGroups_r17 := func() *RB_SetGroup_r17 {
			return new(RB_SetGroup_r17)
		}
		if err = tmp_Rb_SetGroups_r17.Decode(r, fn_Rb_SetGroups_r17); err != nil {
			return utils.WrapError("Decode Rb_SetGroups_r17", err)
		}
		ie.Rb_SetGroups_r17 = []RB_SetGroup_r17{}
		for _, i := range tmp_Rb_SetGroups_r17.Value {
			ie.Rb_SetGroups_r17 = append(ie.Rb_SetGroups_r17, *i)
		}
	}
	if ResourceAvailability_r17Present {
		tmp_ResourceAvailability_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofResourceAvailabilityPerCombination_r16}, false)
		fn_ResourceAvailability_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 7}, false)
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
	return nil
}
