package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CGI_InfoEUTRALogging struct {
	Plmn_Identity_eutra_5gc    *PLMN_Identity    `optional`
	TrackingAreaCode_eutra_5gc *TrackingAreaCode `optional`
	CellIdentity_eutra_5gc     *aper.BitString   `lb:28,ub:28,optional`
	Plmn_Identity_eutra_epc    *PLMN_Identity    `optional`
	TrackingAreaCode_eutra_epc *aper.BitString   `lb:16,ub:16,optional`
	CellIdentity_eutra_epc     *aper.BitString   `lb:28,ub:28,optional`
}

func (ie *CGI_InfoEUTRALogging) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Plmn_Identity_eutra_5gc != nil, ie.TrackingAreaCode_eutra_5gc != nil, ie.CellIdentity_eutra_5gc != nil, ie.Plmn_Identity_eutra_epc != nil, ie.TrackingAreaCode_eutra_epc != nil, ie.CellIdentity_eutra_epc != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Plmn_Identity_eutra_5gc != nil {
		if err = ie.Plmn_Identity_eutra_5gc.Encode(w); err != nil {
			return utils.WrapError("Encode Plmn_Identity_eutra_5gc", err)
		}
	}
	if ie.TrackingAreaCode_eutra_5gc != nil {
		if err = ie.TrackingAreaCode_eutra_5gc.Encode(w); err != nil {
			return utils.WrapError("Encode TrackingAreaCode_eutra_5gc", err)
		}
	}
	if ie.CellIdentity_eutra_5gc != nil {
		if err = w.WriteBitString(ie.CellIdentity_eutra_5gc.Bytes, uint(ie.CellIdentity_eutra_5gc.NumBits), &aper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
			return utils.WrapError("Encode CellIdentity_eutra_5gc", err)
		}
	}
	if ie.Plmn_Identity_eutra_epc != nil {
		if err = ie.Plmn_Identity_eutra_epc.Encode(w); err != nil {
			return utils.WrapError("Encode Plmn_Identity_eutra_epc", err)
		}
	}
	if ie.TrackingAreaCode_eutra_epc != nil {
		if err = w.WriteBitString(ie.TrackingAreaCode_eutra_epc.Bytes, uint(ie.TrackingAreaCode_eutra_epc.NumBits), &aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode TrackingAreaCode_eutra_epc", err)
		}
	}
	if ie.CellIdentity_eutra_epc != nil {
		if err = w.WriteBitString(ie.CellIdentity_eutra_epc.Bytes, uint(ie.CellIdentity_eutra_epc.NumBits), &aper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
			return utils.WrapError("Encode CellIdentity_eutra_epc", err)
		}
	}
	return nil
}

func (ie *CGI_InfoEUTRALogging) Decode(r *aper.AperReader) error {
	var err error
	var Plmn_Identity_eutra_5gcPresent bool
	if Plmn_Identity_eutra_5gcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TrackingAreaCode_eutra_5gcPresent bool
	if TrackingAreaCode_eutra_5gcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CellIdentity_eutra_5gcPresent bool
	if CellIdentity_eutra_5gcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Plmn_Identity_eutra_epcPresent bool
	if Plmn_Identity_eutra_epcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TrackingAreaCode_eutra_epcPresent bool
	if TrackingAreaCode_eutra_epcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CellIdentity_eutra_epcPresent bool
	if CellIdentity_eutra_epcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Plmn_Identity_eutra_5gcPresent {
		ie.Plmn_Identity_eutra_5gc = new(PLMN_Identity)
		if err = ie.Plmn_Identity_eutra_5gc.Decode(r); err != nil {
			return utils.WrapError("Decode Plmn_Identity_eutra_5gc", err)
		}
	}
	if TrackingAreaCode_eutra_5gcPresent {
		ie.TrackingAreaCode_eutra_5gc = new(TrackingAreaCode)
		if err = ie.TrackingAreaCode_eutra_5gc.Decode(r); err != nil {
			return utils.WrapError("Decode TrackingAreaCode_eutra_5gc", err)
		}
	}
	if CellIdentity_eutra_5gcPresent {
		var tmp_bs_CellIdentity_eutra_5gc []byte
		var n_CellIdentity_eutra_5gc uint
		if tmp_bs_CellIdentity_eutra_5gc, n_CellIdentity_eutra_5gc, err = r.ReadBitString(&aper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
			return utils.WrapError("Decode CellIdentity_eutra_5gc", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_CellIdentity_eutra_5gc,
			NumBits: uint64(n_CellIdentity_eutra_5gc),
		}
		ie.CellIdentity_eutra_5gc = &tmp_bitstring
	}
	if Plmn_Identity_eutra_epcPresent {
		ie.Plmn_Identity_eutra_epc = new(PLMN_Identity)
		if err = ie.Plmn_Identity_eutra_epc.Decode(r); err != nil {
			return utils.WrapError("Decode Plmn_Identity_eutra_epc", err)
		}
	}
	if TrackingAreaCode_eutra_epcPresent {
		var tmp_bs_TrackingAreaCode_eutra_epc []byte
		var n_TrackingAreaCode_eutra_epc uint
		if tmp_bs_TrackingAreaCode_eutra_epc, n_TrackingAreaCode_eutra_epc, err = r.ReadBitString(&aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode TrackingAreaCode_eutra_epc", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_TrackingAreaCode_eutra_epc,
			NumBits: uint64(n_TrackingAreaCode_eutra_epc),
		}
		ie.TrackingAreaCode_eutra_epc = &tmp_bitstring
	}
	if CellIdentity_eutra_epcPresent {
		var tmp_bs_CellIdentity_eutra_epc []byte
		var n_CellIdentity_eutra_epc uint
		if tmp_bs_CellIdentity_eutra_epc, n_CellIdentity_eutra_epc, err = r.ReadBitString(&aper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
			return utils.WrapError("Decode CellIdentity_eutra_epc", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_CellIdentity_eutra_epc,
			NumBits: uint64(n_CellIdentity_eutra_epc),
		}
		ie.CellIdentity_eutra_epc = &tmp_bitstring
	}
	return nil
}
