package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RF_ParametersMRDC_supportedBandCombinationListNEDC_Only_v1720 struct {
	SupportedBandCombinationList_v1700 *BandCombinationList_v1700 `optional`
	SupportedBandCombinationList_v1720 *BandCombinationList_v1720 `optional`
}

func (ie *RF_ParametersMRDC_supportedBandCombinationListNEDC_Only_v1720) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SupportedBandCombinationList_v1700 != nil, ie.SupportedBandCombinationList_v1720 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SupportedBandCombinationList_v1700 != nil {
		if err = ie.SupportedBandCombinationList_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedBandCombinationList_v1700", err)
		}
	}
	if ie.SupportedBandCombinationList_v1720 != nil {
		if err = ie.SupportedBandCombinationList_v1720.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedBandCombinationList_v1720", err)
		}
	}
	return nil
}

func (ie *RF_ParametersMRDC_supportedBandCombinationListNEDC_Only_v1720) Decode(r *aper.AperReader) error {
	var err error
	var SupportedBandCombinationList_v1700Present bool
	if SupportedBandCombinationList_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedBandCombinationList_v1720Present bool
	if SupportedBandCombinationList_v1720Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SupportedBandCombinationList_v1700Present {
		ie.SupportedBandCombinationList_v1700 = new(BandCombinationList_v1700)
		if err = ie.SupportedBandCombinationList_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedBandCombinationList_v1700", err)
		}
	}
	if SupportedBandCombinationList_v1720Present {
		ie.SupportedBandCombinationList_v1720 = new(BandCombinationList_v1720)
		if err = ie.SupportedBandCombinationList_v1720.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedBandCombinationList_v1720", err)
		}
	}
	return nil
}
