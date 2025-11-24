package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConfigRestrictInfoSCG_servCellIndexRangeSCG struct {
	LowBound ServCellIndex `madatory`
	UpBound  ServCellIndex `madatory`
}

func (ie *ConfigRestrictInfoSCG_servCellIndexRangeSCG) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.LowBound.Encode(w); err != nil {
		return utils.WrapError("Encode LowBound", err)
	}
	if err = ie.UpBound.Encode(w); err != nil {
		return utils.WrapError("Encode UpBound", err)
	}
	return nil
}

func (ie *ConfigRestrictInfoSCG_servCellIndexRangeSCG) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.LowBound.Decode(r); err != nil {
		return utils.WrapError("Decode LowBound", err)
	}
	if err = ie.UpBound.Decode(r); err != nil {
		return utils.WrapError("Decode UpBound", err)
	}
	return nil
}
