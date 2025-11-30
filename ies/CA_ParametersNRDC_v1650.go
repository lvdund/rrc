package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNRDC_v1650 struct {
	SupportedCellGrouping_r16 *aper.BitString `lb:1,ub:maxCellGroupings_r16,optional`
}

func (ie *CA_ParametersNRDC_v1650) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SupportedCellGrouping_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SupportedCellGrouping_r16 != nil {
		if err = w.WriteBitString(ie.SupportedCellGrouping_r16.Bytes, uint(ie.SupportedCellGrouping_r16.NumBits), &aper.Constraint{Lb: 1, Ub: maxCellGroupings_r16}, false); err != nil {
			return utils.WrapError("Encode SupportedCellGrouping_r16", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNRDC_v1650) Decode(r *aper.AperReader) error {
	var err error
	var SupportedCellGrouping_r16Present bool
	if SupportedCellGrouping_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SupportedCellGrouping_r16Present {
		var tmp_bs_SupportedCellGrouping_r16 []byte
		var n_SupportedCellGrouping_r16 uint
		if tmp_bs_SupportedCellGrouping_r16, n_SupportedCellGrouping_r16, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: maxCellGroupings_r16}, false); err != nil {
			return utils.WrapError("Decode SupportedCellGrouping_r16", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_SupportedCellGrouping_r16,
			NumBits: uint64(n_SupportedCellGrouping_r16),
		}
		ie.SupportedCellGrouping_r16 = &tmp_bitstring
	}
	return nil
}
