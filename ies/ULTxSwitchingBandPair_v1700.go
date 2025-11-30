package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ULTxSwitchingBandPair_v1700 struct {
	UplinkTxSwitchingPeriod2T2T_r17 *ULTxSwitchingBandPair_v1700_uplinkTxSwitchingPeriod2T2T_r17 `optional`
}

func (ie *ULTxSwitchingBandPair_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.UplinkTxSwitchingPeriod2T2T_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.UplinkTxSwitchingPeriod2T2T_r17 != nil {
		if err = ie.UplinkTxSwitchingPeriod2T2T_r17.Encode(w); err != nil {
			return utils.WrapError("Encode UplinkTxSwitchingPeriod2T2T_r17", err)
		}
	}
	return nil
}

func (ie *ULTxSwitchingBandPair_v1700) Decode(r *aper.AperReader) error {
	var err error
	var UplinkTxSwitchingPeriod2T2T_r17Present bool
	if UplinkTxSwitchingPeriod2T2T_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if UplinkTxSwitchingPeriod2T2T_r17Present {
		ie.UplinkTxSwitchingPeriod2T2T_r17 = new(ULTxSwitchingBandPair_v1700_uplinkTxSwitchingPeriod2T2T_r17)
		if err = ie.UplinkTxSwitchingPeriod2T2T_r17.Decode(r); err != nil {
			return utils.WrapError("Decode UplinkTxSwitchingPeriod2T2T_r17", err)
		}
	}
	return nil
}
