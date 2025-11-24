package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ULTxSwitchingBandPair_r16 struct {
	BandIndexUL1_r16                      int64                                                 `lb:1,ub:maxSimultaneousBands,madatory`
	BandIndexUL2_r16                      int64                                                 `lb:1,ub:maxSimultaneousBands,madatory`
	UplinkTxSwitchingPeriod_r16           ULTxSwitchingBandPair_r16_uplinkTxSwitchingPeriod_r16 `madatory`
	UplinkTxSwitching_DL_Interruption_r16 *uper.BitString                                       `lb:1,ub:maxSimultaneousBands,optional`
}

func (ie *ULTxSwitchingBandPair_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.UplinkTxSwitching_DL_Interruption_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.BandIndexUL1_r16, &uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false); err != nil {
		return utils.WrapError("WriteInteger BandIndexUL1_r16", err)
	}
	if err = w.WriteInteger(ie.BandIndexUL2_r16, &uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false); err != nil {
		return utils.WrapError("WriteInteger BandIndexUL2_r16", err)
	}
	if err = ie.UplinkTxSwitchingPeriod_r16.Encode(w); err != nil {
		return utils.WrapError("Encode UplinkTxSwitchingPeriod_r16", err)
	}
	if ie.UplinkTxSwitching_DL_Interruption_r16 != nil {
		if err = w.WriteBitString(ie.UplinkTxSwitching_DL_Interruption_r16.Bytes, uint(ie.UplinkTxSwitching_DL_Interruption_r16.NumBits), &uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false); err != nil {
			return utils.WrapError("Encode UplinkTxSwitching_DL_Interruption_r16", err)
		}
	}
	return nil
}

func (ie *ULTxSwitchingBandPair_r16) Decode(r *uper.UperReader) error {
	var err error
	var UplinkTxSwitching_DL_Interruption_r16Present bool
	if UplinkTxSwitching_DL_Interruption_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_BandIndexUL1_r16 int64
	if tmp_int_BandIndexUL1_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false); err != nil {
		return utils.WrapError("ReadInteger BandIndexUL1_r16", err)
	}
	ie.BandIndexUL1_r16 = tmp_int_BandIndexUL1_r16
	var tmp_int_BandIndexUL2_r16 int64
	if tmp_int_BandIndexUL2_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false); err != nil {
		return utils.WrapError("ReadInteger BandIndexUL2_r16", err)
	}
	ie.BandIndexUL2_r16 = tmp_int_BandIndexUL2_r16
	if err = ie.UplinkTxSwitchingPeriod_r16.Decode(r); err != nil {
		return utils.WrapError("Decode UplinkTxSwitchingPeriod_r16", err)
	}
	if UplinkTxSwitching_DL_Interruption_r16Present {
		var tmp_bs_UplinkTxSwitching_DL_Interruption_r16 []byte
		var n_UplinkTxSwitching_DL_Interruption_r16 uint
		if tmp_bs_UplinkTxSwitching_DL_Interruption_r16, n_UplinkTxSwitching_DL_Interruption_r16, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false); err != nil {
			return utils.WrapError("Decode UplinkTxSwitching_DL_Interruption_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_UplinkTxSwitching_DL_Interruption_r16,
			NumBits: uint64(n_UplinkTxSwitching_DL_Interruption_r16),
		}
		ie.UplinkTxSwitching_DL_Interruption_r16 = &tmp_bitstring
	}
	return nil
}
