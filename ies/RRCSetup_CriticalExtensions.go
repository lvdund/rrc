package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCSetup_CriticalExtensions_Choice_nothing uint64 = iota
	RRCSetup_CriticalExtensions_Choice_RrcSetup
	RRCSetup_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type RRCSetup_CriticalExtensions struct {
	Choice                   uint64
	RrcSetup                 *RRCSetup_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *RRCSetup_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCSetup_CriticalExtensions_Choice_RrcSetup:
		if err = ie.RrcSetup.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcSetup", err)
		}
	case RRCSetup_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCSetup_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCSetup_CriticalExtensions_Choice_RrcSetup:
		ie.RrcSetup = new(RRCSetup_IEs)
		if err = ie.RrcSetup.Decode(r); err != nil {
			return utils.WrapError("Decode RrcSetup", err)
		}
	case RRCSetup_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
