package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasObjectRxTxDiff_r17_dl_Ref_r17_Choice_nothing uint64 = iota
	MeasObjectRxTxDiff_r17_dl_Ref_r17_Choice_Prs_Ref_r17
	MeasObjectRxTxDiff_r17_dl_Ref_r17_Choice_Csi_RS_Ref_r17
)

type MeasObjectRxTxDiff_r17_dl_Ref_r17 struct {
	Choice         uint64
	Prs_Ref_r17    uper.NULL `madatory`
	Csi_RS_Ref_r17 uper.NULL `madatory`
}

func (ie *MeasObjectRxTxDiff_r17_dl_Ref_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasObjectRxTxDiff_r17_dl_Ref_r17_Choice_Prs_Ref_r17:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Prs_Ref_r17", err)
		}
	case MeasObjectRxTxDiff_r17_dl_Ref_r17_Choice_Csi_RS_Ref_r17:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Csi_RS_Ref_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasObjectRxTxDiff_r17_dl_Ref_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasObjectRxTxDiff_r17_dl_Ref_r17_Choice_Prs_Ref_r17:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Prs_Ref_r17", err)
		}
	case MeasObjectRxTxDiff_r17_dl_Ref_r17_Choice_Csi_RS_Ref_r17:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Csi_RS_Ref_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
