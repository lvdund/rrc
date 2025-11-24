package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SBCCH_SL_BCH_MessageType_C1_Choice_nothing uint64 = iota
	SBCCH_SL_BCH_MessageType_C1_Choice_MasterInformationBlockSidelink
	SBCCH_SL_BCH_MessageType_C1_Choice_Spare1
)

type SBCCH_SL_BCH_MessageType_C1 struct {
	Choice                         uint64
	MasterInformationBlockSidelink *MasterInformationBlockSidelink
	Spare1                         uper.NULL `madatory`
}

func (ie *SBCCH_SL_BCH_MessageType_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SBCCH_SL_BCH_MessageType_C1_Choice_MasterInformationBlockSidelink:
		if err = ie.MasterInformationBlockSidelink.Encode(w); err != nil {
			err = utils.WrapError("Encode MasterInformationBlockSidelink", err)
		}
	case SBCCH_SL_BCH_MessageType_C1_Choice_Spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SBCCH_SL_BCH_MessageType_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SBCCH_SL_BCH_MessageType_C1_Choice_MasterInformationBlockSidelink:
		ie.MasterInformationBlockSidelink = new(MasterInformationBlockSidelink)
		if err = ie.MasterInformationBlockSidelink.Decode(r); err != nil {
			return utils.WrapError("Decode MasterInformationBlockSidelink", err)
		}
	case SBCCH_SL_BCH_MessageType_C1_Choice_Spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
