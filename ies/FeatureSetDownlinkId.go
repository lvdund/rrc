package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlinkId struct {
	Value uint64 `lb:0,ub:maxDownlinkFeatureSets,madatory`
}

func (ie *FeatureSetDownlinkId) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: maxDownlinkFeatureSets}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlinkId", err)
	}
	return nil
}

func (ie *FeatureSetDownlinkId) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxDownlinkFeatureSets}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlinkId", err)
	}
	ie.Value = uint64(v)
	return nil
}
