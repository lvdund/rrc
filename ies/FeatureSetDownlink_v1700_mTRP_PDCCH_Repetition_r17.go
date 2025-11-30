package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17 struct {
	NumBD_twoPDCCH_r17 int64                                                                 `lb:2,ub:3,madatory`
	MaxNumOverlaps_r17 FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17_maxNumOverlaps_r17 `madatory`
}

func (ie *FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.NumBD_twoPDCCH_r17, &aper.Constraint{Lb: 2, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteInteger NumBD_twoPDCCH_r17", err)
	}
	if err = ie.MaxNumOverlaps_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumOverlaps_r17", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_NumBD_twoPDCCH_r17 int64
	if tmp_int_NumBD_twoPDCCH_r17, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadInteger NumBD_twoPDCCH_r17", err)
	}
	ie.NumBD_twoPDCCH_r17 = tmp_int_NumBD_twoPDCCH_r17
	if err = ie.MaxNumOverlaps_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumOverlaps_r17", err)
	}
	return nil
}
