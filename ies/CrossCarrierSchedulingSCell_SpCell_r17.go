package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CrossCarrierSchedulingSCell_SpCell_r17 struct {
	SupportedSCS_Combinations_r17 *CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17 `lb:1,ub:496,optional`
	Pdcch_MonitoringOccasion_r17  CrossCarrierSchedulingSCell_SpCell_r17_pdcch_MonitoringOccasion_r17   `madatory`
}

func (ie *CrossCarrierSchedulingSCell_SpCell_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SupportedSCS_Combinations_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SupportedSCS_Combinations_r17 != nil {
		if err = ie.SupportedSCS_Combinations_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedSCS_Combinations_r17", err)
		}
	}
	if err = ie.Pdcch_MonitoringOccasion_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Pdcch_MonitoringOccasion_r17", err)
	}
	return nil
}

func (ie *CrossCarrierSchedulingSCell_SpCell_r17) Decode(r *aper.AperReader) error {
	var err error
	var SupportedSCS_Combinations_r17Present bool
	if SupportedSCS_Combinations_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SupportedSCS_Combinations_r17Present {
		ie.SupportedSCS_Combinations_r17 = new(CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17)
		if err = ie.SupportedSCS_Combinations_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedSCS_Combinations_r17", err)
		}
	}
	if err = ie.Pdcch_MonitoringOccasion_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Pdcch_MonitoringOccasion_r17", err)
	}
	return nil
}
