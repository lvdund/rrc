package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type IntraFreqExcludedCellList struct {
	Value []PCI_Range `lb:1,ub:maxCellExcluded,madatory`
}

func (ie *IntraFreqExcludedCellList) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*PCI_Range]([]*PCI_Range{}, aper.Constraint{Lb: 1, Ub: maxCellExcluded}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode IntraFreqExcludedCellList", err)
	}
	return nil
}

func (ie *IntraFreqExcludedCellList) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*PCI_Range]([]*PCI_Range{}, aper.Constraint{Lb: 1, Ub: maxCellExcluded}, false)
	fn := func() *PCI_Range {
		return new(PCI_Range)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode IntraFreqExcludedCellList", err)
	}
	ie.Value = []PCI_Range{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
