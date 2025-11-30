package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16 struct {
	Scs_15kHz_r16 *FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16_scs_15kHz_r16 `optional`
	Scs_30kHz_r16 *FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16_scs_30kHz_r16 `optional`
	Scs_60kHz_r16 *FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16_scs_60kHz_r16 `optional`
}

func (ie *FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs_15kHz_r16 != nil, ie.Scs_30kHz_r16 != nil, ie.Scs_60kHz_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Scs_15kHz_r16 != nil {
		if err = ie.Scs_15kHz_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_15kHz_r16", err)
		}
	}
	if ie.Scs_30kHz_r16 != nil {
		if err = ie.Scs_30kHz_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_30kHz_r16", err)
		}
	}
	if ie.Scs_60kHz_r16 != nil {
		if err = ie.Scs_60kHz_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_60kHz_r16", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16) Decode(r *aper.AperReader) error {
	var err error
	var Scs_15kHz_r16Present bool
	if Scs_15kHz_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_30kHz_r16Present bool
	if Scs_30kHz_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_60kHz_r16Present bool
	if Scs_60kHz_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Scs_15kHz_r16Present {
		ie.Scs_15kHz_r16 = new(FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16_scs_15kHz_r16)
		if err = ie.Scs_15kHz_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_15kHz_r16", err)
		}
	}
	if Scs_30kHz_r16Present {
		ie.Scs_30kHz_r16 = new(FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16_scs_30kHz_r16)
		if err = ie.Scs_30kHz_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_30kHz_r16", err)
		}
	}
	if Scs_60kHz_r16Present {
		ie.Scs_60kHz_r16 = new(FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16_scs_60kHz_r16)
		if err = ie.Scs_60kHz_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_60kHz_r16", err)
		}
	}
	return nil
}
