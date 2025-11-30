package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1610_intraFreqDAPS_UL_r16 struct {
	Dummy                     *FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy                     `optional`
	IntraFreqTwoTAGs_DAPS_r16 *FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_intraFreqTwoTAGs_DAPS_r16 `optional`
	Dummy1                    *FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy1                    `optional`
	Dummy2                    *FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy2                    `optional`
	Dummy3                    *FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy3                    `optional`
}

func (ie *FeatureSetUplink_v1610_intraFreqDAPS_UL_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Dummy != nil, ie.IntraFreqTwoTAGs_DAPS_r16 != nil, ie.Dummy1 != nil, ie.Dummy2 != nil, ie.Dummy3 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Dummy != nil {
		if err = ie.Dummy.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy", err)
		}
	}
	if ie.IntraFreqTwoTAGs_DAPS_r16 != nil {
		if err = ie.IntraFreqTwoTAGs_DAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IntraFreqTwoTAGs_DAPS_r16", err)
		}
	}
	if ie.Dummy1 != nil {
		if err = ie.Dummy1.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy1", err)
		}
	}
	if ie.Dummy2 != nil {
		if err = ie.Dummy2.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy2", err)
		}
	}
	if ie.Dummy3 != nil {
		if err = ie.Dummy3.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy3", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_v1610_intraFreqDAPS_UL_r16) Decode(r *aper.AperReader) error {
	var err error
	var DummyPresent bool
	if DummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var IntraFreqTwoTAGs_DAPS_r16Present bool
	if IntraFreqTwoTAGs_DAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy1Present bool
	if Dummy1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy2Present bool
	if Dummy2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy3Present bool
	if Dummy3Present, err = r.ReadBool(); err != nil {
		return err
	}
	if DummyPresent {
		ie.Dummy = new(FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy)
		if err = ie.Dummy.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy", err)
		}
	}
	if IntraFreqTwoTAGs_DAPS_r16Present {
		ie.IntraFreqTwoTAGs_DAPS_r16 = new(FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_intraFreqTwoTAGs_DAPS_r16)
		if err = ie.IntraFreqTwoTAGs_DAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IntraFreqTwoTAGs_DAPS_r16", err)
		}
	}
	if Dummy1Present {
		ie.Dummy1 = new(FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy1)
		if err = ie.Dummy1.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy1", err)
		}
	}
	if Dummy2Present {
		ie.Dummy2 = new(FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy2)
		if err = ie.Dummy2.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy2", err)
		}
	}
	if Dummy3Present {
		ie.Dummy3 = new(FeatureSetUplink_v1610_intraFreqDAPS_UL_r16_dummy3)
		if err = ie.Dummy3.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy3", err)
		}
	}
	return nil
}
