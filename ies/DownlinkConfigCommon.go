package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DownlinkConfigCommon struct {
	FrequencyInfoDL               *FrequencyInfoDL    `optional`
	InitialDownlinkBWP            *BWP_DownlinkCommon `optional`
	InitialDownlinkBWP_RedCap_r17 *BWP_DownlinkCommon `optional,ext-1`
}

func (ie *DownlinkConfigCommon) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.InitialDownlinkBWP_RedCap_r17 != nil
	preambleBits := []bool{hasExtensions, ie.FrequencyInfoDL != nil, ie.InitialDownlinkBWP != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FrequencyInfoDL != nil {
		if err = ie.FrequencyInfoDL.Encode(w); err != nil {
			return utils.WrapError("Encode FrequencyInfoDL", err)
		}
	}
	if ie.InitialDownlinkBWP != nil {
		if err = ie.InitialDownlinkBWP.Encode(w); err != nil {
			return utils.WrapError("Encode InitialDownlinkBWP", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.InitialDownlinkBWP_RedCap_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap DownlinkConfigCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.InitialDownlinkBWP_RedCap_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
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

func (ie *DownlinkConfigCommon) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var FrequencyInfoDLPresent bool
	if FrequencyInfoDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var InitialDownlinkBWPPresent bool
	if InitialDownlinkBWPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if FrequencyInfoDLPresent {
		ie.FrequencyInfoDL = new(FrequencyInfoDL)
		if err = ie.FrequencyInfoDL.Decode(r); err != nil {
			return utils.WrapError("Decode FrequencyInfoDL", err)
		}
	}
	if InitialDownlinkBWPPresent {
		ie.InitialDownlinkBWP = new(BWP_DownlinkCommon)
		if err = ie.InitialDownlinkBWP.Decode(r); err != nil {
			return utils.WrapError("Decode InitialDownlinkBWP", err)
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			InitialDownlinkBWP_RedCap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
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
