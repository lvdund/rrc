package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MobilityFromNRCommand_CriticalExtensions_Choice_nothing uint64 = iota
	MobilityFromNRCommand_CriticalExtensions_Choice_MobilityFromNRCommand
	MobilityFromNRCommand_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type MobilityFromNRCommand_CriticalExtensions struct {
	Choice                   uint64
	MobilityFromNRCommand    *MobilityFromNRCommand_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *MobilityFromNRCommand_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MobilityFromNRCommand_CriticalExtensions_Choice_MobilityFromNRCommand:
		if err = ie.MobilityFromNRCommand.Encode(w); err != nil {
			err = utils.WrapError("Encode MobilityFromNRCommand", err)
		}
	case MobilityFromNRCommand_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MobilityFromNRCommand_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MobilityFromNRCommand_CriticalExtensions_Choice_MobilityFromNRCommand:
		ie.MobilityFromNRCommand = new(MobilityFromNRCommand_IEs)
		if err = ie.MobilityFromNRCommand.Decode(r); err != nil {
			return utils.WrapError("Decode MobilityFromNRCommand", err)
		}
	case MobilityFromNRCommand_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
