package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SystemInformation_CriticalExtensions_Choice_nothing uint64 = iota
	SystemInformation_CriticalExtensions_Choice_SystemInformation
	SystemInformation_CriticalExtensions_Choice_CriticalExtensionsFuture_r16_SystemInformation
)

type SystemInformation_CriticalExtensions struct {
	Choice                                         uint64
	SystemInformation                              *SystemInformation_IEs
	CriticalExtensionsFuture_r16_SystemInformation *SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation
}

func (ie *SystemInformation_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SystemInformation_CriticalExtensions_Choice_SystemInformation:
		if err = ie.SystemInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode SystemInformation", err)
		}
	case SystemInformation_CriticalExtensions_Choice_CriticalExtensionsFuture_r16_SystemInformation:
		if err = ie.CriticalExtensionsFuture_r16_SystemInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode CriticalExtensionsFuture_r16_SystemInformation", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SystemInformation_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SystemInformation_CriticalExtensions_Choice_SystemInformation:
		ie.SystemInformation = new(SystemInformation_IEs)
		if err = ie.SystemInformation.Decode(r); err != nil {
			return utils.WrapError("Decode SystemInformation", err)
		}
	case SystemInformation_CriticalExtensions_Choice_CriticalExtensionsFuture_r16_SystemInformation:
		ie.CriticalExtensionsFuture_r16_SystemInformation = new(SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation)
		if err = ie.CriticalExtensionsFuture_r16_SystemInformation.Decode(r); err != nil {
			return utils.WrapError("Decode CriticalExtensionsFuture_r16_SystemInformation", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
