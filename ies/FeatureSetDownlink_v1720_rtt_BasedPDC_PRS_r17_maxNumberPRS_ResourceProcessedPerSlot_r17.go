package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17 struct {
	Scs_15kHz_r17  *FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_15kHz_r17  `optional`
	Scs_30kHz_r17  *FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_30kHz_r17  `optional`
	Scs_60kHz_r17  *FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_60kHz_r17  `optional`
	Scs_120kHz_r17 *FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_120kHz_r17 `optional`
}

func (ie *FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs_15kHz_r17 != nil, ie.Scs_30kHz_r17 != nil, ie.Scs_60kHz_r17 != nil, ie.Scs_120kHz_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Scs_15kHz_r17 != nil {
		if err = ie.Scs_15kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_15kHz_r17", err)
		}
	}
	if ie.Scs_30kHz_r17 != nil {
		if err = ie.Scs_30kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_30kHz_r17", err)
		}
	}
	if ie.Scs_60kHz_r17 != nil {
		if err = ie.Scs_60kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_60kHz_r17", err)
		}
	}
	if ie.Scs_120kHz_r17 != nil {
		if err = ie.Scs_120kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_120kHz_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17) Decode(r *uper.UperReader) error {
	var err error
	var Scs_15kHz_r17Present bool
	if Scs_15kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_30kHz_r17Present bool
	if Scs_30kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_60kHz_r17Present bool
	if Scs_60kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_120kHz_r17Present bool
	if Scs_120kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Scs_15kHz_r17Present {
		ie.Scs_15kHz_r17 = new(FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_15kHz_r17)
		if err = ie.Scs_15kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_15kHz_r17", err)
		}
	}
	if Scs_30kHz_r17Present {
		ie.Scs_30kHz_r17 = new(FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_30kHz_r17)
		if err = ie.Scs_30kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_30kHz_r17", err)
		}
	}
	if Scs_60kHz_r17Present {
		ie.Scs_60kHz_r17 = new(FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_60kHz_r17)
		if err = ie.Scs_60kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_60kHz_r17", err)
		}
	}
	if Scs_120kHz_r17Present {
		ie.Scs_120kHz_r17 = new(FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17_scs_120kHz_r17)
		if err = ie.Scs_120kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_120kHz_r17", err)
		}
	}
	return nil
}
