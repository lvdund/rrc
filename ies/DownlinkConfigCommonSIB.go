package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DownlinkConfigCommonSIB struct {
	FrequencyInfoDL               FrequencyInfoDL_SIB `madatory`
	InitialDownlinkBWP            BWP_DownlinkCommon  `madatory`
	Bcch_Config                   BCCH_Config         `madatory`
	Pcch_Config                   PCCH_Config         `madatory`
	Pei_Config_r17                *PEI_Config_r17     `optional,ext-1`
	InitialDownlinkBWP_RedCap_r17 *BWP_DownlinkCommon `optional,ext-1`
}

func (ie *DownlinkConfigCommonSIB) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Pei_Config_r17 != nil || ie.InitialDownlinkBWP_RedCap_r17 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.FrequencyInfoDL.Encode(w); err != nil {
		return utils.WrapError("Encode FrequencyInfoDL", err)
	}
	if err = ie.InitialDownlinkBWP.Encode(w); err != nil {
		return utils.WrapError("Encode InitialDownlinkBWP", err)
	}
	if err = ie.Bcch_Config.Encode(w); err != nil {
		return utils.WrapError("Encode Bcch_Config", err)
	}
	if err = ie.Pcch_Config.Encode(w); err != nil {
		return utils.WrapError("Encode Pcch_Config", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Pei_Config_r17 != nil || ie.InitialDownlinkBWP_RedCap_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap DownlinkConfigCommonSIB", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Pei_Config_r17 != nil, ie.InitialDownlinkBWP_RedCap_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Pei_Config_r17 optional
			if ie.Pei_Config_r17 != nil {
				if err = ie.Pei_Config_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pei_Config_r17", err)
				}
			}
			// encode InitialDownlinkBWP_RedCap_r17 optional
			if ie.InitialDownlinkBWP_RedCap_r17 != nil {
				if err = ie.InitialDownlinkBWP_RedCap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode InitialDownlinkBWP_RedCap_r17", err)
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

func (ie *DownlinkConfigCommonSIB) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.FrequencyInfoDL.Decode(r); err != nil {
		return utils.WrapError("Decode FrequencyInfoDL", err)
	}
	if err = ie.InitialDownlinkBWP.Decode(r); err != nil {
		return utils.WrapError("Decode InitialDownlinkBWP", err)
	}
	if err = ie.Bcch_Config.Decode(r); err != nil {
		return utils.WrapError("Decode Bcch_Config", err)
	}
	if err = ie.Pcch_Config.Decode(r); err != nil {
		return utils.WrapError("Decode Pcch_Config", err)
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

			Pei_Config_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			InitialDownlinkBWP_RedCap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Pei_Config_r17 optional
			if Pei_Config_r17Present {
				ie.Pei_Config_r17 = new(PEI_Config_r17)
				if err = ie.Pei_Config_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pei_Config_r17", err)
				}
			}
			// decode InitialDownlinkBWP_RedCap_r17 optional
			if InitialDownlinkBWP_RedCap_r17Present {
				ie.InitialDownlinkBWP_RedCap_r17 = new(BWP_DownlinkCommon)
				if err = ie.InitialDownlinkBWP_RedCap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode InitialDownlinkBWP_RedCap_r17", err)
				}
			}
		}
	}
	return nil
}
