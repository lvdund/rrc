package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SSB_ToMeasure_Choice_nothing uint64 = iota
	SSB_ToMeasure_Choice_ShortBitmap
	SSB_ToMeasure_Choice_MediumBitmap
	SSB_ToMeasure_Choice_LongBitmap
)

type SSB_ToMeasure struct {
	Choice       uint64
	ShortBitmap  aper.BitString `lb:4,ub:4,madatory`
	MediumBitmap aper.BitString `lb:8,ub:8,madatory`
	LongBitmap   aper.BitString `lb:64,ub:64,madatory`
}

func (ie *SSB_ToMeasure) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SSB_ToMeasure_Choice_ShortBitmap:
		if err = w.WriteBitString(ie.ShortBitmap.Bytes, uint(ie.ShortBitmap.NumBits), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode ShortBitmap", err)
		}
	case SSB_ToMeasure_Choice_MediumBitmap:
		if err = w.WriteBitString(ie.MediumBitmap.Bytes, uint(ie.MediumBitmap.NumBits), &aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode MediumBitmap", err)
		}
	case SSB_ToMeasure_Choice_LongBitmap:
		if err = w.WriteBitString(ie.LongBitmap.Bytes, uint(ie.LongBitmap.NumBits), &aper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			err = utils.WrapError("Encode LongBitmap", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SSB_ToMeasure) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SSB_ToMeasure_Choice_ShortBitmap:
		var tmp_bs_ShortBitmap []byte
		var n_ShortBitmap uint
		if tmp_bs_ShortBitmap, n_ShortBitmap, err = r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode ShortBitmap", err)
		}
		ie.ShortBitmap = aper.BitString{
			Bytes:   tmp_bs_ShortBitmap,
			NumBits: uint64(n_ShortBitmap),
		}
	case SSB_ToMeasure_Choice_MediumBitmap:
		var tmp_bs_MediumBitmap []byte
		var n_MediumBitmap uint
		if tmp_bs_MediumBitmap, n_MediumBitmap, err = r.ReadBitString(&aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode MediumBitmap", err)
		}
		ie.MediumBitmap = aper.BitString{
			Bytes:   tmp_bs_MediumBitmap,
			NumBits: uint64(n_MediumBitmap),
		}
	case SSB_ToMeasure_Choice_LongBitmap:
		var tmp_bs_LongBitmap []byte
		var n_LongBitmap uint
		if tmp_bs_LongBitmap, n_LongBitmap, err = r.ReadBitString(&aper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode LongBitmap", err)
		}
		ie.LongBitmap = aper.BitString{
			Bytes:   tmp_bs_LongBitmap,
			NumBits: uint64(n_LongBitmap),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
