package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandParameters_Choice_nothing uint64 = iota
	BandParameters_Choice_Eutra
	BandParameters_Choice_Nr
)

type BandParameters struct {
	Choice uint64
	Eutra  *BandParameters_eutra
	Nr     *BandParameters_nr
}

func (ie *BandParameters) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParameters_Choice_Eutra:
		if err = ie.Eutra.Encode(w); err != nil {
			err = utils.WrapError("Encode Eutra", err)
		}
	case BandParameters_Choice_Nr:
		if err = ie.Nr.Encode(w); err != nil {
			err = utils.WrapError("Encode Nr", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BandParameters) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParameters_Choice_Eutra:
		ie.Eutra = new(BandParameters_eutra)
		if err = ie.Eutra.Decode(r); err != nil {
			return utils.WrapError("Decode Eutra", err)
		}
	case BandParameters_Choice_Nr:
		ie.Nr = new(BandParameters_nr)
		if err = ie.Nr.Decode(r); err != nil {
			return utils.WrapError("Decode Nr", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
