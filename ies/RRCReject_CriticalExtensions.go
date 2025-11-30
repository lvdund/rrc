package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCReject_CriticalExtensions_Choice_nothing uint64 = iota
	RRCReject_CriticalExtensions_Choice_RrcReject
	RRCReject_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type RRCReject_CriticalExtensions struct {
	Choice                   uint64
	RrcReject                *RRCReject_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *RRCReject_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReject_CriticalExtensions_Choice_RrcReject:
		if err = ie.RrcReject.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcReject", err)
		}
	case RRCReject_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCReject_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReject_CriticalExtensions_Choice_RrcReject:
		ie.RrcReject = new(RRCReject_IEs)
		if err = ie.RrcReject.Decode(r); err != nil {
			return utils.WrapError("Decode RrcReject", err)
		}
	case RRCReject_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
