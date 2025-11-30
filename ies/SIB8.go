package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB8 struct {
	MessageIdentifier             aper.BitString                 `lb:16,ub:16,madatory`
	SerialNumber                  aper.BitString                 `lb:16,ub:16,madatory`
	WarningMessageSegmentType     SIB8_warningMessageSegmentType `madatory`
	WarningMessageSegmentNumber   int64                          `lb:0,ub:63,madatory`
	WarningMessageSegment         []byte                         `madatory`
	DataCodingScheme              *[]byte                        `lb:1,ub:1,optional`
	WarningAreaCoordinatesSegment *[]byte                        `optional`
	LateNonCriticalExtension      *[]byte                        `optional`
}

func (ie *SIB8) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.DataCodingScheme != nil, ie.WarningAreaCoordinatesSegment != nil, ie.LateNonCriticalExtension != nil}
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
	if err = ie.WarningMessageSegmentType.Encode(w); err != nil {
		return utils.WrapError("Encode WarningMessageSegmentType", err)
	}
	if err = w.WriteInteger(ie.WarningMessageSegmentNumber, &aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger WarningMessageSegmentNumber", err)
	}
	if err = w.WriteOctetString(ie.WarningMessageSegment, nil, false); err != nil {
		return utils.WrapError("WriteOctetString WarningMessageSegment", err)
	}
	if ie.DataCodingScheme != nil {
		if err = w.WriteOctetString(*ie.DataCodingScheme, &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
			return utils.WrapError("Encode DataCodingScheme", err)
		}
	}
	if ie.WarningAreaCoordinatesSegment != nil {
		if err = w.WriteOctetString(*ie.WarningAreaCoordinatesSegment, nil, false); err != nil {
			return utils.WrapError("Encode WarningAreaCoordinatesSegment", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB8) Decode(r *aper.AperReader) error {
	var err error
	var DataCodingSchemePresent bool
	if DataCodingSchemePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var WarningAreaCoordinatesSegmentPresent bool
	if WarningAreaCoordinatesSegmentPresent, err = r.ReadBool(); err != nil {
		return err
	}
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
	if err = ie.WarningMessageSegmentType.Decode(r); err != nil {
		return utils.WrapError("Decode WarningMessageSegmentType", err)
	}
	var tmp_int_WarningMessageSegmentNumber int64
	if tmp_int_WarningMessageSegmentNumber, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger WarningMessageSegmentNumber", err)
	}
	ie.WarningMessageSegmentNumber = tmp_int_WarningMessageSegmentNumber
	var tmp_os_WarningMessageSegment []byte
	if tmp_os_WarningMessageSegment, err = r.ReadOctetString(nil, false); err != nil {
		return utils.WrapError("ReadOctetString WarningMessageSegment", err)
	}
	ie.WarningMessageSegment = tmp_os_WarningMessageSegment
	if DataCodingSchemePresent {
		var tmp_os_DataCodingScheme []byte
		if tmp_os_DataCodingScheme, err = r.ReadOctetString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode DataCodingScheme", err)
		}
		ie.DataCodingScheme = &tmp_os_DataCodingScheme
	}
	if WarningAreaCoordinatesSegmentPresent {
		var tmp_os_WarningAreaCoordinatesSegment []byte
		if tmp_os_WarningAreaCoordinatesSegment, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode WarningAreaCoordinatesSegment", err)
		}
		ie.WarningAreaCoordinatesSegment = &tmp_os_WarningAreaCoordinatesSegment
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
