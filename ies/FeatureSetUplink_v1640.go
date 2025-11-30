package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1640 struct {
	TwoHARQ_ACK_Codebook_type1_r16                            *SubSlot_Config_r16                                                               `optional`
	TwoHARQ_ACK_Codebook_type2_r16                            *SubSlot_Config_r16                                                               `optional`
	OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16 *FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16 `optional`
}

func (ie *FeatureSetUplink_v1640) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.TwoHARQ_ACK_Codebook_type1_r16 != nil, ie.TwoHARQ_ACK_Codebook_type2_r16 != nil, ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.TwoHARQ_ACK_Codebook_type1_r16 != nil {
		if err = ie.TwoHARQ_ACK_Codebook_type1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TwoHARQ_ACK_Codebook_type1_r16", err)
		}
	}
	if ie.TwoHARQ_ACK_Codebook_type2_r16 != nil {
		if err = ie.TwoHARQ_ACK_Codebook_type2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TwoHARQ_ACK_Codebook_type2_r16", err)
		}
	}
	if ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16 != nil {
		if err = ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_v1640) Decode(r *aper.AperReader) error {
	var err error
	var TwoHARQ_ACK_Codebook_type1_r16Present bool
	if TwoHARQ_ACK_Codebook_type1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoHARQ_ACK_Codebook_type2_r16Present bool
	if TwoHARQ_ACK_Codebook_type2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16Present bool
	if OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if TwoHARQ_ACK_Codebook_type1_r16Present {
		ie.TwoHARQ_ACK_Codebook_type1_r16 = new(SubSlot_Config_r16)
		if err = ie.TwoHARQ_ACK_Codebook_type1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TwoHARQ_ACK_Codebook_type1_r16", err)
		}
	}
	if TwoHARQ_ACK_Codebook_type2_r16Present {
		ie.TwoHARQ_ACK_Codebook_type2_r16 = new(SubSlot_Config_r16)
		if err = ie.TwoHARQ_ACK_Codebook_type2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TwoHARQ_ACK_Codebook_type2_r16", err)
		}
	}
	if OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16Present {
		ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16 = new(FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16)
		if err = ie.OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode OffsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16", err)
		}
	}
	return nil
}
