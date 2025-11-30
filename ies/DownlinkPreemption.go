package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DownlinkPreemption struct {
	Int_RNTI                        RNTI_Value                          `madatory`
	TimeFrequencySet                DownlinkPreemption_timeFrequencySet `madatory`
	Dci_PayloadSize                 int64                               `lb:0,ub:maxINT_DCI_PayloadSize,madatory`
	Int_ConfigurationPerServingCell []INT_ConfigurationPerServingCell   `lb:1,ub:maxNrofServingCells,madatory`
}

func (ie *DownlinkPreemption) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Int_RNTI.Encode(w); err != nil {
		return utils.WrapError("Encode Int_RNTI", err)
	}
	if err = ie.TimeFrequencySet.Encode(w); err != nil {
		return utils.WrapError("Encode TimeFrequencySet", err)
	}
	if err = w.WriteInteger(ie.Dci_PayloadSize, &aper.Constraint{Lb: 0, Ub: maxINT_DCI_PayloadSize}, false); err != nil {
		return utils.WrapError("WriteInteger Dci_PayloadSize", err)
	}
	tmp_Int_ConfigurationPerServingCell := utils.NewSequence[*INT_ConfigurationPerServingCell]([]*INT_ConfigurationPerServingCell{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.Int_ConfigurationPerServingCell {
		tmp_Int_ConfigurationPerServingCell.Value = append(tmp_Int_ConfigurationPerServingCell.Value, &i)
	}
	if err = tmp_Int_ConfigurationPerServingCell.Encode(w); err != nil {
		return utils.WrapError("Encode Int_ConfigurationPerServingCell", err)
	}
	return nil
}

func (ie *DownlinkPreemption) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Int_RNTI.Decode(r); err != nil {
		return utils.WrapError("Decode Int_RNTI", err)
	}
	if err = ie.TimeFrequencySet.Decode(r); err != nil {
		return utils.WrapError("Decode TimeFrequencySet", err)
	}
	var tmp_int_Dci_PayloadSize int64
	if tmp_int_Dci_PayloadSize, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxINT_DCI_PayloadSize}, false); err != nil {
		return utils.WrapError("ReadInteger Dci_PayloadSize", err)
	}
	ie.Dci_PayloadSize = tmp_int_Dci_PayloadSize
	tmp_Int_ConfigurationPerServingCell := utils.NewSequence[*INT_ConfigurationPerServingCell]([]*INT_ConfigurationPerServingCell{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn_Int_ConfigurationPerServingCell := func() *INT_ConfigurationPerServingCell {
		return new(INT_ConfigurationPerServingCell)
	}
	if err = tmp_Int_ConfigurationPerServingCell.Decode(r, fn_Int_ConfigurationPerServingCell); err != nil {
		return utils.WrapError("Decode Int_ConfigurationPerServingCell", err)
	}
	ie.Int_ConfigurationPerServingCell = []INT_ConfigurationPerServingCell{}
	for _, i := range tmp_Int_ConfigurationPerServingCell.Value {
		ie.Int_ConfigurationPerServingCell = append(ie.Int_ConfigurationPerServingCell, *i)
	}
	return nil
}
