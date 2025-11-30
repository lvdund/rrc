package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_searchSpaceSharingCA_DL_Enum_supported aper.Enumerated = 0
)

type FeatureSetDownlink_searchSpaceSharingCA_DL struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *FeatureSetDownlink_searchSpaceSharingCA_DL) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_searchSpaceSharingCA_DL", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_searchSpaceSharingCA_DL) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_searchSpaceSharingCA_DL", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
