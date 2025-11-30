package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SDAP_Parameters struct {
	As_ReflectiveQoS  *SDAP_Parameters_as_ReflectiveQoS  `optional`
	Sdap_QOS_IAB_r16  *SDAP_Parameters_sdap_QOS_IAB_r16  `optional,ext-1`
	SdapHeaderIAB_r16 *SDAP_Parameters_sdapHeaderIAB_r16 `optional,ext-1`
}

func (ie *SDAP_Parameters) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Sdap_QOS_IAB_r16 != nil || ie.SdapHeaderIAB_r16 != nil
	preambleBits := []bool{hasExtensions, ie.As_ReflectiveQoS != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.As_ReflectiveQoS != nil {
		if err = ie.As_ReflectiveQoS.Encode(w); err != nil {
			return utils.WrapError("Encode As_ReflectiveQoS", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Sdap_QOS_IAB_r16 != nil || ie.SdapHeaderIAB_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SDAP_Parameters", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Sdap_QOS_IAB_r16 != nil, ie.SdapHeaderIAB_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sdap_QOS_IAB_r16 optional
			if ie.Sdap_QOS_IAB_r16 != nil {
				if err = ie.Sdap_QOS_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sdap_QOS_IAB_r16", err)
				}
			}
			// encode SdapHeaderIAB_r16 optional
			if ie.SdapHeaderIAB_r16 != nil {
				if err = ie.SdapHeaderIAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SdapHeaderIAB_r16", err)
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

func (ie *SDAP_Parameters) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var As_ReflectiveQoSPresent bool
	if As_ReflectiveQoSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if As_ReflectiveQoSPresent {
		ie.As_ReflectiveQoS = new(SDAP_Parameters_as_ReflectiveQoS)
		if err = ie.As_ReflectiveQoS.Decode(r); err != nil {
			return utils.WrapError("Decode As_ReflectiveQoS", err)
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

			Sdap_QOS_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SdapHeaderIAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sdap_QOS_IAB_r16 optional
			if Sdap_QOS_IAB_r16Present {
				ie.Sdap_QOS_IAB_r16 = new(SDAP_Parameters_sdap_QOS_IAB_r16)
				if err = ie.Sdap_QOS_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sdap_QOS_IAB_r16", err)
				}
			}
			// decode SdapHeaderIAB_r16 optional
			if SdapHeaderIAB_r16Present {
				ie.SdapHeaderIAB_r16 = new(SDAP_Parameters_sdapHeaderIAB_r16)
				if err = ie.SdapHeaderIAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SdapHeaderIAB_r16", err)
				}
			}
		}
	}
	return nil
}
