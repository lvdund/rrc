package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type QuantityConfigNR struct {
	QuantityConfigCell     QuantityConfigRS  `madatory`
	QuantityConfigRS_Index *QuantityConfigRS `optional`
}

func (ie *QuantityConfigNR) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.QuantityConfigRS_Index != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.QuantityConfigCell.Encode(w); err != nil {
		return utils.WrapError("Encode QuantityConfigCell", err)
	}
	if ie.QuantityConfigRS_Index != nil {
		if err = ie.QuantityConfigRS_Index.Encode(w); err != nil {
			return utils.WrapError("Encode QuantityConfigRS_Index", err)
		}
	}
	return nil
}

func (ie *QuantityConfigNR) Decode(r *uper.UperReader) error {
	var err error
	var QuantityConfigRS_IndexPresent bool
	if QuantityConfigRS_IndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.QuantityConfigCell.Decode(r); err != nil {
		return utils.WrapError("Decode QuantityConfigCell", err)
	}
	if QuantityConfigRS_IndexPresent {
		ie.QuantityConfigRS_Index = new(QuantityConfigRS)
		if err = ie.QuantityConfigRS_Index.Decode(r); err != nil {
			return utils.WrapError("Decode QuantityConfigRS_Index", err)
		}
	}
	return nil
}
