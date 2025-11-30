package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DMRS_DownlinkConfig struct {
	Dmrs_Type               *DMRS_DownlinkConfig_dmrs_Type               `optional`
	Dmrs_AdditionalPosition *DMRS_DownlinkConfig_dmrs_AdditionalPosition `optional`
	MaxLength               *DMRS_DownlinkConfig_maxLength               `optional`
	ScramblingID0           *int64                                       `lb:0,ub:65535,optional`
	ScramblingID1           *int64                                       `lb:0,ub:65535,optional`
	PhaseTrackingRS         *PTRS_DownlinkConfig                         `optional,setuprelease`
	Dmrs_Downlink_r16       *DMRS_DownlinkConfig_dmrs_Downlink_r16       `optional,ext-1`
}

func (ie *DMRS_DownlinkConfig) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Dmrs_Downlink_r16 != nil
	preambleBits := []bool{hasExtensions, ie.Dmrs_Type != nil, ie.Dmrs_AdditionalPosition != nil, ie.MaxLength != nil, ie.ScramblingID0 != nil, ie.ScramblingID1 != nil, ie.PhaseTrackingRS != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Dmrs_Type != nil {
		if err = ie.Dmrs_Type.Encode(w); err != nil {
			return utils.WrapError("Encode Dmrs_Type", err)
		}
	}
	if ie.Dmrs_AdditionalPosition != nil {
		if err = ie.Dmrs_AdditionalPosition.Encode(w); err != nil {
			return utils.WrapError("Encode Dmrs_AdditionalPosition", err)
		}
	}
	if ie.MaxLength != nil {
		if err = ie.MaxLength.Encode(w); err != nil {
			return utils.WrapError("Encode MaxLength", err)
		}
	}
	if ie.ScramblingID0 != nil {
		if err = w.WriteInteger(*ie.ScramblingID0, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode ScramblingID0", err)
		}
	}
	if ie.ScramblingID1 != nil {
		if err = w.WriteInteger(*ie.ScramblingID1, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode ScramblingID1", err)
		}
	}
	if ie.PhaseTrackingRS != nil {
		tmp_PhaseTrackingRS := utils.SetupRelease[*PTRS_DownlinkConfig]{
			Setup: ie.PhaseTrackingRS,
		}
		if err = tmp_PhaseTrackingRS.Encode(w); err != nil {
			return utils.WrapError("Encode PhaseTrackingRS", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Dmrs_Downlink_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap DMRS_DownlinkConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Dmrs_Downlink_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Dmrs_Downlink_r16 optional
			if ie.Dmrs_Downlink_r16 != nil {
				if err = ie.Dmrs_Downlink_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dmrs_Downlink_r16", err)
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

func (ie *DMRS_DownlinkConfig) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Dmrs_TypePresent bool
	if Dmrs_TypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dmrs_AdditionalPositionPresent bool
	if Dmrs_AdditionalPositionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxLengthPresent bool
	if MaxLengthPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ScramblingID0Present bool
	if ScramblingID0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ScramblingID1Present bool
	if ScramblingID1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PhaseTrackingRSPresent bool
	if PhaseTrackingRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Dmrs_TypePresent {
		ie.Dmrs_Type = new(DMRS_DownlinkConfig_dmrs_Type)
		if err = ie.Dmrs_Type.Decode(r); err != nil {
			return utils.WrapError("Decode Dmrs_Type", err)
		}
	}
	if Dmrs_AdditionalPositionPresent {
		ie.Dmrs_AdditionalPosition = new(DMRS_DownlinkConfig_dmrs_AdditionalPosition)
		if err = ie.Dmrs_AdditionalPosition.Decode(r); err != nil {
			return utils.WrapError("Decode Dmrs_AdditionalPosition", err)
		}
	}
	if MaxLengthPresent {
		ie.MaxLength = new(DMRS_DownlinkConfig_maxLength)
		if err = ie.MaxLength.Decode(r); err != nil {
			return utils.WrapError("Decode MaxLength", err)
		}
	}
	if ScramblingID0Present {
		var tmp_int_ScramblingID0 int64
		if tmp_int_ScramblingID0, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode ScramblingID0", err)
		}
		ie.ScramblingID0 = &tmp_int_ScramblingID0
	}
	if ScramblingID1Present {
		var tmp_int_ScramblingID1 int64
		if tmp_int_ScramblingID1, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode ScramblingID1", err)
		}
		ie.ScramblingID1 = &tmp_int_ScramblingID1
	}
	if PhaseTrackingRSPresent {
		tmp_PhaseTrackingRS := utils.SetupRelease[*PTRS_DownlinkConfig]{}
		if err = tmp_PhaseTrackingRS.Decode(r); err != nil {
			return utils.WrapError("Decode PhaseTrackingRS", err)
		}
		ie.PhaseTrackingRS = tmp_PhaseTrackingRS.Setup
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Dmrs_Downlink_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Dmrs_Downlink_r16 optional
			if Dmrs_Downlink_r16Present {
				ie.Dmrs_Downlink_r16 = new(DMRS_DownlinkConfig_dmrs_Downlink_r16)
				if err = ie.Dmrs_Downlink_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dmrs_Downlink_r16", err)
				}
			}
		}
	}
	return nil
}
