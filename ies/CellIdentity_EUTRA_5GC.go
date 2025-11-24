package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CellIdentity_EUTRA_5GC_Choice_nothing uint64 = iota
	CellIdentity_EUTRA_5GC_Choice_CellIdentity_EUTRA
	CellIdentity_EUTRA_5GC_Choice_CellId_index
)

type CellIdentity_EUTRA_5GC struct {
	Choice             uint64
	CellIdentity_EUTRA uper.BitString `lb:28,ub:28,madatory`
	CellId_index       int64          `lb:1,ub:maxPLMN,madatory`
}

func (ie *CellIdentity_EUTRA_5GC) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CellIdentity_EUTRA_5GC_Choice_CellIdentity_EUTRA:
		if err = w.WriteBitString(ie.CellIdentity_EUTRA.Bytes, uint(ie.CellIdentity_EUTRA.NumBits), &uper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
			err = utils.WrapError("Encode CellIdentity_EUTRA", err)
		}
	case CellIdentity_EUTRA_5GC_Choice_CellId_index:
		if err = w.WriteInteger(int64(ie.CellId_index), &uper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
			err = utils.WrapError("Encode CellId_index", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CellIdentity_EUTRA_5GC) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CellIdentity_EUTRA_5GC_Choice_CellIdentity_EUTRA:
		var tmp_bs_CellIdentity_EUTRA []byte
		var n_CellIdentity_EUTRA uint
		if tmp_bs_CellIdentity_EUTRA, n_CellIdentity_EUTRA, err = r.ReadBitString(&uper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
			return utils.WrapError("Decode CellIdentity_EUTRA", err)
		}
		ie.CellIdentity_EUTRA = uper.BitString{
			Bytes:   tmp_bs_CellIdentity_EUTRA,
			NumBits: uint64(n_CellIdentity_EUTRA),
		}
	case CellIdentity_EUTRA_5GC_Choice_CellId_index:
		var tmp_int_CellId_index int64
		if tmp_int_CellId_index, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
			return utils.WrapError("Decode CellId_index", err)
		}
		ie.CellId_index = tmp_int_CellId_index
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
