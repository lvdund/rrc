package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CounterCheck_CriticalExtensions_Choice_nothing uint64 = iota
	CounterCheck_CriticalExtensions_Choice_CounterCheck
	CounterCheck_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type CounterCheck_CriticalExtensions struct {
	Choice                   uint64
	CounterCheck             *CounterCheck_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *CounterCheck_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CounterCheck_CriticalExtensions_Choice_CounterCheck:
		if err = ie.CounterCheck.Encode(w); err != nil {
			err = utils.WrapError("Encode CounterCheck", err)
		}
	case CounterCheck_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CounterCheck_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CounterCheck_CriticalExtensions_Choice_CounterCheck:
		ie.CounterCheck = new(CounterCheck_IEs)
		if err = ie.CounterCheck.Decode(r); err != nil {
			return utils.WrapError("Decode CounterCheck", err)
		}
	case CounterCheck_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
