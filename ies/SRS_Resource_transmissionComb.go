package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_Resource_transmissionComb_Choice_nothing uint64 = iota
	SRS_Resource_transmissionComb_Choice_N2
	SRS_Resource_transmissionComb_Choice_N4
)

type SRS_Resource_transmissionComb struct {
	Choice uint64
	N2     *SRS_Resource_transmissionComb_n2
	N4     *SRS_Resource_transmissionComb_n4
}

func (ie *SRS_Resource_transmissionComb) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_Resource_transmissionComb_Choice_N2:
		if err = ie.N2.Encode(w); err != nil {
			err = utils.WrapError("Encode N2", err)
		}
	case SRS_Resource_transmissionComb_Choice_N4:
		if err = ie.N4.Encode(w); err != nil {
			err = utils.WrapError("Encode N4", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_Resource_transmissionComb) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_Resource_transmissionComb_Choice_N2:
		ie.N2 = new(SRS_Resource_transmissionComb_n2)
		if err = ie.N2.Decode(r); err != nil {
			return utils.WrapError("Decode N2", err)
		}
	case SRS_Resource_transmissionComb_Choice_N4:
		ie.N4 = new(SRS_Resource_transmissionComb_n4)
		if err = ie.N4.Decode(r); err != nil {
			return utils.WrapError("Decode N4", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
