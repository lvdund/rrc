package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PosSI_SchedulingInfo_r16 struct {
	PosSchedulingInfoList_r16     []PosSchedulingInfo_r16 `lb:1,ub:maxSI_Message,madatory`
	PosSI_RequestConfig_r16       *SI_RequestConfig       `optional`
	PosSI_RequestConfigSUL_r16    *SI_RequestConfig       `optional`
	PosSI_RequestConfigRedCap_r17 *SI_RequestConfig       `optional,ext-1`
}

func (ie *PosSI_SchedulingInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.PosSI_RequestConfigRedCap_r17 != nil
	preambleBits := []bool{hasExtensions, ie.PosSI_RequestConfig_r16 != nil, ie.PosSI_RequestConfigSUL_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_PosSchedulingInfoList_r16 := utils.NewSequence[*PosSchedulingInfo_r16]([]*PosSchedulingInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxSI_Message}, false)
	for _, i := range ie.PosSchedulingInfoList_r16 {
		tmp_PosSchedulingInfoList_r16.Value = append(tmp_PosSchedulingInfoList_r16.Value, &i)
	}
	if err = tmp_PosSchedulingInfoList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PosSchedulingInfoList_r16", err)
	}
	if ie.PosSI_RequestConfig_r16 != nil {
		if err = ie.PosSI_RequestConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PosSI_RequestConfig_r16", err)
		}
	}
	if ie.PosSI_RequestConfigSUL_r16 != nil {
		if err = ie.PosSI_RequestConfigSUL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PosSI_RequestConfigSUL_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.PosSI_RequestConfigRedCap_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PosSI_SchedulingInfo_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.PosSI_RequestConfigRedCap_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode PosSI_RequestConfigRedCap_r17 optional
			if ie.PosSI_RequestConfigRedCap_r17 != nil {
				if err = ie.PosSI_RequestConfigRedCap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PosSI_RequestConfigRedCap_r17", err)
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

func (ie *PosSI_SchedulingInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var PosSI_RequestConfig_r16Present bool
	if PosSI_RequestConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PosSI_RequestConfigSUL_r16Present bool
	if PosSI_RequestConfigSUL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_PosSchedulingInfoList_r16 := utils.NewSequence[*PosSchedulingInfo_r16]([]*PosSchedulingInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxSI_Message}, false)
	fn_PosSchedulingInfoList_r16 := func() *PosSchedulingInfo_r16 {
		return new(PosSchedulingInfo_r16)
	}
	if err = tmp_PosSchedulingInfoList_r16.Decode(r, fn_PosSchedulingInfoList_r16); err != nil {
		return utils.WrapError("Decode PosSchedulingInfoList_r16", err)
	}
	ie.PosSchedulingInfoList_r16 = []PosSchedulingInfo_r16{}
	for _, i := range tmp_PosSchedulingInfoList_r16.Value {
		ie.PosSchedulingInfoList_r16 = append(ie.PosSchedulingInfoList_r16, *i)
	}
	if PosSI_RequestConfig_r16Present {
		ie.PosSI_RequestConfig_r16 = new(SI_RequestConfig)
		if err = ie.PosSI_RequestConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSI_RequestConfig_r16", err)
		}
	}
	if PosSI_RequestConfigSUL_r16Present {
		ie.PosSI_RequestConfigSUL_r16 = new(SI_RequestConfig)
		if err = ie.PosSI_RequestConfigSUL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSI_RequestConfigSUL_r16", err)
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

			PosSI_RequestConfigRedCap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode PosSI_RequestConfigRedCap_r17 optional
			if PosSI_RequestConfigRedCap_r17Present {
				ie.PosSI_RequestConfigRedCap_r17 = new(SI_RequestConfig)
				if err = ie.PosSI_RequestConfigRedCap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode PosSI_RequestConfigRedCap_r17", err)
				}
			}
		}
	}
	return nil
}
