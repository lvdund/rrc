package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	Resourcetype_r16_SRS_PosResource_r16_Choice_nothing uint64 = iota
	Resourcetype_r16_SRS_PosResource_r16_Choice_Aperiodic_r16
	Resourcetype_r16_SRS_PosResource_r16_Choice_Semi_persistent_r16
	Resourcetype_r16_SRS_PosResource_r16_Choice_Periodic_r16
)

type Resourcetype_r16_SRS_PosResource_r16 struct {
	Choice              uint64
	Aperiodic_r16       *Resourcetype_r16_SRS_PosResource_r16_aperiodic_r16
	Semi_persistent_r16 *Resourcetype_r16_SRS_PosResource_r16_semi_persistent_r16
	Periodic_r16        *Resourcetype_r16_SRS_PosResource_r16_periodic_r16
}

func (ie *Resourcetype_r16_SRS_PosResource_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case Resourcetype_r16_SRS_PosResource_r16_Choice_Aperiodic_r16:
		if err = ie.Aperiodic_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Aperiodic_r16", err)
		}
	case Resourcetype_r16_SRS_PosResource_r16_Choice_Semi_persistent_r16:
		if err = ie.Semi_persistent_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Semi_persistent_r16", err)
		}
	case Resourcetype_r16_SRS_PosResource_r16_Choice_Periodic_r16:
		if err = ie.Periodic_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Periodic_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *Resourcetype_r16_SRS_PosResource_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case Resourcetype_r16_SRS_PosResource_r16_Choice_Aperiodic_r16:
		ie.Aperiodic_r16 = new(Resourcetype_r16_SRS_PosResource_r16_aperiodic_r16)
		if err = ie.Aperiodic_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Aperiodic_r16", err)
		}
	case Resourcetype_r16_SRS_PosResource_r16_Choice_Semi_persistent_r16:
		ie.Semi_persistent_r16 = new(Resourcetype_r16_SRS_PosResource_r16_semi_persistent_r16)
		if err = ie.Semi_persistent_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Semi_persistent_r16", err)
		}
	case Resourcetype_r16_SRS_PosResource_r16_Choice_Periodic_r16:
		ie.Periodic_r16 = new(Resourcetype_r16_SRS_PosResource_r16_periodic_r16)
		if err = ie.Periodic_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Periodic_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
