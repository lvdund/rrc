package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type AvailabilityCombination_r16 struct {
	AvailabilityCombinationId_r16 AvailabilityCombinationId_r16 `madatory`
	ResourceAvailability_r16      []int64                       `lb:1,ub:maxNrofResourceAvailabilityPerCombination_r16,e_lb:0,e_ub:7,madatory`
}

func (ie *AvailabilityCombination_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.AvailabilityCombinationId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode AvailabilityCombinationId_r16", err)
	}
	tmp_ResourceAvailability_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofResourceAvailabilityPerCombination_r16}, false)
	for _, i := range ie.ResourceAvailability_r16 {
		tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 7}, false)
		tmp_ResourceAvailability_r16.Value = append(tmp_ResourceAvailability_r16.Value, &tmp_ie)
	}
	if err = tmp_ResourceAvailability_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ResourceAvailability_r16", err)
	}
	return nil
}

func (ie *AvailabilityCombination_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.AvailabilityCombinationId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode AvailabilityCombinationId_r16", err)
	}
	tmp_ResourceAvailability_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofResourceAvailabilityPerCombination_r16}, false)
	fn_ResourceAvailability_r16 := func() *utils.INTEGER {
		ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 7}, false)
		return &ie
	}
	if err = tmp_ResourceAvailability_r16.Decode(r, fn_ResourceAvailability_r16); err != nil {
		return utils.WrapError("Decode ResourceAvailability_r16", err)
	}
	ie.ResourceAvailability_r16 = []int64{}
	for _, i := range tmp_ResourceAvailability_r16.Value {
		ie.ResourceAvailability_r16 = append(ie.ResourceAvailability_r16, int64(i.Value))
	}
	return nil
}
