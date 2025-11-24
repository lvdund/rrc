package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	HandoverCommand_CriticalExtensions_C1_Choice_nothing uint64 = iota
	HandoverCommand_CriticalExtensions_C1_Choice_HandoverCommand
	HandoverCommand_CriticalExtensions_C1_Choice_Spare3
	HandoverCommand_CriticalExtensions_C1_Choice_Spare2
	HandoverCommand_CriticalExtensions_C1_Choice_Spare1
)

type HandoverCommand_CriticalExtensions_C1 struct {
	Choice          uint64
	HandoverCommand *HandoverCommand_IEs
	Spare3          uper.NULL `madatory`
	Spare2          uper.NULL `madatory`
	Spare1          uper.NULL `madatory`
}

func (ie *HandoverCommand_CriticalExtensions_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case HandoverCommand_CriticalExtensions_C1_Choice_HandoverCommand:
		if err = ie.HandoverCommand.Encode(w); err != nil {
			err = utils.WrapError("Encode HandoverCommand", err)
		}
	case HandoverCommand_CriticalExtensions_C1_Choice_Spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare3", err)
		}
	case HandoverCommand_CriticalExtensions_C1_Choice_Spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare2", err)
		}
	case HandoverCommand_CriticalExtensions_C1_Choice_Spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *HandoverCommand_CriticalExtensions_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case HandoverCommand_CriticalExtensions_C1_Choice_HandoverCommand:
		ie.HandoverCommand = new(HandoverCommand_IEs)
		if err = ie.HandoverCommand.Decode(r); err != nil {
			return utils.WrapError("Decode HandoverCommand", err)
		}
	case HandoverCommand_CriticalExtensions_C1_Choice_Spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare3", err)
		}
	case HandoverCommand_CriticalExtensions_C1_Choice_Spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare2", err)
		}
	case HandoverCommand_CriticalExtensions_C1_Choice_Spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
