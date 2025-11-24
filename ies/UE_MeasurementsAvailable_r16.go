package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_MeasurementsAvailable_r16 struct {
	LogMeasAvailable_r16          *UE_MeasurementsAvailable_r16_logMeasAvailable_r16         `optional`
	LogMeasAvailableBT_r16        *UE_MeasurementsAvailable_r16_logMeasAvailableBT_r16       `optional`
	LogMeasAvailableWLAN_r16      *UE_MeasurementsAvailable_r16_logMeasAvailableWLAN_r16     `optional`
	ConnEstFailInfoAvailable_r16  *UE_MeasurementsAvailable_r16_connEstFailInfoAvailable_r16 `optional`
	Rlf_InfoAvailable_r16         *UE_MeasurementsAvailable_r16_rlf_InfoAvailable_r16        `optional`
	SuccessHO_InfoAvailable_r17   *UE_MeasurementsAvailable_r16_successHO_InfoAvailable_r17  `optional,ext-1`
	SigLogMeasConfigAvailable_r17 *bool                                                      `optional,ext-1`
}

func (ie *UE_MeasurementsAvailable_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.SuccessHO_InfoAvailable_r17 != nil || ie.SigLogMeasConfigAvailable_r17 != nil
	preambleBits := []bool{hasExtensions, ie.LogMeasAvailable_r16 != nil, ie.LogMeasAvailableBT_r16 != nil, ie.LogMeasAvailableWLAN_r16 != nil, ie.ConnEstFailInfoAvailable_r16 != nil, ie.Rlf_InfoAvailable_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.LogMeasAvailable_r16 != nil {
		if err = ie.LogMeasAvailable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode LogMeasAvailable_r16", err)
		}
	}
	if ie.LogMeasAvailableBT_r16 != nil {
		if err = ie.LogMeasAvailableBT_r16.Encode(w); err != nil {
			return utils.WrapError("Encode LogMeasAvailableBT_r16", err)
		}
	}
	if ie.LogMeasAvailableWLAN_r16 != nil {
		if err = ie.LogMeasAvailableWLAN_r16.Encode(w); err != nil {
			return utils.WrapError("Encode LogMeasAvailableWLAN_r16", err)
		}
	}
	if ie.ConnEstFailInfoAvailable_r16 != nil {
		if err = ie.ConnEstFailInfoAvailable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ConnEstFailInfoAvailable_r16", err)
		}
	}
	if ie.Rlf_InfoAvailable_r16 != nil {
		if err = ie.Rlf_InfoAvailable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Rlf_InfoAvailable_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.SuccessHO_InfoAvailable_r17 != nil || ie.SigLogMeasConfigAvailable_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap UE_MeasurementsAvailable_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.SuccessHO_InfoAvailable_r17 != nil, ie.SigLogMeasConfigAvailable_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SuccessHO_InfoAvailable_r17 optional
			if ie.SuccessHO_InfoAvailable_r17 != nil {
				if err = ie.SuccessHO_InfoAvailable_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SuccessHO_InfoAvailable_r17", err)
				}
			}
			// encode SigLogMeasConfigAvailable_r17 optional
			if ie.SigLogMeasConfigAvailable_r17 != nil {
				if err = extWriter.WriteBoolean(*ie.SigLogMeasConfigAvailable_r17); err != nil {
					return utils.WrapError("Encode SigLogMeasConfigAvailable_r17", err)
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

func (ie *UE_MeasurementsAvailable_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var LogMeasAvailable_r16Present bool
	if LogMeasAvailable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LogMeasAvailableBT_r16Present bool
	if LogMeasAvailableBT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LogMeasAvailableWLAN_r16Present bool
	if LogMeasAvailableWLAN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ConnEstFailInfoAvailable_r16Present bool
	if ConnEstFailInfoAvailable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rlf_InfoAvailable_r16Present bool
	if Rlf_InfoAvailable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if LogMeasAvailable_r16Present {
		ie.LogMeasAvailable_r16 = new(UE_MeasurementsAvailable_r16_logMeasAvailable_r16)
		if err = ie.LogMeasAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LogMeasAvailable_r16", err)
		}
	}
	if LogMeasAvailableBT_r16Present {
		ie.LogMeasAvailableBT_r16 = new(UE_MeasurementsAvailable_r16_logMeasAvailableBT_r16)
		if err = ie.LogMeasAvailableBT_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LogMeasAvailableBT_r16", err)
		}
	}
	if LogMeasAvailableWLAN_r16Present {
		ie.LogMeasAvailableWLAN_r16 = new(UE_MeasurementsAvailable_r16_logMeasAvailableWLAN_r16)
		if err = ie.LogMeasAvailableWLAN_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LogMeasAvailableWLAN_r16", err)
		}
	}
	if ConnEstFailInfoAvailable_r16Present {
		ie.ConnEstFailInfoAvailable_r16 = new(UE_MeasurementsAvailable_r16_connEstFailInfoAvailable_r16)
		if err = ie.ConnEstFailInfoAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ConnEstFailInfoAvailable_r16", err)
		}
	}
	if Rlf_InfoAvailable_r16Present {
		ie.Rlf_InfoAvailable_r16 = new(UE_MeasurementsAvailable_r16_rlf_InfoAvailable_r16)
		if err = ie.Rlf_InfoAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Rlf_InfoAvailable_r16", err)
		}
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

			SuccessHO_InfoAvailable_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SigLogMeasConfigAvailable_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SuccessHO_InfoAvailable_r17 optional
			if SuccessHO_InfoAvailable_r17Present {
				ie.SuccessHO_InfoAvailable_r17 = new(UE_MeasurementsAvailable_r16_successHO_InfoAvailable_r17)
				if err = ie.SuccessHO_InfoAvailable_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SuccessHO_InfoAvailable_r17", err)
				}
			}
			// decode SigLogMeasConfigAvailable_r17 optional
			if SigLogMeasConfigAvailable_r17Present {
				var tmp_bool_SigLogMeasConfigAvailable_r17 bool
				if tmp_bool_SigLogMeasConfigAvailable_r17, err = extReader.ReadBoolean(); err != nil {
					return utils.WrapError("Decode SigLogMeasConfigAvailable_r17", err)
				}
				ie.SigLogMeasConfigAvailable_r17 = &tmp_bool_SigLogMeasConfigAvailable_r17
			}
		}
	}
	return nil
}
