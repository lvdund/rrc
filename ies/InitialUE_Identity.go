package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	InitialUE_Identity_Choice_nothing uint64 = iota
	InitialUE_Identity_Choice_Ng_5G_S_TMSI_Part1
	InitialUE_Identity_Choice_RandomValue
)

type InitialUE_Identity struct {
	Choice             uint64
	Ng_5G_S_TMSI_Part1 uper.BitString `lb:39,ub:39,madatory`
	RandomValue        uper.BitString `lb:39,ub:39,madatory`
}

func (ie *InitialUE_Identity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case InitialUE_Identity_Choice_Ng_5G_S_TMSI_Part1:
		if err = w.WriteBitString(ie.Ng_5G_S_TMSI_Part1.Bytes, uint(ie.Ng_5G_S_TMSI_Part1.NumBits), &uper.Constraint{Lb: 39, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode Ng_5G_S_TMSI_Part1", err)
		}
	case InitialUE_Identity_Choice_RandomValue:
		if err = w.WriteBitString(ie.RandomValue.Bytes, uint(ie.RandomValue.NumBits), &uper.Constraint{Lb: 39, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode RandomValue", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *InitialUE_Identity) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case InitialUE_Identity_Choice_Ng_5G_S_TMSI_Part1:
		var tmp_bs_Ng_5G_S_TMSI_Part1 []byte
		var n_Ng_5G_S_TMSI_Part1 uint
		if tmp_bs_Ng_5G_S_TMSI_Part1, n_Ng_5G_S_TMSI_Part1, err = r.ReadBitString(&uper.Constraint{Lb: 39, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode Ng_5G_S_TMSI_Part1", err)
		}
		ie.Ng_5G_S_TMSI_Part1 = uper.BitString{
			Bytes:   tmp_bs_Ng_5G_S_TMSI_Part1,
			NumBits: uint64(n_Ng_5G_S_TMSI_Part1),
		}
	case InitialUE_Identity_Choice_RandomValue:
		var tmp_bs_RandomValue []byte
		var n_RandomValue uint
		if tmp_bs_RandomValue, n_RandomValue, err = r.ReadBitString(&uper.Constraint{Lb: 39, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode RandomValue", err)
		}
		ie.RandomValue = uper.BitString{
			Bytes:   tmp_bs_RandomValue,
			NumBits: uint64(n_RandomValue),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
