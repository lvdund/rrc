package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetUplink_twoPUCCH_Group_Enum_supported aper.Enumerated = 0
)

type FeatureSetUplink_twoPUCCH_Group struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *FeatureSetUplink_twoPUCCH_Group) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode FeatureSetUplink_twoPUCCH_Group", err)
	}
	return nil
}

func (ie *FeatureSetUplink_twoPUCCH_Group) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode FeatureSetUplink_twoPUCCH_Group", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
