package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_nothing uint64 = iota
	RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N2
	RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N4
	RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N5
	RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N8
	RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N10
	RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N20
	RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N40
)

type RateMatchPattern_patternType_bitmaps_periodicityAndPattern struct {
	Choice uint64
	N2     uper.BitString `lb:2,ub:2,madatory`
	N4     uper.BitString `lb:4,ub:4,madatory`
	N5     uper.BitString `lb:5,ub:5,madatory`
	N8     uper.BitString `lb:8,ub:8,madatory`
	N10    uper.BitString `lb:10,ub:10,madatory`
	N20    uper.BitString `lb:20,ub:20,madatory`
	N40    uper.BitString `lb:40,ub:40,madatory`
}

func (ie *RateMatchPattern_patternType_bitmaps_periodicityAndPattern) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 7, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N2:
		if err = w.WriteBitString(ie.N2.Bytes, uint(ie.N2.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			err = utils.WrapError("Encode N2", err)
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N4:
		if err = w.WriteBitString(ie.N4.Bytes, uint(ie.N4.NumBits), &uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode N4", err)
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N5:
		if err = w.WriteBitString(ie.N5.Bytes, uint(ie.N5.NumBits), &uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
			err = utils.WrapError("Encode N5", err)
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N8:
		if err = w.WriteBitString(ie.N8.Bytes, uint(ie.N8.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode N8", err)
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N10:
		if err = w.WriteBitString(ie.N10.Bytes, uint(ie.N10.NumBits), &uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			err = utils.WrapError("Encode N10", err)
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N20:
		if err = w.WriteBitString(ie.N20.Bytes, uint(ie.N20.NumBits), &uper.Constraint{Lb: 20, Ub: 20}, false); err != nil {
			err = utils.WrapError("Encode N20", err)
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N40:
		if err = w.WriteBitString(ie.N40.Bytes, uint(ie.N40.NumBits), &uper.Constraint{Lb: 40, Ub: 40}, false); err != nil {
			err = utils.WrapError("Encode N40", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RateMatchPattern_patternType_bitmaps_periodicityAndPattern) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(7, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N2:
		var tmp_bs_N2 []byte
		var n_N2 uint
		if tmp_bs_N2, n_N2, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode N2", err)
		}
		ie.N2 = uper.BitString{
			Bytes:   tmp_bs_N2,
			NumBits: uint64(n_N2),
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N4:
		var tmp_bs_N4 []byte
		var n_N4 uint
		if tmp_bs_N4, n_N4, err = r.ReadBitString(&uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode N4", err)
		}
		ie.N4 = uper.BitString{
			Bytes:   tmp_bs_N4,
			NumBits: uint64(n_N4),
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N5:
		var tmp_bs_N5 []byte
		var n_N5 uint
		if tmp_bs_N5, n_N5, err = r.ReadBitString(&uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
			return utils.WrapError("Decode N5", err)
		}
		ie.N5 = uper.BitString{
			Bytes:   tmp_bs_N5,
			NumBits: uint64(n_N5),
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N8:
		var tmp_bs_N8 []byte
		var n_N8 uint
		if tmp_bs_N8, n_N8, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode N8", err)
		}
		ie.N8 = uper.BitString{
			Bytes:   tmp_bs_N8,
			NumBits: uint64(n_N8),
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N10:
		var tmp_bs_N10 []byte
		var n_N10 uint
		if tmp_bs_N10, n_N10, err = r.ReadBitString(&uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			return utils.WrapError("Decode N10", err)
		}
		ie.N10 = uper.BitString{
			Bytes:   tmp_bs_N10,
			NumBits: uint64(n_N10),
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N20:
		var tmp_bs_N20 []byte
		var n_N20 uint
		if tmp_bs_N20, n_N20, err = r.ReadBitString(&uper.Constraint{Lb: 20, Ub: 20}, false); err != nil {
			return utils.WrapError("Decode N20", err)
		}
		ie.N20 = uper.BitString{
			Bytes:   tmp_bs_N20,
			NumBits: uint64(n_N20),
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_N40:
		var tmp_bs_N40 []byte
		var n_N40 uint
		if tmp_bs_N40, n_N40, err = r.ReadBitString(&uper.Constraint{Lb: 40, Ub: 40}, false); err != nil {
			return utils.WrapError("Decode N40", err)
		}
		ie.N40 = uper.BitString{
			Bytes:   tmp_bs_N40,
			NumBits: uint64(n_N40),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
