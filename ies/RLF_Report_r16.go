package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLF_Report_r16_Choice_nothing uint64 = iota
	RLF_Report_r16_Choice_Nr_RLF_Report_r16
	RLF_Report_r16_Choice_Eutra_RLF_Report_r16
)

type RLF_Report_r16 struct {
	Choice               uint64
	Nr_RLF_Report_r16    *RLF_Report_r16_nr_RLF_Report_r16
	Eutra_RLF_Report_r16 *RLF_Report_r16_eutra_RLF_Report_r16
}

func (ie *RLF_Report_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RLF_Report_r16_Choice_Nr_RLF_Report_r16:
		if err = ie.Nr_RLF_Report_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Nr_RLF_Report_r16", err)
		}
	case RLF_Report_r16_Choice_Eutra_RLF_Report_r16:
		if err = ie.Eutra_RLF_Report_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Eutra_RLF_Report_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RLF_Report_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RLF_Report_r16_Choice_Nr_RLF_Report_r16:
		ie.Nr_RLF_Report_r16 = new(RLF_Report_r16_nr_RLF_Report_r16)
		if err = ie.Nr_RLF_Report_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Nr_RLF_Report_r16", err)
		}
	case RLF_Report_r16_Choice_Eutra_RLF_Report_r16:
		ie.Eutra_RLF_Report_r16 = new(RLF_Report_r16_eutra_RLF_Report_r16)
		if err = ie.Eutra_RLF_Report_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Eutra_RLF_Report_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
