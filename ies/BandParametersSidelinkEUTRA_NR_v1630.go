package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandParametersSidelinkEUTRA_NR_v1630_Choice_nothing uint64 = iota
	BandParametersSidelinkEUTRA_NR_v1630_Choice_Eutra
	BandParametersSidelinkEUTRA_NR_v1630_Choice_Nr
)

type BandParametersSidelinkEUTRA_NR_v1630 struct {
	Choice uint64
	Eutra  uper.NULL `madatory`
	Nr     *BandParametersSidelinkEUTRA_NR_v1630_nr
}

func (ie *BandParametersSidelinkEUTRA_NR_v1630) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParametersSidelinkEUTRA_NR_v1630_Choice_Eutra:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Eutra", err)
		}
	case BandParametersSidelinkEUTRA_NR_v1630_Choice_Nr:
		if err = ie.Nr.Encode(w); err != nil {
			err = utils.WrapError("Encode Nr", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BandParametersSidelinkEUTRA_NR_v1630) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParametersSidelinkEUTRA_NR_v1630_Choice_Eutra:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Eutra", err)
		}
	case BandParametersSidelinkEUTRA_NR_v1630_Choice_Nr:
		ie.Nr = new(BandParametersSidelinkEUTRA_NR_v1630_nr)
		if err = ie.Nr.Decode(r); err != nil {
			return utils.WrapError("Decode Nr", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
