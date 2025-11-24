package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportConfig_reportQuantity_Choice_nothing uint64 = iota
	CSI_ReportConfig_reportQuantity_Choice_None
	CSI_ReportConfig_reportQuantity_Choice_Cri_RI_PMI_CQI
	CSI_ReportConfig_reportQuantity_Choice_Cri_RI_i1
	CSI_ReportConfig_reportQuantity_Choice_Cri_RI_i1_CQI
	CSI_ReportConfig_reportQuantity_Choice_Cri_RI_CQI
	CSI_ReportConfig_reportQuantity_Choice_Cri_RSRP
	CSI_ReportConfig_reportQuantity_Choice_Ssb_Index_RSRP
	CSI_ReportConfig_reportQuantity_Choice_Cri_RI_LI_PMI_CQI
)

type CSI_ReportConfig_reportQuantity struct {
	Choice            uint64
	None              uper.NULL `madatory`
	Cri_RI_PMI_CQI    uper.NULL `madatory`
	Cri_RI_i1         uper.NULL `madatory`
	Cri_RI_i1_CQI     *CSI_ReportConfig_reportQuantity_cri_RI_i1_CQI
	Cri_RI_CQI        uper.NULL `madatory`
	Cri_RSRP          uper.NULL `madatory`
	Ssb_Index_RSRP    uper.NULL `madatory`
	Cri_RI_LI_PMI_CQI uper.NULL `madatory`
}

func (ie *CSI_ReportConfig_reportQuantity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportConfig_reportQuantity_Choice_None:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode None", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_Cri_RI_PMI_CQI:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Cri_RI_PMI_CQI", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_Cri_RI_i1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Cri_RI_i1", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_Cri_RI_i1_CQI:
		if err = ie.Cri_RI_i1_CQI.Encode(w); err != nil {
			err = utils.WrapError("Encode Cri_RI_i1_CQI", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_Cri_RI_CQI:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Cri_RI_CQI", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_Cri_RSRP:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Cri_RSRP", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_Ssb_Index_RSRP:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Ssb_Index_RSRP", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_Cri_RI_LI_PMI_CQI:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Cri_RI_LI_PMI_CQI", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_ReportConfig_reportQuantity) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportConfig_reportQuantity_Choice_None:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode None", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_Cri_RI_PMI_CQI:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Cri_RI_PMI_CQI", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_Cri_RI_i1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Cri_RI_i1", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_Cri_RI_i1_CQI:
		ie.Cri_RI_i1_CQI = new(CSI_ReportConfig_reportQuantity_cri_RI_i1_CQI)
		if err = ie.Cri_RI_i1_CQI.Decode(r); err != nil {
			return utils.WrapError("Decode Cri_RI_i1_CQI", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_Cri_RI_CQI:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Cri_RI_CQI", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_Cri_RSRP:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Cri_RSRP", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_Ssb_Index_RSRP:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Ssb_Index_RSRP", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_Cri_RI_LI_PMI_CQI:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Cri_RI_LI_PMI_CQI", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
