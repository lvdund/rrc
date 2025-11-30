package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FailureInformation_CriticalExtensions_Choice_nothing uint64 = iota
	FailureInformation_CriticalExtensions_Choice_FailureInformation
	FailureInformation_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type FailureInformation_CriticalExtensions struct {
	Choice                   uint64
	FailureInformation       *FailureInformation_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *FailureInformation_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case FailureInformation_CriticalExtensions_Choice_FailureInformation:
		if err = ie.FailureInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode FailureInformation", err)
		}
	case FailureInformation_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *FailureInformation_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case FailureInformation_CriticalExtensions_Choice_FailureInformation:
		ie.FailureInformation = new(FailureInformation_IEs)
		if err = ie.FailureInformation.Decode(r); err != nil {
			return utils.WrapError("Decode FailureInformation", err)
		}
	case FailureInformation_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
