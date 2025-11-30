package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_v1710_srs_AntennaSwitchingBeyond4RX_r17 struct {
	SupportedSRS_TxPortSwitchBeyond4Rx_r17 aper.BitString `lb:11,ub:11,madatory`
	EntryNumberAffectBeyond4Rx_r17         *int64         `lb:1,ub:32,optional`
	EntryNumberSwitchBeyond4Rx_r17         *int64         `lb:1,ub:32,optional`
}

func (ie *BandParameters_v1710_srs_AntennaSwitchingBeyond4RX_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.EntryNumberAffectBeyond4Rx_r17 != nil, ie.EntryNumberSwitchBeyond4Rx_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteBitString(ie.SupportedSRS_TxPortSwitchBeyond4Rx_r17.Bytes, uint(ie.SupportedSRS_TxPortSwitchBeyond4Rx_r17.NumBits), &aper.Constraint{Lb: 11, Ub: 11}, false); err != nil {
		return utils.WrapError("WriteBitString SupportedSRS_TxPortSwitchBeyond4Rx_r17", err)
	}
	if ie.EntryNumberAffectBeyond4Rx_r17 != nil {
		if err = w.WriteInteger(*ie.EntryNumberAffectBeyond4Rx_r17, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode EntryNumberAffectBeyond4Rx_r17", err)
		}
	}
	if ie.EntryNumberSwitchBeyond4Rx_r17 != nil {
		if err = w.WriteInteger(*ie.EntryNumberSwitchBeyond4Rx_r17, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode EntryNumberSwitchBeyond4Rx_r17", err)
		}
	}
	return nil
}

func (ie *BandParameters_v1710_srs_AntennaSwitchingBeyond4RX_r17) Decode(r *aper.AperReader) error {
	var err error
	var EntryNumberAffectBeyond4Rx_r17Present bool
	if EntryNumberAffectBeyond4Rx_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var EntryNumberSwitchBeyond4Rx_r17Present bool
	if EntryNumberSwitchBeyond4Rx_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_bs_SupportedSRS_TxPortSwitchBeyond4Rx_r17 []byte
	var n_SupportedSRS_TxPortSwitchBeyond4Rx_r17 uint
	if tmp_bs_SupportedSRS_TxPortSwitchBeyond4Rx_r17, n_SupportedSRS_TxPortSwitchBeyond4Rx_r17, err = r.ReadBitString(&aper.Constraint{Lb: 11, Ub: 11}, false); err != nil {
		return utils.WrapError("ReadBitString SupportedSRS_TxPortSwitchBeyond4Rx_r17", err)
	}
	ie.SupportedSRS_TxPortSwitchBeyond4Rx_r17 = aper.BitString{
		Bytes:   tmp_bs_SupportedSRS_TxPortSwitchBeyond4Rx_r17,
		NumBits: uint64(n_SupportedSRS_TxPortSwitchBeyond4Rx_r17),
	}
	if EntryNumberAffectBeyond4Rx_r17Present {
		var tmp_int_EntryNumberAffectBeyond4Rx_r17 int64
		if tmp_int_EntryNumberAffectBeyond4Rx_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode EntryNumberAffectBeyond4Rx_r17", err)
		}
		ie.EntryNumberAffectBeyond4Rx_r17 = &tmp_int_EntryNumberAffectBeyond4Rx_r17
	}
	if EntryNumberSwitchBeyond4Rx_r17Present {
		var tmp_int_EntryNumberSwitchBeyond4Rx_r17 int64
		if tmp_int_EntryNumberSwitchBeyond4Rx_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode EntryNumberSwitchBeyond4Rx_r17", err)
		}
		ie.EntryNumberSwitchBeyond4Rx_r17 = &tmp_int_EntryNumberSwitchBeyond4Rx_r17
	}
	return nil
}
