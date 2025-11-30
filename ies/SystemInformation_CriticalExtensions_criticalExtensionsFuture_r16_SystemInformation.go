package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation_Choice_nothing uint64 = iota
	SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation_Choice_PosSystemInformation_r16
	SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation_Choice_CriticalExtensionsFuture
)

type SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation struct {
	Choice                   uint64
	PosSystemInformation_r16 *PosSystemInformation_r16_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation_Choice_PosSystemInformation_r16:
		if err = ie.PosSystemInformation_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSystemInformation_r16", err)
		}
	case SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation_Choice_PosSystemInformation_r16:
		ie.PosSystemInformation_r16 = new(PosSystemInformation_r16_IEs)
		if err = ie.PosSystemInformation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSystemInformation_r16", err)
		}
	case SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
