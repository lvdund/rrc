package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17 struct {
	MaxNumberPRS_Resource_r17                 FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17                  `madatory`
	MaxNumberPRS_ResourceProcessedPerSlot_r17 *FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17 `optional`
}

func (ie *FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MaxNumberPRS_ResourceProcessedPerSlot_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.MaxNumberPRS_Resource_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberPRS_Resource_r17", err)
	}
	if ie.MaxNumberPRS_ResourceProcessedPerSlot_r17 != nil {
		if err = ie.MaxNumberPRS_ResourceProcessedPerSlot_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumberPRS_ResourceProcessedPerSlot_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17) Decode(r *uper.UperReader) error {
	var err error
	var MaxNumberPRS_ResourceProcessedPerSlot_r17Present bool
	if MaxNumberPRS_ResourceProcessedPerSlot_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.MaxNumberPRS_Resource_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberPRS_Resource_r17", err)
	}
	if MaxNumberPRS_ResourceProcessedPerSlot_r17Present {
		ie.MaxNumberPRS_ResourceProcessedPerSlot_r17 = new(FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_ResourceProcessedPerSlot_r17)
		if err = ie.MaxNumberPRS_ResourceProcessedPerSlot_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumberPRS_ResourceProcessedPerSlot_r17", err)
		}
	}
	return nil
}
