package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SI_RequestConfig struct {
	Rach_OccasionsSI    *SI_RequestConfig_rach_OccasionsSI `optional`
	Si_RequestPeriod    *SI_RequestConfig_si_RequestPeriod `optional`
	Si_RequestResources []SI_RequestResources              `lb:1,ub:maxSI_Message,madatory`
}

func (ie *SI_RequestConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Rach_OccasionsSI != nil, ie.Si_RequestPeriod != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Rach_OccasionsSI != nil {
		if err = ie.Rach_OccasionsSI.Encode(w); err != nil {
			return utils.WrapError("Encode Rach_OccasionsSI", err)
		}
	}
	if ie.Si_RequestPeriod != nil {
		if err = ie.Si_RequestPeriod.Encode(w); err != nil {
			return utils.WrapError("Encode Si_RequestPeriod", err)
		}
	}
	tmp_Si_RequestResources := utils.NewSequence[*SI_RequestResources]([]*SI_RequestResources{}, uper.Constraint{Lb: 1, Ub: maxSI_Message}, false)
	for _, i := range ie.Si_RequestResources {
		tmp_Si_RequestResources.Value = append(tmp_Si_RequestResources.Value, &i)
	}
	if err = tmp_Si_RequestResources.Encode(w); err != nil {
		return utils.WrapError("Encode Si_RequestResources", err)
	}
	return nil
}

func (ie *SI_RequestConfig) Decode(r *uper.UperReader) error {
	var err error
	var Rach_OccasionsSIPresent bool
	if Rach_OccasionsSIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Si_RequestPeriodPresent bool
	if Si_RequestPeriodPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Rach_OccasionsSIPresent {
		ie.Rach_OccasionsSI = new(SI_RequestConfig_rach_OccasionsSI)
		if err = ie.Rach_OccasionsSI.Decode(r); err != nil {
			return utils.WrapError("Decode Rach_OccasionsSI", err)
		}
	}
	if Si_RequestPeriodPresent {
		ie.Si_RequestPeriod = new(SI_RequestConfig_si_RequestPeriod)
		if err = ie.Si_RequestPeriod.Decode(r); err != nil {
			return utils.WrapError("Decode Si_RequestPeriod", err)
		}
	}
	tmp_Si_RequestResources := utils.NewSequence[*SI_RequestResources]([]*SI_RequestResources{}, uper.Constraint{Lb: 1, Ub: maxSI_Message}, false)
	fn_Si_RequestResources := func() *SI_RequestResources {
		return new(SI_RequestResources)
	}
	if err = tmp_Si_RequestResources.Decode(r, fn_Si_RequestResources); err != nil {
		return utils.WrapError("Decode Si_RequestResources", err)
	}
	ie.Si_RequestResources = []SI_RequestResources{}
	for _, i := range tmp_Si_RequestResources.Value {
		ie.Si_RequestResources = append(ie.Si_RequestResources, *i)
	}
	return nil
}
