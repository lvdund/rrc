package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RA_PrioritizationSliceInfoList_r17 struct {
	Value []RA_PrioritizationSliceInfo_r17 `lb:1,ub:maxSliceInfo_r17,madatory`
}

func (ie *RA_PrioritizationSliceInfoList_r17) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*RA_PrioritizationSliceInfo_r17]([]*RA_PrioritizationSliceInfo_r17{}, aper.Constraint{Lb: 1, Ub: maxSliceInfo_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode RA_PrioritizationSliceInfoList_r17", err)
	}
	return nil
}

func (ie *RA_PrioritizationSliceInfoList_r17) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*RA_PrioritizationSliceInfo_r17]([]*RA_PrioritizationSliceInfo_r17{}, aper.Constraint{Lb: 1, Ub: maxSliceInfo_r17}, false)
	fn := func() *RA_PrioritizationSliceInfo_r17 {
		return new(RA_PrioritizationSliceInfo_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode RA_PrioritizationSliceInfoList_r17", err)
	}
	ie.Value = []RA_PrioritizationSliceInfo_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
