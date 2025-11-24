package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_UplinkTxSwitch_v1630 struct {
	BandCombination_v1630 *BandCombination_v1630 `optional`
}

func (ie *BandCombination_UplinkTxSwitch_v1630) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.BandCombination_v1630 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.BandCombination_v1630 != nil {
		if err = ie.BandCombination_v1630.Encode(w); err != nil {
			return utils.WrapError("Encode BandCombination_v1630", err)
		}
	}
	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1630) Decode(r *uper.UperReader) error {
	var err error
	var BandCombination_v1630Present bool
	if BandCombination_v1630Present, err = r.ReadBool(); err != nil {
		return err
	}
	if BandCombination_v1630Present {
		ie.BandCombination_v1630 = new(BandCombination_v1630)
		if err = ie.BandCombination_v1630.Decode(r); err != nil {
			return utils.WrapError("Decode BandCombination_v1630", err)
		}
	}
	return nil
}
