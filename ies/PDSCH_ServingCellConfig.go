package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_ServingCellConfig struct {
	CodeBlockGroupTransmission               *PDSCH_CodeBlockGroupTransmission                         `optional,setuprelease`
	XOverhead                                *PDSCH_ServingCellConfig_xOverhead                        `optional`
	NrofHARQ_ProcessesForPDSCH               *PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH       `optional`
	Pucch_Cell                               *ServCellIndex                                            `optional`
	MaxMIMO_Layers                           *int64                                                    `lb:1,ub:8,optional,ext-1`
	ProcessingType2Enabled                   *bool                                                     `optional,ext-1`
	Pdsch_CodeBlockGroupTransmissionList_r16 *PDSCH_CodeBlockGroupTransmissionList_r16                 `optional,ext-2,setuprelease`
	DownlinkHARQ_FeedbackDisabled_r17        *DownlinkHARQ_FeedbackDisabled_r17                        `optional,ext-3,setuprelease`
	NrofHARQ_ProcessesForPDSCH_v1700         *PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_v1700 `optional,ext-3`
}

func (ie *PDSCH_ServingCellConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.MaxMIMO_Layers != nil || ie.ProcessingType2Enabled != nil || ie.Pdsch_CodeBlockGroupTransmissionList_r16 != nil || ie.DownlinkHARQ_FeedbackDisabled_r17 != nil || ie.NrofHARQ_ProcessesForPDSCH_v1700 != nil
	preambleBits := []bool{hasExtensions, ie.CodeBlockGroupTransmission != nil, ie.XOverhead != nil, ie.NrofHARQ_ProcessesForPDSCH != nil, ie.Pucch_Cell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CodeBlockGroupTransmission != nil {
		tmp_CodeBlockGroupTransmission := utils.SetupRelease[*PDSCH_CodeBlockGroupTransmission]{
			Setup: ie.CodeBlockGroupTransmission,
		}
		if err = tmp_CodeBlockGroupTransmission.Encode(w); err != nil {
			return utils.WrapError("Encode CodeBlockGroupTransmission", err)
		}
	}
	if ie.XOverhead != nil {
		if err = ie.XOverhead.Encode(w); err != nil {
			return utils.WrapError("Encode XOverhead", err)
		}
	}
	if ie.NrofHARQ_ProcessesForPDSCH != nil {
		if err = ie.NrofHARQ_ProcessesForPDSCH.Encode(w); err != nil {
			return utils.WrapError("Encode NrofHARQ_ProcessesForPDSCH", err)
		}
	}
	if ie.Pucch_Cell != nil {
		if err = ie.Pucch_Cell.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_Cell", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.MaxMIMO_Layers != nil || ie.ProcessingType2Enabled != nil, ie.Pdsch_CodeBlockGroupTransmissionList_r16 != nil, ie.DownlinkHARQ_FeedbackDisabled_r17 != nil || ie.NrofHARQ_ProcessesForPDSCH_v1700 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PDSCH_ServingCellConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.MaxMIMO_Layers != nil, ie.ProcessingType2Enabled != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MaxMIMO_Layers optional
			if ie.MaxMIMO_Layers != nil {
				if err = extWriter.WriteInteger(*ie.MaxMIMO_Layers, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
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
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.Pdsch_CodeBlockGroupTransmissionList_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Pdsch_CodeBlockGroupTransmissionList_r16 optional
			if ie.Pdsch_CodeBlockGroupTransmissionList_r16 != nil {
				tmp_Pdsch_CodeBlockGroupTransmissionList_r16 := utils.SetupRelease[*PDSCH_CodeBlockGroupTransmissionList_r16]{
					Setup: ie.Pdsch_CodeBlockGroupTransmissionList_r16,
				}
				if err = tmp_Pdsch_CodeBlockGroupTransmissionList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_CodeBlockGroupTransmissionList_r16", err)
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
			optionals_ext_3 := []bool{ie.DownlinkHARQ_FeedbackDisabled_r17 != nil, ie.NrofHARQ_ProcessesForPDSCH_v1700 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode DownlinkHARQ_FeedbackDisabled_r17 optional
			if ie.DownlinkHARQ_FeedbackDisabled_r17 != nil {
				tmp_DownlinkHARQ_FeedbackDisabled_r17 := utils.SetupRelease[*DownlinkHARQ_FeedbackDisabled_r17]{
					Setup: ie.DownlinkHARQ_FeedbackDisabled_r17,
				}
				if err = tmp_DownlinkHARQ_FeedbackDisabled_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DownlinkHARQ_FeedbackDisabled_r17", err)
				}
			}
			// encode NrofHARQ_ProcessesForPDSCH_v1700 optional
			if ie.NrofHARQ_ProcessesForPDSCH_v1700 != nil {
				if err = ie.NrofHARQ_ProcessesForPDSCH_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode NrofHARQ_ProcessesForPDSCH_v1700", err)
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

func (ie *PDSCH_ServingCellConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var CodeBlockGroupTransmissionPresent bool
	if CodeBlockGroupTransmissionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var XOverheadPresent bool
	if XOverheadPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NrofHARQ_ProcessesForPDSCHPresent bool
	if NrofHARQ_ProcessesForPDSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_CellPresent bool
	if Pucch_CellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CodeBlockGroupTransmissionPresent {
		tmp_CodeBlockGroupTransmission := utils.SetupRelease[*PDSCH_CodeBlockGroupTransmission]{}
		if err = tmp_CodeBlockGroupTransmission.Decode(r); err != nil {
			return utils.WrapError("Decode CodeBlockGroupTransmission", err)
		}
		ie.CodeBlockGroupTransmission = tmp_CodeBlockGroupTransmission.Setup
	}
	if XOverheadPresent {
		ie.XOverhead = new(PDSCH_ServingCellConfig_xOverhead)
		if err = ie.XOverhead.Decode(r); err != nil {
			return utils.WrapError("Decode XOverhead", err)
		}
	}
	if NrofHARQ_ProcessesForPDSCHPresent {
		ie.NrofHARQ_ProcessesForPDSCH = new(PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH)
		if err = ie.NrofHARQ_ProcessesForPDSCH.Decode(r); err != nil {
			return utils.WrapError("Decode NrofHARQ_ProcessesForPDSCH", err)
		}
	}
	if Pucch_CellPresent {
		ie.Pucch_Cell = new(ServCellIndex)
		if err = ie.Pucch_Cell.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_Cell", err)
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
				if tmp_int_MaxMIMO_Layers, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
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

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Pdsch_CodeBlockGroupTransmissionList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Pdsch_CodeBlockGroupTransmissionList_r16 optional
			if Pdsch_CodeBlockGroupTransmissionList_r16Present {
				tmp_Pdsch_CodeBlockGroupTransmissionList_r16 := utils.SetupRelease[*PDSCH_CodeBlockGroupTransmissionList_r16]{}
				if err = tmp_Pdsch_CodeBlockGroupTransmissionList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_CodeBlockGroupTransmissionList_r16", err)
				}
				ie.Pdsch_CodeBlockGroupTransmissionList_r16 = tmp_Pdsch_CodeBlockGroupTransmissionList_r16.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			DownlinkHARQ_FeedbackDisabled_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			NrofHARQ_ProcessesForPDSCH_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode DownlinkHARQ_FeedbackDisabled_r17 optional
			if DownlinkHARQ_FeedbackDisabled_r17Present {
				tmp_DownlinkHARQ_FeedbackDisabled_r17 := utils.SetupRelease[*DownlinkHARQ_FeedbackDisabled_r17]{}
				if err = tmp_DownlinkHARQ_FeedbackDisabled_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode DownlinkHARQ_FeedbackDisabled_r17", err)
				}
				ie.DownlinkHARQ_FeedbackDisabled_r17 = tmp_DownlinkHARQ_FeedbackDisabled_r17.Setup
			}
			// decode NrofHARQ_ProcessesForPDSCH_v1700 optional
			if NrofHARQ_ProcessesForPDSCH_v1700Present {
				ie.NrofHARQ_ProcessesForPDSCH_v1700 = new(PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_v1700)
				if err = ie.NrofHARQ_ProcessesForPDSCH_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode NrofHARQ_ProcessesForPDSCH_v1700", err)
				}
			}
		}
	}
	return nil
}
