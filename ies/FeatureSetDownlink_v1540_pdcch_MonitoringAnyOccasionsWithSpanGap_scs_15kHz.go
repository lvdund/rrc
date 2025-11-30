package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap_scs_15kHz_Enum_set1 aper.Enumerated = 0
	FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap_scs_15kHz_Enum_set2 aper.Enumerated = 1
	FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap_scs_15kHz_Enum_set3 aper.Enumerated = 2
)

type FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap_scs_15kHz struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap_scs_15kHz) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap_scs_15kHz", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap_scs_15kHz) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap_scs_15kHz", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
