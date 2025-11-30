package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type TrackingAreaIdentityList_r16 struct {
	Value []TrackingAreaIdentity_r16 `lb:1,ub:8,madatory`
}

func (ie *TrackingAreaIdentityList_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*TrackingAreaIdentity_r16]([]*TrackingAreaIdentity_r16{}, aper.Constraint{Lb: 1, Ub: 8}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode TrackingAreaIdentityList_r16", err)
	}
	return nil
}

func (ie *TrackingAreaIdentityList_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*TrackingAreaIdentity_r16]([]*TrackingAreaIdentity_r16{}, aper.Constraint{Lb: 1, Ub: 8}, false)
	fn := func() *TrackingAreaIdentity_r16 {
		return new(TrackingAreaIdentity_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode TrackingAreaIdentityList_r16", err)
	}
	ie.Value = []TrackingAreaIdentity_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
