package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplinkPerCC_v1700 struct {
	SupportedMinBandwidthUL_r17    *SupportedBandwidth_v1700                                   `optional`
	MTRP_PUSCH_RepetitionTypeB_r17 *FeatureSetUplinkPerCC_v1700_mTRP_PUSCH_RepetitionTypeB_r17 `optional`
	MTRP_PUSCH_TypeB_CB_r17        *FeatureSetUplinkPerCC_v1700_mTRP_PUSCH_TypeB_CB_r17        `optional`
	SupportedBandwidthUL_v1710     *SupportedBandwidth_v1700                                   `optional`
}

func (ie *FeatureSetUplinkPerCC_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SupportedMinBandwidthUL_r17 != nil, ie.MTRP_PUSCH_RepetitionTypeB_r17 != nil, ie.MTRP_PUSCH_TypeB_CB_r17 != nil, ie.SupportedBandwidthUL_v1710 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SupportedMinBandwidthUL_r17 != nil {
		if err = ie.SupportedMinBandwidthUL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedMinBandwidthUL_r17", err)
		}
	}
	if ie.MTRP_PUSCH_RepetitionTypeB_r17 != nil {
		if err = ie.MTRP_PUSCH_RepetitionTypeB_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MTRP_PUSCH_RepetitionTypeB_r17", err)
		}
	}
	if ie.MTRP_PUSCH_TypeB_CB_r17 != nil {
		if err = ie.MTRP_PUSCH_TypeB_CB_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MTRP_PUSCH_TypeB_CB_r17", err)
		}
	}
	if ie.SupportedBandwidthUL_v1710 != nil {
		if err = ie.SupportedBandwidthUL_v1710.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedBandwidthUL_v1710", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplinkPerCC_v1700) Decode(r *aper.AperReader) error {
	var err error
	var SupportedMinBandwidthUL_r17Present bool
	if SupportedMinBandwidthUL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MTRP_PUSCH_RepetitionTypeB_r17Present bool
	if MTRP_PUSCH_RepetitionTypeB_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MTRP_PUSCH_TypeB_CB_r17Present bool
	if MTRP_PUSCH_TypeB_CB_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedBandwidthUL_v1710Present bool
	if SupportedBandwidthUL_v1710Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SupportedMinBandwidthUL_r17Present {
		ie.SupportedMinBandwidthUL_r17 = new(SupportedBandwidth_v1700)
		if err = ie.SupportedMinBandwidthUL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedMinBandwidthUL_r17", err)
		}
	}
	if MTRP_PUSCH_RepetitionTypeB_r17Present {
		ie.MTRP_PUSCH_RepetitionTypeB_r17 = new(FeatureSetUplinkPerCC_v1700_mTRP_PUSCH_RepetitionTypeB_r17)
		if err = ie.MTRP_PUSCH_RepetitionTypeB_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MTRP_PUSCH_RepetitionTypeB_r17", err)
		}
	}
	if MTRP_PUSCH_TypeB_CB_r17Present {
		ie.MTRP_PUSCH_TypeB_CB_r17 = new(FeatureSetUplinkPerCC_v1700_mTRP_PUSCH_TypeB_CB_r17)
		if err = ie.MTRP_PUSCH_TypeB_CB_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MTRP_PUSCH_TypeB_CB_r17", err)
		}
	}
	if SupportedBandwidthUL_v1710Present {
		ie.SupportedBandwidthUL_v1710 = new(SupportedBandwidth_v1700)
		if err = ie.SupportedBandwidthUL_v1710.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedBandwidthUL_v1710", err)
		}
	}
	return nil
}
