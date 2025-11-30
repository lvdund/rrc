package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17 struct {
	Pusch_PreparationLowPriority_r17 FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_pusch_PreparationLowPriority_r17 `madatory`
	AdditionalCancellationTime_r17   *FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17  `optional`
	MaxNumberCarriers_r17            int64                                                                                                      `lb:1,ub:16,madatory`
}

func (ie *FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.AdditionalCancellationTime_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Pusch_PreparationLowPriority_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Pusch_PreparationLowPriority_r17", err)
	}
	if ie.AdditionalCancellationTime_r17 != nil {
		if err = ie.AdditionalCancellationTime_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AdditionalCancellationTime_r17", err)
		}
	}
	if err = w.WriteInteger(ie.MaxNumberCarriers_r17, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberCarriers_r17", err)
	}
	return nil
}

func (ie *FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17) Decode(r *aper.AperReader) error {
	var err error
	var AdditionalCancellationTime_r17Present bool
	if AdditionalCancellationTime_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Pusch_PreparationLowPriority_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Pusch_PreparationLowPriority_r17", err)
	}
	if AdditionalCancellationTime_r17Present {
		ie.AdditionalCancellationTime_r17 = new(FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17)
		if err = ie.AdditionalCancellationTime_r17.Decode(r); err != nil {
			return utils.WrapError("Decode AdditionalCancellationTime_r17", err)
		}
	}
	var tmp_int_MaxNumberCarriers_r17 int64
	if tmp_int_MaxNumberCarriers_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberCarriers_r17", err)
	}
	ie.MaxNumberCarriers_r17 = tmp_int_MaxNumberCarriers_r17
	return nil
}
