package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MBS_NeighbourCell_r17 struct {
	PhysCellId_r17  PhysCellId     `madatory`
	CarrierFreq_r17 *ARFCN_ValueNR `optional`
}

func (ie *MBS_NeighbourCell_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CarrierFreq_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.PhysCellId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode PhysCellId_r17", err)
	}
	if ie.CarrierFreq_r17 != nil {
		if err = ie.CarrierFreq_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CarrierFreq_r17", err)
		}
	}
	return nil
}

func (ie *MBS_NeighbourCell_r17) Decode(r *uper.UperReader) error {
	var err error
	var CarrierFreq_r17Present bool
	if CarrierFreq_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.PhysCellId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode PhysCellId_r17", err)
	}
	if CarrierFreq_r17Present {
		ie.CarrierFreq_r17 = new(ARFCN_ValueNR)
		if err = ie.CarrierFreq_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CarrierFreq_r17", err)
		}
	}
	return nil
}
