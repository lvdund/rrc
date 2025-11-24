package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasResults_measResultNeighCells_Choice_nothing uint64 = iota
	MeasResults_measResultNeighCells_Choice_MeasResultListNR
	MeasResults_measResultNeighCells_Choice_MeasResultListEUTRA
	MeasResults_measResultNeighCells_Choice_MeasResultListUTRA_FDD_r16
	MeasResults_measResultNeighCells_Choice_Sl_MeasResultsCandRelay_r17
)

type MeasResults_measResultNeighCells struct {
	Choice                      uint64
	MeasResultListNR            *MeasResultListNR
	MeasResultListEUTRA         *MeasResultListEUTRA
	MeasResultListUTRA_FDD_r16  *MeasResultListUTRA_FDD_r16
	Sl_MeasResultsCandRelay_r17 []byte `madatory`
}

func (ie *MeasResults_measResultNeighCells) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasResults_measResultNeighCells_Choice_MeasResultListNR:
		if err = ie.MeasResultListNR.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasResultListNR", err)
		}
	case MeasResults_measResultNeighCells_Choice_MeasResultListEUTRA:
		if err = ie.MeasResultListEUTRA.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasResultListEUTRA", err)
		}
	case MeasResults_measResultNeighCells_Choice_MeasResultListUTRA_FDD_r16:
		if err = ie.MeasResultListUTRA_FDD_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasResultListUTRA_FDD_r16", err)
		}
	case MeasResults_measResultNeighCells_Choice_Sl_MeasResultsCandRelay_r17:
		if err = w.WriteOctetString(ie.Sl_MeasResultsCandRelay_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			err = utils.WrapError("Encode Sl_MeasResultsCandRelay_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasResults_measResultNeighCells) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasResults_measResultNeighCells_Choice_MeasResultListNR:
		ie.MeasResultListNR = new(MeasResultListNR)
		if err = ie.MeasResultListNR.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultListNR", err)
		}
	case MeasResults_measResultNeighCells_Choice_MeasResultListEUTRA:
		ie.MeasResultListEUTRA = new(MeasResultListEUTRA)
		if err = ie.MeasResultListEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultListEUTRA", err)
		}
	case MeasResults_measResultNeighCells_Choice_MeasResultListUTRA_FDD_r16:
		ie.MeasResultListUTRA_FDD_r16 = new(MeasResultListUTRA_FDD_r16)
		if err = ie.MeasResultListUTRA_FDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultListUTRA_FDD_r16", err)
		}
	case MeasResults_measResultNeighCells_Choice_Sl_MeasResultsCandRelay_r17:
		var tmp_os_Sl_MeasResultsCandRelay_r17 []byte
		if tmp_os_Sl_MeasResultsCandRelay_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Sl_MeasResultsCandRelay_r17", err)
		}
		ie.Sl_MeasResultsCandRelay_r17 = tmp_os_Sl_MeasResultsCandRelay_r17
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
