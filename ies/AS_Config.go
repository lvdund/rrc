package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type AS_Config struct {
	RrcReconfiguration     []byte                          `madatory`
	SourceRB_SN_Config     *[]byte                         `optional,ext-1`
	SourceSCG_NR_Config    *[]byte                         `optional,ext-1`
	SourceSCG_EUTRA_Config *[]byte                         `optional,ext-1`
	SourceSCG_Configured   *AS_Config_sourceSCG_Configured `optional,ext-2`
	Sdt_Config_r17         *SDT_Config_r17                 `optional,ext-3`
}

func (ie *AS_Config) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.SourceRB_SN_Config != nil || ie.SourceSCG_NR_Config != nil || ie.SourceSCG_EUTRA_Config != nil || ie.SourceSCG_Configured != nil || ie.Sdt_Config_r17 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteOctetString(ie.RrcReconfiguration, nil, false); err != nil {
		return utils.WrapError("WriteOctetString RrcReconfiguration", err)
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.SourceRB_SN_Config != nil || ie.SourceSCG_NR_Config != nil || ie.SourceSCG_EUTRA_Config != nil, ie.SourceSCG_Configured != nil, ie.Sdt_Config_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap AS_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.SourceRB_SN_Config != nil, ie.SourceSCG_NR_Config != nil, ie.SourceSCG_EUTRA_Config != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SourceRB_SN_Config optional
			if ie.SourceRB_SN_Config != nil {
				if err = extWriter.WriteOctetString(*ie.SourceRB_SN_Config, nil, false); err != nil {
					return utils.WrapError("Encode SourceRB_SN_Config", err)
				}
			}
			// encode SourceSCG_NR_Config optional
			if ie.SourceSCG_NR_Config != nil {
				if err = extWriter.WriteOctetString(*ie.SourceSCG_NR_Config, nil, false); err != nil {
					return utils.WrapError("Encode SourceSCG_NR_Config", err)
				}
			}
			// encode SourceSCG_EUTRA_Config optional
			if ie.SourceSCG_EUTRA_Config != nil {
				if err = extWriter.WriteOctetString(*ie.SourceSCG_EUTRA_Config, nil, false); err != nil {
					return utils.WrapError("Encode SourceSCG_EUTRA_Config", err)
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
			optionals_ext_2 := []bool{ie.SourceSCG_Configured != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SourceSCG_Configured optional
			if ie.SourceSCG_Configured != nil {
				if err = ie.SourceSCG_Configured.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SourceSCG_Configured", err)
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
			optionals_ext_3 := []bool{ie.Sdt_Config_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sdt_Config_r17 optional
			if ie.Sdt_Config_r17 != nil {
				if err = ie.Sdt_Config_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sdt_Config_r17", err)
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

func (ie *AS_Config) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_os_RrcReconfiguration []byte
	if tmp_os_RrcReconfiguration, err = r.ReadOctetString(nil, false); err != nil {
		return utils.WrapError("ReadOctetString RrcReconfiguration", err)
	}
	ie.RrcReconfiguration = tmp_os_RrcReconfiguration

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

			SourceRB_SN_ConfigPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SourceSCG_NR_ConfigPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SourceSCG_EUTRA_ConfigPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SourceRB_SN_Config optional
			if SourceRB_SN_ConfigPresent {
				var tmp_os_SourceRB_SN_Config []byte
				if tmp_os_SourceRB_SN_Config, err = extReader.ReadOctetString(nil, false); err != nil {
					return utils.WrapError("Decode SourceRB_SN_Config", err)
				}
				ie.SourceRB_SN_Config = &tmp_os_SourceRB_SN_Config
			}
			// decode SourceSCG_NR_Config optional
			if SourceSCG_NR_ConfigPresent {
				var tmp_os_SourceSCG_NR_Config []byte
				if tmp_os_SourceSCG_NR_Config, err = extReader.ReadOctetString(nil, false); err != nil {
					return utils.WrapError("Decode SourceSCG_NR_Config", err)
				}
				ie.SourceSCG_NR_Config = &tmp_os_SourceSCG_NR_Config
			}
			// decode SourceSCG_EUTRA_Config optional
			if SourceSCG_EUTRA_ConfigPresent {
				var tmp_os_SourceSCG_EUTRA_Config []byte
				if tmp_os_SourceSCG_EUTRA_Config, err = extReader.ReadOctetString(nil, false); err != nil {
					return utils.WrapError("Decode SourceSCG_EUTRA_Config", err)
				}
				ie.SourceSCG_EUTRA_Config = &tmp_os_SourceSCG_EUTRA_Config
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SourceSCG_ConfiguredPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SourceSCG_Configured optional
			if SourceSCG_ConfiguredPresent {
				ie.SourceSCG_Configured = new(AS_Config_sourceSCG_Configured)
				if err = ie.SourceSCG_Configured.Decode(extReader); err != nil {
					return utils.WrapError("Decode SourceSCG_Configured", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Sdt_Config_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sdt_Config_r17 optional
			if Sdt_Config_r17Present {
				ie.Sdt_Config_r17 = new(SDT_Config_r17)
				if err = ie.Sdt_Config_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sdt_Config_r17", err)
				}
			}
		}
	}
	return nil
}
