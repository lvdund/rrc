package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_pdcch_MonitoringAnyOccasions_Enum_withoutDCI_Gap aper.Enumerated = 0
	FeatureSetDownlink_pdcch_MonitoringAnyOccasions_Enum_withDCI_Gap    aper.Enumerated = 1
)

type FeatureSetDownlink_pdcch_MonitoringAnyOccasions struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *FeatureSetDownlink_pdcch_MonitoringAnyOccasions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_pdcch_MonitoringAnyOccasions", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_pdcch_MonitoringAnyOccasions) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_pdcch_MonitoringAnyOccasions", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
