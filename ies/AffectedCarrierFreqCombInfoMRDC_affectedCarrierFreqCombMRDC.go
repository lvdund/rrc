package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type AffectedCarrierFreqCombInfoMRDC_affectedCarrierFreqCombMRDC struct {
	AffectedCarrierFreqCombEUTRA *AffectedCarrierFreqCombEUTRA `optional`
	AffectedCarrierFreqCombNR    AffectedCarrierFreqCombNR     `madatory`
}

func (ie *AffectedCarrierFreqCombInfoMRDC_affectedCarrierFreqCombMRDC) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.AffectedCarrierFreqCombEUTRA != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.AffectedCarrierFreqCombEUTRA != nil {
		if err = ie.AffectedCarrierFreqCombEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode AffectedCarrierFreqCombEUTRA", err)
		}
	}
	if err = ie.AffectedCarrierFreqCombNR.Encode(w); err != nil {
		return utils.WrapError("Encode AffectedCarrierFreqCombNR", err)
	}
	return nil
}

func (ie *AffectedCarrierFreqCombInfoMRDC_affectedCarrierFreqCombMRDC) Decode(r *aper.AperReader) error {
	var err error
	var AffectedCarrierFreqCombEUTRAPresent bool
	if AffectedCarrierFreqCombEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if AffectedCarrierFreqCombEUTRAPresent {
		ie.AffectedCarrierFreqCombEUTRA = new(AffectedCarrierFreqCombEUTRA)
		if err = ie.AffectedCarrierFreqCombEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode AffectedCarrierFreqCombEUTRA", err)
		}
	}
	if err = ie.AffectedCarrierFreqCombNR.Decode(r); err != nil {
		return utils.WrapError("Decode AffectedCarrierFreqCombNR", err)
	}
	return nil
}
