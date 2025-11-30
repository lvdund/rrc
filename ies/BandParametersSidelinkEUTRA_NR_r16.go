package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandParametersSidelinkEUTRA_NR_r16_Choice_nothing uint64 = iota
	BandParametersSidelinkEUTRA_NR_r16_Choice_Eutra
	BandParametersSidelinkEUTRA_NR_r16_Choice_Nr
)

type BandParametersSidelinkEUTRA_NR_r16 struct {
	Choice uint64
	Eutra  *BandParametersSidelinkEUTRA_NR_r16_eutra
	Nr     *BandParametersSidelinkEUTRA_NR_r16_nr
}

func (ie *BandParametersSidelinkEUTRA_NR_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParametersSidelinkEUTRA_NR_r16_Choice_Eutra:
		if err = ie.Eutra.Encode(w); err != nil {
			err = utils.WrapError("Encode Eutra", err)
		}
	case BandParametersSidelinkEUTRA_NR_r16_Choice_Nr:
		if err = ie.Nr.Encode(w); err != nil {
			err = utils.WrapError("Encode Nr", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BandParametersSidelinkEUTRA_NR_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParametersSidelinkEUTRA_NR_r16_Choice_Eutra:
		ie.Eutra = new(BandParametersSidelinkEUTRA_NR_r16_eutra)
		if err = ie.Eutra.Decode(r); err != nil {
			return utils.WrapError("Decode Eutra", err)
		}
	case BandParametersSidelinkEUTRA_NR_r16_Choice_Nr:
		ie.Nr = new(BandParametersSidelinkEUTRA_NR_r16_nr)
		if err = ie.Nr.Decode(r); err != nil {
			return utils.WrapError("Decode Nr", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
