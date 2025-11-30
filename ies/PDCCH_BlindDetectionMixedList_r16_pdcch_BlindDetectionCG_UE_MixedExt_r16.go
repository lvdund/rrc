package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCG_UE_MixedExt_r16 struct {
	Pdcch_BlindDetectionMCG_UE_Mixed_v16a0 PDCCH_BlindDetectionCG_UE_MixedExt_r16 `madatory`
	Pdcch_BlindDetectionSCG_UE_Mixed_v16a0 PDCCH_BlindDetectionCG_UE_MixedExt_r16 `madatory`
}

func (ie *PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCG_UE_MixedExt_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Pdcch_BlindDetectionMCG_UE_Mixed_v16a0.Encode(w); err != nil {
		return utils.WrapError("Encode Pdcch_BlindDetectionMCG_UE_Mixed_v16a0", err)
	}
	if err = ie.Pdcch_BlindDetectionSCG_UE_Mixed_v16a0.Encode(w); err != nil {
		return utils.WrapError("Encode Pdcch_BlindDetectionSCG_UE_Mixed_v16a0", err)
	}
	return nil
}

func (ie *PDCCH_BlindDetectionMixedList_r16_pdcch_BlindDetectionCG_UE_MixedExt_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Pdcch_BlindDetectionMCG_UE_Mixed_v16a0.Decode(r); err != nil {
		return utils.WrapError("Decode Pdcch_BlindDetectionMCG_UE_Mixed_v16a0", err)
	}
	if err = ie.Pdcch_BlindDetectionSCG_UE_Mixed_v16a0.Decode(r); err != nil {
		return utils.WrapError("Decode Pdcch_BlindDetectionSCG_UE_Mixed_v16a0", err)
	}
	return nil
}
