package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RACH_ConfigGenericTwoStepRA_r16 struct {
	MsgA_PRACH_ConfigurationIndex_r16    *int64                                                             `lb:0,ub:262,optional`
	MsgA_RO_FDM_r16                      *RACH_ConfigGenericTwoStepRA_r16_msgA_RO_FDM_r16                   `optional`
	MsgA_RO_FrequencyStart_r16           *int64                                                             `lb:0,ub:maxNrofPhysicalResourceBlocks_1,optional`
	MsgA_ZeroCorrelationZoneConfig_r16   *int64                                                             `lb:0,ub:15,optional`
	MsgA_PreamblePowerRampingStep_r16    *RACH_ConfigGenericTwoStepRA_r16_msgA_PreamblePowerRampingStep_r16 `optional`
	MsgA_PreambleReceivedTargetPower_r16 *int64                                                             `lb:-202,ub:-60,optional`
	MsgB_ResponseWindow_r16              *RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_r16           `optional`
	PreambleTransMax_r16                 *RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16              `optional`
	MsgB_ResponseWindow_v1700            *RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_v1700         `optional,ext-1`
}

func (ie *RACH_ConfigGenericTwoStepRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.MsgB_ResponseWindow_v1700 != nil
	preambleBits := []bool{hasExtensions, ie.MsgA_PRACH_ConfigurationIndex_r16 != nil, ie.MsgA_RO_FDM_r16 != nil, ie.MsgA_RO_FrequencyStart_r16 != nil, ie.MsgA_ZeroCorrelationZoneConfig_r16 != nil, ie.MsgA_PreamblePowerRampingStep_r16 != nil, ie.MsgA_PreambleReceivedTargetPower_r16 != nil, ie.MsgB_ResponseWindow_r16 != nil, ie.PreambleTransMax_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MsgA_PRACH_ConfigurationIndex_r16 != nil {
		if err = w.WriteInteger(*ie.MsgA_PRACH_ConfigurationIndex_r16, &uper.Constraint{Lb: 0, Ub: 262}, false); err != nil {
			return utils.WrapError("Encode MsgA_PRACH_ConfigurationIndex_r16", err)
		}
	}
	if ie.MsgA_RO_FDM_r16 != nil {
		if err = ie.MsgA_RO_FDM_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_RO_FDM_r16", err)
		}
	}
	if ie.MsgA_RO_FrequencyStart_r16 != nil {
		if err = w.WriteInteger(*ie.MsgA_RO_FrequencyStart_r16, &uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Encode MsgA_RO_FrequencyStart_r16", err)
		}
	}
	if ie.MsgA_ZeroCorrelationZoneConfig_r16 != nil {
		if err = w.WriteInteger(*ie.MsgA_ZeroCorrelationZoneConfig_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode MsgA_ZeroCorrelationZoneConfig_r16", err)
		}
	}
	if ie.MsgA_PreamblePowerRampingStep_r16 != nil {
		if err = ie.MsgA_PreamblePowerRampingStep_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_PreamblePowerRampingStep_r16", err)
		}
	}
	if ie.MsgA_PreambleReceivedTargetPower_r16 != nil {
		if err = w.WriteInteger(*ie.MsgA_PreambleReceivedTargetPower_r16, &uper.Constraint{Lb: -202, Ub: -60}, false); err != nil {
			return utils.WrapError("Encode MsgA_PreambleReceivedTargetPower_r16", err)
		}
	}
	if ie.MsgB_ResponseWindow_r16 != nil {
		if err = ie.MsgB_ResponseWindow_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgB_ResponseWindow_r16", err)
		}
	}
	if ie.PreambleTransMax_r16 != nil {
		if err = ie.PreambleTransMax_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PreambleTransMax_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.MsgB_ResponseWindow_v1700 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RACH_ConfigGenericTwoStepRA_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.MsgB_ResponseWindow_v1700 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MsgB_ResponseWindow_v1700 optional
			if ie.MsgB_ResponseWindow_v1700 != nil {
				if err = ie.MsgB_ResponseWindow_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MsgB_ResponseWindow_v1700", err)
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

func (ie *RACH_ConfigGenericTwoStepRA_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_PRACH_ConfigurationIndex_r16Present bool
	if MsgA_PRACH_ConfigurationIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_RO_FDM_r16Present bool
	if MsgA_RO_FDM_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_RO_FrequencyStart_r16Present bool
	if MsgA_RO_FrequencyStart_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_ZeroCorrelationZoneConfig_r16Present bool
	if MsgA_ZeroCorrelationZoneConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_PreamblePowerRampingStep_r16Present bool
	if MsgA_PreamblePowerRampingStep_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_PreambleReceivedTargetPower_r16Present bool
	if MsgA_PreambleReceivedTargetPower_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgB_ResponseWindow_r16Present bool
	if MsgB_ResponseWindow_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PreambleTransMax_r16Present bool
	if PreambleTransMax_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MsgA_PRACH_ConfigurationIndex_r16Present {
		var tmp_int_MsgA_PRACH_ConfigurationIndex_r16 int64
		if tmp_int_MsgA_PRACH_ConfigurationIndex_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 262}, false); err != nil {
			return utils.WrapError("Decode MsgA_PRACH_ConfigurationIndex_r16", err)
		}
		ie.MsgA_PRACH_ConfigurationIndex_r16 = &tmp_int_MsgA_PRACH_ConfigurationIndex_r16
	}
	if MsgA_RO_FDM_r16Present {
		ie.MsgA_RO_FDM_r16 = new(RACH_ConfigGenericTwoStepRA_r16_msgA_RO_FDM_r16)
		if err = ie.MsgA_RO_FDM_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_RO_FDM_r16", err)
		}
	}
	if MsgA_RO_FrequencyStart_r16Present {
		var tmp_int_MsgA_RO_FrequencyStart_r16 int64
		if tmp_int_MsgA_RO_FrequencyStart_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Decode MsgA_RO_FrequencyStart_r16", err)
		}
		ie.MsgA_RO_FrequencyStart_r16 = &tmp_int_MsgA_RO_FrequencyStart_r16
	}
	if MsgA_ZeroCorrelationZoneConfig_r16Present {
		var tmp_int_MsgA_ZeroCorrelationZoneConfig_r16 int64
		if tmp_int_MsgA_ZeroCorrelationZoneConfig_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode MsgA_ZeroCorrelationZoneConfig_r16", err)
		}
		ie.MsgA_ZeroCorrelationZoneConfig_r16 = &tmp_int_MsgA_ZeroCorrelationZoneConfig_r16
	}
	if MsgA_PreamblePowerRampingStep_r16Present {
		ie.MsgA_PreamblePowerRampingStep_r16 = new(RACH_ConfigGenericTwoStepRA_r16_msgA_PreamblePowerRampingStep_r16)
		if err = ie.MsgA_PreamblePowerRampingStep_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_PreamblePowerRampingStep_r16", err)
		}
	}
	if MsgA_PreambleReceivedTargetPower_r16Present {
		var tmp_int_MsgA_PreambleReceivedTargetPower_r16 int64
		if tmp_int_MsgA_PreambleReceivedTargetPower_r16, err = r.ReadInteger(&uper.Constraint{Lb: -202, Ub: -60}, false); err != nil {
			return utils.WrapError("Decode MsgA_PreambleReceivedTargetPower_r16", err)
		}
		ie.MsgA_PreambleReceivedTargetPower_r16 = &tmp_int_MsgA_PreambleReceivedTargetPower_r16
	}
	if MsgB_ResponseWindow_r16Present {
		ie.MsgB_ResponseWindow_r16 = new(RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_r16)
		if err = ie.MsgB_ResponseWindow_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgB_ResponseWindow_r16", err)
		}
	}
	if PreambleTransMax_r16Present {
		ie.PreambleTransMax_r16 = new(RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16)
		if err = ie.PreambleTransMax_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PreambleTransMax_r16", err)
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

			MsgB_ResponseWindow_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MsgB_ResponseWindow_v1700 optional
			if MsgB_ResponseWindow_v1700Present {
				ie.MsgB_ResponseWindow_v1700 = new(RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_v1700)
				if err = ie.MsgB_ResponseWindow_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode MsgB_ResponseWindow_v1700", err)
				}
			}
		}
	}
	return nil
}
