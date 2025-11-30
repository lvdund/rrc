package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16 struct {
	MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16 *MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16 `optional`
	MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16 *MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC2_r16 `optional`
}

func (ie *MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16 != nil, ie.MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16 != nil {
		if err = ie.MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16", err)
		}
	}
	if ie.MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16 != nil {
		if err = ie.MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16", err)
		}
	}
	return nil
}

func (ie *MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16) Decode(r *aper.AperReader) error {
	var err error
	var MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16Present bool
	if MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16Present bool
	if MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16Present {
		ie.MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16 = new(MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16)
		if err = ie.MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxUplinkDutyCycle_FDD_TDD_EN_DC1_r16", err)
		}
	}
	if MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16Present {
		ie.MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16 = new(MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC2_r16)
		if err = ie.MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxUplinkDutyCycle_FDD_TDD_EN_DC2_r16", err)
		}
	}
	return nil
}
