package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17_Enum_n1  aper.Enumerated = 0
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17_Enum_n2  aper.Enumerated = 1
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17_Enum_n4  aper.Enumerated = 2
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17_Enum_n8  aper.Enumerated = 3
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17_Enum_n16 aper.Enumerated = 4
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17_Enum_n32 aper.Enumerated = 5
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17_Enum_n64 aper.Enumerated = 6
)

type FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17 struct {
	Value aper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
