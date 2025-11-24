package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	S_NSSAI_Choice_nothing uint64 = iota
	S_NSSAI_Choice_Sst
	S_NSSAI_Choice_Sst_SD
)

type S_NSSAI struct {
	Choice uint64
	Sst    uper.BitString `lb:8,ub:8,madatory`
	Sst_SD uper.BitString `lb:32,ub:32,madatory`
}

func (ie *S_NSSAI) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case S_NSSAI_Choice_Sst:
		if err = w.WriteBitString(ie.Sst.Bytes, uint(ie.Sst.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode Sst", err)
		}
	case S_NSSAI_Choice_Sst_SD:
		if err = w.WriteBitString(ie.Sst_SD.Bytes, uint(ie.Sst_SD.NumBits), &uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			err = utils.WrapError("Encode Sst_SD", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *S_NSSAI) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case S_NSSAI_Choice_Sst:
		var tmp_bs_Sst []byte
		var n_Sst uint
		if tmp_bs_Sst, n_Sst, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Sst", err)
		}
		ie.Sst = uper.BitString{
			Bytes:   tmp_bs_Sst,
			NumBits: uint64(n_Sst),
		}
	case S_NSSAI_Choice_Sst_SD:
		var tmp_bs_Sst_SD []byte
		var n_Sst_SD uint
		if tmp_bs_Sst_SD, n_Sst_SD, err = r.ReadBitString(&uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode Sst_SD", err)
		}
		ie.Sst_SD = uper.BitString{
			Bytes:   tmp_bs_Sst_SD,
			NumBits: uint64(n_Sst_SD),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
