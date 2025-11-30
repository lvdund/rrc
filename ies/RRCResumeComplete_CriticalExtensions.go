package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCResumeComplete_CriticalExtensions_Choice_nothing uint64 = iota
	RRCResumeComplete_CriticalExtensions_Choice_RrcResumeComplete
	RRCResumeComplete_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type RRCResumeComplete_CriticalExtensions struct {
	Choice                   uint64
	RrcResumeComplete        *RRCResumeComplete_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *RRCResumeComplete_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCResumeComplete_CriticalExtensions_Choice_RrcResumeComplete:
		if err = ie.RrcResumeComplete.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcResumeComplete", err)
		}
	case RRCResumeComplete_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCResumeComplete_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCResumeComplete_CriticalExtensions_Choice_RrcResumeComplete:
		ie.RrcResumeComplete = new(RRCResumeComplete_IEs)
		if err = ie.RrcResumeComplete.Decode(r); err != nil {
			return utils.WrapError("Decode RrcResumeComplete", err)
		}
	case RRCResumeComplete_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
