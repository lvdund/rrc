package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB7 struct {
	MessageIdentifier           uper.BitString                 `lb:16,ub:16,madatory`
	SerialNumber                uper.BitString                 `lb:16,ub:16,madatory`
	WarningMessageSegmentType   SIB7_warningMessageSegmentType `madatory`
	WarningMessageSegmentNumber int64                          `lb:0,ub:63,madatory`
	WarningMessageSegment       []byte                         `madatory`
	DataCodingScheme            *[]byte                        `lb:1,ub:1,optional`
	LateNonCriticalExtension    *[]byte                        `optional`
}

func (ie *SIB7) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.DataCodingScheme != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteBitString(ie.MessageIdentifier.Bytes, uint(ie.MessageIdentifier.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteBitString MessageIdentifier", err)
	}
	if err = w.WriteBitString(ie.SerialNumber.Bytes, uint(ie.SerialNumber.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteBitString SerialNumber", err)
	}
	if err = ie.WarningMessageSegmentType.Encode(w); err != nil {
		return utils.WrapError("Encode WarningMessageSegmentType", err)
	}
	if err = w.WriteInteger(ie.WarningMessageSegmentNumber, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger WarningMessageSegmentNumber", err)
	}
	if err = w.WriteOctetString(ie.WarningMessageSegment, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString WarningMessageSegment", err)
	}
	if ie.DataCodingScheme != nil {
		if err = w.WriteOctetString(*ie.DataCodingScheme, &uper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
			return utils.WrapError("Encode DataCodingScheme", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB7) Decode(r *uper.UperReader) error {
	var err error
	var DataCodingSchemePresent bool
	if DataCodingSchemePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_bs_MessageIdentifier []byte
	var n_MessageIdentifier uint
	if tmp_bs_MessageIdentifier, n_MessageIdentifier, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadBitString MessageIdentifier", err)
	}
	ie.MessageIdentifier = uper.BitString{
		Bytes:   tmp_bs_MessageIdentifier,
		NumBits: uint64(n_MessageIdentifier),
	}
	var tmp_bs_SerialNumber []byte
	var n_SerialNumber uint
	if tmp_bs_SerialNumber, n_SerialNumber, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadBitString SerialNumber", err)
	}
	ie.SerialNumber = uper.BitString{
		Bytes:   tmp_bs_SerialNumber,
		NumBits: uint64(n_SerialNumber),
	}
	if err = ie.WarningMessageSegmentType.Decode(r); err != nil {
		return utils.WrapError("Decode WarningMessageSegmentType", err)
	}
	var tmp_int_WarningMessageSegmentNumber int64
	if tmp_int_WarningMessageSegmentNumber, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger WarningMessageSegmentNumber", err)
	}
	ie.WarningMessageSegmentNumber = tmp_int_WarningMessageSegmentNumber
	var tmp_os_WarningMessageSegment []byte
	if tmp_os_WarningMessageSegment, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString WarningMessageSegment", err)
	}
	ie.WarningMessageSegment = tmp_os_WarningMessageSegment
	if DataCodingSchemePresent {
		var tmp_os_DataCodingScheme []byte
		if tmp_os_DataCodingScheme, err = r.ReadOctetString(&uper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode DataCodingScheme", err)
		}
		ie.DataCodingScheme = &tmp_os_DataCodingScheme
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
