package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB4 struct {
	InterFreqCarrierFreqList       InterFreqCarrierFreqList        `madatory`
	LateNonCriticalExtension       *[]byte                         `optional`
	InterFreqCarrierFreqList_v1610 *InterFreqCarrierFreqList_v1610 `optional,ext-1`
	InterFreqCarrierFreqList_v1700 *InterFreqCarrierFreqList_v1700 `optional,ext-2`
	InterFreqCarrierFreqList_v1720 *InterFreqCarrierFreqList_v1720 `optional,ext-3`
	InterFreqCarrierFreqList_v1730 *InterFreqCarrierFreqList_v1730 `optional,ext-4`
}

func (ie *SIB4) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.InterFreqCarrierFreqList_v1610 != nil || ie.InterFreqCarrierFreqList_v1700 != nil || ie.InterFreqCarrierFreqList_v1720 != nil || ie.InterFreqCarrierFreqList_v1730 != nil
	preambleBits := []bool{hasExtensions, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.InterFreqCarrierFreqList.Encode(w); err != nil {
		return utils.WrapError("Encode InterFreqCarrierFreqList", err)
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 4 bits for 4 extension groups
		extBitmap := []bool{ie.InterFreqCarrierFreqList_v1610 != nil, ie.InterFreqCarrierFreqList_v1700 != nil, ie.InterFreqCarrierFreqList_v1720 != nil, ie.InterFreqCarrierFreqList_v1730 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SIB4", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.InterFreqCarrierFreqList_v1610 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode InterFreqCarrierFreqList_v1610 optional
			if ie.InterFreqCarrierFreqList_v1610 != nil {
				if err = ie.InterFreqCarrierFreqList_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode InterFreqCarrierFreqList_v1610", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.InterFreqCarrierFreqList_v1700 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode InterFreqCarrierFreqList_v1700 optional
			if ie.InterFreqCarrierFreqList_v1700 != nil {
				if err = ie.InterFreqCarrierFreqList_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode InterFreqCarrierFreqList_v1700", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.InterFreqCarrierFreqList_v1720 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode InterFreqCarrierFreqList_v1720 optional
			if ie.InterFreqCarrierFreqList_v1720 != nil {
				if err = ie.InterFreqCarrierFreqList_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode InterFreqCarrierFreqList_v1720", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{ie.InterFreqCarrierFreqList_v1730 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode InterFreqCarrierFreqList_v1730 optional
			if ie.InterFreqCarrierFreqList_v1730 != nil {
				if err = ie.InterFreqCarrierFreqList_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode InterFreqCarrierFreqList_v1730", err)
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

func (ie *SIB4) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.InterFreqCarrierFreqList.Decode(r); err != nil {
		return utils.WrapError("Decode InterFreqCarrierFreqList", err)
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}

	if extensionBit {
		// Read extension bitmap: 4 bits for 4 extension groups
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			InterFreqCarrierFreqList_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode InterFreqCarrierFreqList_v1610 optional
			if InterFreqCarrierFreqList_v1610Present {
				ie.InterFreqCarrierFreqList_v1610 = new(InterFreqCarrierFreqList_v1610)
				if err = ie.InterFreqCarrierFreqList_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode InterFreqCarrierFreqList_v1610", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			InterFreqCarrierFreqList_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode InterFreqCarrierFreqList_v1700 optional
			if InterFreqCarrierFreqList_v1700Present {
				ie.InterFreqCarrierFreqList_v1700 = new(InterFreqCarrierFreqList_v1700)
				if err = ie.InterFreqCarrierFreqList_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode InterFreqCarrierFreqList_v1700", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			InterFreqCarrierFreqList_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode InterFreqCarrierFreqList_v1720 optional
			if InterFreqCarrierFreqList_v1720Present {
				ie.InterFreqCarrierFreqList_v1720 = new(InterFreqCarrierFreqList_v1720)
				if err = ie.InterFreqCarrierFreqList_v1720.Decode(extReader); err != nil {
					return utils.WrapError("Decode InterFreqCarrierFreqList_v1720", err)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			InterFreqCarrierFreqList_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode InterFreqCarrierFreqList_v1730 optional
			if InterFreqCarrierFreqList_v1730Present {
				ie.InterFreqCarrierFreqList_v1730 = new(InterFreqCarrierFreqList_v1730)
				if err = ie.InterFreqCarrierFreqList_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode InterFreqCarrierFreqList_v1730", err)
				}
			}
		}
	}
	return nil
}
