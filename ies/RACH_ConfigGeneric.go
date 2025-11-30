package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RACH_ConfigGeneric struct {
	Prach_ConfigurationIndex                 int64                                                        `lb:0,ub:255,madatory`
	Msg1_FDM                                 RACH_ConfigGeneric_msg1_FDM                                  `madatory`
	Msg1_FrequencyStart                      int64                                                        `lb:0,ub:maxNrofPhysicalResourceBlocks_1,madatory`
	ZeroCorrelationZoneConfig                int64                                                        `lb:0,ub:15,madatory`
	PreambleReceivedTargetPower              int64                                                        `lb:-202,ub:-60,madatory`
	PreambleTransMax                         RACH_ConfigGeneric_preambleTransMax                          `madatory`
	PowerRampingStep                         RACH_ConfigGeneric_powerRampingStep                          `madatory`
	Ra_ResponseWindow                        RACH_ConfigGeneric_ra_ResponseWindow                         `madatory`
	Prach_ConfigurationPeriodScaling_IAB_r16 *RACH_ConfigGeneric_prach_ConfigurationPeriodScaling_IAB_r16 `optional,ext-1`
	Prach_ConfigurationFrameOffset_IAB_r16   *int64                                                       `lb:0,ub:63,optional,ext-1`
	Prach_ConfigurationSOffset_IAB_r16       *int64                                                       `lb:0,ub:39,optional,ext-1`
	Ra_ResponseWindow_v1610                  *RACH_ConfigGeneric_ra_ResponseWindow_v1610                  `optional,ext-1`
	Prach_ConfigurationIndex_v1610           *int64                                                       `lb:256,ub:262,optional,ext-1`
	Ra_ResponseWindow_v1700                  *RACH_ConfigGeneric_ra_ResponseWindow_v1700                  `optional,ext-2`
}

func (ie *RACH_ConfigGeneric) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Prach_ConfigurationPeriodScaling_IAB_r16 != nil || ie.Prach_ConfigurationFrameOffset_IAB_r16 != nil || ie.Prach_ConfigurationSOffset_IAB_r16 != nil || ie.Ra_ResponseWindow_v1610 != nil || ie.Prach_ConfigurationIndex_v1610 != nil || ie.Ra_ResponseWindow_v1700 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.Prach_ConfigurationIndex, &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return utils.WrapError("WriteInteger Prach_ConfigurationIndex", err)
	}
	if err = ie.Msg1_FDM.Encode(w); err != nil {
		return utils.WrapError("Encode Msg1_FDM", err)
	}
	if err = w.WriteInteger(ie.Msg1_FrequencyStart, &aper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
		return utils.WrapError("WriteInteger Msg1_FrequencyStart", err)
	}
	if err = w.WriteInteger(ie.ZeroCorrelationZoneConfig, &aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger ZeroCorrelationZoneConfig", err)
	}
	if err = w.WriteInteger(ie.PreambleReceivedTargetPower, &aper.Constraint{Lb: -202, Ub: -60}, false); err != nil {
		return utils.WrapError("WriteInteger PreambleReceivedTargetPower", err)
	}
	if err = ie.PreambleTransMax.Encode(w); err != nil {
		return utils.WrapError("Encode PreambleTransMax", err)
	}
	if err = ie.PowerRampingStep.Encode(w); err != nil {
		return utils.WrapError("Encode PowerRampingStep", err)
	}
	if err = ie.Ra_ResponseWindow.Encode(w); err != nil {
		return utils.WrapError("Encode Ra_ResponseWindow", err)
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.Prach_ConfigurationPeriodScaling_IAB_r16 != nil || ie.Prach_ConfigurationFrameOffset_IAB_r16 != nil || ie.Prach_ConfigurationSOffset_IAB_r16 != nil || ie.Ra_ResponseWindow_v1610 != nil || ie.Prach_ConfigurationIndex_v1610 != nil, ie.Ra_ResponseWindow_v1700 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RACH_ConfigGeneric", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Prach_ConfigurationPeriodScaling_IAB_r16 != nil, ie.Prach_ConfigurationFrameOffset_IAB_r16 != nil, ie.Prach_ConfigurationSOffset_IAB_r16 != nil, ie.Ra_ResponseWindow_v1610 != nil, ie.Prach_ConfigurationIndex_v1610 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Prach_ConfigurationPeriodScaling_IAB_r16 optional
			if ie.Prach_ConfigurationPeriodScaling_IAB_r16 != nil {
				if err = ie.Prach_ConfigurationPeriodScaling_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Prach_ConfigurationPeriodScaling_IAB_r16", err)
				}
			}
			// encode Prach_ConfigurationFrameOffset_IAB_r16 optional
			if ie.Prach_ConfigurationFrameOffset_IAB_r16 != nil {
				if err = extWriter.WriteInteger(*ie.Prach_ConfigurationFrameOffset_IAB_r16, &aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
					return utils.WrapError("Encode Prach_ConfigurationFrameOffset_IAB_r16", err)
				}
			}
			// encode Prach_ConfigurationSOffset_IAB_r16 optional
			if ie.Prach_ConfigurationSOffset_IAB_r16 != nil {
				if err = extWriter.WriteInteger(*ie.Prach_ConfigurationSOffset_IAB_r16, &aper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
					return utils.WrapError("Encode Prach_ConfigurationSOffset_IAB_r16", err)
				}
			}
			// encode Ra_ResponseWindow_v1610 optional
			if ie.Ra_ResponseWindow_v1610 != nil {
				if err = ie.Ra_ResponseWindow_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ra_ResponseWindow_v1610", err)
				}
			}
			// encode Prach_ConfigurationIndex_v1610 optional
			if ie.Prach_ConfigurationIndex_v1610 != nil {
				if err = extWriter.WriteInteger(*ie.Prach_ConfigurationIndex_v1610, &aper.Constraint{Lb: 256, Ub: 262}, false); err != nil {
					return utils.WrapError("Encode Prach_ConfigurationIndex_v1610", err)
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
			optionals_ext_2 := []bool{ie.Ra_ResponseWindow_v1700 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Ra_ResponseWindow_v1700 optional
			if ie.Ra_ResponseWindow_v1700 != nil {
				if err = ie.Ra_ResponseWindow_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ra_ResponseWindow_v1700", err)
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

func (ie *RACH_ConfigGeneric) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_Prach_ConfigurationIndex int64
	if tmp_int_Prach_ConfigurationIndex, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return utils.WrapError("ReadInteger Prach_ConfigurationIndex", err)
	}
	ie.Prach_ConfigurationIndex = tmp_int_Prach_ConfigurationIndex
	if err = ie.Msg1_FDM.Decode(r); err != nil {
		return utils.WrapError("Decode Msg1_FDM", err)
	}
	var tmp_int_Msg1_FrequencyStart int64
	if tmp_int_Msg1_FrequencyStart, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
		return utils.WrapError("ReadInteger Msg1_FrequencyStart", err)
	}
	ie.Msg1_FrequencyStart = tmp_int_Msg1_FrequencyStart
	var tmp_int_ZeroCorrelationZoneConfig int64
	if tmp_int_ZeroCorrelationZoneConfig, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger ZeroCorrelationZoneConfig", err)
	}
	ie.ZeroCorrelationZoneConfig = tmp_int_ZeroCorrelationZoneConfig
	var tmp_int_PreambleReceivedTargetPower int64
	if tmp_int_PreambleReceivedTargetPower, err = r.ReadInteger(&aper.Constraint{Lb: -202, Ub: -60}, false); err != nil {
		return utils.WrapError("ReadInteger PreambleReceivedTargetPower", err)
	}
	ie.PreambleReceivedTargetPower = tmp_int_PreambleReceivedTargetPower
	if err = ie.PreambleTransMax.Decode(r); err != nil {
		return utils.WrapError("Decode PreambleTransMax", err)
	}
	if err = ie.PowerRampingStep.Decode(r); err != nil {
		return utils.WrapError("Decode PowerRampingStep", err)
	}
	if err = ie.Ra_ResponseWindow.Decode(r); err != nil {
		return utils.WrapError("Decode Ra_ResponseWindow", err)
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

			Prach_ConfigurationPeriodScaling_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Prach_ConfigurationFrameOffset_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Prach_ConfigurationSOffset_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ra_ResponseWindow_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Prach_ConfigurationIndex_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Prach_ConfigurationPeriodScaling_IAB_r16 optional
			if Prach_ConfigurationPeriodScaling_IAB_r16Present {
				ie.Prach_ConfigurationPeriodScaling_IAB_r16 = new(RACH_ConfigGeneric_prach_ConfigurationPeriodScaling_IAB_r16)
				if err = ie.Prach_ConfigurationPeriodScaling_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Prach_ConfigurationPeriodScaling_IAB_r16", err)
				}
			}
			// decode Prach_ConfigurationFrameOffset_IAB_r16 optional
			if Prach_ConfigurationFrameOffset_IAB_r16Present {
				var tmp_int_Prach_ConfigurationFrameOffset_IAB_r16 int64
				if tmp_int_Prach_ConfigurationFrameOffset_IAB_r16, err = extReader.ReadInteger(&aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
					return utils.WrapError("Decode Prach_ConfigurationFrameOffset_IAB_r16", err)
				}
				ie.Prach_ConfigurationFrameOffset_IAB_r16 = &tmp_int_Prach_ConfigurationFrameOffset_IAB_r16
			}
			// decode Prach_ConfigurationSOffset_IAB_r16 optional
			if Prach_ConfigurationSOffset_IAB_r16Present {
				var tmp_int_Prach_ConfigurationSOffset_IAB_r16 int64
				if tmp_int_Prach_ConfigurationSOffset_IAB_r16, err = extReader.ReadInteger(&aper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
					return utils.WrapError("Decode Prach_ConfigurationSOffset_IAB_r16", err)
				}
				ie.Prach_ConfigurationSOffset_IAB_r16 = &tmp_int_Prach_ConfigurationSOffset_IAB_r16
			}
			// decode Ra_ResponseWindow_v1610 optional
			if Ra_ResponseWindow_v1610Present {
				ie.Ra_ResponseWindow_v1610 = new(RACH_ConfigGeneric_ra_ResponseWindow_v1610)
				if err = ie.Ra_ResponseWindow_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ra_ResponseWindow_v1610", err)
				}
			}
			// decode Prach_ConfigurationIndex_v1610 optional
			if Prach_ConfigurationIndex_v1610Present {
				var tmp_int_Prach_ConfigurationIndex_v1610 int64
				if tmp_int_Prach_ConfigurationIndex_v1610, err = extReader.ReadInteger(&aper.Constraint{Lb: 256, Ub: 262}, false); err != nil {
					return utils.WrapError("Decode Prach_ConfigurationIndex_v1610", err)
				}
				ie.Prach_ConfigurationIndex_v1610 = &tmp_int_Prach_ConfigurationIndex_v1610
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Ra_ResponseWindow_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Ra_ResponseWindow_v1700 optional
			if Ra_ResponseWindow_v1700Present {
				ie.Ra_ResponseWindow_v1700 = new(RACH_ConfigGeneric_ra_ResponseWindow_v1700)
				if err = ie.Ra_ResponseWindow_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ra_ResponseWindow_v1700", err)
				}
			}
		}
	}
	return nil
}
