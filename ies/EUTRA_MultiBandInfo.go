package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_MultiBandInfo struct {
	Eutra_FreqBandIndicator FreqBandIndicatorEUTRA `madatory`
	Eutra_NS_PmaxList       *EUTRA_NS_PmaxList     `optional`
}

func (ie *EUTRA_MultiBandInfo) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Eutra_NS_PmaxList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Eutra_FreqBandIndicator.Encode(w); err != nil {
		return utils.WrapError("Encode Eutra_FreqBandIndicator", err)
	}
	if ie.Eutra_NS_PmaxList != nil {
		if err = ie.Eutra_NS_PmaxList.Encode(w); err != nil {
			return utils.WrapError("Encode Eutra_NS_PmaxList", err)
		}
	}
	return nil
}

func (ie *EUTRA_MultiBandInfo) Decode(r *aper.AperReader) error {
	var err error
	var Eutra_NS_PmaxListPresent bool
	if Eutra_NS_PmaxListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Eutra_FreqBandIndicator.Decode(r); err != nil {
		return utils.WrapError("Decode Eutra_FreqBandIndicator", err)
	}
	if Eutra_NS_PmaxListPresent {
		ie.Eutra_NS_PmaxList = new(EUTRA_NS_PmaxList)
		if err = ie.Eutra_NS_PmaxList.Decode(r); err != nil {
			return utils.WrapError("Decode Eutra_NS_PmaxList", err)
		}
	}
	return nil
}
