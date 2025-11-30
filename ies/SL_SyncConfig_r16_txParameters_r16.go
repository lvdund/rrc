package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_SyncConfig_r16_txParameters_r16 struct {
	SyncTxThreshIC_r16   *SL_RSRP_Range_r16 `optional`
	SyncTxThreshOoC_r16  *SL_RSRP_Range_r16 `optional`
	SyncInfoReserved_r16 *aper.BitString    `lb:2,ub:2,optional`
}

func (ie *SL_SyncConfig_r16_txParameters_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SyncTxThreshIC_r16 != nil, ie.SyncTxThreshOoC_r16 != nil, ie.SyncInfoReserved_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SyncTxThreshIC_r16 != nil {
		if err = ie.SyncTxThreshIC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SyncTxThreshIC_r16", err)
		}
	}
	if ie.SyncTxThreshOoC_r16 != nil {
		if err = ie.SyncTxThreshOoC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SyncTxThreshOoC_r16", err)
		}
	}
	if ie.SyncInfoReserved_r16 != nil {
		if err = w.WriteBitString(ie.SyncInfoReserved_r16.Bytes, uint(ie.SyncInfoReserved_r16.NumBits), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode SyncInfoReserved_r16", err)
		}
	}
	return nil
}

func (ie *SL_SyncConfig_r16_txParameters_r16) Decode(r *aper.AperReader) error {
	var err error
	var SyncTxThreshIC_r16Present bool
	if SyncTxThreshIC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SyncTxThreshOoC_r16Present bool
	if SyncTxThreshOoC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SyncInfoReserved_r16Present bool
	if SyncInfoReserved_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SyncTxThreshIC_r16Present {
		ie.SyncTxThreshIC_r16 = new(SL_RSRP_Range_r16)
		if err = ie.SyncTxThreshIC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SyncTxThreshIC_r16", err)
		}
	}
	if SyncTxThreshOoC_r16Present {
		ie.SyncTxThreshOoC_r16 = new(SL_RSRP_Range_r16)
		if err = ie.SyncTxThreshOoC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SyncTxThreshOoC_r16", err)
		}
	}
	if SyncInfoReserved_r16Present {
		var tmp_bs_SyncInfoReserved_r16 []byte
		var n_SyncInfoReserved_r16 uint
		if tmp_bs_SyncInfoReserved_r16, n_SyncInfoReserved_r16, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode SyncInfoReserved_r16", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_SyncInfoReserved_r16,
			NumBits: uint64(n_SyncInfoReserved_r16),
		}
		ie.SyncInfoReserved_r16 = &tmp_bitstring
	}
	return nil
}
