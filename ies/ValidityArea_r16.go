package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ValidityArea_r16 struct {
	CarrierFreq_r16      ARFCN_ValueNR     `madatory`
	ValidityCellList_r16 *ValidityCellList `optional`
}

func (ie *ValidityArea_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ValidityCellList_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.CarrierFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreq_r16", err)
	}
	if ie.ValidityCellList_r16 != nil {
		if err = ie.ValidityCellList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ValidityCellList_r16", err)
		}
	}
	return nil
}

func (ie *ValidityArea_r16) Decode(r *uper.UperReader) error {
	var err error
	var ValidityCellList_r16Present bool
	if ValidityCellList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.CarrierFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreq_r16", err)
	}
	if ValidityCellList_r16Present {
		ie.ValidityCellList_r16 = new(ValidityCellList)
		if err = ie.ValidityCellList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ValidityCellList_r16", err)
		}
	}
	return nil
}
