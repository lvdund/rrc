package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RF_Parameters_v16a0 struct {
	SupportedBandCombinationList_v16a0                *BandCombinationList_v16a0                `optional`
	SupportedBandCombinationList_UplinkTxSwitch_v16a0 *BandCombinationList_UplinkTxSwitch_v16a0 `optional`
}

func (ie *RF_Parameters_v16a0) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SupportedBandCombinationList_v16a0 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v16a0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SupportedBandCombinationList_v16a0 != nil {
		if err = ie.SupportedBandCombinationList_v16a0.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedBandCombinationList_v16a0", err)
		}
	}
	if ie.SupportedBandCombinationList_UplinkTxSwitch_v16a0 != nil {
		if err = ie.SupportedBandCombinationList_UplinkTxSwitch_v16a0.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedBandCombinationList_UplinkTxSwitch_v16a0", err)
		}
	}
	return nil
}

func (ie *RF_Parameters_v16a0) Decode(r *aper.AperReader) error {
	var err error
	var SupportedBandCombinationList_v16a0Present bool
	if SupportedBandCombinationList_v16a0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedBandCombinationList_UplinkTxSwitch_v16a0Present bool
	if SupportedBandCombinationList_UplinkTxSwitch_v16a0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SupportedBandCombinationList_v16a0Present {
		ie.SupportedBandCombinationList_v16a0 = new(BandCombinationList_v16a0)
		if err = ie.SupportedBandCombinationList_v16a0.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedBandCombinationList_v16a0", err)
		}
	}
	if SupportedBandCombinationList_UplinkTxSwitch_v16a0Present {
		ie.SupportedBandCombinationList_UplinkTxSwitch_v16a0 = new(BandCombinationList_UplinkTxSwitch_v16a0)
		if err = ie.SupportedBandCombinationList_UplinkTxSwitch_v16a0.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedBandCombinationList_UplinkTxSwitch_v16a0", err)
		}
	}
	return nil
}
