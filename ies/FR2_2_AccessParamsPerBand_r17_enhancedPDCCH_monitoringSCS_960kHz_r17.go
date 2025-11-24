package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17 struct {
	Pdcch_monitoring4_1_r17 *FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17_pdcch_monitoring4_1_r17 `optional`
	Pdcch_monitoring4_2_r17 *FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17_pdcch_monitoring4_2_r17 `optional`
	Pdcch_monitoring8_4_r17 *FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17_pdcch_monitoring8_4_r17 `optional`
}

func (ie *FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Pdcch_monitoring4_1_r17 != nil, ie.Pdcch_monitoring4_2_r17 != nil, ie.Pdcch_monitoring8_4_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pdcch_monitoring4_1_r17 != nil {
		if err = ie.Pdcch_monitoring4_1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_monitoring4_1_r17", err)
		}
	}
	if ie.Pdcch_monitoring4_2_r17 != nil {
		if err = ie.Pdcch_monitoring4_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_monitoring4_2_r17", err)
		}
	}
	if ie.Pdcch_monitoring8_4_r17 != nil {
		if err = ie.Pdcch_monitoring8_4_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_monitoring8_4_r17", err)
		}
	}
	return nil
}

func (ie *FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17) Decode(r *uper.UperReader) error {
	var err error
	var Pdcch_monitoring4_1_r17Present bool
	if Pdcch_monitoring4_1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_monitoring4_2_r17Present bool
	if Pdcch_monitoring4_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_monitoring8_4_r17Present bool
	if Pdcch_monitoring8_4_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Pdcch_monitoring4_1_r17Present {
		ie.Pdcch_monitoring4_1_r17 = new(FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17_pdcch_monitoring4_1_r17)
		if err = ie.Pdcch_monitoring4_1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_monitoring4_1_r17", err)
		}
	}
	if Pdcch_monitoring4_2_r17Present {
		ie.Pdcch_monitoring4_2_r17 = new(FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17_pdcch_monitoring4_2_r17)
		if err = ie.Pdcch_monitoring4_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_monitoring4_2_r17", err)
		}
	}
	if Pdcch_monitoring8_4_r17Present {
		ie.Pdcch_monitoring8_4_r17 = new(FR2_2_AccessParamsPerBand_r17_enhancedPDCCH_monitoringSCS_960kHz_r17_pdcch_monitoring8_4_r17)
		if err = ie.Pdcch_monitoring8_4_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_monitoring8_4_r17", err)
		}
	}
	return nil
}
