package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ReestabNCellInfoList struct {
	Value []ReestabNCellInfo `lb:1,ub:maxCellPrep,madatory`
}

func (ie *ReestabNCellInfoList) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*ReestabNCellInfo]([]*ReestabNCellInfo{}, aper.Constraint{Lb: 1, Ub: maxCellPrep}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode ReestabNCellInfoList", err)
	}
	return nil
}

func (ie *ReestabNCellInfoList) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*ReestabNCellInfo]([]*ReestabNCellInfo{}, aper.Constraint{Lb: 1, Ub: maxCellPrep}, false)
	fn := func() *ReestabNCellInfo {
		return new(ReestabNCellInfo)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode ReestabNCellInfoList", err)
	}
	ie.Value = []ReestabNCellInfo{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
