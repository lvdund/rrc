package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	EphemerisInfo_r17_Choice_nothing uint64 = iota
	EphemerisInfo_r17_Choice_PositionVelocity_r17
	EphemerisInfo_r17_Choice_Orbital_r17
)

type EphemerisInfo_r17 struct {
	Choice               uint64
	PositionVelocity_r17 *PositionVelocity_r17
	Orbital_r17          *Orbital_r17
}

func (ie *EphemerisInfo_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EphemerisInfo_r17_Choice_PositionVelocity_r17:
		if err = ie.PositionVelocity_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode PositionVelocity_r17", err)
		}
	case EphemerisInfo_r17_Choice_Orbital_r17:
		if err = ie.Orbital_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Orbital_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *EphemerisInfo_r17) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EphemerisInfo_r17_Choice_PositionVelocity_r17:
		ie.PositionVelocity_r17 = new(PositionVelocity_r17)
		if err = ie.PositionVelocity_r17.Decode(r); err != nil {
			return utils.WrapError("Decode PositionVelocity_r17", err)
		}
	case EphemerisInfo_r17_Choice_Orbital_r17:
		ie.Orbital_r17 = new(Orbital_r17)
		if err = ie.Orbital_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Orbital_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
