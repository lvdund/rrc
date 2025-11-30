package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1610_intraFreqDAPS_r16 struct {
	IntraFreqDiffSCS_DAPS_r16 *FeatureSetDownlink_v1610_intraFreqDAPS_r16_intraFreqDiffSCS_DAPS_r16 `optional`
	IntraFreqAsyncDAPS_r16    *FeatureSetDownlink_v1610_intraFreqDAPS_r16_intraFreqAsyncDAPS_r16    `optional`
}

func (ie *FeatureSetDownlink_v1610_intraFreqDAPS_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.IntraFreqDiffSCS_DAPS_r16 != nil, ie.IntraFreqAsyncDAPS_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.IntraFreqDiffSCS_DAPS_r16 != nil {
		if err = ie.IntraFreqDiffSCS_DAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IntraFreqDiffSCS_DAPS_r16", err)
		}
	}
	if ie.IntraFreqAsyncDAPS_r16 != nil {
		if err = ie.IntraFreqAsyncDAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IntraFreqAsyncDAPS_r16", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_v1610_intraFreqDAPS_r16) Decode(r *aper.AperReader) error {
	var err error
	var IntraFreqDiffSCS_DAPS_r16Present bool
	if IntraFreqDiffSCS_DAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var IntraFreqAsyncDAPS_r16Present bool
	if IntraFreqAsyncDAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if IntraFreqDiffSCS_DAPS_r16Present {
		ie.IntraFreqDiffSCS_DAPS_r16 = new(FeatureSetDownlink_v1610_intraFreqDAPS_r16_intraFreqDiffSCS_DAPS_r16)
		if err = ie.IntraFreqDiffSCS_DAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IntraFreqDiffSCS_DAPS_r16", err)
		}
	}
	if IntraFreqAsyncDAPS_r16Present {
		ie.IntraFreqAsyncDAPS_r16 = new(FeatureSetDownlink_v1610_intraFreqDAPS_r16_intraFreqAsyncDAPS_r16)
		if err = ie.IntraFreqAsyncDAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IntraFreqAsyncDAPS_r16", err)
		}
	}
	return nil
}
