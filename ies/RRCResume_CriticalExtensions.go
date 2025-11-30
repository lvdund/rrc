package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCResume_CriticalExtensions_Choice_nothing uint64 = iota
	RRCResume_CriticalExtensions_Choice_RrcResume
	RRCResume_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type RRCResume_CriticalExtensions struct {
	Choice                   uint64
	RrcResume                *RRCResume_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *RRCResume_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCResume_CriticalExtensions_Choice_RrcResume:
		if err = ie.RrcResume.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcResume", err)
		}
	case RRCResume_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCResume_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCResume_CriticalExtensions_Choice_RrcResume:
		ie.RrcResume = new(RRCResume_IEs)
		if err = ie.RrcResume.Decode(r); err != nil {
			return utils.WrapError("Decode RrcResume", err)
		}
	case RRCResume_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
