package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1610_pdcch_Monitoring_r16 struct {
	Pdsch_ProcessingType1_r16 *Pdsch_ProcessingType_r16 `optional`
	Pdsch_ProcessingType2_r16 *Pdsch_ProcessingType_r16 `optional`
}

func (ie *FeatureSetDownlink_v1610_pdcch_Monitoring_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Pdsch_ProcessingType1_r16 != nil, ie.Pdsch_ProcessingType2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pdsch_ProcessingType1_r16 != nil {
		if err = ie.Pdsch_ProcessingType1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_ProcessingType1_r16", err)
		}
	}
	if ie.Pdsch_ProcessingType2_r16 != nil {
		if err = ie.Pdsch_ProcessingType2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_ProcessingType2_r16", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_v1610_pdcch_Monitoring_r16) Decode(r *uper.UperReader) error {
	var err error
	var Pdsch_ProcessingType1_r16Present bool
	if Pdsch_ProcessingType1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_ProcessingType2_r16Present bool
	if Pdsch_ProcessingType2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Pdsch_ProcessingType1_r16Present {
		ie.Pdsch_ProcessingType1_r16 = new(Pdsch_ProcessingType_r16)
		if err = ie.Pdsch_ProcessingType1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_ProcessingType1_r16", err)
		}
	}
	if Pdsch_ProcessingType2_r16Present {
		ie.Pdsch_ProcessingType2_r16 = new(Pdsch_ProcessingType_r16)
		if err = ie.Pdsch_ProcessingType2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_ProcessingType2_r16", err)
		}
	}
	return nil
}
