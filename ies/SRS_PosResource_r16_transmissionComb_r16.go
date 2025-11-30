package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_PosResource_r16_transmissionComb_r16_Choice_nothing uint64 = iota
	SRS_PosResource_r16_transmissionComb_r16_Choice_N2_r16
	SRS_PosResource_r16_transmissionComb_r16_Choice_N4_r16
	SRS_PosResource_r16_transmissionComb_r16_Choice_N8_r16
)

type SRS_PosResource_r16_transmissionComb_r16 struct {
	Choice uint64
	N2_r16 *SRS_PosResource_r16_transmissionComb_r16_n2_r16
	N4_r16 *SRS_PosResource_r16_transmissionComb_r16_n4_r16
	N8_r16 *SRS_PosResource_r16_transmissionComb_r16_n8_r16
}

func (ie *SRS_PosResource_r16_transmissionComb_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_PosResource_r16_transmissionComb_r16_Choice_N2_r16:
		if err = ie.N2_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode N2_r16", err)
		}
	case SRS_PosResource_r16_transmissionComb_r16_Choice_N4_r16:
		if err = ie.N4_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode N4_r16", err)
		}
	case SRS_PosResource_r16_transmissionComb_r16_Choice_N8_r16:
		if err = ie.N8_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode N8_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_PosResource_r16_transmissionComb_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_PosResource_r16_transmissionComb_r16_Choice_N2_r16:
		ie.N2_r16 = new(SRS_PosResource_r16_transmissionComb_r16_n2_r16)
		if err = ie.N2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode N2_r16", err)
		}
	case SRS_PosResource_r16_transmissionComb_r16_Choice_N4_r16:
		ie.N4_r16 = new(SRS_PosResource_r16_transmissionComb_r16_n4_r16)
		if err = ie.N4_r16.Decode(r); err != nil {
			return utils.WrapError("Decode N4_r16", err)
		}
	case SRS_PosResource_r16_transmissionComb_r16_Choice_N8_r16:
		ie.N8_r16 = new(SRS_PosResource_r16_transmissionComb_r16_n8_r16)
		if err = ie.N8_r16.Decode(r); err != nil {
			return utils.WrapError("Decode N8_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
