package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RLC_Parameters struct {
	Am_WithShortSN               *RLC_Parameters_am_WithShortSN               `optional`
	Um_WithShortSN               *RLC_Parameters_um_WithShortSN               `optional`
	Um_WithLongSN                *RLC_Parameters_um_WithLongSN                `optional`
	ExtendedT_PollRetransmit_r16 *RLC_Parameters_extendedT_PollRetransmit_r16 `optional,ext-1`
	ExtendedT_StatusProhibit_r16 *RLC_Parameters_extendedT_StatusProhibit_r16 `optional,ext-1`
	Am_WithLongSN_RedCap_r17     *RLC_Parameters_am_WithLongSN_RedCap_r17     `optional,ext-2`
}

func (ie *RLC_Parameters) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.ExtendedT_PollRetransmit_r16 != nil || ie.ExtendedT_StatusProhibit_r16 != nil || ie.Am_WithLongSN_RedCap_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Am_WithShortSN != nil, ie.Um_WithShortSN != nil, ie.Um_WithLongSN != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Am_WithShortSN != nil {
		if err = ie.Am_WithShortSN.Encode(w); err != nil {
			return utils.WrapError("Encode Am_WithShortSN", err)
		}
	}
	if ie.Um_WithShortSN != nil {
		if err = ie.Um_WithShortSN.Encode(w); err != nil {
			return utils.WrapError("Encode Um_WithShortSN", err)
		}
	}
	if ie.Um_WithLongSN != nil {
		if err = ie.Um_WithLongSN.Encode(w); err != nil {
			return utils.WrapError("Encode Um_WithLongSN", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.ExtendedT_PollRetransmit_r16 != nil || ie.ExtendedT_StatusProhibit_r16 != nil, ie.Am_WithLongSN_RedCap_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RLC_Parameters", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ExtendedT_PollRetransmit_r16 != nil, ie.ExtendedT_StatusProhibit_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ExtendedT_PollRetransmit_r16 optional
			if ie.ExtendedT_PollRetransmit_r16 != nil {
				if err = ie.ExtendedT_PollRetransmit_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ExtendedT_PollRetransmit_r16", err)
				}
			}
			// encode ExtendedT_StatusProhibit_r16 optional
			if ie.ExtendedT_StatusProhibit_r16 != nil {
				if err = ie.ExtendedT_StatusProhibit_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ExtendedT_StatusProhibit_r16", err)
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
			optionals_ext_2 := []bool{ie.Am_WithLongSN_RedCap_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Am_WithLongSN_RedCap_r17 optional
			if ie.Am_WithLongSN_RedCap_r17 != nil {
				if err = ie.Am_WithLongSN_RedCap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Am_WithLongSN_RedCap_r17", err)
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

func (ie *RLC_Parameters) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Am_WithShortSNPresent bool
	if Am_WithShortSNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Um_WithShortSNPresent bool
	if Um_WithShortSNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Um_WithLongSNPresent bool
	if Um_WithLongSNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Am_WithShortSNPresent {
		ie.Am_WithShortSN = new(RLC_Parameters_am_WithShortSN)
		if err = ie.Am_WithShortSN.Decode(r); err != nil {
			return utils.WrapError("Decode Am_WithShortSN", err)
		}
	}
	if Um_WithShortSNPresent {
		ie.Um_WithShortSN = new(RLC_Parameters_um_WithShortSN)
		if err = ie.Um_WithShortSN.Decode(r); err != nil {
			return utils.WrapError("Decode Um_WithShortSN", err)
		}
	}
	if Um_WithLongSNPresent {
		ie.Um_WithLongSN = new(RLC_Parameters_um_WithLongSN)
		if err = ie.Um_WithLongSN.Decode(r); err != nil {
			return utils.WrapError("Decode Um_WithLongSN", err)
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			ExtendedT_PollRetransmit_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ExtendedT_StatusProhibit_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ExtendedT_PollRetransmit_r16 optional
			if ExtendedT_PollRetransmit_r16Present {
				ie.ExtendedT_PollRetransmit_r16 = new(RLC_Parameters_extendedT_PollRetransmit_r16)
				if err = ie.ExtendedT_PollRetransmit_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ExtendedT_PollRetransmit_r16", err)
				}
			}
			// decode ExtendedT_StatusProhibit_r16 optional
			if ExtendedT_StatusProhibit_r16Present {
				ie.ExtendedT_StatusProhibit_r16 = new(RLC_Parameters_extendedT_StatusProhibit_r16)
				if err = ie.ExtendedT_StatusProhibit_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ExtendedT_StatusProhibit_r16", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Am_WithLongSN_RedCap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Am_WithLongSN_RedCap_r17 optional
			if Am_WithLongSN_RedCap_r17Present {
				ie.Am_WithLongSN_RedCap_r17 = new(RLC_Parameters_am_WithLongSN_RedCap_r17)
				if err = ie.Am_WithLongSN_RedCap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Am_WithLongSN_RedCap_r17", err)
				}
			}
		}
	}
	return nil
}
