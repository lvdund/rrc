package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SCGFailureInformationEUTRA_CriticalExtensions_Choice_nothing uint64 = iota
	SCGFailureInformationEUTRA_CriticalExtensions_Choice_ScgFailureInformationEUTRA
	SCGFailureInformationEUTRA_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type SCGFailureInformationEUTRA_CriticalExtensions struct {
	Choice                     uint64
	ScgFailureInformationEUTRA *SCGFailureInformationEUTRA_IEs
	CriticalExtensionsFuture   interface{} `madatory`
}

func (ie *SCGFailureInformationEUTRA_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCGFailureInformationEUTRA_CriticalExtensions_Choice_ScgFailureInformationEUTRA:
		if err = ie.ScgFailureInformationEUTRA.Encode(w); err != nil {
			err = utils.WrapError("Encode ScgFailureInformationEUTRA", err)
		}
	case SCGFailureInformationEUTRA_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SCGFailureInformationEUTRA_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCGFailureInformationEUTRA_CriticalExtensions_Choice_ScgFailureInformationEUTRA:
		ie.ScgFailureInformationEUTRA = new(SCGFailureInformationEUTRA_IEs)
		if err = ie.ScgFailureInformationEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode ScgFailureInformationEUTRA", err)
		}
	case SCGFailureInformationEUTRA_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
