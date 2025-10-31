package ies

// RN-SubframeConfig-r10-rpdcch-Config-r10-pucch-Config-r10-tdd ::= CHOICE
const (
	RnSubframeconfigR10RpdcchConfigR10PucchConfigR10TddChoiceNothing = iota
	RnSubframeconfigR10RpdcchConfigR10PucchConfigR10TddChoiceChannelselectionmultiplexingbundling
	RnSubframeconfigR10RpdcchConfigR10PucchConfigR10TddChoiceFallbackforformat3
)

type RnSubframeconfigR10RpdcchConfigR10PucchConfigR10Tdd struct {
	Choice                               uint64
	Channelselectionmultiplexingbundling *RnSubframeconfigR10RpdcchConfigR10PucchConfigR10TddChannelselectionmultiplexingbundling
	Fallbackforformat3                   *RnSubframeconfigR10RpdcchConfigR10PucchConfigR10TddFallbackforformat3
}
