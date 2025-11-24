package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB5 struct {
	CarrierFreqListEUTRA          *CarrierFreqListEUTRA               `optional`
	T_ReselectionEUTRA            T_Reselection                       `madatory`
	T_ReselectionEUTRA_SF         *SpeedStateScaleFactors             `optional`
	LateNonCriticalExtension      *[]byte                             `optional`
	CarrierFreqListEUTRA_v1610    *CarrierFreqListEUTRA_v1610         `optional,ext-1`
	CarrierFreqListEUTRA_v1700    *CarrierFreqListEUTRA_v1700         `optional,ext-2`
	IdleModeMeasVoiceFallback_r17 *SIB5_idleModeMeasVoiceFallback_r17 `optional,ext-2`
}

func (ie *SIB5) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.CarrierFreqListEUTRA_v1610 != nil || ie.CarrierFreqListEUTRA_v1700 != nil || ie.IdleModeMeasVoiceFallback_r17 != nil
	preambleBits := []bool{hasExtensions, ie.CarrierFreqListEUTRA != nil, ie.T_ReselectionEUTRA_SF != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CarrierFreqListEUTRA != nil {
		if err = ie.CarrierFreqListEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode CarrierFreqListEUTRA", err)
		}
	}
	if err = ie.T_ReselectionEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode T_ReselectionEUTRA", err)
	}
	if ie.T_ReselectionEUTRA_SF != nil {
		if err = ie.T_ReselectionEUTRA_SF.Encode(w); err != nil {
			return utils.WrapError("Encode T_ReselectionEUTRA_SF", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.CarrierFreqListEUTRA_v1610 != nil, ie.CarrierFreqListEUTRA_v1700 != nil || ie.IdleModeMeasVoiceFallback_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SIB5", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.CarrierFreqListEUTRA_v1610 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode CarrierFreqListEUTRA_v1610 optional
			if ie.CarrierFreqListEUTRA_v1610 != nil {
				if err = ie.CarrierFreqListEUTRA_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CarrierFreqListEUTRA_v1610", err)
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
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.CarrierFreqListEUTRA_v1700 != nil, ie.IdleModeMeasVoiceFallback_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode CarrierFreqListEUTRA_v1700 optional
			if ie.CarrierFreqListEUTRA_v1700 != nil {
				if err = ie.CarrierFreqListEUTRA_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CarrierFreqListEUTRA_v1700", err)
				}
			}
			// encode IdleModeMeasVoiceFallback_r17 optional
			if ie.IdleModeMeasVoiceFallback_r17 != nil {
				if err = ie.IdleModeMeasVoiceFallback_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IdleModeMeasVoiceFallback_r17", err)
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

func (ie *SIB5) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var CarrierFreqListEUTRAPresent bool
	if CarrierFreqListEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var T_ReselectionEUTRA_SFPresent bool
	if T_ReselectionEUTRA_SFPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CarrierFreqListEUTRAPresent {
		ie.CarrierFreqListEUTRA = new(CarrierFreqListEUTRA)
		if err = ie.CarrierFreqListEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode CarrierFreqListEUTRA", err)
		}
	}
	if err = ie.T_ReselectionEUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode T_ReselectionEUTRA", err)
	}
	if T_ReselectionEUTRA_SFPresent {
		ie.T_ReselectionEUTRA_SF = new(SpeedStateScaleFactors)
		if err = ie.T_ReselectionEUTRA_SF.Decode(r); err != nil {
			return utils.WrapError("Decode T_ReselectionEUTRA_SF", err)
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
		// Read extension bitmap: 2 bits for 2 extension groups
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

			CarrierFreqListEUTRA_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode CarrierFreqListEUTRA_v1610 optional
			if CarrierFreqListEUTRA_v1610Present {
				ie.CarrierFreqListEUTRA_v1610 = new(CarrierFreqListEUTRA_v1610)
				if err = ie.CarrierFreqListEUTRA_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode CarrierFreqListEUTRA_v1610", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			CarrierFreqListEUTRA_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IdleModeMeasVoiceFallback_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode CarrierFreqListEUTRA_v1700 optional
			if CarrierFreqListEUTRA_v1700Present {
				ie.CarrierFreqListEUTRA_v1700 = new(CarrierFreqListEUTRA_v1700)
				if err = ie.CarrierFreqListEUTRA_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode CarrierFreqListEUTRA_v1700", err)
				}
			}
			// decode IdleModeMeasVoiceFallback_r17 optional
			if IdleModeMeasVoiceFallback_r17Present {
				ie.IdleModeMeasVoiceFallback_r17 = new(SIB5_idleModeMeasVoiceFallback_r17)
				if err = ie.IdleModeMeasVoiceFallback_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode IdleModeMeasVoiceFallback_r17", err)
				}
			}
		}
	}
	return nil
}
