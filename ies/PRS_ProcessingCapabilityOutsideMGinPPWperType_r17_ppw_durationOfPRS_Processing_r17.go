package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_Choice_nothing uint64 = iota
	PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_Choice_Ppw_durationOfPRS_Processing1_r17
	PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_Choice_Ppw_durationOfPRS_Processing2_r17
)

type PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17 struct {
	Choice                            uint64
	Ppw_durationOfPRS_Processing1_r17 *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_ppw_durationOfPRS_Processing1_r17
	Ppw_durationOfPRS_Processing2_r17 *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_ppw_durationOfPRS_Processing2_r17
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_Choice_Ppw_durationOfPRS_Processing1_r17:
		if err = ie.Ppw_durationOfPRS_Processing1_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Ppw_durationOfPRS_Processing1_r17", err)
		}
	case PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_Choice_Ppw_durationOfPRS_Processing2_r17:
		if err = ie.Ppw_durationOfPRS_Processing2_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Ppw_durationOfPRS_Processing2_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_Choice_Ppw_durationOfPRS_Processing1_r17:
		ie.Ppw_durationOfPRS_Processing1_r17 = new(PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_ppw_durationOfPRS_Processing1_r17)
		if err = ie.Ppw_durationOfPRS_Processing1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ppw_durationOfPRS_Processing1_r17", err)
		}
	case PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_Choice_Ppw_durationOfPRS_Processing2_r17:
		ie.Ppw_durationOfPRS_Processing2_r17 = new(PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_ppw_durationOfPRS_Processing2_r17)
		if err = ie.Ppw_durationOfPRS_Processing2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ppw_durationOfPRS_Processing2_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
