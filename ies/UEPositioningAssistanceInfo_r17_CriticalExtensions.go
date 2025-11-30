package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UEPositioningAssistanceInfo_r17_CriticalExtensions_Choice_nothing uint64 = iota
	UEPositioningAssistanceInfo_r17_CriticalExtensions_Choice_UePositioningAssistanceInfo_r17
	UEPositioningAssistanceInfo_r17_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type UEPositioningAssistanceInfo_r17_CriticalExtensions struct {
	Choice                          uint64
	UePositioningAssistanceInfo_r17 *UEPositioningAssistanceInfo_r17_IEs
	CriticalExtensionsFuture        interface{} `madatory`
}

func (ie *UEPositioningAssistanceInfo_r17_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UEPositioningAssistanceInfo_r17_CriticalExtensions_Choice_UePositioningAssistanceInfo_r17:
		if err = ie.UePositioningAssistanceInfo_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode UePositioningAssistanceInfo_r17", err)
		}
	case UEPositioningAssistanceInfo_r17_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UEPositioningAssistanceInfo_r17_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UEPositioningAssistanceInfo_r17_CriticalExtensions_Choice_UePositioningAssistanceInfo_r17:
		ie.UePositioningAssistanceInfo_r17 = new(UEPositioningAssistanceInfo_r17_IEs)
		if err = ie.UePositioningAssistanceInfo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode UePositioningAssistanceInfo_r17", err)
		}
	case UEPositioningAssistanceInfo_r17_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
