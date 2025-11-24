package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandNR_enhancedType3_HARQ_CodebookFeedback_r17 struct {
	EnhancedType3_HARQ_Codebooks_r17 BandNR_enhancedType3_HARQ_CodebookFeedback_r17_enhancedType3_HARQ_Codebooks_r17 `madatory`
	MaxNumberPUCCH_Transmissions_r17 BandNR_enhancedType3_HARQ_CodebookFeedback_r17_maxNumberPUCCH_Transmissions_r17 `madatory`
}

func (ie *BandNR_enhancedType3_HARQ_CodebookFeedback_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.EnhancedType3_HARQ_Codebooks_r17.Encode(w); err != nil {
		return utils.WrapError("Encode EnhancedType3_HARQ_Codebooks_r17", err)
	}
	if err = ie.MaxNumberPUCCH_Transmissions_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberPUCCH_Transmissions_r17", err)
	}
	return nil
}

func (ie *BandNR_enhancedType3_HARQ_CodebookFeedback_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.EnhancedType3_HARQ_Codebooks_r17.Decode(r); err != nil {
		return utils.WrapError("Decode EnhancedType3_HARQ_Codebooks_r17", err)
	}
	if err = ie.MaxNumberPUCCH_Transmissions_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberPUCCH_Transmissions_r17", err)
	}
	return nil
}
