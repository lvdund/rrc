package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_ppw_durationOfPRS_Processing2_r17 struct {
	Ppw_durationOfPRS_ProcessingSymbolsN2_r17 PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_ppw_durationOfPRS_Processing2_r17_ppw_durationOfPRS_ProcessingSymbolsN2_r17 `madatory`
	Ppw_durationOfPRS_ProcessingSymbolsT2_r17 PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_ppw_durationOfPRS_Processing2_r17_ppw_durationOfPRS_ProcessingSymbolsT2_r17 `madatory`
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_ppw_durationOfPRS_Processing2_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Ppw_durationOfPRS_ProcessingSymbolsN2_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Ppw_durationOfPRS_ProcessingSymbolsN2_r17", err)
	}
	if err = ie.Ppw_durationOfPRS_ProcessingSymbolsT2_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Ppw_durationOfPRS_ProcessingSymbolsT2_r17", err)
	}
	return nil
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17_ppw_durationOfPRS_Processing2_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Ppw_durationOfPRS_ProcessingSymbolsN2_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Ppw_durationOfPRS_ProcessingSymbolsN2_r17", err)
	}
	if err = ie.Ppw_durationOfPRS_ProcessingSymbolsT2_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Ppw_durationOfPRS_ProcessingSymbolsT2_r17", err)
	}
	return nil
}
