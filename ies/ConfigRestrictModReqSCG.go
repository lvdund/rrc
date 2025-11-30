package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ConfigRestrictModReqSCG struct {
	RequestedBC_MRDC                   *BandCombinationInfoSN `optional`
	RequestedP_MaxFR1                  *P_Max                 `optional`
	RequestedPDCCH_BlindDetectionSCG   *int64                 `lb:1,ub:15,optional,ext-1`
	RequestedP_MaxEUTRA                *P_Max                 `optional,ext-1`
	RequestedP_MaxFR2_r16              *P_Max                 `optional,ext-2`
	RequestedMaxInterFreqMeasIdSCG_r16 *int64                 `lb:1,ub:maxMeasIdentitiesMN,optional,ext-2`
	RequestedMaxIntraFreqMeasIdSCG_r16 *int64                 `lb:1,ub:maxMeasIdentitiesMN,optional,ext-2`
	RequestedToffset_r16               *T_Offset_r16          `optional,ext-2`
}

func (ie *ConfigRestrictModReqSCG) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.RequestedPDCCH_BlindDetectionSCG != nil || ie.RequestedP_MaxEUTRA != nil || ie.RequestedP_MaxFR2_r16 != nil || ie.RequestedMaxInterFreqMeasIdSCG_r16 != nil || ie.RequestedMaxIntraFreqMeasIdSCG_r16 != nil || ie.RequestedToffset_r16 != nil
	preambleBits := []bool{hasExtensions, ie.RequestedBC_MRDC != nil, ie.RequestedP_MaxFR1 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.RequestedBC_MRDC != nil {
		if err = ie.RequestedBC_MRDC.Encode(w); err != nil {
			return utils.WrapError("Encode RequestedBC_MRDC", err)
		}
	}
	if ie.RequestedP_MaxFR1 != nil {
		if err = ie.RequestedP_MaxFR1.Encode(w); err != nil {
			return utils.WrapError("Encode RequestedP_MaxFR1", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.RequestedPDCCH_BlindDetectionSCG != nil || ie.RequestedP_MaxEUTRA != nil, ie.RequestedP_MaxFR2_r16 != nil || ie.RequestedMaxInterFreqMeasIdSCG_r16 != nil || ie.RequestedMaxIntraFreqMeasIdSCG_r16 != nil || ie.RequestedToffset_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ConfigRestrictModReqSCG", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.RequestedPDCCH_BlindDetectionSCG != nil, ie.RequestedP_MaxEUTRA != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode RequestedPDCCH_BlindDetectionSCG optional
			if ie.RequestedPDCCH_BlindDetectionSCG != nil {
				if err = extWriter.WriteInteger(*ie.RequestedPDCCH_BlindDetectionSCG, &aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
					return utils.WrapError("Encode RequestedPDCCH_BlindDetectionSCG", err)
				}
			}
			// encode RequestedP_MaxEUTRA optional
			if ie.RequestedP_MaxEUTRA != nil {
				if err = ie.RequestedP_MaxEUTRA.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RequestedP_MaxEUTRA", err)
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
			optionals_ext_2 := []bool{ie.RequestedP_MaxFR2_r16 != nil, ie.RequestedMaxInterFreqMeasIdSCG_r16 != nil, ie.RequestedMaxIntraFreqMeasIdSCG_r16 != nil, ie.RequestedToffset_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode RequestedP_MaxFR2_r16 optional
			if ie.RequestedP_MaxFR2_r16 != nil {
				if err = ie.RequestedP_MaxFR2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RequestedP_MaxFR2_r16", err)
				}
			}
			// encode RequestedMaxInterFreqMeasIdSCG_r16 optional
			if ie.RequestedMaxInterFreqMeasIdSCG_r16 != nil {
				if err = extWriter.WriteInteger(*ie.RequestedMaxInterFreqMeasIdSCG_r16, &aper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
					return utils.WrapError("Encode RequestedMaxInterFreqMeasIdSCG_r16", err)
				}
			}
			// encode RequestedMaxIntraFreqMeasIdSCG_r16 optional
			if ie.RequestedMaxIntraFreqMeasIdSCG_r16 != nil {
				if err = extWriter.WriteInteger(*ie.RequestedMaxIntraFreqMeasIdSCG_r16, &aper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
					return utils.WrapError("Encode RequestedMaxIntraFreqMeasIdSCG_r16", err)
				}
			}
			// encode RequestedToffset_r16 optional
			if ie.RequestedToffset_r16 != nil {
				if err = ie.RequestedToffset_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RequestedToffset_r16", err)
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

func (ie *ConfigRestrictModReqSCG) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var RequestedBC_MRDCPresent bool
	if RequestedBC_MRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RequestedP_MaxFR1Present bool
	if RequestedP_MaxFR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	if RequestedBC_MRDCPresent {
		ie.RequestedBC_MRDC = new(BandCombinationInfoSN)
		if err = ie.RequestedBC_MRDC.Decode(r); err != nil {
			return utils.WrapError("Decode RequestedBC_MRDC", err)
		}
	}
	if RequestedP_MaxFR1Present {
		ie.RequestedP_MaxFR1 = new(P_Max)
		if err = ie.RequestedP_MaxFR1.Decode(r); err != nil {
			return utils.WrapError("Decode RequestedP_MaxFR1", err)
		}
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

			RequestedPDCCH_BlindDetectionSCGPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			RequestedP_MaxEUTRAPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode RequestedPDCCH_BlindDetectionSCG optional
			if RequestedPDCCH_BlindDetectionSCGPresent {
				var tmp_int_RequestedPDCCH_BlindDetectionSCG int64
				if tmp_int_RequestedPDCCH_BlindDetectionSCG, err = extReader.ReadInteger(&aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
					return utils.WrapError("Decode RequestedPDCCH_BlindDetectionSCG", err)
				}
				ie.RequestedPDCCH_BlindDetectionSCG = &tmp_int_RequestedPDCCH_BlindDetectionSCG
			}
			// decode RequestedP_MaxEUTRA optional
			if RequestedP_MaxEUTRAPresent {
				ie.RequestedP_MaxEUTRA = new(P_Max)
				if err = ie.RequestedP_MaxEUTRA.Decode(extReader); err != nil {
					return utils.WrapError("Decode RequestedP_MaxEUTRA", err)
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

			RequestedP_MaxFR2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			RequestedMaxInterFreqMeasIdSCG_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			RequestedMaxIntraFreqMeasIdSCG_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			RequestedToffset_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode RequestedP_MaxFR2_r16 optional
			if RequestedP_MaxFR2_r16Present {
				ie.RequestedP_MaxFR2_r16 = new(P_Max)
				if err = ie.RequestedP_MaxFR2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode RequestedP_MaxFR2_r16", err)
				}
			}
			// decode RequestedMaxInterFreqMeasIdSCG_r16 optional
			if RequestedMaxInterFreqMeasIdSCG_r16Present {
				var tmp_int_RequestedMaxInterFreqMeasIdSCG_r16 int64
				if tmp_int_RequestedMaxInterFreqMeasIdSCG_r16, err = extReader.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
					return utils.WrapError("Decode RequestedMaxInterFreqMeasIdSCG_r16", err)
				}
				ie.RequestedMaxInterFreqMeasIdSCG_r16 = &tmp_int_RequestedMaxInterFreqMeasIdSCG_r16
			}
			// decode RequestedMaxIntraFreqMeasIdSCG_r16 optional
			if RequestedMaxIntraFreqMeasIdSCG_r16Present {
				var tmp_int_RequestedMaxIntraFreqMeasIdSCG_r16 int64
				if tmp_int_RequestedMaxIntraFreqMeasIdSCG_r16, err = extReader.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
					return utils.WrapError("Decode RequestedMaxIntraFreqMeasIdSCG_r16", err)
				}
				ie.RequestedMaxIntraFreqMeasIdSCG_r16 = &tmp_int_RequestedMaxIntraFreqMeasIdSCG_r16
			}
			// decode RequestedToffset_r16 optional
			if RequestedToffset_r16Present {
				ie.RequestedToffset_r16 = new(T_Offset_r16)
				if err = ie.RequestedToffset_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode RequestedToffset_r16", err)
				}
			}
		}
	}
	return nil
}
