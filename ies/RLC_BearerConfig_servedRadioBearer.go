package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLC_BearerConfig_servedRadioBearer_Choice_nothing uint64 = iota
	RLC_BearerConfig_servedRadioBearer_Choice_Srb_Identity
	RLC_BearerConfig_servedRadioBearer_Choice_Drb_Identity
)

type RLC_BearerConfig_servedRadioBearer struct {
	Choice       uint64
	Srb_Identity *SRB_Identity
	Drb_Identity *DRB_Identity
}

func (ie *RLC_BearerConfig_servedRadioBearer) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RLC_BearerConfig_servedRadioBearer_Choice_Srb_Identity:
		if err = ie.Srb_Identity.Encode(w); err != nil {
			err = utils.WrapError("Encode Srb_Identity", err)
		}
	case RLC_BearerConfig_servedRadioBearer_Choice_Drb_Identity:
		if err = ie.Drb_Identity.Encode(w); err != nil {
			err = utils.WrapError("Encode Drb_Identity", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RLC_BearerConfig_servedRadioBearer) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RLC_BearerConfig_servedRadioBearer_Choice_Srb_Identity:
		ie.Srb_Identity = new(SRB_Identity)
		if err = ie.Srb_Identity.Decode(r); err != nil {
			return utils.WrapError("Decode Srb_Identity", err)
		}
	case RLC_BearerConfig_servedRadioBearer_Choice_Drb_Identity:
		ie.Drb_Identity = new(DRB_Identity)
		if err = ie.Drb_Identity.Decode(r); err != nil {
			return utils.WrapError("Decode Drb_Identity", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
