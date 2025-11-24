package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCReconfigurationFailureSidelink_CriticalExtensions_Choice_nothing uint64 = iota
	RRCReconfigurationFailureSidelink_CriticalExtensions_Choice_RrcReconfigurationFailureSidelink_r16
	RRCReconfigurationFailureSidelink_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type RRCReconfigurationFailureSidelink_CriticalExtensions struct {
	Choice                                uint64
	RrcReconfigurationFailureSidelink_r16 *RRCReconfigurationFailureSidelink_r16_IEs
	CriticalExtensionsFuture              interface{} `madatory`
}

func (ie *RRCReconfigurationFailureSidelink_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReconfigurationFailureSidelink_CriticalExtensions_Choice_RrcReconfigurationFailureSidelink_r16:
		if err = ie.RrcReconfigurationFailureSidelink_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcReconfigurationFailureSidelink_r16", err)
		}
	case RRCReconfigurationFailureSidelink_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCReconfigurationFailureSidelink_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReconfigurationFailureSidelink_CriticalExtensions_Choice_RrcReconfigurationFailureSidelink_r16:
		ie.RrcReconfigurationFailureSidelink_r16 = new(RRCReconfigurationFailureSidelink_r16_IEs)
		if err = ie.RrcReconfigurationFailureSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode RrcReconfigurationFailureSidelink_r16", err)
		}
	case RRCReconfigurationFailureSidelink_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
