package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CFRA_resources_Choice_nothing uint64 = iota
	CFRA_resources_Choice_Ssb
	CFRA_resources_Choice_Csirs
)

type CFRA_resources struct {
	Choice uint64
	Ssb    *CFRA_resources_ssb
	Csirs  *CFRA_resources_csirs
}

func (ie *CFRA_resources) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CFRA_resources_Choice_Ssb:
		if err = ie.Ssb.Encode(w); err != nil {
			err = utils.WrapError("Encode Ssb", err)
		}
	case CFRA_resources_Choice_Csirs:
		if err = ie.Csirs.Encode(w); err != nil {
			err = utils.WrapError("Encode Csirs", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CFRA_resources) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CFRA_resources_Choice_Ssb:
		ie.Ssb = new(CFRA_resources_ssb)
		if err = ie.Ssb.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb", err)
		}
	case CFRA_resources_Choice_Csirs:
		ie.Csirs = new(CFRA_resources_csirs)
		if err = ie.Csirs.Decode(r); err != nil {
			return utils.WrapError("Decode Csirs", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
