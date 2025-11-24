package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB9 struct {
	TimeInfo                 *SIB9_timeInfo         `lb:2,ub:2,optional`
	LateNonCriticalExtension *[]byte                `optional`
	ReferenceTimeInfo_r16    *ReferenceTimeInfo_r16 `optional,ext-1`
}

func (ie *SIB9) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.ReferenceTimeInfo_r16 != nil
	preambleBits := []bool{hasExtensions, ie.TimeInfo != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.TimeInfo != nil {
		if err = ie.TimeInfo.Encode(w); err != nil {
			return utils.WrapError("Encode TimeInfo", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.ReferenceTimeInfo_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SIB9", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ReferenceTimeInfo_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ReferenceTimeInfo_r16 optional
			if ie.ReferenceTimeInfo_r16 != nil {
				if err = ie.ReferenceTimeInfo_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ReferenceTimeInfo_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *SIB9) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var TimeInfoPresent bool
	if TimeInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if TimeInfoPresent {
		ie.TimeInfo = new(SIB9_timeInfo)
		if err = ie.TimeInfo.Decode(r); err != nil {
			return utils.WrapError("Decode TimeInfo", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ReferenceTimeInfo_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ReferenceTimeInfo_r16 optional
			if ReferenceTimeInfo_r16Present {
				ie.ReferenceTimeInfo_r16 = new(ReferenceTimeInfo_r16)
				if err = ie.ReferenceTimeInfo_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ReferenceTimeInfo_r16", err)
				}
			}
		}
	}
	return nil
}
