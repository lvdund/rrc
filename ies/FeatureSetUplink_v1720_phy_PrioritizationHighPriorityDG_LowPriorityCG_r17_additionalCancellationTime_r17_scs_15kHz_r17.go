package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_15kHz_r17_Enum_sym0 aper.Enumerated = 0
	FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_15kHz_r17_Enum_sym1 aper.Enumerated = 1
	FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_15kHz_r17_Enum_sym2 aper.Enumerated = 2
)

type FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_15kHz_r17 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_15kHz_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_15kHz_r17", err)
	}
	return nil
}

func (ie *FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_15kHz_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_15kHz_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
