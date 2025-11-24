package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SCellConfig struct {
	SCellIndex                       SCellIndex                                `madatory`
	SCellConfigCommon                *ServingCellConfigCommon                  `optional`
	SCellConfigDedicated             *ServingCellConfig                        `optional`
	Smtc                             *SSB_MTC                                  `optional,ext-1`
	SCellState_r16                   *SCellConfig_sCellState_r16               `optional,ext-2`
	SecondaryDRX_GroupConfig_r16     *SCellConfig_secondaryDRX_GroupConfig_r16 `optional,ext-2`
	PreConfGapStatus_r17             *uper.BitString                           `lb:maxNrofGapId_r17,ub:maxNrofGapId_r17,optional,ext-3`
	GoodServingCellEvaluationBFD_r17 *GoodServingCellEvaluation_r17            `optional,ext-3`
	SCellSIB20_r17                   *SCellSIB20_r17                           `optional,ext-3,setuprelease`
}

func (ie *SCellConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Smtc != nil || ie.SCellState_r16 != nil || ie.SecondaryDRX_GroupConfig_r16 != nil || ie.PreConfGapStatus_r17 != nil || ie.GoodServingCellEvaluationBFD_r17 != nil || ie.SCellSIB20_r17 != nil
	preambleBits := []bool{hasExtensions, ie.SCellConfigCommon != nil, ie.SCellConfigDedicated != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.SCellIndex.Encode(w); err != nil {
		return utils.WrapError("Encode SCellIndex", err)
	}
	if ie.SCellConfigCommon != nil {
		if err = ie.SCellConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode SCellConfigCommon", err)
		}
	}
	if ie.SCellConfigDedicated != nil {
		if err = ie.SCellConfigDedicated.Encode(w); err != nil {
			return utils.WrapError("Encode SCellConfigDedicated", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.Smtc != nil, ie.SCellState_r16 != nil || ie.SecondaryDRX_GroupConfig_r16 != nil, ie.PreConfGapStatus_r17 != nil || ie.GoodServingCellEvaluationBFD_r17 != nil || ie.SCellSIB20_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SCellConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Smtc != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Smtc optional
			if ie.Smtc != nil {
				if err = ie.Smtc.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Smtc", err)
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
			optionals_ext_2 := []bool{ie.SCellState_r16 != nil, ie.SecondaryDRX_GroupConfig_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SCellState_r16 optional
			if ie.SCellState_r16 != nil {
				if err = ie.SCellState_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SCellState_r16", err)
				}
			}
			// encode SecondaryDRX_GroupConfig_r16 optional
			if ie.SecondaryDRX_GroupConfig_r16 != nil {
				if err = ie.SecondaryDRX_GroupConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SecondaryDRX_GroupConfig_r16", err)
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
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.PreConfGapStatus_r17 != nil, ie.GoodServingCellEvaluationBFD_r17 != nil, ie.SCellSIB20_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode PreConfGapStatus_r17 optional
			if ie.PreConfGapStatus_r17 != nil {
				if err = extWriter.WriteBitString(ie.PreConfGapStatus_r17.Bytes, uint(ie.PreConfGapStatus_r17.NumBits), &uper.Constraint{Lb: maxNrofGapId_r17, Ub: maxNrofGapId_r17}, false); err != nil {
					return utils.WrapError("Encode PreConfGapStatus_r17", err)
				}
			}
			// encode GoodServingCellEvaluationBFD_r17 optional
			if ie.GoodServingCellEvaluationBFD_r17 != nil {
				if err = ie.GoodServingCellEvaluationBFD_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode GoodServingCellEvaluationBFD_r17", err)
				}
			}
			// encode SCellSIB20_r17 optional
			if ie.SCellSIB20_r17 != nil {
				tmp_SCellSIB20_r17 := utils.SetupRelease[*SCellSIB20_r17]{
					Setup: ie.SCellSIB20_r17,
				}
				if err = tmp_SCellSIB20_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SCellSIB20_r17", err)
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

func (ie *SCellConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var SCellConfigCommonPresent bool
	if SCellConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SCellConfigDedicatedPresent bool
	if SCellConfigDedicatedPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.SCellIndex.Decode(r); err != nil {
		return utils.WrapError("Decode SCellIndex", err)
	}
	if SCellConfigCommonPresent {
		ie.SCellConfigCommon = new(ServingCellConfigCommon)
		if err = ie.SCellConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode SCellConfigCommon", err)
		}
	}
	if SCellConfigDedicatedPresent {
		ie.SCellConfigDedicated = new(ServingCellConfig)
		if err = ie.SCellConfigDedicated.Decode(r); err != nil {
			return utils.WrapError("Decode SCellConfigDedicated", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
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

			SmtcPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Smtc optional
			if SmtcPresent {
				ie.Smtc = new(SSB_MTC)
				if err = ie.Smtc.Decode(extReader); err != nil {
					return utils.WrapError("Decode Smtc", err)
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

			SCellState_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SecondaryDRX_GroupConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SCellState_r16 optional
			if SCellState_r16Present {
				ie.SCellState_r16 = new(SCellConfig_sCellState_r16)
				if err = ie.SCellState_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SCellState_r16", err)
				}
			}
			// decode SecondaryDRX_GroupConfig_r16 optional
			if SecondaryDRX_GroupConfig_r16Present {
				ie.SecondaryDRX_GroupConfig_r16 = new(SCellConfig_secondaryDRX_GroupConfig_r16)
				if err = ie.SecondaryDRX_GroupConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SecondaryDRX_GroupConfig_r16", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			PreConfGapStatus_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			GoodServingCellEvaluationBFD_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SCellSIB20_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode PreConfGapStatus_r17 optional
			if PreConfGapStatus_r17Present {
				var tmp_bs_PreConfGapStatus_r17 []byte
				var n_PreConfGapStatus_r17 uint
				if tmp_bs_PreConfGapStatus_r17, n_PreConfGapStatus_r17, err = extReader.ReadBitString(&uper.Constraint{Lb: maxNrofGapId_r17, Ub: maxNrofGapId_r17}, false); err != nil {
					return utils.WrapError("Decode PreConfGapStatus_r17", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_PreConfGapStatus_r17,
					NumBits: uint64(n_PreConfGapStatus_r17),
				}
				ie.PreConfGapStatus_r17 = &tmp_bitstring
			}
			// decode GoodServingCellEvaluationBFD_r17 optional
			if GoodServingCellEvaluationBFD_r17Present {
				ie.GoodServingCellEvaluationBFD_r17 = new(GoodServingCellEvaluation_r17)
				if err = ie.GoodServingCellEvaluationBFD_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode GoodServingCellEvaluationBFD_r17", err)
				}
			}
			// decode SCellSIB20_r17 optional
			if SCellSIB20_r17Present {
				tmp_SCellSIB20_r17 := utils.SetupRelease[*SCellSIB20_r17]{}
				if err = tmp_SCellSIB20_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SCellSIB20_r17", err)
				}
				ie.SCellSIB20_r17 = tmp_SCellSIB20_r17.Setup
			}
		}
	}
	return nil
}
