package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_UplinkTxSwitch_v1650 struct {
	BandCombination_v1650 *BandCombination_v1650 `optional`
}

func (ie *BandCombination_UplinkTxSwitch_v1650) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.BandCombination_v1650 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.BandCombination_v1650 != nil {
		if err = ie.BandCombination_v1650.Encode(w); err != nil {
			return utils.WrapError("Encode BandCombination_v1650", err)
		}
	}
	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1650) Decode(r *aper.AperReader) error {
	var err error
	var BandCombination_v1650Present bool
	if BandCombination_v1650Present, err = r.ReadBool(); err != nil {
		return err
	}
	if BandCombination_v1650Present {
		ie.BandCombination_v1650 = new(BandCombination_v1650)
		if err = ie.BandCombination_v1650.Decode(r); err != nil {
			return utils.WrapError("Decode BandCombination_v1650", err)
		}
	}
	return nil
}
