package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationDL_r16 struct {
	SupportPDCCH_ToPDSCH_r16    *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationDL_r16_supportPDCCH_ToPDSCH_r16    `optional`
	SupportPDSCH_ToHARQ_ACK_r16 *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationDL_r16_supportPDSCH_ToHARQ_ACK_r16 `optional`
}

func (ie *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationDL_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SupportPDCCH_ToPDSCH_r16 != nil, ie.SupportPDSCH_ToHARQ_ACK_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SupportPDCCH_ToPDSCH_r16 != nil {
		if err = ie.SupportPDCCH_ToPDSCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SupportPDCCH_ToPDSCH_r16", err)
		}
	}
	if ie.SupportPDSCH_ToHARQ_ACK_r16 != nil {
		if err = ie.SupportPDSCH_ToHARQ_ACK_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SupportPDSCH_ToHARQ_ACK_r16", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationDL_r16) Decode(r *aper.AperReader) error {
	var err error
	var SupportPDCCH_ToPDSCH_r16Present bool
	if SupportPDCCH_ToPDSCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportPDSCH_ToHARQ_ACK_r16Present bool
	if SupportPDSCH_ToHARQ_ACK_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SupportPDCCH_ToPDSCH_r16Present {
		ie.SupportPDCCH_ToPDSCH_r16 = new(MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationDL_r16_supportPDCCH_ToPDSCH_r16)
		if err = ie.SupportPDCCH_ToPDSCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SupportPDCCH_ToPDSCH_r16", err)
		}
	}
	if SupportPDSCH_ToHARQ_ACK_r16Present {
		ie.SupportPDSCH_ToHARQ_ACK_r16 = new(MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationDL_r16_supportPDSCH_ToHARQ_ACK_r16)
		if err = ie.SupportPDSCH_ToHARQ_ACK_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SupportPDSCH_ToHARQ_ACK_r16", err)
		}
	}
	return nil
}
