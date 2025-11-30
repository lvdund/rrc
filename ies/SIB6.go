package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB6 struct {
	MessageIdentifier        aper.BitString `lb:16,ub:16,madatory`
	SerialNumber             aper.BitString `lb:16,ub:16,madatory`
	WarningType              []byte         `lb:2,ub:2,madatory`
	LateNonCriticalExtension *[]byte        `optional`
}

func (ie *SIB6) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteBitString(ie.MessageIdentifier.Bytes, uint(ie.MessageIdentifier.NumBits), &aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteBitString MessageIdentifier", err)
	}
	if err = w.WriteBitString(ie.SerialNumber.Bytes, uint(ie.SerialNumber.NumBits), &aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteBitString SerialNumber", err)
	}
	if err = w.WriteOctetString(ie.WarningType, &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteOctetString WarningType", err)
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB6) Decode(r *aper.AperReader) error {
	var err error
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_bs_MessageIdentifier []byte
	var n_MessageIdentifier uint
	if tmp_bs_MessageIdentifier, n_MessageIdentifier, err = r.ReadBitString(&aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadBitString MessageIdentifier", err)
	}
	ie.MessageIdentifier = aper.BitString{
		Bytes:   tmp_bs_MessageIdentifier,
		NumBits: uint64(n_MessageIdentifier),
	}
	var tmp_bs_SerialNumber []byte
	var n_SerialNumber uint
	if tmp_bs_SerialNumber, n_SerialNumber, err = r.ReadBitString(&aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadBitString SerialNumber", err)
	}
	ie.SerialNumber = aper.BitString{
		Bytes:   tmp_bs_SerialNumber,
		NumBits: uint64(n_SerialNumber),
	}
	var tmp_os_WarningType []byte
	if tmp_os_WarningType, err = r.ReadOctetString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadOctetString WarningType", err)
	}
	ie.WarningType = tmp_os_WarningType
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
