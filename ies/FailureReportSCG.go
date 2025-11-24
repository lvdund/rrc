package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FailureReportSCG struct {
	FailureType           FailureReportSCG_failureType        `madatory`
	MeasResultFreqList    *MeasResultFreqList                 `optional`
	MeasResultSCG_Failure *[]byte                             `optional`
	LocationInfo_r16      *LocationInfo_r16                   `optional,ext-1`
	FailureType_v1610     *FailureReportSCG_failureType_v1610 `optional,ext-1`
	PreviousPSCellId_r17  *PSCellId_r17                       `optional,ext-2`
	FailedPSCellId_r17    *PSCellId_r17                       `optional,ext-2`
	TimeSCGFailure_r17    *int64                              `lb:0,ub:1023,optional,ext-2`
	PerRAInfoList_r17     *PerRAInfoList_r16                  `optional,ext-2`
}

func (ie *FailureReportSCG) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.LocationInfo_r16 != nil || ie.FailureType_v1610 != nil || ie.PreviousPSCellId_r17 != nil || ie.FailedPSCellId_r17 != nil || ie.TimeSCGFailure_r17 != nil || ie.PerRAInfoList_r17 != nil
	preambleBits := []bool{hasExtensions, ie.MeasResultFreqList != nil, ie.MeasResultSCG_Failure != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.FailureType.Encode(w); err != nil {
		return utils.WrapError("Encode FailureType", err)
	}
	if ie.MeasResultFreqList != nil {
		if err = ie.MeasResultFreqList.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultFreqList", err)
		}
	}
	if ie.MeasResultSCG_Failure != nil {
		if err = w.WriteOctetString(*ie.MeasResultSCG_Failure, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode MeasResultSCG_Failure", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.LocationInfo_r16 != nil || ie.FailureType_v1610 != nil, ie.PreviousPSCellId_r17 != nil || ie.FailedPSCellId_r17 != nil || ie.TimeSCGFailure_r17 != nil || ie.PerRAInfoList_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap FailureReportSCG", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.LocationInfo_r16 != nil, ie.FailureType_v1610 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode LocationInfo_r16 optional
			if ie.LocationInfo_r16 != nil {
				if err = ie.LocationInfo_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode LocationInfo_r16", err)
				}
			}
			// encode FailureType_v1610 optional
			if ie.FailureType_v1610 != nil {
				if err = ie.FailureType_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FailureType_v1610", err)
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
			optionals_ext_2 := []bool{ie.PreviousPSCellId_r17 != nil, ie.FailedPSCellId_r17 != nil, ie.TimeSCGFailure_r17 != nil, ie.PerRAInfoList_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode PreviousPSCellId_r17 optional
			if ie.PreviousPSCellId_r17 != nil {
				if err = ie.PreviousPSCellId_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PreviousPSCellId_r17", err)
				}
			}
			// encode FailedPSCellId_r17 optional
			if ie.FailedPSCellId_r17 != nil {
				if err = ie.FailedPSCellId_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FailedPSCellId_r17", err)
				}
			}
			// encode TimeSCGFailure_r17 optional
			if ie.TimeSCGFailure_r17 != nil {
				if err = extWriter.WriteInteger(*ie.TimeSCGFailure_r17, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
					return utils.WrapError("Encode TimeSCGFailure_r17", err)
				}
			}
			// encode PerRAInfoList_r17 optional
			if ie.PerRAInfoList_r17 != nil {
				if err = ie.PerRAInfoList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PerRAInfoList_r17", err)
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

func (ie *FailureReportSCG) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultFreqListPresent bool
	if MeasResultFreqListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultSCG_FailurePresent bool
	if MeasResultSCG_FailurePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.FailureType.Decode(r); err != nil {
		return utils.WrapError("Decode FailureType", err)
	}
	if MeasResultFreqListPresent {
		ie.MeasResultFreqList = new(MeasResultFreqList)
		if err = ie.MeasResultFreqList.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultFreqList", err)
		}
	}
	if MeasResultSCG_FailurePresent {
		var tmp_os_MeasResultSCG_Failure []byte
		if tmp_os_MeasResultSCG_Failure, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode MeasResultSCG_Failure", err)
		}
		ie.MeasResultSCG_Failure = &tmp_os_MeasResultSCG_Failure
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

			LocationInfo_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FailureType_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode LocationInfo_r16 optional
			if LocationInfo_r16Present {
				ie.LocationInfo_r16 = new(LocationInfo_r16)
				if err = ie.LocationInfo_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode LocationInfo_r16", err)
				}
			}
			// decode FailureType_v1610 optional
			if FailureType_v1610Present {
				ie.FailureType_v1610 = new(FailureReportSCG_failureType_v1610)
				if err = ie.FailureType_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode FailureType_v1610", err)
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

			PreviousPSCellId_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FailedPSCellId_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			TimeSCGFailure_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PerRAInfoList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode PreviousPSCellId_r17 optional
			if PreviousPSCellId_r17Present {
				ie.PreviousPSCellId_r17 = new(PSCellId_r17)
				if err = ie.PreviousPSCellId_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode PreviousPSCellId_r17", err)
				}
			}
			// decode FailedPSCellId_r17 optional
			if FailedPSCellId_r17Present {
				ie.FailedPSCellId_r17 = new(PSCellId_r17)
				if err = ie.FailedPSCellId_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode FailedPSCellId_r17", err)
				}
			}
			// decode TimeSCGFailure_r17 optional
			if TimeSCGFailure_r17Present {
				var tmp_int_TimeSCGFailure_r17 int64
				if tmp_int_TimeSCGFailure_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
					return utils.WrapError("Decode TimeSCGFailure_r17", err)
				}
				ie.TimeSCGFailure_r17 = &tmp_int_TimeSCGFailure_r17
			}
			// decode PerRAInfoList_r17 optional
			if PerRAInfoList_r17Present {
				ie.PerRAInfoList_r17 = new(PerRAInfoList_r16)
				if err = ie.PerRAInfoList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode PerRAInfoList_r17", err)
				}
			}
		}
	}
	return nil
}
