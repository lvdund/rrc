package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MCGFailureInformation_r16_CriticalExtensions_Choice_nothing uint64 = iota
	MCGFailureInformation_r16_CriticalExtensions_Choice_McgFailureInformation_r16
	MCGFailureInformation_r16_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type MCGFailureInformation_r16_CriticalExtensions struct {
	Choice                    uint64
	McgFailureInformation_r16 *MCGFailureInformation_r16_IEs
	CriticalExtensionsFuture  interface{} `madatory`
}

func (ie *MCGFailureInformation_r16_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MCGFailureInformation_r16_CriticalExtensions_Choice_McgFailureInformation_r16:
		if err = ie.McgFailureInformation_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode McgFailureInformation_r16", err)
		}
	case MCGFailureInformation_r16_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MCGFailureInformation_r16_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MCGFailureInformation_r16_CriticalExtensions_Choice_McgFailureInformation_r16:
		ie.McgFailureInformation_r16 = new(MCGFailureInformation_r16_IEs)
		if err = ie.McgFailureInformation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode McgFailureInformation_r16", err)
		}
	case MCGFailureInformation_r16_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
