package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SSB_MTC_AdditionalPCI_r17 struct {
	AdditionalPCIIndex_r17   AdditionalPCIIndex_r17                             `madatory`
	AdditionalPCI_r17        PhysCellId                                         `madatory`
	Periodicity_r17          SSB_MTC_AdditionalPCI_r17_periodicity_r17          `madatory`
	Ssb_PositionsInBurst_r17 SSB_MTC_AdditionalPCI_r17_ssb_PositionsInBurst_r17 `lb:4,ub:4,madatory`
	Ss_PBCH_BlockPower_r17   int64                                              `lb:-60,ub:50,madatory`
}

func (ie *SSB_MTC_AdditionalPCI_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.AdditionalPCIIndex_r17.Encode(w); err != nil {
		return utils.WrapError("Encode AdditionalPCIIndex_r17", err)
	}
	if err = ie.AdditionalPCI_r17.Encode(w); err != nil {
		return utils.WrapError("Encode AdditionalPCI_r17", err)
	}
	if err = ie.Periodicity_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Periodicity_r17", err)
	}
	if err = ie.Ssb_PositionsInBurst_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Ssb_PositionsInBurst_r17", err)
	}
	if err = w.WriteInteger(ie.Ss_PBCH_BlockPower_r17, &aper.Constraint{Lb: -60, Ub: 50}, false); err != nil {
		return utils.WrapError("WriteInteger Ss_PBCH_BlockPower_r17", err)
	}
	return nil
}

func (ie *SSB_MTC_AdditionalPCI_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.AdditionalPCIIndex_r17.Decode(r); err != nil {
		return utils.WrapError("Decode AdditionalPCIIndex_r17", err)
	}
	if err = ie.AdditionalPCI_r17.Decode(r); err != nil {
		return utils.WrapError("Decode AdditionalPCI_r17", err)
	}
	if err = ie.Periodicity_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Periodicity_r17", err)
	}
	if err = ie.Ssb_PositionsInBurst_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Ssb_PositionsInBurst_r17", err)
	}
	var tmp_int_Ss_PBCH_BlockPower_r17 int64
	if tmp_int_Ss_PBCH_BlockPower_r17, err = r.ReadInteger(&aper.Constraint{Lb: -60, Ub: 50}, false); err != nil {
		return utils.WrapError("ReadInteger Ss_PBCH_BlockPower_r17", err)
	}
	ie.Ss_PBCH_BlockPower_r17 = tmp_int_Ss_PBCH_BlockPower_r17
	return nil
}
