package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_BlindDetectionMixed1_r17 struct {
	Pdcch_BlindDetectionCA_Mixed1_r17    *PDCCH_BlindDetectionCA_Mixed1_r17                                   `optional`
	Pdcch_BlindDetectionCG_UE_Mixed1_r17 *PDCCH_BlindDetectionMixed1_r17_pdcch_BlindDetectionCG_UE_Mixed1_r17 `optional`
}

func (ie *PDCCH_BlindDetectionMixed1_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Pdcch_BlindDetectionCA_Mixed1_r17 != nil, ie.Pdcch_BlindDetectionCG_UE_Mixed1_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pdcch_BlindDetectionCA_Mixed1_r17 != nil {
		if err = ie.Pdcch_BlindDetectionCA_Mixed1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_BlindDetectionCA_Mixed1_r17", err)
		}
	}
	if ie.Pdcch_BlindDetectionCG_UE_Mixed1_r17 != nil {
		if err = ie.Pdcch_BlindDetectionCG_UE_Mixed1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_BlindDetectionCG_UE_Mixed1_r17", err)
		}
	}
	return nil
}

func (ie *PDCCH_BlindDetectionMixed1_r17) Decode(r *aper.AperReader) error {
	var err error
	var Pdcch_BlindDetectionCA_Mixed1_r17Present bool
	if Pdcch_BlindDetectionCA_Mixed1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_BlindDetectionCG_UE_Mixed1_r17Present bool
	if Pdcch_BlindDetectionCG_UE_Mixed1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Pdcch_BlindDetectionCA_Mixed1_r17Present {
		ie.Pdcch_BlindDetectionCA_Mixed1_r17 = new(PDCCH_BlindDetectionCA_Mixed1_r17)
		if err = ie.Pdcch_BlindDetectionCA_Mixed1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_BlindDetectionCA_Mixed1_r17", err)
		}
	}
	if Pdcch_BlindDetectionCG_UE_Mixed1_r17Present {
		ie.Pdcch_BlindDetectionCG_UE_Mixed1_r17 = new(PDCCH_BlindDetectionMixed1_r17_pdcch_BlindDetectionCG_UE_Mixed1_r17)
		if err = ie.Pdcch_BlindDetectionCG_UE_Mixed1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_BlindDetectionCG_UE_Mixed1_r17", err)
		}
	}
	return nil
}
