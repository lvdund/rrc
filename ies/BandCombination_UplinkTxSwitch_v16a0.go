package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_UplinkTxSwitch_v16a0 struct {
	BandCombination_v16a0 *BandCombination_v16a0 `optional`
}

func (ie *BandCombination_UplinkTxSwitch_v16a0) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.BandCombination_v16a0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.BandCombination_v16a0 != nil {
		if err = ie.BandCombination_v16a0.Encode(w); err != nil {
			return utils.WrapError("Encode BandCombination_v16a0", err)
		}
	}
	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v16a0) Decode(r *aper.AperReader) error {
	var err error
	var BandCombination_v16a0Present bool
	if BandCombination_v16a0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if BandCombination_v16a0Present {
		ie.BandCombination_v16a0 = new(BandCombination_v16a0)
		if err = ie.BandCombination_v16a0.Decode(r); err != nil {
			return utils.WrapError("Decode BandCombination_v16a0", err)
		}
	}
	return nil
}
