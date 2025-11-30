package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCReestablishment_CriticalExtensions_Choice_nothing uint64 = iota
	RRCReestablishment_CriticalExtensions_Choice_RrcReestablishment
	RRCReestablishment_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type RRCReestablishment_CriticalExtensions struct {
	Choice                   uint64
	RrcReestablishment       *RRCReestablishment_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *RRCReestablishment_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReestablishment_CriticalExtensions_Choice_RrcReestablishment:
		if err = ie.RrcReestablishment.Encode(w); err != nil {
			err = utils.WrapError("Encode RrcReestablishment", err)
		}
	case RRCReestablishment_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCReestablishment_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReestablishment_CriticalExtensions_Choice_RrcReestablishment:
		ie.RrcReestablishment = new(RRCReestablishment_IEs)
		if err = ie.RrcReestablishment.Decode(r); err != nil {
			return utils.WrapError("Decode RrcReestablishment", err)
		}
	case RRCReestablishment_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
