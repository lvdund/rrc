package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCReconfigurationCompleteSidelink_CriticalExtensions_Choice_nothing uint64 = iota
	RRCReconfigurationCompleteSidelink_CriticalExtensions_Choice_RrcReconfigurationCompleteSidelink_r16
	RRCReconfigurationCompleteSidelink_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type RRCReconfigurationCompleteSidelink_CriticalExtensions struct {
	Choice                                 uint64
	RrcReconfigurationCompleteSidelink_r16 *RRCReconfigurationCompleteSidelink_r16_IEs
	CriticalExtensionsFuture               interface{} `madatory`
}

func (ie *RRCReconfigurationCompleteSidelink_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReconfigurationCompleteSidelink_CriticalExtensions_Choice_RrcReconfigurationCompleteSidelink_r16:
		if err = ie.RrcReconfigurationCompleteSidelink_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcReconfigurationCompleteSidelink_r16", err)
		}
	case RRCReconfigurationCompleteSidelink_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCReconfigurationCompleteSidelink_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReconfigurationCompleteSidelink_CriticalExtensions_Choice_RrcReconfigurationCompleteSidelink_r16:
		ie.RrcReconfigurationCompleteSidelink_r16 = new(RRCReconfigurationCompleteSidelink_r16_IEs)
		if err = ie.RrcReconfigurationCompleteSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode RrcReconfigurationCompleteSidelink_r16", err)
		}
	case RRCReconfigurationCompleteSidelink_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
