package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCReconfiguration_CriticalExtensions_Choice_nothing uint64 = iota
	RRCReconfiguration_CriticalExtensions_Choice_RrcReconfiguration
	RRCReconfiguration_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type RRCReconfiguration_CriticalExtensions struct {
	Choice                   uint64
	RrcReconfiguration       *RRCReconfiguration_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *RRCReconfiguration_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReconfiguration_CriticalExtensions_Choice_RrcReconfiguration:
		if err = ie.RrcReconfiguration.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcReconfiguration", err)
		}
	case RRCReconfiguration_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCReconfiguration_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReconfiguration_CriticalExtensions_Choice_RrcReconfiguration:
		ie.RrcReconfiguration = new(RRCReconfiguration_IEs)
		if err = ie.RrcReconfiguration.Decode(r); err != nil {
			return utils.WrapError("Decode RrcReconfiguration", err)
		}
	case RRCReconfiguration_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
