package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PRACH_ResourceDedicatedBFR_Choice_nothing uint64 = iota
	PRACH_ResourceDedicatedBFR_Choice_Ssb
	PRACH_ResourceDedicatedBFR_Choice_Csi_RS
)

type PRACH_ResourceDedicatedBFR struct {
	Choice uint64
	Ssb    *BFR_SSB_Resource
	Csi_RS *BFR_CSIRS_Resource
}

func (ie *PRACH_ResourceDedicatedBFR) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PRACH_ResourceDedicatedBFR_Choice_Ssb:
		if err = ie.Ssb.Encode(w); err != nil {
			err = utils.WrapError("Encode Ssb", err)
		}
	case PRACH_ResourceDedicatedBFR_Choice_Csi_RS:
		if err = ie.Csi_RS.Encode(w); err != nil {
			err = utils.WrapError("Encode Csi_RS", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PRACH_ResourceDedicatedBFR) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PRACH_ResourceDedicatedBFR_Choice_Ssb:
		ie.Ssb = new(BFR_SSB_Resource)
		if err = ie.Ssb.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb", err)
		}
	case PRACH_ResourceDedicatedBFR_Choice_Csi_RS:
		ie.Csi_RS = new(BFR_CSIRS_Resource)
		if err = ie.Csi_RS.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RS", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
