package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MasterInformationBlockSidelink struct {
	Sl_TDD_Config_r16     aper.BitString `lb:12,ub:12,madatory`
	InCoverage_r16        bool           `madatory`
	DirectFrameNumber_r16 aper.BitString `lb:10,ub:10,madatory`
	SlotIndex_r16         aper.BitString `lb:7,ub:7,madatory`
	ReservedBits_r16      aper.BitString `lb:2,ub:2,madatory`
}

func (ie *MasterInformationBlockSidelink) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Sl_TDD_Config_r16.Bytes, uint(ie.Sl_TDD_Config_r16.NumBits), &aper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
		return utils.WrapError("WriteBitString Sl_TDD_Config_r16", err)
	}
	if err = w.WriteBoolean(ie.InCoverage_r16); err != nil {
		return utils.WrapError("WriteBoolean InCoverage_r16", err)
	}
	if err = w.WriteBitString(ie.DirectFrameNumber_r16.Bytes, uint(ie.DirectFrameNumber_r16.NumBits), &aper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
		return utils.WrapError("WriteBitString DirectFrameNumber_r16", err)
	}
	if err = w.WriteBitString(ie.SlotIndex_r16.Bytes, uint(ie.SlotIndex_r16.NumBits), &aper.Constraint{Lb: 7, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteBitString SlotIndex_r16", err)
	}
	if err = w.WriteBitString(ie.ReservedBits_r16.Bytes, uint(ie.ReservedBits_r16.NumBits), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteBitString ReservedBits_r16", err)
	}
	return nil
}

func (ie *MasterInformationBlockSidelink) Decode(r *aper.AperReader) error {
	var err error
	var tmp_bs_Sl_TDD_Config_r16 []byte
	var n_Sl_TDD_Config_r16 uint
	if tmp_bs_Sl_TDD_Config_r16, n_Sl_TDD_Config_r16, err = r.ReadBitString(&aper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
		return utils.WrapError("ReadBitString Sl_TDD_Config_r16", err)
	}
	ie.Sl_TDD_Config_r16 = aper.BitString{
		Bytes:   tmp_bs_Sl_TDD_Config_r16,
		NumBits: uint64(n_Sl_TDD_Config_r16),
	}
	var tmp_bool_InCoverage_r16 bool
	if tmp_bool_InCoverage_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean InCoverage_r16", err)
	}
	ie.InCoverage_r16 = tmp_bool_InCoverage_r16
	var tmp_bs_DirectFrameNumber_r16 []byte
	var n_DirectFrameNumber_r16 uint
	if tmp_bs_DirectFrameNumber_r16, n_DirectFrameNumber_r16, err = r.ReadBitString(&aper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
		return utils.WrapError("ReadBitString DirectFrameNumber_r16", err)
	}
	ie.DirectFrameNumber_r16 = aper.BitString{
		Bytes:   tmp_bs_DirectFrameNumber_r16,
		NumBits: uint64(n_DirectFrameNumber_r16),
	}
	var tmp_bs_SlotIndex_r16 []byte
	var n_SlotIndex_r16 uint
	if tmp_bs_SlotIndex_r16, n_SlotIndex_r16, err = r.ReadBitString(&aper.Constraint{Lb: 7, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadBitString SlotIndex_r16", err)
	}
	ie.SlotIndex_r16 = aper.BitString{
		Bytes:   tmp_bs_SlotIndex_r16,
		NumBits: uint64(n_SlotIndex_r16),
	}
	var tmp_bs_ReservedBits_r16 []byte
	var n_ReservedBits_r16 uint
	if tmp_bs_ReservedBits_r16, n_ReservedBits_r16, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadBitString ReservedBits_r16", err)
	}
	ie.ReservedBits_r16 = aper.BitString{
		Bytes:   tmp_bs_ReservedBits_r16,
		NumBits: uint64(n_ReservedBits_r16),
	}
	return nil
}
