package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_Parameters_v1630 struct {
	MaxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16 *MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16 `optional`
	InterBandMRDC_WithOverlapDL_Bands_r16            *MRDC_Parameters_v1630_interBandMRDC_WithOverlapDL_Bands_r16            `optional`
}

func (ie *MRDC_Parameters_v1630) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MaxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16 != nil, ie.InterBandMRDC_WithOverlapDL_Bands_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MaxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16 != nil {
		if err = ie.MaxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16", err)
		}
	}
	if ie.InterBandMRDC_WithOverlapDL_Bands_r16 != nil {
		if err = ie.InterBandMRDC_WithOverlapDL_Bands_r16.Encode(w); err != nil {
			return utils.WrapError("Encode InterBandMRDC_WithOverlapDL_Bands_r16", err)
		}
	}
	return nil
}

func (ie *MRDC_Parameters_v1630) Decode(r *aper.AperReader) error {
	var err error
	var MaxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16Present bool
	if MaxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InterBandMRDC_WithOverlapDL_Bands_r16Present bool
	if InterBandMRDC_WithOverlapDL_Bands_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MaxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16Present {
		ie.MaxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16 = new(MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16)
		if err = ie.MaxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16", err)
		}
	}
	if InterBandMRDC_WithOverlapDL_Bands_r16Present {
		ie.InterBandMRDC_WithOverlapDL_Bands_r16 = new(MRDC_Parameters_v1630_interBandMRDC_WithOverlapDL_Bands_r16)
		if err = ie.InterBandMRDC_WithOverlapDL_Bands_r16.Decode(r); err != nil {
			return utils.WrapError("Decode InterBandMRDC_WithOverlapDL_Bands_r16", err)
		}
	}
	return nil
}
