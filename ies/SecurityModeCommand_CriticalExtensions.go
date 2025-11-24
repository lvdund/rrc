package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SecurityModeCommand_CriticalExtensions_Choice_nothing uint64 = iota
	SecurityModeCommand_CriticalExtensions_Choice_SecurityModeCommand
	SecurityModeCommand_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type SecurityModeCommand_CriticalExtensions struct {
	Choice                   uint64
	SecurityModeCommand      *SecurityModeCommand_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *SecurityModeCommand_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SecurityModeCommand_CriticalExtensions_Choice_SecurityModeCommand:
		if err = ie.SecurityModeCommand.Encode(w); err != nil {
			err = utils.WrapError("Encode SecurityModeCommand", err)
		}
	case SecurityModeCommand_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SecurityModeCommand_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SecurityModeCommand_CriticalExtensions_Choice_SecurityModeCommand:
		ie.SecurityModeCommand = new(SecurityModeCommand_IEs)
		if err = ie.SecurityModeCommand.Decode(r); err != nil {
			return utils.WrapError("Decode SecurityModeCommand", err)
		}
	case SecurityModeCommand_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
