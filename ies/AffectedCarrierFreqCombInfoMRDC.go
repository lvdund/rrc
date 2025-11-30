package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type AffectedCarrierFreqCombInfoMRDC struct {
	VictimSystemType            VictimSystemType                                             `madatory`
	InterferenceDirectionMRDC   AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC    `madatory`
	AffectedCarrierFreqCombMRDC *AffectedCarrierFreqCombInfoMRDC_affectedCarrierFreqCombMRDC `optional`
}

func (ie *AffectedCarrierFreqCombInfoMRDC) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.AffectedCarrierFreqCombMRDC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.VictimSystemType.Encode(w); err != nil {
		return utils.WrapError("Encode VictimSystemType", err)
	}
	if err = ie.InterferenceDirectionMRDC.Encode(w); err != nil {
		return utils.WrapError("Encode InterferenceDirectionMRDC", err)
	}
	if ie.AffectedCarrierFreqCombMRDC != nil {
		if err = ie.AffectedCarrierFreqCombMRDC.Encode(w); err != nil {
			return utils.WrapError("Encode AffectedCarrierFreqCombMRDC", err)
		}
	}
	return nil
}

func (ie *AffectedCarrierFreqCombInfoMRDC) Decode(r *aper.AperReader) error {
	var err error
	var AffectedCarrierFreqCombMRDCPresent bool
	if AffectedCarrierFreqCombMRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.VictimSystemType.Decode(r); err != nil {
		return utils.WrapError("Decode VictimSystemType", err)
	}
	if err = ie.InterferenceDirectionMRDC.Decode(r); err != nil {
		return utils.WrapError("Decode InterferenceDirectionMRDC", err)
	}
	if AffectedCarrierFreqCombMRDCPresent {
		ie.AffectedCarrierFreqCombMRDC = new(AffectedCarrierFreqCombInfoMRDC_affectedCarrierFreqCombMRDC)
		if err = ie.AffectedCarrierFreqCombMRDC.Decode(r); err != nil {
			return utils.WrapError("Decode AffectedCarrierFreqCombMRDC", err)
		}
	}
	return nil
}
