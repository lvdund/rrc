package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_30kHz_r17_Enum_n1  aper.Enumerated = 0
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_30kHz_r17_Enum_n2  aper.Enumerated = 1
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_30kHz_r17_Enum_n4  aper.Enumerated = 2
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_30kHz_r17_Enum_n6  aper.Enumerated = 3
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_30kHz_r17_Enum_n8  aper.Enumerated = 4
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_30kHz_r17_Enum_n12 aper.Enumerated = 5
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_30kHz_r17_Enum_n16 aper.Enumerated = 6
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_30kHz_r17_Enum_n24 aper.Enumerated = 7
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_30kHz_r17_Enum_n32 aper.Enumerated = 8
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_30kHz_r17_Enum_n48 aper.Enumerated = 9
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_30kHz_r17_Enum_n64 aper.Enumerated = 10
)

type FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_30kHz_r17 struct {
	Value aper.Enumerated `lb:0,ub:10,madatory`
}

func (ie *FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_30kHz_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_30kHz_r17", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_30kHz_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_30kHz_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
