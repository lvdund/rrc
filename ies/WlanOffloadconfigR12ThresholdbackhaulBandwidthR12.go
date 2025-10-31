package ies

// WLAN-OffloadConfig-r12-thresholdBackhaul-Bandwidth-r12 ::= SEQUENCE
type WlanOffloadconfigR12ThresholdbackhaulBandwidthR12 struct {
	ThresholdbackhauldlBandwidthlowR12  WlanBackhaulrateR12
	ThresholdbackhauldlBandwidthhighR12 WlanBackhaulrateR12
	ThresholdbackhaululBandwidthlowR12  WlanBackhaulrateR12
	ThresholdbackhaululBandwidthhighR12 WlanBackhaulrateR12
}
