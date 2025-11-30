package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CounterCheckResponse_CriticalExtensions_Choice_nothing uint64 = iota
	CounterCheckResponse_CriticalExtensions_Choice_CounterCheckResponse
	CounterCheckResponse_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type CounterCheckResponse_CriticalExtensions struct {
	Choice                   uint64
	CounterCheckResponse     *CounterCheckResponse_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *CounterCheckResponse_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CounterCheckResponse_CriticalExtensions_Choice_CounterCheckResponse:
		if err = ie.CounterCheckResponse.Encode(w); err != nil {
			err = utils.WrapError("Encode CounterCheckResponse", err)
		}
	case CounterCheckResponse_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CounterCheckResponse_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CounterCheckResponse_CriticalExtensions_Choice_CounterCheckResponse:
		ie.CounterCheckResponse = new(CounterCheckResponse_IEs)
		if err = ie.CounterCheckResponse.Decode(r); err != nil {
			return utils.WrapError("Decode CounterCheckResponse", err)
		}
	case CounterCheckResponse_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
