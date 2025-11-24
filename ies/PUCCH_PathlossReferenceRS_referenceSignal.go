package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_PathlossReferenceRS_referenceSignal_Choice_nothing uint64 = iota
	PUCCH_PathlossReferenceRS_referenceSignal_Choice_Ssb_Index
	PUCCH_PathlossReferenceRS_referenceSignal_Choice_Csi_RS_Index
)

type PUCCH_PathlossReferenceRS_referenceSignal struct {
	Choice       uint64
	Ssb_Index    *SSB_Index
	Csi_RS_Index *NZP_CSI_RS_ResourceId
}

func (ie *PUCCH_PathlossReferenceRS_referenceSignal) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_PathlossReferenceRS_referenceSignal_Choice_Ssb_Index:
		if err = ie.Ssb_Index.Encode(w); err != nil {
			err = utils.WrapError("Encode Ssb_Index", err)
		}
	case PUCCH_PathlossReferenceRS_referenceSignal_Choice_Csi_RS_Index:
		if err = ie.Csi_RS_Index.Encode(w); err != nil {
			err = utils.WrapError("Encode Csi_RS_Index", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PUCCH_PathlossReferenceRS_referenceSignal) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_PathlossReferenceRS_referenceSignal_Choice_Ssb_Index:
		ie.Ssb_Index = new(SSB_Index)
		if err = ie.Ssb_Index.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_Index", err)
		}
	case PUCCH_PathlossReferenceRS_referenceSignal_Choice_Csi_RS_Index:
		ie.Csi_RS_Index = new(NZP_CSI_RS_ResourceId)
		if err = ie.Csi_RS_Index.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RS_Index", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
