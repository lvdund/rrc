package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCRelease_CriticalExtensions_Choice_nothing uint64 = iota
	RRCRelease_CriticalExtensions_Choice_RrcRelease
	RRCRelease_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type RRCRelease_CriticalExtensions struct {
	Choice                   uint64
	RrcRelease               *RRCRelease_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *RRCRelease_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCRelease_CriticalExtensions_Choice_RrcRelease:
		if err = ie.RrcRelease.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcRelease", err)
		}
	case RRCRelease_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCRelease_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCRelease_CriticalExtensions_Choice_RrcRelease:
		ie.RrcRelease = new(RRCRelease_IEs)
		if err = ie.RrcRelease.Decode(r); err != nil {
			return utils.WrapError("Decode RrcRelease", err)
		}
	case RRCRelease_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
