package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FreqPriorityEUTRA struct {
	CarrierFreq                ARFCN_ValueEUTRA            `madatory`
	CellReselectionPriority    CellReselectionPriority     `madatory`
	CellReselectionSubPriority *CellReselectionSubPriority `optional`
}

func (ie *FreqPriorityEUTRA) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.CellReselectionSubPriority != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.CarrierFreq.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreq", err)
	}
	if err = ie.CellReselectionPriority.Encode(w); err != nil {
		return utils.WrapError("Encode CellReselectionPriority", err)
	}
	if ie.CellReselectionSubPriority != nil {
		if err = ie.CellReselectionSubPriority.Encode(w); err != nil {
			return utils.WrapError("Encode CellReselectionSubPriority", err)
		}
	}
	return nil
}

func (ie *FreqPriorityEUTRA) Decode(r *aper.AperReader) error {
	var err error
	var CellReselectionSubPriorityPresent bool
	if CellReselectionSubPriorityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.CarrierFreq.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreq", err)
	}
	if err = ie.CellReselectionPriority.Decode(r); err != nil {
		return utils.WrapError("Decode CellReselectionPriority", err)
	}
	if CellReselectionSubPriorityPresent {
		ie.CellReselectionSubPriority = new(CellReselectionSubPriority)
		if err = ie.CellReselectionSubPriority.Decode(r); err != nil {
			return utils.WrapError("Decode CellReselectionSubPriority", err)
		}
	}
	return nil
}
