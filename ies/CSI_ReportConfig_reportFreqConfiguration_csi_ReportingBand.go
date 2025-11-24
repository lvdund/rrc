package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_nothing uint64 = iota
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands3
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands4
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands5
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands6
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands7
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands8
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands9
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands10
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands11
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands12
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands13
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands14
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands15
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands16
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands17
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands18
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands19_v1530
)

type CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand struct {
	Choice           uint64
	Subbands3        uper.BitString `lb:3,ub:3,madatory`
	Subbands4        uper.BitString `lb:4,ub:4,madatory`
	Subbands5        uper.BitString `lb:5,ub:5,madatory`
	Subbands6        uper.BitString `lb:6,ub:6,madatory`
	Subbands7        uper.BitString `lb:7,ub:7,madatory`
	Subbands8        uper.BitString `lb:8,ub:8,madatory`
	Subbands9        uper.BitString `lb:9,ub:9,madatory`
	Subbands10       uper.BitString `lb:10,ub:10,madatory`
	Subbands11       uper.BitString `lb:11,ub:11,madatory`
	Subbands12       uper.BitString `lb:12,ub:12,madatory`
	Subbands13       uper.BitString `lb:13,ub:13,madatory`
	Subbands14       uper.BitString `lb:14,ub:14,madatory`
	Subbands15       uper.BitString `lb:15,ub:15,madatory`
	Subbands16       uper.BitString `lb:16,ub:16,madatory`
	Subbands17       uper.BitString `lb:17,ub:17,madatory`
	Subbands18       uper.BitString `lb:18,ub:18,madatory`
	Subbands19_v1530 uper.BitString `lb:19,ub:19,madatory`
}

func (ie *CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 17, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands3:
		if err = w.WriteBitString(ie.Subbands3.Bytes, uint(ie.Subbands3.NumBits), &uper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode Subbands3", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands4:
		if err = w.WriteBitString(ie.Subbands4.Bytes, uint(ie.Subbands4.NumBits), &uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode Subbands4", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands5:
		if err = w.WriteBitString(ie.Subbands5.Bytes, uint(ie.Subbands5.NumBits), &uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
			err = utils.WrapError("Encode Subbands5", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands6:
		if err = w.WriteBitString(ie.Subbands6.Bytes, uint(ie.Subbands6.NumBits), &uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			err = utils.WrapError("Encode Subbands6", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands7:
		if err = w.WriteBitString(ie.Subbands7.Bytes, uint(ie.Subbands7.NumBits), &uper.Constraint{Lb: 7, Ub: 7}, false); err != nil {
			err = utils.WrapError("Encode Subbands7", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands8:
		if err = w.WriteBitString(ie.Subbands8.Bytes, uint(ie.Subbands8.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode Subbands8", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands9:
		if err = w.WriteBitString(ie.Subbands9.Bytes, uint(ie.Subbands9.NumBits), &uper.Constraint{Lb: 9, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode Subbands9", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands10:
		if err = w.WriteBitString(ie.Subbands10.Bytes, uint(ie.Subbands10.NumBits), &uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			err = utils.WrapError("Encode Subbands10", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands11:
		if err = w.WriteBitString(ie.Subbands11.Bytes, uint(ie.Subbands11.NumBits), &uper.Constraint{Lb: 11, Ub: 11}, false); err != nil {
			err = utils.WrapError("Encode Subbands11", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands12:
		if err = w.WriteBitString(ie.Subbands12.Bytes, uint(ie.Subbands12.NumBits), &uper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
			err = utils.WrapError("Encode Subbands12", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands13:
		if err = w.WriteBitString(ie.Subbands13.Bytes, uint(ie.Subbands13.NumBits), &uper.Constraint{Lb: 13, Ub: 13}, false); err != nil {
			err = utils.WrapError("Encode Subbands13", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands14:
		if err = w.WriteBitString(ie.Subbands14.Bytes, uint(ie.Subbands14.NumBits), &uper.Constraint{Lb: 14, Ub: 14}, false); err != nil {
			err = utils.WrapError("Encode Subbands14", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands15:
		if err = w.WriteBitString(ie.Subbands15.Bytes, uint(ie.Subbands15.NumBits), &uper.Constraint{Lb: 15, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode Subbands15", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands16:
		if err = w.WriteBitString(ie.Subbands16.Bytes, uint(ie.Subbands16.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			err = utils.WrapError("Encode Subbands16", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands17:
		if err = w.WriteBitString(ie.Subbands17.Bytes, uint(ie.Subbands17.NumBits), &uper.Constraint{Lb: 17, Ub: 17}, false); err != nil {
			err = utils.WrapError("Encode Subbands17", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands18:
		if err = w.WriteBitString(ie.Subbands18.Bytes, uint(ie.Subbands18.NumBits), &uper.Constraint{Lb: 18, Ub: 18}, false); err != nil {
			err = utils.WrapError("Encode Subbands18", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands19_v1530:
		if err = w.WriteBitString(ie.Subbands19_v1530.Bytes, uint(ie.Subbands19_v1530.NumBits), &uper.Constraint{Lb: 19, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode Subbands19_v1530", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(17, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands3:
		var tmp_bs_Subbands3 []byte
		var n_Subbands3 uint
		if tmp_bs_Subbands3, n_Subbands3, err = r.ReadBitString(&uper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode Subbands3", err)
		}
		ie.Subbands3 = uper.BitString{
			Bytes:   tmp_bs_Subbands3,
			NumBits: uint64(n_Subbands3),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands4:
		var tmp_bs_Subbands4 []byte
		var n_Subbands4 uint
		if tmp_bs_Subbands4, n_Subbands4, err = r.ReadBitString(&uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode Subbands4", err)
		}
		ie.Subbands4 = uper.BitString{
			Bytes:   tmp_bs_Subbands4,
			NumBits: uint64(n_Subbands4),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands5:
		var tmp_bs_Subbands5 []byte
		var n_Subbands5 uint
		if tmp_bs_Subbands5, n_Subbands5, err = r.ReadBitString(&uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
			return utils.WrapError("Decode Subbands5", err)
		}
		ie.Subbands5 = uper.BitString{
			Bytes:   tmp_bs_Subbands5,
			NumBits: uint64(n_Subbands5),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands6:
		var tmp_bs_Subbands6 []byte
		var n_Subbands6 uint
		if tmp_bs_Subbands6, n_Subbands6, err = r.ReadBitString(&uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode Subbands6", err)
		}
		ie.Subbands6 = uper.BitString{
			Bytes:   tmp_bs_Subbands6,
			NumBits: uint64(n_Subbands6),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands7:
		var tmp_bs_Subbands7 []byte
		var n_Subbands7 uint
		if tmp_bs_Subbands7, n_Subbands7, err = r.ReadBitString(&uper.Constraint{Lb: 7, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode Subbands7", err)
		}
		ie.Subbands7 = uper.BitString{
			Bytes:   tmp_bs_Subbands7,
			NumBits: uint64(n_Subbands7),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands8:
		var tmp_bs_Subbands8 []byte
		var n_Subbands8 uint
		if tmp_bs_Subbands8, n_Subbands8, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Subbands8", err)
		}
		ie.Subbands8 = uper.BitString{
			Bytes:   tmp_bs_Subbands8,
			NumBits: uint64(n_Subbands8),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands9:
		var tmp_bs_Subbands9 []byte
		var n_Subbands9 uint
		if tmp_bs_Subbands9, n_Subbands9, err = r.ReadBitString(&uper.Constraint{Lb: 9, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Subbands9", err)
		}
		ie.Subbands9 = uper.BitString{
			Bytes:   tmp_bs_Subbands9,
			NumBits: uint64(n_Subbands9),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands10:
		var tmp_bs_Subbands10 []byte
		var n_Subbands10 uint
		if tmp_bs_Subbands10, n_Subbands10, err = r.ReadBitString(&uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			return utils.WrapError("Decode Subbands10", err)
		}
		ie.Subbands10 = uper.BitString{
			Bytes:   tmp_bs_Subbands10,
			NumBits: uint64(n_Subbands10),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands11:
		var tmp_bs_Subbands11 []byte
		var n_Subbands11 uint
		if tmp_bs_Subbands11, n_Subbands11, err = r.ReadBitString(&uper.Constraint{Lb: 11, Ub: 11}, false); err != nil {
			return utils.WrapError("Decode Subbands11", err)
		}
		ie.Subbands11 = uper.BitString{
			Bytes:   tmp_bs_Subbands11,
			NumBits: uint64(n_Subbands11),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands12:
		var tmp_bs_Subbands12 []byte
		var n_Subbands12 uint
		if tmp_bs_Subbands12, n_Subbands12, err = r.ReadBitString(&uper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
			return utils.WrapError("Decode Subbands12", err)
		}
		ie.Subbands12 = uper.BitString{
			Bytes:   tmp_bs_Subbands12,
			NumBits: uint64(n_Subbands12),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands13:
		var tmp_bs_Subbands13 []byte
		var n_Subbands13 uint
		if tmp_bs_Subbands13, n_Subbands13, err = r.ReadBitString(&uper.Constraint{Lb: 13, Ub: 13}, false); err != nil {
			return utils.WrapError("Decode Subbands13", err)
		}
		ie.Subbands13 = uper.BitString{
			Bytes:   tmp_bs_Subbands13,
			NumBits: uint64(n_Subbands13),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands14:
		var tmp_bs_Subbands14 []byte
		var n_Subbands14 uint
		if tmp_bs_Subbands14, n_Subbands14, err = r.ReadBitString(&uper.Constraint{Lb: 14, Ub: 14}, false); err != nil {
			return utils.WrapError("Decode Subbands14", err)
		}
		ie.Subbands14 = uper.BitString{
			Bytes:   tmp_bs_Subbands14,
			NumBits: uint64(n_Subbands14),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands15:
		var tmp_bs_Subbands15 []byte
		var n_Subbands15 uint
		if tmp_bs_Subbands15, n_Subbands15, err = r.ReadBitString(&uper.Constraint{Lb: 15, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Subbands15", err)
		}
		ie.Subbands15 = uper.BitString{
			Bytes:   tmp_bs_Subbands15,
			NumBits: uint64(n_Subbands15),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands16:
		var tmp_bs_Subbands16 []byte
		var n_Subbands16 uint
		if tmp_bs_Subbands16, n_Subbands16, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Subbands16", err)
		}
		ie.Subbands16 = uper.BitString{
			Bytes:   tmp_bs_Subbands16,
			NumBits: uint64(n_Subbands16),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands17:
		var tmp_bs_Subbands17 []byte
		var n_Subbands17 uint
		if tmp_bs_Subbands17, n_Subbands17, err = r.ReadBitString(&uper.Constraint{Lb: 17, Ub: 17}, false); err != nil {
			return utils.WrapError("Decode Subbands17", err)
		}
		ie.Subbands17 = uper.BitString{
			Bytes:   tmp_bs_Subbands17,
			NumBits: uint64(n_Subbands17),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands18:
		var tmp_bs_Subbands18 []byte
		var n_Subbands18 uint
		if tmp_bs_Subbands18, n_Subbands18, err = r.ReadBitString(&uper.Constraint{Lb: 18, Ub: 18}, false); err != nil {
			return utils.WrapError("Decode Subbands18", err)
		}
		ie.Subbands18 = uper.BitString{
			Bytes:   tmp_bs_Subbands18,
			NumBits: uint64(n_Subbands18),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_Subbands19_v1530:
		var tmp_bs_Subbands19_v1530 []byte
		var n_Subbands19_v1530 uint
		if tmp_bs_Subbands19_v1530, n_Subbands19_v1530, err = r.ReadBitString(&uper.Constraint{Lb: 19, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode Subbands19_v1530", err)
		}
		ie.Subbands19_v1530 = uper.BitString{
			Bytes:   tmp_bs_Subbands19_v1530,
			NumBits: uint64(n_Subbands19_v1530),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
