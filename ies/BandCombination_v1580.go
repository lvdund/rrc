package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1580 struct {
	Mrdc_Parameters_v1580 MRDC_Parameters_v1580 `madatory`
}

func (ie *BandCombination_v1580) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Mrdc_Parameters_v1580.Encode(w); err != nil {
		return utils.WrapError("Encode Mrdc_Parameters_v1580", err)
	}
	return nil
}

func (ie *BandCombination_v1580) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Mrdc_Parameters_v1580.Decode(r); err != nil {
		return utils.WrapError("Decode Mrdc_Parameters_v1580", err)
	}
	return nil
}
