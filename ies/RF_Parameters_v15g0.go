package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RF_Parameters_v15g0 struct {
	SupportedBandCombinationList_v15g0 *BandCombinationList_v15g0 `optional`
}

func (ie *RF_Parameters_v15g0) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SupportedBandCombinationList_v15g0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SupportedBandCombinationList_v15g0 != nil {
		if err = ie.SupportedBandCombinationList_v15g0.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedBandCombinationList_v15g0", err)
		}
	}
	return nil
}

func (ie *RF_Parameters_v15g0) Decode(r *aper.AperReader) error {
	var err error
	var SupportedBandCombinationList_v15g0Present bool
	if SupportedBandCombinationList_v15g0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SupportedBandCombinationList_v15g0Present {
		ie.SupportedBandCombinationList_v15g0 = new(BandCombinationList_v15g0)
		if err = ie.SupportedBandCombinationList_v15g0.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedBandCombinationList_v15g0", err)
		}
	}
	return nil
}
