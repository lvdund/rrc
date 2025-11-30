package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_BlindDetectionMixed_r17_pdcch_BlindDetectionCG_UE_Mixed_r17 struct {
	Pdcch_BlindDetectionMCG_UE_Mixed_v17 PDCCH_BlindDetectionCG_UE_Mixed_r17 `madatory`
	Pdcch_BlindDetectionSCG_UE_Mixed_v17 PDCCH_BlindDetectionCG_UE_Mixed_r17 `madatory`
}

func (ie *PDCCH_BlindDetectionMixed_r17_pdcch_BlindDetectionCG_UE_Mixed_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Pdcch_BlindDetectionMCG_UE_Mixed_v17.Encode(w); err != nil {
		return utils.WrapError("Encode Pdcch_BlindDetectionMCG_UE_Mixed_v17", err)
	}
	if err = ie.Pdcch_BlindDetectionSCG_UE_Mixed_v17.Encode(w); err != nil {
		return utils.WrapError("Encode Pdcch_BlindDetectionSCG_UE_Mixed_v17", err)
	}
	return nil
}

func (ie *PDCCH_BlindDetectionMixed_r17_pdcch_BlindDetectionCG_UE_Mixed_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Pdcch_BlindDetectionMCG_UE_Mixed_v17.Decode(r); err != nil {
		return utils.WrapError("Decode Pdcch_BlindDetectionMCG_UE_Mixed_v17", err)
	}
	if err = ie.Pdcch_BlindDetectionSCG_UE_Mixed_v17.Decode(r); err != nil {
		return utils.WrapError("Decode Pdcch_BlindDetectionSCG_UE_Mixed_v17", err)
	}
	return nil
}
