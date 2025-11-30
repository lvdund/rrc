package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NR_MultiBandInfo struct {
	FreqBandIndicatorNR *FreqBandIndicatorNR `optional`
	Nr_NS_PmaxList      *NR_NS_PmaxList      `optional`
}

func (ie *NR_MultiBandInfo) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.FreqBandIndicatorNR != nil, ie.Nr_NS_PmaxList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FreqBandIndicatorNR != nil {
		if err = ie.FreqBandIndicatorNR.Encode(w); err != nil {
			return utils.WrapError("Encode FreqBandIndicatorNR", err)
		}
	}
	if ie.Nr_NS_PmaxList != nil {
		if err = ie.Nr_NS_PmaxList.Encode(w); err != nil {
			return utils.WrapError("Encode Nr_NS_PmaxList", err)
		}
	}
	return nil
}

func (ie *NR_MultiBandInfo) Decode(r *aper.AperReader) error {
	var err error
	var FreqBandIndicatorNRPresent bool
	if FreqBandIndicatorNRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Nr_NS_PmaxListPresent bool
	if Nr_NS_PmaxListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if FreqBandIndicatorNRPresent {
		ie.FreqBandIndicatorNR = new(FreqBandIndicatorNR)
		if err = ie.FreqBandIndicatorNR.Decode(r); err != nil {
			return utils.WrapError("Decode FreqBandIndicatorNR", err)
		}
	}
	if Nr_NS_PmaxListPresent {
		ie.Nr_NS_PmaxList = new(NR_NS_PmaxList)
		if err = ie.Nr_NS_PmaxList.Decode(r); err != nil {
			return utils.WrapError("Decode Nr_NS_PmaxList", err)
		}
	}
	return nil
}
