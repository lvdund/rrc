package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UplinkDataCompression_r17_Choice_nothing uint64 = iota
	UplinkDataCompression_r17_Choice_NewSetup
	UplinkDataCompression_r17_Choice_Drb_ContinueUDC
)

type UplinkDataCompression_r17 struct {
	Choice          uint64
	NewSetup        *UplinkDataCompression_r17_newSetup
	Drb_ContinueUDC uper.NULL `madatory`
}

func (ie *UplinkDataCompression_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UplinkDataCompression_r17_Choice_NewSetup:
		if err = ie.NewSetup.Encode(w); err != nil {
			err = utils.WrapError("Encode NewSetup", err)
		}
	case UplinkDataCompression_r17_Choice_Drb_ContinueUDC:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Drb_ContinueUDC", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UplinkDataCompression_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UplinkDataCompression_r17_Choice_NewSetup:
		ie.NewSetup = new(UplinkDataCompression_r17_newSetup)
		if err = ie.NewSetup.Decode(r); err != nil {
			return utils.WrapError("Decode NewSetup", err)
		}
	case UplinkDataCompression_r17_Choice_Drb_ContinueUDC:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Drb_ContinueUDC", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
