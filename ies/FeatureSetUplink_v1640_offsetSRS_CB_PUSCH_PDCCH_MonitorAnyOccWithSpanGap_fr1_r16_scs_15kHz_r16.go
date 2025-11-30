package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16_scs_15kHz_r16_Enum_set1 aper.Enumerated = 0
	FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16_scs_15kHz_r16_Enum_set2 aper.Enumerated = 1
	FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16_scs_15kHz_r16_Enum_set3 aper.Enumerated = 2
)

type FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16_scs_15kHz_r16 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16_scs_15kHz_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16_scs_15kHz_r16", err)
	}
	return nil
}

func (ie *FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16_scs_15kHz_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16_scs_15kHz_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
