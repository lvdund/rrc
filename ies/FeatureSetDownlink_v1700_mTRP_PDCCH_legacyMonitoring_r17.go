package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1700_mTRP_PDCCH_legacyMonitoring_r17 struct {
	Scs_15kHz_r17 *PDCCH_RepetitionParameters_r17 `optional`
	Scs_30kHz_r17 *PDCCH_RepetitionParameters_r17 `optional`
}

func (ie *FeatureSetDownlink_v1700_mTRP_PDCCH_legacyMonitoring_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs_15kHz_r17 != nil, ie.Scs_30kHz_r17 != nil}
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
	return nil
}

func (ie *FeatureSetDownlink_v1700_mTRP_PDCCH_legacyMonitoring_r17) Decode(r *aper.AperReader) error {
	var err error
	var Scs_15kHz_r17Present bool
	if Scs_15kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_30kHz_r17Present bool
	if Scs_30kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Scs_15kHz_r17Present {
		ie.Scs_15kHz_r17 = new(PDCCH_RepetitionParameters_r17)
		if err = ie.Scs_15kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_15kHz_r17", err)
		}
	}
	if Scs_30kHz_r17Present {
		ie.Scs_30kHz_r17 = new(PDCCH_RepetitionParameters_r17)
		if err = ie.Scs_30kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_30kHz_r17", err)
		}
	}
	return nil
}
