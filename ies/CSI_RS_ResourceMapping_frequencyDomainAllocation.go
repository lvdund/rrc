package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_nothing uint64 = iota
	CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_Row1
	CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_Row2
	CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_Row4
	CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_Other
)

type CSI_RS_ResourceMapping_frequencyDomainAllocation struct {
	Choice uint64
	Row1   aper.BitString `lb:4,ub:4,madatory`
	Row2   aper.BitString `lb:12,ub:12,madatory`
	Row4   aper.BitString `lb:3,ub:3,madatory`
	Other  aper.BitString `lb:6,ub:6,madatory`
}

func (ie *CSI_RS_ResourceMapping_frequencyDomainAllocation) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_Row1:
		if err = w.WriteBitString(ie.Row1.Bytes, uint(ie.Row1.NumBits), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode Row1", err)
		}
	case CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_Row2:
		if err = w.WriteBitString(ie.Row2.Bytes, uint(ie.Row2.NumBits), &aper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
			err = utils.WrapError("Encode Row2", err)
		}
	case CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_Row4:
		if err = w.WriteBitString(ie.Row4.Bytes, uint(ie.Row4.NumBits), &aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode Row4", err)
		}
	case CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_Other:
		if err = w.WriteBitString(ie.Other.Bytes, uint(ie.Other.NumBits), &aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			err = utils.WrapError("Encode Other", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_RS_ResourceMapping_frequencyDomainAllocation) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_Row1:
		var tmp_bs_Row1 []byte
		var n_Row1 uint
		if tmp_bs_Row1, n_Row1, err = r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode Row1", err)
		}
		ie.Row1 = aper.BitString{
			Bytes:   tmp_bs_Row1,
			NumBits: uint64(n_Row1),
		}
	case CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_Row2:
		var tmp_bs_Row2 []byte
		var n_Row2 uint
		if tmp_bs_Row2, n_Row2, err = r.ReadBitString(&aper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
			return utils.WrapError("Decode Row2", err)
		}
		ie.Row2 = aper.BitString{
			Bytes:   tmp_bs_Row2,
			NumBits: uint64(n_Row2),
		}
	case CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_Row4:
		var tmp_bs_Row4 []byte
		var n_Row4 uint
		if tmp_bs_Row4, n_Row4, err = r.ReadBitString(&aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode Row4", err)
		}
		ie.Row4 = aper.BitString{
			Bytes:   tmp_bs_Row4,
			NumBits: uint64(n_Row4),
		}
	case CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_Other:
		var tmp_bs_Other []byte
		var n_Other uint
		if tmp_bs_Other, n_Other, err = r.ReadBitString(&aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode Other", err)
		}
		ie.Other = aper.BitString{
			Bytes:   tmp_bs_Other,
			NumBits: uint64(n_Other),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
