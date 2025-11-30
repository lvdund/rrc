package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RMTC_Config_r16 struct {
	Rmtc_Periodicity_r16      RMTC_Config_r16_rmtc_Periodicity_r16       `madatory`
	Rmtc_SubframeOffset_r16   *int64                                     `lb:0,ub:639,optional`
	MeasDurationSymbols_r16   RMTC_Config_r16_measDurationSymbols_r16    `madatory`
	Rmtc_Frequency_r16        ARFCN_ValueNR                              `madatory`
	Ref_SCS_CP_r16            RMTC_Config_r16_ref_SCS_CP_r16             `madatory`
	Rmtc_Bandwidth_r17        *RMTC_Config_r16_rmtc_Bandwidth_r17        `optional,ext-1`
	MeasDurationSymbols_v1700 *RMTC_Config_r16_measDurationSymbols_v1700 `optional,ext-1`
	Ref_SCS_CP_v1700          *RMTC_Config_r16_ref_SCS_CP_v1700          `optional,ext-1`
	Tci_StateInfo_r17         *RMTC_Config_r16_tci_StateInfo_r17         `optional,ext-1`
	Ref_BWPId_r17             *BWP_Id                                    `optional,ext-2`
}

func (ie *RMTC_Config_r16) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Rmtc_Bandwidth_r17 != nil || ie.MeasDurationSymbols_v1700 != nil || ie.Ref_SCS_CP_v1700 != nil || ie.Tci_StateInfo_r17 != nil || ie.Ref_BWPId_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Rmtc_SubframeOffset_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Rmtc_Periodicity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Rmtc_Periodicity_r16", err)
	}
	if ie.Rmtc_SubframeOffset_r16 != nil {
		if err = w.WriteInteger(*ie.Rmtc_SubframeOffset_r16, &aper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Encode Rmtc_SubframeOffset_r16", err)
		}
	}
	if err = ie.MeasDurationSymbols_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MeasDurationSymbols_r16", err)
	}
	if err = ie.Rmtc_Frequency_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Rmtc_Frequency_r16", err)
	}
	if err = ie.Ref_SCS_CP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Ref_SCS_CP_r16", err)
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.Rmtc_Bandwidth_r17 != nil || ie.MeasDurationSymbols_v1700 != nil || ie.Ref_SCS_CP_v1700 != nil || ie.Tci_StateInfo_r17 != nil, ie.Ref_BWPId_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RMTC_Config_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Rmtc_Bandwidth_r17 != nil, ie.MeasDurationSymbols_v1700 != nil, ie.Ref_SCS_CP_v1700 != nil, ie.Tci_StateInfo_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Rmtc_Bandwidth_r17 optional
			if ie.Rmtc_Bandwidth_r17 != nil {
				if err = ie.Rmtc_Bandwidth_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Rmtc_Bandwidth_r17", err)
				}
			}
			// encode MeasDurationSymbols_v1700 optional
			if ie.MeasDurationSymbols_v1700 != nil {
				if err = ie.MeasDurationSymbols_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MeasDurationSymbols_v1700", err)
				}
			}
			// encode Ref_SCS_CP_v1700 optional
			if ie.Ref_SCS_CP_v1700 != nil {
				if err = ie.Ref_SCS_CP_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ref_SCS_CP_v1700", err)
				}
			}
			// encode Tci_StateInfo_r17 optional
			if ie.Tci_StateInfo_r17 != nil {
				if err = ie.Tci_StateInfo_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Tci_StateInfo_r17", err)
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
			optionals_ext_2 := []bool{ie.Ref_BWPId_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Ref_BWPId_r17 optional
			if ie.Ref_BWPId_r17 != nil {
				if err = ie.Ref_BWPId_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ref_BWPId_r17", err)
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

func (ie *RMTC_Config_r16) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Rmtc_SubframeOffset_r16Present bool
	if Rmtc_SubframeOffset_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Rmtc_Periodicity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Rmtc_Periodicity_r16", err)
	}
	if Rmtc_SubframeOffset_r16Present {
		var tmp_int_Rmtc_SubframeOffset_r16 int64
		if tmp_int_Rmtc_SubframeOffset_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode Rmtc_SubframeOffset_r16", err)
		}
		ie.Rmtc_SubframeOffset_r16 = &tmp_int_Rmtc_SubframeOffset_r16
	}
	if err = ie.MeasDurationSymbols_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MeasDurationSymbols_r16", err)
	}
	if err = ie.Rmtc_Frequency_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Rmtc_Frequency_r16", err)
	}
	if err = ie.Ref_SCS_CP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Ref_SCS_CP_r16", err)
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Rmtc_Bandwidth_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MeasDurationSymbols_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ref_SCS_CP_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Tci_StateInfo_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Rmtc_Bandwidth_r17 optional
			if Rmtc_Bandwidth_r17Present {
				ie.Rmtc_Bandwidth_r17 = new(RMTC_Config_r16_rmtc_Bandwidth_r17)
				if err = ie.Rmtc_Bandwidth_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Rmtc_Bandwidth_r17", err)
				}
			}
			// decode MeasDurationSymbols_v1700 optional
			if MeasDurationSymbols_v1700Present {
				ie.MeasDurationSymbols_v1700 = new(RMTC_Config_r16_measDurationSymbols_v1700)
				if err = ie.MeasDurationSymbols_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode MeasDurationSymbols_v1700", err)
				}
			}
			// decode Ref_SCS_CP_v1700 optional
			if Ref_SCS_CP_v1700Present {
				ie.Ref_SCS_CP_v1700 = new(RMTC_Config_r16_ref_SCS_CP_v1700)
				if err = ie.Ref_SCS_CP_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ref_SCS_CP_v1700", err)
				}
			}
			// decode Tci_StateInfo_r17 optional
			if Tci_StateInfo_r17Present {
				ie.Tci_StateInfo_r17 = new(RMTC_Config_r16_tci_StateInfo_r17)
				if err = ie.Tci_StateInfo_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Tci_StateInfo_r17", err)
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

			Ref_BWPId_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Ref_BWPId_r17 optional
			if Ref_BWPId_r17Present {
				ie.Ref_BWPId_r17 = new(BWP_Id)
				if err = ie.Ref_BWPId_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ref_BWPId_r17", err)
				}
			}
		}
	}
	return nil
}
