package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReconfigurationWithSync struct {
	SpCellConfigCommon         *ServingCellConfigCommon                      `optional`
	NewUE_Identity             RNTI_Value                                    `madatory`
	T304                       ReconfigurationWithSync_t304                  `madatory`
	Rach_ConfigDedicated       *ReconfigurationWithSync_rach_ConfigDedicated `optional`
	Smtc                       *SSB_MTC                                      `optional,ext-1`
	Daps_UplinkPowerConfig_r16 *DAPS_UplinkPowerConfig_r16                   `optional,ext-2`
	Sl_PathSwitchConfig_r17    *SL_PathSwitchConfig_r17                      `optional,ext-3`
}

func (ie *ReconfigurationWithSync) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Smtc != nil || ie.Daps_UplinkPowerConfig_r16 != nil || ie.Sl_PathSwitchConfig_r17 != nil
	preambleBits := []bool{hasExtensions, ie.SpCellConfigCommon != nil, ie.Rach_ConfigDedicated != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SpCellConfigCommon != nil {
		if err = ie.SpCellConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode SpCellConfigCommon", err)
		}
	}
	if err = ie.NewUE_Identity.Encode(w); err != nil {
		return utils.WrapError("Encode NewUE_Identity", err)
	}
	if err = ie.T304.Encode(w); err != nil {
		return utils.WrapError("Encode T304", err)
	}
	if ie.Rach_ConfigDedicated != nil {
		if err = ie.Rach_ConfigDedicated.Encode(w); err != nil {
			return utils.WrapError("Encode Rach_ConfigDedicated", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.Smtc != nil, ie.Daps_UplinkPowerConfig_r16 != nil, ie.Sl_PathSwitchConfig_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ReconfigurationWithSync", err)
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
			optionals_ext_2 := []bool{ie.Daps_UplinkPowerConfig_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Daps_UplinkPowerConfig_r16 optional
			if ie.Daps_UplinkPowerConfig_r16 != nil {
				if err = ie.Daps_UplinkPowerConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Daps_UplinkPowerConfig_r16", err)
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
			optionals_ext_3 := []bool{ie.Sl_PathSwitchConfig_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sl_PathSwitchConfig_r17 optional
			if ie.Sl_PathSwitchConfig_r17 != nil {
				if err = ie.Sl_PathSwitchConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_PathSwitchConfig_r17", err)
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

func (ie *ReconfigurationWithSync) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var SpCellConfigCommonPresent bool
	if SpCellConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Rach_ConfigDedicatedPresent bool
	if Rach_ConfigDedicatedPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SpCellConfigCommonPresent {
		ie.SpCellConfigCommon = new(ServingCellConfigCommon)
		if err = ie.SpCellConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode SpCellConfigCommon", err)
		}
	}
	if err = ie.NewUE_Identity.Decode(r); err != nil {
		return utils.WrapError("Decode NewUE_Identity", err)
	}
	if err = ie.T304.Decode(r); err != nil {
		return utils.WrapError("Decode T304", err)
	}
	if Rach_ConfigDedicatedPresent {
		ie.Rach_ConfigDedicated = new(ReconfigurationWithSync_rach_ConfigDedicated)
		if err = ie.Rach_ConfigDedicated.Decode(r); err != nil {
			return utils.WrapError("Decode Rach_ConfigDedicated", err)
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

			Daps_UplinkPowerConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Daps_UplinkPowerConfig_r16 optional
			if Daps_UplinkPowerConfig_r16Present {
				ie.Daps_UplinkPowerConfig_r16 = new(DAPS_UplinkPowerConfig_r16)
				if err = ie.Daps_UplinkPowerConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Daps_UplinkPowerConfig_r16", err)
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

			Sl_PathSwitchConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sl_PathSwitchConfig_r17 optional
			if Sl_PathSwitchConfig_r17Present {
				ie.Sl_PathSwitchConfig_r17 = new(SL_PathSwitchConfig_r17)
				if err = ie.Sl_PathSwitchConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_PathSwitchConfig_r17", err)
				}
			}
		}
	}
	return nil
}
