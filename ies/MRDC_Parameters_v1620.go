package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_Parameters_v1620 struct {
	MaxUplinkDutyCycle_interBandENDC_TDD_PC2_r16 *MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16 `optional`
	Tdm_restrictionTDD_endc_r16                  *MRDC_Parameters_v1620_tdm_restrictionTDD_endc_r16                  `optional`
	Tdm_restrictionFDD_endc_r16                  *MRDC_Parameters_v1620_tdm_restrictionFDD_endc_r16                  `optional`
	SingleUL_HARQ_offsetTDD_PCell_r16            *MRDC_Parameters_v1620_singleUL_HARQ_offsetTDD_PCell_r16            `optional`
	Tdm_restrictionDualTX_FDD_endc_r16           *MRDC_Parameters_v1620_tdm_restrictionDualTX_FDD_endc_r16           `optional`
}

func (ie *MRDC_Parameters_v1620) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MaxUplinkDutyCycle_interBandENDC_TDD_PC2_r16 != nil, ie.Tdm_restrictionTDD_endc_r16 != nil, ie.Tdm_restrictionFDD_endc_r16 != nil, ie.SingleUL_HARQ_offsetTDD_PCell_r16 != nil, ie.Tdm_restrictionDualTX_FDD_endc_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MaxUplinkDutyCycle_interBandENDC_TDD_PC2_r16 != nil {
		if err = ie.MaxUplinkDutyCycle_interBandENDC_TDD_PC2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxUplinkDutyCycle_interBandENDC_TDD_PC2_r16", err)
		}
	}
	if ie.Tdm_restrictionTDD_endc_r16 != nil {
		if err = ie.Tdm_restrictionTDD_endc_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Tdm_restrictionTDD_endc_r16", err)
		}
	}
	if ie.Tdm_restrictionFDD_endc_r16 != nil {
		if err = ie.Tdm_restrictionFDD_endc_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Tdm_restrictionFDD_endc_r16", err)
		}
	}
	if ie.SingleUL_HARQ_offsetTDD_PCell_r16 != nil {
		if err = ie.SingleUL_HARQ_offsetTDD_PCell_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SingleUL_HARQ_offsetTDD_PCell_r16", err)
		}
	}
	if ie.Tdm_restrictionDualTX_FDD_endc_r16 != nil {
		if err = ie.Tdm_restrictionDualTX_FDD_endc_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Tdm_restrictionDualTX_FDD_endc_r16", err)
		}
	}
	return nil
}

func (ie *MRDC_Parameters_v1620) Decode(r *aper.AperReader) error {
	var err error
	var MaxUplinkDutyCycle_interBandENDC_TDD_PC2_r16Present bool
	if MaxUplinkDutyCycle_interBandENDC_TDD_PC2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Tdm_restrictionTDD_endc_r16Present bool
	if Tdm_restrictionTDD_endc_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Tdm_restrictionFDD_endc_r16Present bool
	if Tdm_restrictionFDD_endc_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SingleUL_HARQ_offsetTDD_PCell_r16Present bool
	if SingleUL_HARQ_offsetTDD_PCell_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Tdm_restrictionDualTX_FDD_endc_r16Present bool
	if Tdm_restrictionDualTX_FDD_endc_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MaxUplinkDutyCycle_interBandENDC_TDD_PC2_r16Present {
		ie.MaxUplinkDutyCycle_interBandENDC_TDD_PC2_r16 = new(MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16)
		if err = ie.MaxUplinkDutyCycle_interBandENDC_TDD_PC2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxUplinkDutyCycle_interBandENDC_TDD_PC2_r16", err)
		}
	}
	if Tdm_restrictionTDD_endc_r16Present {
		ie.Tdm_restrictionTDD_endc_r16 = new(MRDC_Parameters_v1620_tdm_restrictionTDD_endc_r16)
		if err = ie.Tdm_restrictionTDD_endc_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Tdm_restrictionTDD_endc_r16", err)
		}
	}
	if Tdm_restrictionFDD_endc_r16Present {
		ie.Tdm_restrictionFDD_endc_r16 = new(MRDC_Parameters_v1620_tdm_restrictionFDD_endc_r16)
		if err = ie.Tdm_restrictionFDD_endc_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Tdm_restrictionFDD_endc_r16", err)
		}
	}
	if SingleUL_HARQ_offsetTDD_PCell_r16Present {
		ie.SingleUL_HARQ_offsetTDD_PCell_r16 = new(MRDC_Parameters_v1620_singleUL_HARQ_offsetTDD_PCell_r16)
		if err = ie.SingleUL_HARQ_offsetTDD_PCell_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SingleUL_HARQ_offsetTDD_PCell_r16", err)
		}
	}
	if Tdm_restrictionDualTX_FDD_endc_r16Present {
		ie.Tdm_restrictionDualTX_FDD_endc_r16 = new(MRDC_Parameters_v1620_tdm_restrictionDualTX_FDD_endc_r16)
		if err = ie.Tdm_restrictionDualTX_FDD_endc_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Tdm_restrictionDualTX_FDD_endc_r16", err)
		}
	}
	return nil
}
