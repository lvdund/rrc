package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RLC_BearerConfig struct {
	LogicalChannelIdentity        LogicalChannelIdentity              `madatory`
	ServedRadioBearer             *RLC_BearerConfig_servedRadioBearer `optional`
	ReestablishRLC                *RLC_BearerConfig_reestablishRLC    `optional`
	Rlc_Config                    *RLC_Config                         `optional`
	Mac_LogicalChannelConfig      *LogicalChannelConfig               `optional`
	Rlc_Config_v1610              *RLC_Config_v1610                   `optional,ext-1`
	Rlc_Config_v1700              *RLC_Config_v1700                   `optional,ext-2`
	LogicalChannelIdentityExt_r17 *LogicalChannelIdentityExt_r17      `optional,ext-2`
	MulticastRLC_BearerConfig_r17 *MulticastRLC_BearerConfig_r17      `optional,ext-2`
	ServedRadioBearerSRB4_r17     *SRB_Identity_v1700                 `optional,ext-2`
}

func (ie *RLC_BearerConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Rlc_Config_v1610 != nil || ie.Rlc_Config_v1700 != nil || ie.LogicalChannelIdentityExt_r17 != nil || ie.MulticastRLC_BearerConfig_r17 != nil || ie.ServedRadioBearerSRB4_r17 != nil
	preambleBits := []bool{hasExtensions, ie.ServedRadioBearer != nil, ie.ReestablishRLC != nil, ie.Rlc_Config != nil, ie.Mac_LogicalChannelConfig != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.LogicalChannelIdentity.Encode(w); err != nil {
		return utils.WrapError("Encode LogicalChannelIdentity", err)
	}
	if ie.ServedRadioBearer != nil {
		if err = ie.ServedRadioBearer.Encode(w); err != nil {
			return utils.WrapError("Encode ServedRadioBearer", err)
		}
	}
	if ie.ReestablishRLC != nil {
		if err = ie.ReestablishRLC.Encode(w); err != nil {
			return utils.WrapError("Encode ReestablishRLC", err)
		}
	}
	if ie.Rlc_Config != nil {
		if err = ie.Rlc_Config.Encode(w); err != nil {
			return utils.WrapError("Encode Rlc_Config", err)
		}
	}
	if ie.Mac_LogicalChannelConfig != nil {
		if err = ie.Mac_LogicalChannelConfig.Encode(w); err != nil {
			return utils.WrapError("Encode Mac_LogicalChannelConfig", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.Rlc_Config_v1610 != nil, ie.Rlc_Config_v1700 != nil || ie.LogicalChannelIdentityExt_r17 != nil || ie.MulticastRLC_BearerConfig_r17 != nil || ie.ServedRadioBearerSRB4_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RLC_BearerConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Rlc_Config_v1610 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Rlc_Config_v1610 optional
			if ie.Rlc_Config_v1610 != nil {
				if err = ie.Rlc_Config_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Rlc_Config_v1610", err)
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
			optionals_ext_2 := []bool{ie.Rlc_Config_v1700 != nil, ie.LogicalChannelIdentityExt_r17 != nil, ie.MulticastRLC_BearerConfig_r17 != nil, ie.ServedRadioBearerSRB4_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Rlc_Config_v1700 optional
			if ie.Rlc_Config_v1700 != nil {
				if err = ie.Rlc_Config_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Rlc_Config_v1700", err)
				}
			}
			// encode LogicalChannelIdentityExt_r17 optional
			if ie.LogicalChannelIdentityExt_r17 != nil {
				if err = ie.LogicalChannelIdentityExt_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode LogicalChannelIdentityExt_r17", err)
				}
			}
			// encode MulticastRLC_BearerConfig_r17 optional
			if ie.MulticastRLC_BearerConfig_r17 != nil {
				if err = ie.MulticastRLC_BearerConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MulticastRLC_BearerConfig_r17", err)
				}
			}
			// encode ServedRadioBearerSRB4_r17 optional
			if ie.ServedRadioBearerSRB4_r17 != nil {
				if err = ie.ServedRadioBearerSRB4_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ServedRadioBearerSRB4_r17", err)
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

func (ie *RLC_BearerConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ServedRadioBearerPresent bool
	if ServedRadioBearerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ReestablishRLCPresent bool
	if ReestablishRLCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Rlc_ConfigPresent bool
	if Rlc_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Mac_LogicalChannelConfigPresent bool
	if Mac_LogicalChannelConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.LogicalChannelIdentity.Decode(r); err != nil {
		return utils.WrapError("Decode LogicalChannelIdentity", err)
	}
	if ServedRadioBearerPresent {
		ie.ServedRadioBearer = new(RLC_BearerConfig_servedRadioBearer)
		if err = ie.ServedRadioBearer.Decode(r); err != nil {
			return utils.WrapError("Decode ServedRadioBearer", err)
		}
	}
	if ReestablishRLCPresent {
		ie.ReestablishRLC = new(RLC_BearerConfig_reestablishRLC)
		if err = ie.ReestablishRLC.Decode(r); err != nil {
			return utils.WrapError("Decode ReestablishRLC", err)
		}
	}
	if Rlc_ConfigPresent {
		ie.Rlc_Config = new(RLC_Config)
		if err = ie.Rlc_Config.Decode(r); err != nil {
			return utils.WrapError("Decode Rlc_Config", err)
		}
	}
	if Mac_LogicalChannelConfigPresent {
		ie.Mac_LogicalChannelConfig = new(LogicalChannelConfig)
		if err = ie.Mac_LogicalChannelConfig.Decode(r); err != nil {
			return utils.WrapError("Decode Mac_LogicalChannelConfig", err)
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

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Rlc_Config_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Rlc_Config_v1610 optional
			if Rlc_Config_v1610Present {
				ie.Rlc_Config_v1610 = new(RLC_Config_v1610)
				if err = ie.Rlc_Config_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode Rlc_Config_v1610", err)
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

			Rlc_Config_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			LogicalChannelIdentityExt_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MulticastRLC_BearerConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ServedRadioBearerSRB4_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Rlc_Config_v1700 optional
			if Rlc_Config_v1700Present {
				ie.Rlc_Config_v1700 = new(RLC_Config_v1700)
				if err = ie.Rlc_Config_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode Rlc_Config_v1700", err)
				}
			}
			// decode LogicalChannelIdentityExt_r17 optional
			if LogicalChannelIdentityExt_r17Present {
				ie.LogicalChannelIdentityExt_r17 = new(LogicalChannelIdentityExt_r17)
				if err = ie.LogicalChannelIdentityExt_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode LogicalChannelIdentityExt_r17", err)
				}
			}
			// decode MulticastRLC_BearerConfig_r17 optional
			if MulticastRLC_BearerConfig_r17Present {
				ie.MulticastRLC_BearerConfig_r17 = new(MulticastRLC_BearerConfig_r17)
				if err = ie.MulticastRLC_BearerConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MulticastRLC_BearerConfig_r17", err)
				}
			}
			// decode ServedRadioBearerSRB4_r17 optional
			if ServedRadioBearerSRB4_r17Present {
				ie.ServedRadioBearerSRB4_r17 = new(SRB_Identity_v1700)
				if err = ie.ServedRadioBearerSRB4_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ServedRadioBearerSRB4_r17", err)
				}
			}
		}
	}
	return nil
}
