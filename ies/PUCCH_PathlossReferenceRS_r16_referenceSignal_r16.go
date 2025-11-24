package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_PathlossReferenceRS_r16_referenceSignal_r16_Choice_nothing uint64 = iota
	PUCCH_PathlossReferenceRS_r16_referenceSignal_r16_Choice_Ssb_Index_r16
	PUCCH_PathlossReferenceRS_r16_referenceSignal_r16_Choice_Csi_RS_Index_r16
)

type PUCCH_PathlossReferenceRS_r16_referenceSignal_r16 struct {
	Choice           uint64
	Ssb_Index_r16    *SSB_Index
	Csi_RS_Index_r16 *NZP_CSI_RS_ResourceId
}

func (ie *PUCCH_PathlossReferenceRS_r16_referenceSignal_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_PathlossReferenceRS_r16_referenceSignal_r16_Choice_Ssb_Index_r16:
		if err = ie.Ssb_Index_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Ssb_Index_r16", err)
		}
	case PUCCH_PathlossReferenceRS_r16_referenceSignal_r16_Choice_Csi_RS_Index_r16:
		if err = ie.Csi_RS_Index_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Csi_RS_Index_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PUCCH_PathlossReferenceRS_r16_referenceSignal_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_PathlossReferenceRS_r16_referenceSignal_r16_Choice_Ssb_Index_r16:
		ie.Ssb_Index_r16 = new(SSB_Index)
		if err = ie.Ssb_Index_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_Index_r16", err)
		}
	case PUCCH_PathlossReferenceRS_r16_referenceSignal_r16_Choice_Csi_RS_Index_r16:
		ie.Csi_RS_Index_r16 = new(NZP_CSI_RS_ResourceId)
		if err = ie.Csi_RS_Index_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RS_Index_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
