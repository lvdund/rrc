package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MBSInterestIndication_r17_CriticalExtensions_Choice_nothing uint64 = iota
	MBSInterestIndication_r17_CriticalExtensions_Choice_MbsInterestIndication_r17
	MBSInterestIndication_r17_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type MBSInterestIndication_r17_CriticalExtensions struct {
	Choice                    uint64
	MbsInterestIndication_r17 *MBSInterestIndication_r17_IEs
	CriticalExtensionsFuture  interface{} `madatory`
}

func (ie *MBSInterestIndication_r17_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MBSInterestIndication_r17_CriticalExtensions_Choice_MbsInterestIndication_r17:
		if err = ie.MbsInterestIndication_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode MbsInterestIndication_r17", err)
		}
	case MBSInterestIndication_r17_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MBSInterestIndication_r17_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MBSInterestIndication_r17_CriticalExtensions_Choice_MbsInterestIndication_r17:
		ie.MbsInterestIndication_r17 = new(MBSInterestIndication_r17_IEs)
		if err = ie.MbsInterestIndication_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MbsInterestIndication_r17", err)
		}
	case MBSInterestIndication_r17_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
