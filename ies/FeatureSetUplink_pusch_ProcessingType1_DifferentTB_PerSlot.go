package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot struct {
	Scs_15kHz  *FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot_scs_15kHz  `optional`
	Scs_30kHz  *FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot_scs_30kHz  `optional`
	Scs_60kHz  *FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot_scs_60kHz  `optional`
	Scs_120kHz *FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot_scs_120kHz `optional`
}

func (ie *FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs_15kHz != nil, ie.Scs_30kHz != nil, ie.Scs_60kHz != nil, ie.Scs_120kHz != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Scs_15kHz != nil {
		if err = ie.Scs_15kHz.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_15kHz", err)
		}
	}
	if ie.Scs_30kHz != nil {
		if err = ie.Scs_30kHz.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_30kHz", err)
		}
	}
	if ie.Scs_60kHz != nil {
		if err = ie.Scs_60kHz.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_60kHz", err)
		}
	}
	if ie.Scs_120kHz != nil {
		if err = ie.Scs_120kHz.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_120kHz", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot) Decode(r *uper.UperReader) error {
	var err error
	var Scs_15kHzPresent bool
	if Scs_15kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_30kHzPresent bool
	if Scs_30kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_60kHzPresent bool
	if Scs_60kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_120kHzPresent bool
	if Scs_120kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Scs_15kHzPresent {
		ie.Scs_15kHz = new(FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot_scs_15kHz)
		if err = ie.Scs_15kHz.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_15kHz", err)
		}
	}
	if Scs_30kHzPresent {
		ie.Scs_30kHz = new(FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot_scs_30kHz)
		if err = ie.Scs_30kHz.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_30kHz", err)
		}
	}
	if Scs_60kHzPresent {
		ie.Scs_60kHz = new(FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot_scs_60kHz)
		if err = ie.Scs_60kHz.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_60kHz", err)
		}
	}
	if Scs_120kHzPresent {
		ie.Scs_120kHz = new(FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot_scs_120kHz)
		if err = ie.Scs_120kHz.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_120kHz", err)
		}
	}
	return nil
}
