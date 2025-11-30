package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_singleDCI_SDM_scheme_Parameters_r16 struct {
	SupportNewDMRS_Port_r16   *MIMO_ParametersPerBand_singleDCI_SDM_scheme_Parameters_r16_supportNewDMRS_Port_r16   `optional`
	SupportTwoPortDL_PTRS_r16 *MIMO_ParametersPerBand_singleDCI_SDM_scheme_Parameters_r16_supportTwoPortDL_PTRS_r16 `optional`
}

func (ie *MIMO_ParametersPerBand_singleDCI_SDM_scheme_Parameters_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SupportNewDMRS_Port_r16 != nil, ie.SupportTwoPortDL_PTRS_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SupportNewDMRS_Port_r16 != nil {
		if err = ie.SupportNewDMRS_Port_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SupportNewDMRS_Port_r16", err)
		}
	}
	if ie.SupportTwoPortDL_PTRS_r16 != nil {
		if err = ie.SupportTwoPortDL_PTRS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SupportTwoPortDL_PTRS_r16", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_singleDCI_SDM_scheme_Parameters_r16) Decode(r *aper.AperReader) error {
	var err error
	var SupportNewDMRS_Port_r16Present bool
	if SupportNewDMRS_Port_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportTwoPortDL_PTRS_r16Present bool
	if SupportTwoPortDL_PTRS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SupportNewDMRS_Port_r16Present {
		ie.SupportNewDMRS_Port_r16 = new(MIMO_ParametersPerBand_singleDCI_SDM_scheme_Parameters_r16_supportNewDMRS_Port_r16)
		if err = ie.SupportNewDMRS_Port_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SupportNewDMRS_Port_r16", err)
		}
	}
	if SupportTwoPortDL_PTRS_r16Present {
		ie.SupportTwoPortDL_PTRS_r16 = new(MIMO_ParametersPerBand_singleDCI_SDM_scheme_Parameters_r16_supportTwoPortDL_PTRS_r16)
		if err = ie.SupportTwoPortDL_PTRS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SupportTwoPortDL_PTRS_r16", err)
		}
	}
	return nil
}
