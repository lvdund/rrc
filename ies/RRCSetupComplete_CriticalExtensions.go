package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCSetupComplete_CriticalExtensions_Choice_nothing uint64 = iota
	RRCSetupComplete_CriticalExtensions_Choice_RrcSetupComplete
	RRCSetupComplete_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type RRCSetupComplete_CriticalExtensions struct {
	Choice                   uint64
	RrcSetupComplete         *RRCSetupComplete_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *RRCSetupComplete_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCSetupComplete_CriticalExtensions_Choice_RrcSetupComplete:
		if err = ie.RrcSetupComplete.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcSetupComplete", err)
		}
	case RRCSetupComplete_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCSetupComplete_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCSetupComplete_CriticalExtensions_Choice_RrcSetupComplete:
		ie.RrcSetupComplete = new(RRCSetupComplete_IEs)
		if err = ie.RrcSetupComplete.Decode(r); err != nil {
			return utils.WrapError("Decode RrcSetupComplete", err)
		}
	case RRCSetupComplete_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
