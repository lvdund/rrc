package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SliceInfoList_r17 struct {
	Value []SliceInfo_r17 `lb:1,ub:maxSliceInfo_r17,madatory`
}

func (ie *SliceInfoList_r17) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*SliceInfo_r17]([]*SliceInfo_r17{}, aper.Constraint{Lb: 1, Ub: maxSliceInfo_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SliceInfoList_r17", err)
	}
	return nil
}

func (ie *SliceInfoList_r17) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*SliceInfo_r17]([]*SliceInfo_r17{}, aper.Constraint{Lb: 1, Ub: maxSliceInfo_r17}, false)
	fn := func() *SliceInfo_r17 {
		return new(SliceInfo_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SliceInfoList_r17", err)
	}
	ie.Value = []SliceInfo_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
