package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BCCH_BCH_MessageType_Choice_nothing uint64 = iota
	BCCH_BCH_MessageType_Choice_Mib
	BCCH_BCH_MessageType_Choice_MessageClassExtension
)

type BCCH_BCH_MessageType struct {
	Choice                uint64
	Mib                   *MIB
	MessageClassExtension interface{} `madatory`
}

func (ie *BCCH_BCH_MessageType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BCCH_BCH_MessageType_Choice_Mib:
		if err = ie.Mib.Encode(w); err != nil {
			err = utils.WrapError("Encode Mib", err)
		}
	case BCCH_BCH_MessageType_Choice_MessageClassExtension:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BCCH_BCH_MessageType) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BCCH_BCH_MessageType_Choice_Mib:
		ie.Mib = new(MIB)
		if err = ie.Mib.Decode(r); err != nil {
			return utils.WrapError("Decode Mib", err)
		}
	case BCCH_BCH_MessageType_Choice_MessageClassExtension:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
