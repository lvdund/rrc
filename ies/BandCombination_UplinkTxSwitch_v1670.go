package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_UplinkTxSwitch_v1670 struct {
	BandCombination_v15g0 *BandCombination_v15g0 `optional`
}

func (ie *BandCombination_UplinkTxSwitch_v1670) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.BandCombination_v15g0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.BandCombination_v15g0 != nil {
		if err = ie.BandCombination_v15g0.Encode(w); err != nil {
			return utils.WrapError("Encode BandCombination_v15g0", err)
		}
	}
	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1670) Decode(r *uper.UperReader) error {
	var err error
	var BandCombination_v15g0Present bool
	if BandCombination_v15g0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if BandCombination_v15g0Present {
		ie.BandCombination_v15g0 = new(BandCombination_v15g0)
		if err = ie.BandCombination_v15g0.Decode(r); err != nil {
			return utils.WrapError("Decode BandCombination_v15g0", err)
		}
	}
	return nil
}
