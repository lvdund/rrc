package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_SDT_Configuration_r17_sdt_SSB_Subset_r17_Choice_nothing uint64 = iota
	CG_SDT_Configuration_r17_sdt_SSB_Subset_r17_Choice_ShortBitmap_r17
	CG_SDT_Configuration_r17_sdt_SSB_Subset_r17_Choice_MediumBitmap_r17
	CG_SDT_Configuration_r17_sdt_SSB_Subset_r17_Choice_LongBitmap_r17
)

type CG_SDT_Configuration_r17_sdt_SSB_Subset_r17 struct {
	Choice           uint64
	ShortBitmap_r17  uper.BitString `lb:4,ub:4,madatory`
	MediumBitmap_r17 uper.BitString `lb:8,ub:8,madatory`
	LongBitmap_r17   uper.BitString `lb:64,ub:64,madatory`
}

func (ie *CG_SDT_Configuration_r17_sdt_SSB_Subset_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_SDT_Configuration_r17_sdt_SSB_Subset_r17_Choice_ShortBitmap_r17:
		if err = w.WriteBitString(ie.ShortBitmap_r17.Bytes, uint(ie.ShortBitmap_r17.NumBits), &uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode ShortBitmap_r17", err)
		}
	case CG_SDT_Configuration_r17_sdt_SSB_Subset_r17_Choice_MediumBitmap_r17:
		if err = w.WriteBitString(ie.MediumBitmap_r17.Bytes, uint(ie.MediumBitmap_r17.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode MediumBitmap_r17", err)
		}
	case CG_SDT_Configuration_r17_sdt_SSB_Subset_r17_Choice_LongBitmap_r17:
		if err = w.WriteBitString(ie.LongBitmap_r17.Bytes, uint(ie.LongBitmap_r17.NumBits), &uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			err = utils.WrapError("Encode LongBitmap_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CG_SDT_Configuration_r17_sdt_SSB_Subset_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_SDT_Configuration_r17_sdt_SSB_Subset_r17_Choice_ShortBitmap_r17:
		var tmp_bs_ShortBitmap_r17 []byte
		var n_ShortBitmap_r17 uint
		if tmp_bs_ShortBitmap_r17, n_ShortBitmap_r17, err = r.ReadBitString(&uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode ShortBitmap_r17", err)
		}
		ie.ShortBitmap_r17 = uper.BitString{
			Bytes:   tmp_bs_ShortBitmap_r17,
			NumBits: uint64(n_ShortBitmap_r17),
		}
	case CG_SDT_Configuration_r17_sdt_SSB_Subset_r17_Choice_MediumBitmap_r17:
		var tmp_bs_MediumBitmap_r17 []byte
		var n_MediumBitmap_r17 uint
		if tmp_bs_MediumBitmap_r17, n_MediumBitmap_r17, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode MediumBitmap_r17", err)
		}
		ie.MediumBitmap_r17 = uper.BitString{
			Bytes:   tmp_bs_MediumBitmap_r17,
			NumBits: uint64(n_MediumBitmap_r17),
		}
	case CG_SDT_Configuration_r17_sdt_SSB_Subset_r17_Choice_LongBitmap_r17:
		var tmp_bs_LongBitmap_r17 []byte
		var n_LongBitmap_r17 uint
		if tmp_bs_LongBitmap_r17, n_LongBitmap_r17, err = r.ReadBitString(&uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode LongBitmap_r17", err)
		}
		ie.LongBitmap_r17 = uper.BitString{
			Bytes:   tmp_bs_LongBitmap_r17,
			NumBits: uint64(n_LongBitmap_r17),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
