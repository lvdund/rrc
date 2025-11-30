package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSets struct {
	FeatureSetsDownlink            []FeatureSetDownlink            `lb:1,ub:maxDownlinkFeatureSets,optional`
	FeatureSetsDownlinkPerCC       []FeatureSetDownlinkPerCC       `lb:1,ub:maxPerCC_FeatureSets,optional`
	FeatureSetsUplink              []FeatureSetUplink              `lb:1,ub:maxUplinkFeatureSets,optional`
	FeatureSetsUplinkPerCC         []FeatureSetUplinkPerCC         `lb:1,ub:maxPerCC_FeatureSets,optional`
	FeatureSetsDownlink_v1540      []FeatureSetDownlink_v1540      `lb:1,ub:maxDownlinkFeatureSets,optional,ext-1`
	FeatureSetsUplink_v1540        []FeatureSetUplink_v1540        `lb:1,ub:maxUplinkFeatureSets,optional,ext-1`
	FeatureSetsUplinkPerCC_v1540   []FeatureSetUplinkPerCC_v1540   `lb:1,ub:maxPerCC_FeatureSets,optional,ext-1`
	FeatureSetsDownlink_v15a0      []FeatureSetDownlink_v15a0      `lb:1,ub:maxDownlinkFeatureSets,optional,ext-2`
	FeatureSetsDownlink_v1610      []FeatureSetDownlink_v1610      `lb:1,ub:maxDownlinkFeatureSets,optional,ext-3`
	FeatureSetsUplink_v1610        []FeatureSetUplink_v1610        `lb:1,ub:maxUplinkFeatureSets,optional,ext-3`
	FeatureSetDownlinkPerCC_v1620  []FeatureSetDownlinkPerCC_v1620 `lb:1,ub:maxPerCC_FeatureSets,optional,ext-3`
	FeatureSetsUplink_v1630        []FeatureSetUplink_v1630        `lb:1,ub:maxUplinkFeatureSets,optional,ext-4`
	FeatureSetsUplink_v1640        []FeatureSetUplink_v1640        `lb:1,ub:maxUplinkFeatureSets,optional,ext-5`
	FeatureSetsDownlink_v1700      []FeatureSetDownlink_v1700      `lb:1,ub:maxDownlinkFeatureSets,optional,ext-6`
	FeatureSetsDownlinkPerCC_v1700 []FeatureSetDownlinkPerCC_v1700 `lb:1,ub:maxPerCC_FeatureSets,optional,ext-6`
	FeatureSetsUplink_v1710        []FeatureSetUplink_v1710        `lb:1,ub:maxUplinkFeatureSets,optional,ext-6`
	FeatureSetsUplinkPerCC_v1700   []FeatureSetUplinkPerCC_v1700   `lb:1,ub:maxPerCC_FeatureSets,optional,ext-6`
	FeatureSetsDownlink_v1720      []FeatureSetDownlink_v1720      `lb:1,ub:maxDownlinkFeatureSets,optional,ext-7`
	FeatureSetsDownlinkPerCC_v1720 []FeatureSetDownlinkPerCC_v1720 `lb:1,ub:maxPerCC_FeatureSets,optional,ext-7`
	FeatureSetsUplink_v1720        []FeatureSetUplink_v1720        `lb:1,ub:maxUplinkFeatureSets,optional,ext-7`
	FeatureSetsDownlink_v1730      []FeatureSetDownlink_v1730      `lb:1,ub:maxDownlinkFeatureSets,optional,ext-8`
	FeatureSetsDownlinkPerCC_v1730 []FeatureSetDownlinkPerCC_v1730 `lb:1,ub:maxPerCC_FeatureSets,optional,ext-8`
}

func (ie *FeatureSets) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := len(ie.FeatureSetsDownlink_v1540) > 0 || len(ie.FeatureSetsUplink_v1540) > 0 || len(ie.FeatureSetsUplinkPerCC_v1540) > 0 || len(ie.FeatureSetsDownlink_v15a0) > 0 || len(ie.FeatureSetsDownlink_v1610) > 0 || len(ie.FeatureSetsUplink_v1610) > 0 || len(ie.FeatureSetDownlinkPerCC_v1620) > 0 || len(ie.FeatureSetsUplink_v1630) > 0 || len(ie.FeatureSetsUplink_v1640) > 0 || len(ie.FeatureSetsDownlink_v1700) > 0 || len(ie.FeatureSetsDownlinkPerCC_v1700) > 0 || len(ie.FeatureSetsUplink_v1710) > 0 || len(ie.FeatureSetsUplinkPerCC_v1700) > 0 || len(ie.FeatureSetsDownlink_v1720) > 0 || len(ie.FeatureSetsDownlinkPerCC_v1720) > 0 || len(ie.FeatureSetsUplink_v1720) > 0 || len(ie.FeatureSetsDownlink_v1730) > 0 || len(ie.FeatureSetsDownlinkPerCC_v1730) > 0
	preambleBits := []bool{hasExtensions, len(ie.FeatureSetsDownlink) > 0, len(ie.FeatureSetsDownlinkPerCC) > 0, len(ie.FeatureSetsUplink) > 0, len(ie.FeatureSetsUplinkPerCC) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.FeatureSetsDownlink) > 0 {
		tmp_FeatureSetsDownlink := utils.NewSequence[*FeatureSetDownlink]([]*FeatureSetDownlink{}, aper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
		for _, i := range ie.FeatureSetsDownlink {
			tmp_FeatureSetsDownlink.Value = append(tmp_FeatureSetsDownlink.Value, &i)
		}
		if err = tmp_FeatureSetsDownlink.Encode(w); err != nil {
			return utils.WrapError("Encode FeatureSetsDownlink", err)
		}
	}
	if len(ie.FeatureSetsDownlinkPerCC) > 0 {
		tmp_FeatureSetsDownlinkPerCC := utils.NewSequence[*FeatureSetDownlinkPerCC]([]*FeatureSetDownlinkPerCC{}, aper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
		for _, i := range ie.FeatureSetsDownlinkPerCC {
			tmp_FeatureSetsDownlinkPerCC.Value = append(tmp_FeatureSetsDownlinkPerCC.Value, &i)
		}
		if err = tmp_FeatureSetsDownlinkPerCC.Encode(w); err != nil {
			return utils.WrapError("Encode FeatureSetsDownlinkPerCC", err)
		}
	}
	if len(ie.FeatureSetsUplink) > 0 {
		tmp_FeatureSetsUplink := utils.NewSequence[*FeatureSetUplink]([]*FeatureSetUplink{}, aper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
		for _, i := range ie.FeatureSetsUplink {
			tmp_FeatureSetsUplink.Value = append(tmp_FeatureSetsUplink.Value, &i)
		}
		if err = tmp_FeatureSetsUplink.Encode(w); err != nil {
			return utils.WrapError("Encode FeatureSetsUplink", err)
		}
	}
	if len(ie.FeatureSetsUplinkPerCC) > 0 {
		tmp_FeatureSetsUplinkPerCC := utils.NewSequence[*FeatureSetUplinkPerCC]([]*FeatureSetUplinkPerCC{}, aper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
		for _, i := range ie.FeatureSetsUplinkPerCC {
			tmp_FeatureSetsUplinkPerCC.Value = append(tmp_FeatureSetsUplinkPerCC.Value, &i)
		}
		if err = tmp_FeatureSetsUplinkPerCC.Encode(w); err != nil {
			return utils.WrapError("Encode FeatureSetsUplinkPerCC", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 8 bits for 8 extension groups
		extBitmap := []bool{len(ie.FeatureSetsDownlink_v1540) > 0 || len(ie.FeatureSetsUplink_v1540) > 0 || len(ie.FeatureSetsUplinkPerCC_v1540) > 0, len(ie.FeatureSetsDownlink_v15a0) > 0, len(ie.FeatureSetsDownlink_v1610) > 0 || len(ie.FeatureSetsUplink_v1610) > 0 || len(ie.FeatureSetDownlinkPerCC_v1620) > 0, len(ie.FeatureSetsUplink_v1630) > 0, len(ie.FeatureSetsUplink_v1640) > 0, len(ie.FeatureSetsDownlink_v1700) > 0 || len(ie.FeatureSetsDownlinkPerCC_v1700) > 0 || len(ie.FeatureSetsUplink_v1710) > 0 || len(ie.FeatureSetsUplinkPerCC_v1700) > 0, len(ie.FeatureSetsDownlink_v1720) > 0 || len(ie.FeatureSetsDownlinkPerCC_v1720) > 0 || len(ie.FeatureSetsUplink_v1720) > 0, len(ie.FeatureSetsDownlink_v1730) > 0 || len(ie.FeatureSetsDownlinkPerCC_v1730) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap FeatureSets", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.FeatureSetsDownlink_v1540) > 0, len(ie.FeatureSetsUplink_v1540) > 0, len(ie.FeatureSetsUplinkPerCC_v1540) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode FeatureSetsDownlink_v1540 optional
			if len(ie.FeatureSetsDownlink_v1540) > 0 {
				tmp_FeatureSetsDownlink_v1540 := utils.NewSequence[*FeatureSetDownlink_v1540]([]*FeatureSetDownlink_v1540{}, aper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				for _, i := range ie.FeatureSetsDownlink_v1540 {
					tmp_FeatureSetsDownlink_v1540.Value = append(tmp_FeatureSetsDownlink_v1540.Value, &i)
				}
				if err = tmp_FeatureSetsDownlink_v1540.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureSetsDownlink_v1540", err)
				}
			}
			// encode FeatureSetsUplink_v1540 optional
			if len(ie.FeatureSetsUplink_v1540) > 0 {
				tmp_FeatureSetsUplink_v1540 := utils.NewSequence[*FeatureSetUplink_v1540]([]*FeatureSetUplink_v1540{}, aper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				for _, i := range ie.FeatureSetsUplink_v1540 {
					tmp_FeatureSetsUplink_v1540.Value = append(tmp_FeatureSetsUplink_v1540.Value, &i)
				}
				if err = tmp_FeatureSetsUplink_v1540.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureSetsUplink_v1540", err)
				}
			}
			// encode FeatureSetsUplinkPerCC_v1540 optional
			if len(ie.FeatureSetsUplinkPerCC_v1540) > 0 {
				tmp_FeatureSetsUplinkPerCC_v1540 := utils.NewSequence[*FeatureSetUplinkPerCC_v1540]([]*FeatureSetUplinkPerCC_v1540{}, aper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				for _, i := range ie.FeatureSetsUplinkPerCC_v1540 {
					tmp_FeatureSetsUplinkPerCC_v1540.Value = append(tmp_FeatureSetsUplinkPerCC_v1540.Value, &i)
				}
				if err = tmp_FeatureSetsUplinkPerCC_v1540.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureSetsUplinkPerCC_v1540", err)
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
			optionals_ext_2 := []bool{len(ie.FeatureSetsDownlink_v15a0) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode FeatureSetsDownlink_v15a0 optional
			if len(ie.FeatureSetsDownlink_v15a0) > 0 {
				tmp_FeatureSetsDownlink_v15a0 := utils.NewSequence[*FeatureSetDownlink_v15a0]([]*FeatureSetDownlink_v15a0{}, aper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				for _, i := range ie.FeatureSetsDownlink_v15a0 {
					tmp_FeatureSetsDownlink_v15a0.Value = append(tmp_FeatureSetsDownlink_v15a0.Value, &i)
				}
				if err = tmp_FeatureSetsDownlink_v15a0.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureSetsDownlink_v15a0", err)
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
			optionals_ext_3 := []bool{len(ie.FeatureSetsDownlink_v1610) > 0, len(ie.FeatureSetsUplink_v1610) > 0, len(ie.FeatureSetDownlinkPerCC_v1620) > 0}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode FeatureSetsDownlink_v1610 optional
			if len(ie.FeatureSetsDownlink_v1610) > 0 {
				tmp_FeatureSetsDownlink_v1610 := utils.NewSequence[*FeatureSetDownlink_v1610]([]*FeatureSetDownlink_v1610{}, aper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				for _, i := range ie.FeatureSetsDownlink_v1610 {
					tmp_FeatureSetsDownlink_v1610.Value = append(tmp_FeatureSetsDownlink_v1610.Value, &i)
				}
				if err = tmp_FeatureSetsDownlink_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureSetsDownlink_v1610", err)
				}
			}
			// encode FeatureSetsUplink_v1610 optional
			if len(ie.FeatureSetsUplink_v1610) > 0 {
				tmp_FeatureSetsUplink_v1610 := utils.NewSequence[*FeatureSetUplink_v1610]([]*FeatureSetUplink_v1610{}, aper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				for _, i := range ie.FeatureSetsUplink_v1610 {
					tmp_FeatureSetsUplink_v1610.Value = append(tmp_FeatureSetsUplink_v1610.Value, &i)
				}
				if err = tmp_FeatureSetsUplink_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureSetsUplink_v1610", err)
				}
			}
			// encode FeatureSetDownlinkPerCC_v1620 optional
			if len(ie.FeatureSetDownlinkPerCC_v1620) > 0 {
				tmp_FeatureSetDownlinkPerCC_v1620 := utils.NewSequence[*FeatureSetDownlinkPerCC_v1620]([]*FeatureSetDownlinkPerCC_v1620{}, aper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				for _, i := range ie.FeatureSetDownlinkPerCC_v1620 {
					tmp_FeatureSetDownlinkPerCC_v1620.Value = append(tmp_FeatureSetDownlinkPerCC_v1620.Value, &i)
				}
				if err = tmp_FeatureSetDownlinkPerCC_v1620.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureSetDownlinkPerCC_v1620", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{len(ie.FeatureSetsUplink_v1630) > 0}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode FeatureSetsUplink_v1630 optional
			if len(ie.FeatureSetsUplink_v1630) > 0 {
				tmp_FeatureSetsUplink_v1630 := utils.NewSequence[*FeatureSetUplink_v1630]([]*FeatureSetUplink_v1630{}, aper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				for _, i := range ie.FeatureSetsUplink_v1630 {
					tmp_FeatureSetsUplink_v1630.Value = append(tmp_FeatureSetsUplink_v1630.Value, &i)
				}
				if err = tmp_FeatureSetsUplink_v1630.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureSetsUplink_v1630", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 5
		if extBitmap[4] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{len(ie.FeatureSetsUplink_v1640) > 0}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode FeatureSetsUplink_v1640 optional
			if len(ie.FeatureSetsUplink_v1640) > 0 {
				tmp_FeatureSetsUplink_v1640 := utils.NewSequence[*FeatureSetUplink_v1640]([]*FeatureSetUplink_v1640{}, aper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				for _, i := range ie.FeatureSetsUplink_v1640 {
					tmp_FeatureSetsUplink_v1640.Value = append(tmp_FeatureSetsUplink_v1640.Value, &i)
				}
				if err = tmp_FeatureSetsUplink_v1640.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureSetsUplink_v1640", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 6
		if extBitmap[5] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 6
			optionals_ext_6 := []bool{len(ie.FeatureSetsDownlink_v1700) > 0, len(ie.FeatureSetsDownlinkPerCC_v1700) > 0, len(ie.FeatureSetsUplink_v1710) > 0, len(ie.FeatureSetsUplinkPerCC_v1700) > 0}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode FeatureSetsDownlink_v1700 optional
			if len(ie.FeatureSetsDownlink_v1700) > 0 {
				tmp_FeatureSetsDownlink_v1700 := utils.NewSequence[*FeatureSetDownlink_v1700]([]*FeatureSetDownlink_v1700{}, aper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				for _, i := range ie.FeatureSetsDownlink_v1700 {
					tmp_FeatureSetsDownlink_v1700.Value = append(tmp_FeatureSetsDownlink_v1700.Value, &i)
				}
				if err = tmp_FeatureSetsDownlink_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureSetsDownlink_v1700", err)
				}
			}
			// encode FeatureSetsDownlinkPerCC_v1700 optional
			if len(ie.FeatureSetsDownlinkPerCC_v1700) > 0 {
				tmp_FeatureSetsDownlinkPerCC_v1700 := utils.NewSequence[*FeatureSetDownlinkPerCC_v1700]([]*FeatureSetDownlinkPerCC_v1700{}, aper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				for _, i := range ie.FeatureSetsDownlinkPerCC_v1700 {
					tmp_FeatureSetsDownlinkPerCC_v1700.Value = append(tmp_FeatureSetsDownlinkPerCC_v1700.Value, &i)
				}
				if err = tmp_FeatureSetsDownlinkPerCC_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureSetsDownlinkPerCC_v1700", err)
				}
			}
			// encode FeatureSetsUplink_v1710 optional
			if len(ie.FeatureSetsUplink_v1710) > 0 {
				tmp_FeatureSetsUplink_v1710 := utils.NewSequence[*FeatureSetUplink_v1710]([]*FeatureSetUplink_v1710{}, aper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				for _, i := range ie.FeatureSetsUplink_v1710 {
					tmp_FeatureSetsUplink_v1710.Value = append(tmp_FeatureSetsUplink_v1710.Value, &i)
				}
				if err = tmp_FeatureSetsUplink_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureSetsUplink_v1710", err)
				}
			}
			// encode FeatureSetsUplinkPerCC_v1700 optional
			if len(ie.FeatureSetsUplinkPerCC_v1700) > 0 {
				tmp_FeatureSetsUplinkPerCC_v1700 := utils.NewSequence[*FeatureSetUplinkPerCC_v1700]([]*FeatureSetUplinkPerCC_v1700{}, aper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				for _, i := range ie.FeatureSetsUplinkPerCC_v1700 {
					tmp_FeatureSetsUplinkPerCC_v1700.Value = append(tmp_FeatureSetsUplinkPerCC_v1700.Value, &i)
				}
				if err = tmp_FeatureSetsUplinkPerCC_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureSetsUplinkPerCC_v1700", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 7
		if extBitmap[6] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 7
			optionals_ext_7 := []bool{len(ie.FeatureSetsDownlink_v1720) > 0, len(ie.FeatureSetsDownlinkPerCC_v1720) > 0, len(ie.FeatureSetsUplink_v1720) > 0}
			for _, bit := range optionals_ext_7 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode FeatureSetsDownlink_v1720 optional
			if len(ie.FeatureSetsDownlink_v1720) > 0 {
				tmp_FeatureSetsDownlink_v1720 := utils.NewSequence[*FeatureSetDownlink_v1720]([]*FeatureSetDownlink_v1720{}, aper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				for _, i := range ie.FeatureSetsDownlink_v1720 {
					tmp_FeatureSetsDownlink_v1720.Value = append(tmp_FeatureSetsDownlink_v1720.Value, &i)
				}
				if err = tmp_FeatureSetsDownlink_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureSetsDownlink_v1720", err)
				}
			}
			// encode FeatureSetsDownlinkPerCC_v1720 optional
			if len(ie.FeatureSetsDownlinkPerCC_v1720) > 0 {
				tmp_FeatureSetsDownlinkPerCC_v1720 := utils.NewSequence[*FeatureSetDownlinkPerCC_v1720]([]*FeatureSetDownlinkPerCC_v1720{}, aper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				for _, i := range ie.FeatureSetsDownlinkPerCC_v1720 {
					tmp_FeatureSetsDownlinkPerCC_v1720.Value = append(tmp_FeatureSetsDownlinkPerCC_v1720.Value, &i)
				}
				if err = tmp_FeatureSetsDownlinkPerCC_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureSetsDownlinkPerCC_v1720", err)
				}
			}
			// encode FeatureSetsUplink_v1720 optional
			if len(ie.FeatureSetsUplink_v1720) > 0 {
				tmp_FeatureSetsUplink_v1720 := utils.NewSequence[*FeatureSetUplink_v1720]([]*FeatureSetUplink_v1720{}, aper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				for _, i := range ie.FeatureSetsUplink_v1720 {
					tmp_FeatureSetsUplink_v1720.Value = append(tmp_FeatureSetsUplink_v1720.Value, &i)
				}
				if err = tmp_FeatureSetsUplink_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureSetsUplink_v1720", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 8
		if extBitmap[7] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 8
			optionals_ext_8 := []bool{len(ie.FeatureSetsDownlink_v1730) > 0, len(ie.FeatureSetsDownlinkPerCC_v1730) > 0}
			for _, bit := range optionals_ext_8 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode FeatureSetsDownlink_v1730 optional
			if len(ie.FeatureSetsDownlink_v1730) > 0 {
				tmp_FeatureSetsDownlink_v1730 := utils.NewSequence[*FeatureSetDownlink_v1730]([]*FeatureSetDownlink_v1730{}, aper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				for _, i := range ie.FeatureSetsDownlink_v1730 {
					tmp_FeatureSetsDownlink_v1730.Value = append(tmp_FeatureSetsDownlink_v1730.Value, &i)
				}
				if err = tmp_FeatureSetsDownlink_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureSetsDownlink_v1730", err)
				}
			}
			// encode FeatureSetsDownlinkPerCC_v1730 optional
			if len(ie.FeatureSetsDownlinkPerCC_v1730) > 0 {
				tmp_FeatureSetsDownlinkPerCC_v1730 := utils.NewSequence[*FeatureSetDownlinkPerCC_v1730]([]*FeatureSetDownlinkPerCC_v1730{}, aper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				for _, i := range ie.FeatureSetsDownlinkPerCC_v1730 {
					tmp_FeatureSetsDownlinkPerCC_v1730.Value = append(tmp_FeatureSetsDownlinkPerCC_v1730.Value, &i)
				}
				if err = tmp_FeatureSetsDownlinkPerCC_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureSetsDownlinkPerCC_v1730", err)
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

func (ie *FeatureSets) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var FeatureSetsDownlinkPresent bool
	if FeatureSetsDownlinkPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FeatureSetsDownlinkPerCCPresent bool
	if FeatureSetsDownlinkPerCCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FeatureSetsUplinkPresent bool
	if FeatureSetsUplinkPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FeatureSetsUplinkPerCCPresent bool
	if FeatureSetsUplinkPerCCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if FeatureSetsDownlinkPresent {
		tmp_FeatureSetsDownlink := utils.NewSequence[*FeatureSetDownlink]([]*FeatureSetDownlink{}, aper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
		fn_FeatureSetsDownlink := func() *FeatureSetDownlink {
			return new(FeatureSetDownlink)
		}
		if err = tmp_FeatureSetsDownlink.Decode(r, fn_FeatureSetsDownlink); err != nil {
			return utils.WrapError("Decode FeatureSetsDownlink", err)
		}
		ie.FeatureSetsDownlink = []FeatureSetDownlink{}
		for _, i := range tmp_FeatureSetsDownlink.Value {
			ie.FeatureSetsDownlink = append(ie.FeatureSetsDownlink, *i)
		}
	}
	if FeatureSetsDownlinkPerCCPresent {
		tmp_FeatureSetsDownlinkPerCC := utils.NewSequence[*FeatureSetDownlinkPerCC]([]*FeatureSetDownlinkPerCC{}, aper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
		fn_FeatureSetsDownlinkPerCC := func() *FeatureSetDownlinkPerCC {
			return new(FeatureSetDownlinkPerCC)
		}
		if err = tmp_FeatureSetsDownlinkPerCC.Decode(r, fn_FeatureSetsDownlinkPerCC); err != nil {
			return utils.WrapError("Decode FeatureSetsDownlinkPerCC", err)
		}
		ie.FeatureSetsDownlinkPerCC = []FeatureSetDownlinkPerCC{}
		for _, i := range tmp_FeatureSetsDownlinkPerCC.Value {
			ie.FeatureSetsDownlinkPerCC = append(ie.FeatureSetsDownlinkPerCC, *i)
		}
	}
	if FeatureSetsUplinkPresent {
		tmp_FeatureSetsUplink := utils.NewSequence[*FeatureSetUplink]([]*FeatureSetUplink{}, aper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
		fn_FeatureSetsUplink := func() *FeatureSetUplink {
			return new(FeatureSetUplink)
		}
		if err = tmp_FeatureSetsUplink.Decode(r, fn_FeatureSetsUplink); err != nil {
			return utils.WrapError("Decode FeatureSetsUplink", err)
		}
		ie.FeatureSetsUplink = []FeatureSetUplink{}
		for _, i := range tmp_FeatureSetsUplink.Value {
			ie.FeatureSetsUplink = append(ie.FeatureSetsUplink, *i)
		}
	}
	if FeatureSetsUplinkPerCCPresent {
		tmp_FeatureSetsUplinkPerCC := utils.NewSequence[*FeatureSetUplinkPerCC]([]*FeatureSetUplinkPerCC{}, aper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
		fn_FeatureSetsUplinkPerCC := func() *FeatureSetUplinkPerCC {
			return new(FeatureSetUplinkPerCC)
		}
		if err = tmp_FeatureSetsUplinkPerCC.Decode(r, fn_FeatureSetsUplinkPerCC); err != nil {
			return utils.WrapError("Decode FeatureSetsUplinkPerCC", err)
		}
		ie.FeatureSetsUplinkPerCC = []FeatureSetUplinkPerCC{}
		for _, i := range tmp_FeatureSetsUplinkPerCC.Value {
			ie.FeatureSetsUplinkPerCC = append(ie.FeatureSetsUplinkPerCC, *i)
		}
	}

	if extensionBit {
		// Read extension bitmap: 8 bits for 8 extension groups
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

			FeatureSetsDownlink_v1540Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FeatureSetsUplink_v1540Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FeatureSetsUplinkPerCC_v1540Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode FeatureSetsDownlink_v1540 optional
			if FeatureSetsDownlink_v1540Present {
				tmp_FeatureSetsDownlink_v1540 := utils.NewSequence[*FeatureSetDownlink_v1540]([]*FeatureSetDownlink_v1540{}, aper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				fn_FeatureSetsDownlink_v1540 := func() *FeatureSetDownlink_v1540 {
					return new(FeatureSetDownlink_v1540)
				}
				if err = tmp_FeatureSetsDownlink_v1540.Decode(extReader, fn_FeatureSetsDownlink_v1540); err != nil {
					return utils.WrapError("Decode FeatureSetsDownlink_v1540", err)
				}
				ie.FeatureSetsDownlink_v1540 = []FeatureSetDownlink_v1540{}
				for _, i := range tmp_FeatureSetsDownlink_v1540.Value {
					ie.FeatureSetsDownlink_v1540 = append(ie.FeatureSetsDownlink_v1540, *i)
				}
			}
			// decode FeatureSetsUplink_v1540 optional
			if FeatureSetsUplink_v1540Present {
				tmp_FeatureSetsUplink_v1540 := utils.NewSequence[*FeatureSetUplink_v1540]([]*FeatureSetUplink_v1540{}, aper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				fn_FeatureSetsUplink_v1540 := func() *FeatureSetUplink_v1540 {
					return new(FeatureSetUplink_v1540)
				}
				if err = tmp_FeatureSetsUplink_v1540.Decode(extReader, fn_FeatureSetsUplink_v1540); err != nil {
					return utils.WrapError("Decode FeatureSetsUplink_v1540", err)
				}
				ie.FeatureSetsUplink_v1540 = []FeatureSetUplink_v1540{}
				for _, i := range tmp_FeatureSetsUplink_v1540.Value {
					ie.FeatureSetsUplink_v1540 = append(ie.FeatureSetsUplink_v1540, *i)
				}
			}
			// decode FeatureSetsUplinkPerCC_v1540 optional
			if FeatureSetsUplinkPerCC_v1540Present {
				tmp_FeatureSetsUplinkPerCC_v1540 := utils.NewSequence[*FeatureSetUplinkPerCC_v1540]([]*FeatureSetUplinkPerCC_v1540{}, aper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				fn_FeatureSetsUplinkPerCC_v1540 := func() *FeatureSetUplinkPerCC_v1540 {
					return new(FeatureSetUplinkPerCC_v1540)
				}
				if err = tmp_FeatureSetsUplinkPerCC_v1540.Decode(extReader, fn_FeatureSetsUplinkPerCC_v1540); err != nil {
					return utils.WrapError("Decode FeatureSetsUplinkPerCC_v1540", err)
				}
				ie.FeatureSetsUplinkPerCC_v1540 = []FeatureSetUplinkPerCC_v1540{}
				for _, i := range tmp_FeatureSetsUplinkPerCC_v1540.Value {
					ie.FeatureSetsUplinkPerCC_v1540 = append(ie.FeatureSetsUplinkPerCC_v1540, *i)
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

			FeatureSetsDownlink_v15a0Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode FeatureSetsDownlink_v15a0 optional
			if FeatureSetsDownlink_v15a0Present {
				tmp_FeatureSetsDownlink_v15a0 := utils.NewSequence[*FeatureSetDownlink_v15a0]([]*FeatureSetDownlink_v15a0{}, aper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				fn_FeatureSetsDownlink_v15a0 := func() *FeatureSetDownlink_v15a0 {
					return new(FeatureSetDownlink_v15a0)
				}
				if err = tmp_FeatureSetsDownlink_v15a0.Decode(extReader, fn_FeatureSetsDownlink_v15a0); err != nil {
					return utils.WrapError("Decode FeatureSetsDownlink_v15a0", err)
				}
				ie.FeatureSetsDownlink_v15a0 = []FeatureSetDownlink_v15a0{}
				for _, i := range tmp_FeatureSetsDownlink_v15a0.Value {
					ie.FeatureSetsDownlink_v15a0 = append(ie.FeatureSetsDownlink_v15a0, *i)
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

			FeatureSetsDownlink_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FeatureSetsUplink_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FeatureSetDownlinkPerCC_v1620Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode FeatureSetsDownlink_v1610 optional
			if FeatureSetsDownlink_v1610Present {
				tmp_FeatureSetsDownlink_v1610 := utils.NewSequence[*FeatureSetDownlink_v1610]([]*FeatureSetDownlink_v1610{}, aper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				fn_FeatureSetsDownlink_v1610 := func() *FeatureSetDownlink_v1610 {
					return new(FeatureSetDownlink_v1610)
				}
				if err = tmp_FeatureSetsDownlink_v1610.Decode(extReader, fn_FeatureSetsDownlink_v1610); err != nil {
					return utils.WrapError("Decode FeatureSetsDownlink_v1610", err)
				}
				ie.FeatureSetsDownlink_v1610 = []FeatureSetDownlink_v1610{}
				for _, i := range tmp_FeatureSetsDownlink_v1610.Value {
					ie.FeatureSetsDownlink_v1610 = append(ie.FeatureSetsDownlink_v1610, *i)
				}
			}
			// decode FeatureSetsUplink_v1610 optional
			if FeatureSetsUplink_v1610Present {
				tmp_FeatureSetsUplink_v1610 := utils.NewSequence[*FeatureSetUplink_v1610]([]*FeatureSetUplink_v1610{}, aper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				fn_FeatureSetsUplink_v1610 := func() *FeatureSetUplink_v1610 {
					return new(FeatureSetUplink_v1610)
				}
				if err = tmp_FeatureSetsUplink_v1610.Decode(extReader, fn_FeatureSetsUplink_v1610); err != nil {
					return utils.WrapError("Decode FeatureSetsUplink_v1610", err)
				}
				ie.FeatureSetsUplink_v1610 = []FeatureSetUplink_v1610{}
				for _, i := range tmp_FeatureSetsUplink_v1610.Value {
					ie.FeatureSetsUplink_v1610 = append(ie.FeatureSetsUplink_v1610, *i)
				}
			}
			// decode FeatureSetDownlinkPerCC_v1620 optional
			if FeatureSetDownlinkPerCC_v1620Present {
				tmp_FeatureSetDownlinkPerCC_v1620 := utils.NewSequence[*FeatureSetDownlinkPerCC_v1620]([]*FeatureSetDownlinkPerCC_v1620{}, aper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				fn_FeatureSetDownlinkPerCC_v1620 := func() *FeatureSetDownlinkPerCC_v1620 {
					return new(FeatureSetDownlinkPerCC_v1620)
				}
				if err = tmp_FeatureSetDownlinkPerCC_v1620.Decode(extReader, fn_FeatureSetDownlinkPerCC_v1620); err != nil {
					return utils.WrapError("Decode FeatureSetDownlinkPerCC_v1620", err)
				}
				ie.FeatureSetDownlinkPerCC_v1620 = []FeatureSetDownlinkPerCC_v1620{}
				for _, i := range tmp_FeatureSetDownlinkPerCC_v1620.Value {
					ie.FeatureSetDownlinkPerCC_v1620 = append(ie.FeatureSetDownlinkPerCC_v1620, *i)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			FeatureSetsUplink_v1630Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode FeatureSetsUplink_v1630 optional
			if FeatureSetsUplink_v1630Present {
				tmp_FeatureSetsUplink_v1630 := utils.NewSequence[*FeatureSetUplink_v1630]([]*FeatureSetUplink_v1630{}, aper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				fn_FeatureSetsUplink_v1630 := func() *FeatureSetUplink_v1630 {
					return new(FeatureSetUplink_v1630)
				}
				if err = tmp_FeatureSetsUplink_v1630.Decode(extReader, fn_FeatureSetsUplink_v1630); err != nil {
					return utils.WrapError("Decode FeatureSetsUplink_v1630", err)
				}
				ie.FeatureSetsUplink_v1630 = []FeatureSetUplink_v1630{}
				for _, i := range tmp_FeatureSetsUplink_v1630.Value {
					ie.FeatureSetsUplink_v1630 = append(ie.FeatureSetsUplink_v1630, *i)
				}
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			FeatureSetsUplink_v1640Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode FeatureSetsUplink_v1640 optional
			if FeatureSetsUplink_v1640Present {
				tmp_FeatureSetsUplink_v1640 := utils.NewSequence[*FeatureSetUplink_v1640]([]*FeatureSetUplink_v1640{}, aper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				fn_FeatureSetsUplink_v1640 := func() *FeatureSetUplink_v1640 {
					return new(FeatureSetUplink_v1640)
				}
				if err = tmp_FeatureSetsUplink_v1640.Decode(extReader, fn_FeatureSetsUplink_v1640); err != nil {
					return utils.WrapError("Decode FeatureSetsUplink_v1640", err)
				}
				ie.FeatureSetsUplink_v1640 = []FeatureSetUplink_v1640{}
				for _, i := range tmp_FeatureSetsUplink_v1640.Value {
					ie.FeatureSetsUplink_v1640 = append(ie.FeatureSetsUplink_v1640, *i)
				}
			}
		}
		// decode extension group 6
		if len(extBitmap) > 5 && extBitmap[5] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			FeatureSetsDownlink_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FeatureSetsDownlinkPerCC_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FeatureSetsUplink_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FeatureSetsUplinkPerCC_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode FeatureSetsDownlink_v1700 optional
			if FeatureSetsDownlink_v1700Present {
				tmp_FeatureSetsDownlink_v1700 := utils.NewSequence[*FeatureSetDownlink_v1700]([]*FeatureSetDownlink_v1700{}, aper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				fn_FeatureSetsDownlink_v1700 := func() *FeatureSetDownlink_v1700 {
					return new(FeatureSetDownlink_v1700)
				}
				if err = tmp_FeatureSetsDownlink_v1700.Decode(extReader, fn_FeatureSetsDownlink_v1700); err != nil {
					return utils.WrapError("Decode FeatureSetsDownlink_v1700", err)
				}
				ie.FeatureSetsDownlink_v1700 = []FeatureSetDownlink_v1700{}
				for _, i := range tmp_FeatureSetsDownlink_v1700.Value {
					ie.FeatureSetsDownlink_v1700 = append(ie.FeatureSetsDownlink_v1700, *i)
				}
			}
			// decode FeatureSetsDownlinkPerCC_v1700 optional
			if FeatureSetsDownlinkPerCC_v1700Present {
				tmp_FeatureSetsDownlinkPerCC_v1700 := utils.NewSequence[*FeatureSetDownlinkPerCC_v1700]([]*FeatureSetDownlinkPerCC_v1700{}, aper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				fn_FeatureSetsDownlinkPerCC_v1700 := func() *FeatureSetDownlinkPerCC_v1700 {
					return new(FeatureSetDownlinkPerCC_v1700)
				}
				if err = tmp_FeatureSetsDownlinkPerCC_v1700.Decode(extReader, fn_FeatureSetsDownlinkPerCC_v1700); err != nil {
					return utils.WrapError("Decode FeatureSetsDownlinkPerCC_v1700", err)
				}
				ie.FeatureSetsDownlinkPerCC_v1700 = []FeatureSetDownlinkPerCC_v1700{}
				for _, i := range tmp_FeatureSetsDownlinkPerCC_v1700.Value {
					ie.FeatureSetsDownlinkPerCC_v1700 = append(ie.FeatureSetsDownlinkPerCC_v1700, *i)
				}
			}
			// decode FeatureSetsUplink_v1710 optional
			if FeatureSetsUplink_v1710Present {
				tmp_FeatureSetsUplink_v1710 := utils.NewSequence[*FeatureSetUplink_v1710]([]*FeatureSetUplink_v1710{}, aper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				fn_FeatureSetsUplink_v1710 := func() *FeatureSetUplink_v1710 {
					return new(FeatureSetUplink_v1710)
				}
				if err = tmp_FeatureSetsUplink_v1710.Decode(extReader, fn_FeatureSetsUplink_v1710); err != nil {
					return utils.WrapError("Decode FeatureSetsUplink_v1710", err)
				}
				ie.FeatureSetsUplink_v1710 = []FeatureSetUplink_v1710{}
				for _, i := range tmp_FeatureSetsUplink_v1710.Value {
					ie.FeatureSetsUplink_v1710 = append(ie.FeatureSetsUplink_v1710, *i)
				}
			}
			// decode FeatureSetsUplinkPerCC_v1700 optional
			if FeatureSetsUplinkPerCC_v1700Present {
				tmp_FeatureSetsUplinkPerCC_v1700 := utils.NewSequence[*FeatureSetUplinkPerCC_v1700]([]*FeatureSetUplinkPerCC_v1700{}, aper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				fn_FeatureSetsUplinkPerCC_v1700 := func() *FeatureSetUplinkPerCC_v1700 {
					return new(FeatureSetUplinkPerCC_v1700)
				}
				if err = tmp_FeatureSetsUplinkPerCC_v1700.Decode(extReader, fn_FeatureSetsUplinkPerCC_v1700); err != nil {
					return utils.WrapError("Decode FeatureSetsUplinkPerCC_v1700", err)
				}
				ie.FeatureSetsUplinkPerCC_v1700 = []FeatureSetUplinkPerCC_v1700{}
				for _, i := range tmp_FeatureSetsUplinkPerCC_v1700.Value {
					ie.FeatureSetsUplinkPerCC_v1700 = append(ie.FeatureSetsUplinkPerCC_v1700, *i)
				}
			}
		}
		// decode extension group 7
		if len(extBitmap) > 6 && extBitmap[6] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			FeatureSetsDownlink_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FeatureSetsDownlinkPerCC_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FeatureSetsUplink_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode FeatureSetsDownlink_v1720 optional
			if FeatureSetsDownlink_v1720Present {
				tmp_FeatureSetsDownlink_v1720 := utils.NewSequence[*FeatureSetDownlink_v1720]([]*FeatureSetDownlink_v1720{}, aper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				fn_FeatureSetsDownlink_v1720 := func() *FeatureSetDownlink_v1720 {
					return new(FeatureSetDownlink_v1720)
				}
				if err = tmp_FeatureSetsDownlink_v1720.Decode(extReader, fn_FeatureSetsDownlink_v1720); err != nil {
					return utils.WrapError("Decode FeatureSetsDownlink_v1720", err)
				}
				ie.FeatureSetsDownlink_v1720 = []FeatureSetDownlink_v1720{}
				for _, i := range tmp_FeatureSetsDownlink_v1720.Value {
					ie.FeatureSetsDownlink_v1720 = append(ie.FeatureSetsDownlink_v1720, *i)
				}
			}
			// decode FeatureSetsDownlinkPerCC_v1720 optional
			if FeatureSetsDownlinkPerCC_v1720Present {
				tmp_FeatureSetsDownlinkPerCC_v1720 := utils.NewSequence[*FeatureSetDownlinkPerCC_v1720]([]*FeatureSetDownlinkPerCC_v1720{}, aper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				fn_FeatureSetsDownlinkPerCC_v1720 := func() *FeatureSetDownlinkPerCC_v1720 {
					return new(FeatureSetDownlinkPerCC_v1720)
				}
				if err = tmp_FeatureSetsDownlinkPerCC_v1720.Decode(extReader, fn_FeatureSetsDownlinkPerCC_v1720); err != nil {
					return utils.WrapError("Decode FeatureSetsDownlinkPerCC_v1720", err)
				}
				ie.FeatureSetsDownlinkPerCC_v1720 = []FeatureSetDownlinkPerCC_v1720{}
				for _, i := range tmp_FeatureSetsDownlinkPerCC_v1720.Value {
					ie.FeatureSetsDownlinkPerCC_v1720 = append(ie.FeatureSetsDownlinkPerCC_v1720, *i)
				}
			}
			// decode FeatureSetsUplink_v1720 optional
			if FeatureSetsUplink_v1720Present {
				tmp_FeatureSetsUplink_v1720 := utils.NewSequence[*FeatureSetUplink_v1720]([]*FeatureSetUplink_v1720{}, aper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				fn_FeatureSetsUplink_v1720 := func() *FeatureSetUplink_v1720 {
					return new(FeatureSetUplink_v1720)
				}
				if err = tmp_FeatureSetsUplink_v1720.Decode(extReader, fn_FeatureSetsUplink_v1720); err != nil {
					return utils.WrapError("Decode FeatureSetsUplink_v1720", err)
				}
				ie.FeatureSetsUplink_v1720 = []FeatureSetUplink_v1720{}
				for _, i := range tmp_FeatureSetsUplink_v1720.Value {
					ie.FeatureSetsUplink_v1720 = append(ie.FeatureSetsUplink_v1720, *i)
				}
			}
		}
		// decode extension group 8
		if len(extBitmap) > 7 && extBitmap[7] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			FeatureSetsDownlink_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FeatureSetsDownlinkPerCC_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode FeatureSetsDownlink_v1730 optional
			if FeatureSetsDownlink_v1730Present {
				tmp_FeatureSetsDownlink_v1730 := utils.NewSequence[*FeatureSetDownlink_v1730]([]*FeatureSetDownlink_v1730{}, aper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				fn_FeatureSetsDownlink_v1730 := func() *FeatureSetDownlink_v1730 {
					return new(FeatureSetDownlink_v1730)
				}
				if err = tmp_FeatureSetsDownlink_v1730.Decode(extReader, fn_FeatureSetsDownlink_v1730); err != nil {
					return utils.WrapError("Decode FeatureSetsDownlink_v1730", err)
				}
				ie.FeatureSetsDownlink_v1730 = []FeatureSetDownlink_v1730{}
				for _, i := range tmp_FeatureSetsDownlink_v1730.Value {
					ie.FeatureSetsDownlink_v1730 = append(ie.FeatureSetsDownlink_v1730, *i)
				}
			}
			// decode FeatureSetsDownlinkPerCC_v1730 optional
			if FeatureSetsDownlinkPerCC_v1730Present {
				tmp_FeatureSetsDownlinkPerCC_v1730 := utils.NewSequence[*FeatureSetDownlinkPerCC_v1730]([]*FeatureSetDownlinkPerCC_v1730{}, aper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				fn_FeatureSetsDownlinkPerCC_v1730 := func() *FeatureSetDownlinkPerCC_v1730 {
					return new(FeatureSetDownlinkPerCC_v1730)
				}
				if err = tmp_FeatureSetsDownlinkPerCC_v1730.Decode(extReader, fn_FeatureSetsDownlinkPerCC_v1730); err != nil {
					return utils.WrapError("Decode FeatureSetsDownlinkPerCC_v1730", err)
				}
				ie.FeatureSetsDownlinkPerCC_v1730 = []FeatureSetDownlinkPerCC_v1730{}
				for _, i := range tmp_FeatureSetsDownlinkPerCC_v1730.Value {
					ie.FeatureSetsDownlinkPerCC_v1730 = append(ie.FeatureSetsDownlinkPerCC_v1730, *i)
				}
			}
		}
	}
	return nil
}
