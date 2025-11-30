package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type TrackingAreaIdentity_r16 struct {
	Plmn_Identity_r16    PLMN_Identity    `madatory`
	TrackingAreaCode_r16 TrackingAreaCode `madatory`
}

func (ie *TrackingAreaIdentity_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Plmn_Identity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Plmn_Identity_r16", err)
	}
	if err = ie.TrackingAreaCode_r16.Encode(w); err != nil {
		return utils.WrapError("Encode TrackingAreaCode_r16", err)
	}
	return nil
}

func (ie *TrackingAreaIdentity_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Plmn_Identity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Plmn_Identity_r16", err)
	}
	if err = ie.TrackingAreaCode_r16.Decode(r); err != nil {
		return utils.WrapError("Decode TrackingAreaCode_r16", err)
	}
	return nil
}
