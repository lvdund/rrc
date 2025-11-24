package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB2 struct {
	CellReselectionInfoCommon      *SIB2_cellReselectionInfoCommon      `lb:2,ub:maxNrofSS_BlocksToAverage,optional`
	CellReselectionServingFreqInfo *SIB2_cellReselectionServingFreqInfo `optional,ext`
	IntraFreqCellReselectionInfo   *SIB2_intraFreqCellReselectionInfo   `optional,ext`
	RelaxedMeasurement_r16         *SIB2_relaxedMeasurement_r16         `optional,ext-5`
	CellEquivalentSize_r17         *int64                               `lb:2,ub:16,optional,ext-6`
	RelaxedMeasurement_r17         *SIB2_relaxedMeasurement_r17         `optional,ext-6`
}

func (ie *SIB2) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.RelaxedMeasurement_r16 != nil || ie.CellEquivalentSize_r17 != nil || ie.RelaxedMeasurement_r17 != nil
	preambleBits := []bool{hasExtensions, ie.CellReselectionInfoCommon != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CellReselectionInfoCommon != nil {
		if err = ie.CellReselectionInfoCommon.Encode(w); err != nil {
			return utils.WrapError("Encode CellReselectionInfoCommon", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.RelaxedMeasurement_r16 != nil, ie.CellEquivalentSize_r17 != nil || ie.RelaxedMeasurement_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SIB2", err)
		}

		// encode extension group 5
		if extBitmap[4] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{ie.RelaxedMeasurement_r16 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode RelaxedMeasurement_r16 optional
			if ie.RelaxedMeasurement_r16 != nil {
				if err = ie.RelaxedMeasurement_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RelaxedMeasurement_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 6
		if extBitmap[5] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 6
			optionals_ext_6 := []bool{ie.CellEquivalentSize_r17 != nil, ie.RelaxedMeasurement_r17 != nil}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode CellEquivalentSize_r17 optional
			if ie.CellEquivalentSize_r17 != nil {
				if err = extWriter.WriteInteger(*ie.CellEquivalentSize_r17, &uper.Constraint{Lb: 2, Ub: 16}, false); err != nil {
					return utils.WrapError("Encode CellEquivalentSize_r17", err)
				}
			}
			// encode RelaxedMeasurement_r17 optional
			if ie.RelaxedMeasurement_r17 != nil {
				if err = ie.RelaxedMeasurement_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RelaxedMeasurement_r17", err)
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

func (ie *SIB2) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var CellReselectionInfoCommonPresent bool
	if CellReselectionInfoCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CellReselectionInfoCommonPresent {
		ie.CellReselectionInfoCommon = new(SIB2_cellReselectionInfoCommon)
		if err = ie.CellReselectionInfoCommon.Decode(r); err != nil {
			return utils.WrapError("Decode CellReselectionInfoCommon", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			RelaxedMeasurement_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode RelaxedMeasurement_r16 optional
			if RelaxedMeasurement_r16Present {
				ie.RelaxedMeasurement_r16 = new(SIB2_relaxedMeasurement_r16)
				if err = ie.RelaxedMeasurement_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode RelaxedMeasurement_r16", err)
				}
			}
		}
		// decode extension group 6
		if len(extBitmap) > 5 && extBitmap[5] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			CellEquivalentSize_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			RelaxedMeasurement_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode CellEquivalentSize_r17 optional
			if CellEquivalentSize_r17Present {
				var tmp_int_CellEquivalentSize_r17 int64
				if tmp_int_CellEquivalentSize_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 2, Ub: 16}, false); err != nil {
					return utils.WrapError("Decode CellEquivalentSize_r17", err)
				}
				ie.CellEquivalentSize_r17 = &tmp_int_CellEquivalentSize_r17
			}
			// decode RelaxedMeasurement_r17 optional
			if RelaxedMeasurement_r17Present {
				ie.RelaxedMeasurement_r17 = new(SIB2_relaxedMeasurement_r17)
				if err = ie.RelaxedMeasurement_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode RelaxedMeasurement_r17", err)
				}
			}
		}
	}
	return nil
}
