package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_ServingCellConfig struct {
	CodeBlockGroupTransmission     *PUSCH_CodeBlockGroupTransmission                       `optional,setuprelease`
	RateMatching                   *PUSCH_ServingCellConfig_rateMatching                   `optional`
	XOverhead                      *PUSCH_ServingCellConfig_xOverhead                      `optional`
	MaxMIMO_Layers                 *int64                                                  `lb:1,ub:4,optional,ext-1`
	ProcessingType2Enabled         *bool                                                   `optional,ext-1`
	MaxMIMO_LayersDCI_0_2_r16      *MaxMIMO_LayersDCI_0_2_r16                              `optional,ext-2,setuprelease`
	NrofHARQ_ProcessesForPUSCH_r17 *PUSCH_ServingCellConfig_nrofHARQ_ProcessesForPUSCH_r17 `optional,ext-3`
	UplinkHARQ_mode_r17            *UplinkHARQ_mode_r17                                    `optional,ext-3,setuprelease`
}

func (ie *PUSCH_ServingCellConfig) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.MaxMIMO_Layers != nil || ie.ProcessingType2Enabled != nil || ie.MaxMIMO_LayersDCI_0_2_r16 != nil || ie.NrofHARQ_ProcessesForPUSCH_r17 != nil || ie.UplinkHARQ_mode_r17 != nil
	preambleBits := []bool{hasExtensions, ie.CodeBlockGroupTransmission != nil, ie.RateMatching != nil, ie.XOverhead != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CodeBlockGroupTransmission != nil {
		tmp_CodeBlockGroupTransmission := utils.SetupRelease[*PUSCH_CodeBlockGroupTransmission]{
			Setup: ie.CodeBlockGroupTransmission,
		}
		if err = tmp_CodeBlockGroupTransmission.Encode(w); err != nil {
			return utils.WrapError("Encode CodeBlockGroupTransmission", err)
		}
	}
	if ie.RateMatching != nil {
		if err = ie.RateMatching.Encode(w); err != nil {
			return utils.WrapError("Encode RateMatching", err)
		}
	}
	if ie.XOverhead != nil {
		if err = ie.XOverhead.Encode(w); err != nil {
			return utils.WrapError("Encode XOverhead", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.MaxMIMO_Layers != nil || ie.ProcessingType2Enabled != nil, ie.MaxMIMO_LayersDCI_0_2_r16 != nil, ie.NrofHARQ_ProcessesForPUSCH_r17 != nil || ie.UplinkHARQ_mode_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PUSCH_ServingCellConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.MaxMIMO_Layers != nil, ie.ProcessingType2Enabled != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MaxMIMO_Layers optional
			if ie.MaxMIMO_Layers != nil {
				if err = extWriter.WriteInteger(*ie.MaxMIMO_Layers, &aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode MaxMIMO_Layers", err)
				}
			}
			// encode ProcessingType2Enabled optional
			if ie.ProcessingType2Enabled != nil {
				if err = extWriter.WriteBoolean(*ie.ProcessingType2Enabled); err != nil {
					return utils.WrapError("Encode ProcessingType2Enabled", err)
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
			optionals_ext_2 := []bool{ie.MaxMIMO_LayersDCI_0_2_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MaxMIMO_LayersDCI_0_2_r16 optional
			if ie.MaxMIMO_LayersDCI_0_2_r16 != nil {
				tmp_MaxMIMO_LayersDCI_0_2_r16 := utils.SetupRelease[*MaxMIMO_LayersDCI_0_2_r16]{
					Setup: ie.MaxMIMO_LayersDCI_0_2_r16,
				}
				if err = tmp_MaxMIMO_LayersDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxMIMO_LayersDCI_0_2_r16", err)
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
			optionals_ext_3 := []bool{ie.NrofHARQ_ProcessesForPUSCH_r17 != nil, ie.UplinkHARQ_mode_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode NrofHARQ_ProcessesForPUSCH_r17 optional
			if ie.NrofHARQ_ProcessesForPUSCH_r17 != nil {
				if err = ie.NrofHARQ_ProcessesForPUSCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode NrofHARQ_ProcessesForPUSCH_r17", err)
				}
			}
			// encode UplinkHARQ_mode_r17 optional
			if ie.UplinkHARQ_mode_r17 != nil {
				tmp_UplinkHARQ_mode_r17 := utils.SetupRelease[*UplinkHARQ_mode_r17]{
					Setup: ie.UplinkHARQ_mode_r17,
				}
				if err = tmp_UplinkHARQ_mode_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UplinkHARQ_mode_r17", err)
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

func (ie *PUSCH_ServingCellConfig) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var CodeBlockGroupTransmissionPresent bool
	if CodeBlockGroupTransmissionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RateMatchingPresent bool
	if RateMatchingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var XOverheadPresent bool
	if XOverheadPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CodeBlockGroupTransmissionPresent {
		tmp_CodeBlockGroupTransmission := utils.SetupRelease[*PUSCH_CodeBlockGroupTransmission]{}
		if err = tmp_CodeBlockGroupTransmission.Decode(r); err != nil {
			return utils.WrapError("Decode CodeBlockGroupTransmission", err)
		}
		ie.CodeBlockGroupTransmission = tmp_CodeBlockGroupTransmission.Setup
	}
	if RateMatchingPresent {
		ie.RateMatching = new(PUSCH_ServingCellConfig_rateMatching)
		if err = ie.RateMatching.Decode(r); err != nil {
			return utils.WrapError("Decode RateMatching", err)
		}
	}
	if XOverheadPresent {
		ie.XOverhead = new(PUSCH_ServingCellConfig_xOverhead)
		if err = ie.XOverhead.Decode(r); err != nil {
			return utils.WrapError("Decode XOverhead", err)
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			MaxMIMO_LayersPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ProcessingType2EnabledPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MaxMIMO_Layers optional
			if MaxMIMO_LayersPresent {
				var tmp_int_MaxMIMO_Layers int64
				if tmp_int_MaxMIMO_Layers, err = extReader.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode MaxMIMO_Layers", err)
				}
				ie.MaxMIMO_Layers = &tmp_int_MaxMIMO_Layers
			}
			// decode ProcessingType2Enabled optional
			if ProcessingType2EnabledPresent {
				var tmp_bool_ProcessingType2Enabled bool
				if tmp_bool_ProcessingType2Enabled, err = extReader.ReadBoolean(); err != nil {
					return utils.WrapError("Decode ProcessingType2Enabled", err)
				}
				ie.ProcessingType2Enabled = &tmp_bool_ProcessingType2Enabled
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			MaxMIMO_LayersDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MaxMIMO_LayersDCI_0_2_r16 optional
			if MaxMIMO_LayersDCI_0_2_r16Present {
				tmp_MaxMIMO_LayersDCI_0_2_r16 := utils.SetupRelease[*MaxMIMO_LayersDCI_0_2_r16]{}
				if err = tmp_MaxMIMO_LayersDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxMIMO_LayersDCI_0_2_r16", err)
				}
				ie.MaxMIMO_LayersDCI_0_2_r16 = tmp_MaxMIMO_LayersDCI_0_2_r16.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			NrofHARQ_ProcessesForPUSCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UplinkHARQ_mode_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode NrofHARQ_ProcessesForPUSCH_r17 optional
			if NrofHARQ_ProcessesForPUSCH_r17Present {
				ie.NrofHARQ_ProcessesForPUSCH_r17 = new(PUSCH_ServingCellConfig_nrofHARQ_ProcessesForPUSCH_r17)
				if err = ie.NrofHARQ_ProcessesForPUSCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode NrofHARQ_ProcessesForPUSCH_r17", err)
				}
			}
			// decode UplinkHARQ_mode_r17 optional
			if UplinkHARQ_mode_r17Present {
				tmp_UplinkHARQ_mode_r17 := utils.SetupRelease[*UplinkHARQ_mode_r17]{}
				if err = tmp_UplinkHARQ_mode_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UplinkHARQ_mode_r17", err)
				}
				ie.UplinkHARQ_mode_r17 = tmp_UplinkHARQ_mode_r17.Setup
			}
		}
	}
	return nil
}
