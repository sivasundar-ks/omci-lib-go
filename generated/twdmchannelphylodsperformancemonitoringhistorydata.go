/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/*
 * NOTE: This file was generated, manual edits will be overwritten!
 *
 * Generated by 'goCodeGenerator.py':
 *              https://github.com/cboling/OMCI-parser/README.md
 */

package generated

import "github.com/deckarep/golang-set"

// TwdmChannelPhyLodsPerformanceMonitoringHistoryDataClassID is the 16-bit ID for the OMCI
// Managed entity TWDM channel PHY/LODS performance monitoring history data
const TwdmChannelPhyLodsPerformanceMonitoringHistoryDataClassID ClassID = ClassID(444)

var twdmchannelphylodsperformancemonitoringhistorydataBME *ManagedEntityDefinition

// TwdmChannelPhyLodsPerformanceMonitoringHistoryData (class ID #444)
//	This ME collects certain PM data associated with the slot/circuit pack, hosting one or more
//	ANI-G MEs, and a specific TWDM channel. Instances of this ME are created and deleted by the OLT.
//
//	For a complete discussion of generic PM architecture, refer to clause I.4.
//
//	Relationships
//		An instance of this ME is associated with an instance of TWDM channel ME.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the TWDM channel ME. (R,
//			setbycreate) (mandatory) (2-bytes)
//
//		Interval End Time
//			Interval end time: This attribute identifies the most recently finished 15-min interval. (R)
//			(mandatory) (1-byte)
//
//		Threshold Data 1_2 Id
//			Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 and 2 MEs
//			that contains PM threshold values. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Total Received Words Protected By Bit_Interleaved Parity _32 Bip_32
//			Total received words protected by bit-interleaved parity-32 (BIP-32): The count of 4-byte words
//			included in BIP-32 check. This is a product of the number of downstream FS frames received by
//			the size of the downstream FS frame after the FEC parity byte, if any, have been removed. The
//			count applies to the entire downstream data flow, whether or not addressed to that ONT. (R)
//			(mandatory) (8-bytes)
//
//		Bip_32 Bit Error Count
//			BIP-32 bit error count: Count of the bit errors in the received downstream FS frames as measured
//			using BIP-32. If FEC is supported in the downstream direction, the BIP-32 count applies to the
//			downstream FS frame after the FEC correction has been applied and the FEC parity bytes have been
//			removed. (R) (mandatory) (4-bytes)
//
//		Corrected Psbd Hec Error Count
//			Corrected PSBd HEC error count: The count of the errors in either CFC or OCS fields of the PSBd
//			block that have been corrected using the HEC technique. (R) (mandatory) (4-bytes)
//
//		Uncorrectable Psbd Hec Error Count
//			Uncorrectable PSBd HEC error count: The count of the errors in either CFC or OCS fields of the
//			PSBd block that could not be corrected using the HEC technique. (R) (mandatory) (4-bytes)
//
//		Corrected Downstream Fs Header Hec Error Count
//			Corrected downstream FS header HEC error count: The count of the errors in the downstream FS
//			header that have been corrected using the HEC technique. (R) (mandatory) (4-bytes)
//
//		Uncorrectable Downstream Fs Header Hec Error Count
//			Uncorrectable downstream FS header HEC error count: The count of the errors in the downstream FS
//			header that could not be corrected using the HEC technique. (R) (mandatory) (4-bytes)
//
//		Total Number Of Lods Events
//			Total number of LODS events: The count of the state transitions from O5.1/O5.2 to O6, referring
//			to the ONU activation cycle state machine, clause 12 of [ITU-T-G.989.3]. (R) (mandatory)
//			(4-bytes)
//
//		Lods Events Restored In Operating Twdm Channel
//			LODS events restored in operating TWDM channel: The count of LODS events cleared automatically
//			without retuning. (R) (mandatory) (4-bytes)
//
//		Lods Events Restored In Protection Twdm Channel
//			LODS events restored in protection TWDM channel: The count of LODS events resolved by retuning
//			to a pre-configured protection TWDM channel. The event is counted against the original operating
//			channel. (R) (mandatory) (4-bytes)
//
//		Lods Events Restored In Discretionary Twdm Channel
//			LODS events restored in discretionary TWDM channel: The count of LODS events resolved by
//			retuning to a TWDM channel chosen by the ONU, without retuning. Implies that the wavelength
//			channel protection for the operating channel is not active. The event is counted against the
//			original operating channel (R) (mandatory) (4-bytes)
//
//		Lods Events Resulting In Reactivation
//			LODS events resulting in reactivation: The count of LODS events resolved through ONU
//			reactivation; that is, either TO2 (without WLCP) or TO3-+ TO4 (with WLCP) expires before the
//			downstream channel is reacquired, referring to the ONU activation cycle state machine, clause 12
//			of [ITU-T-G.989.3]. The event is counted against the original operating channel (R) (mandatory)
//			(4-bytes)
//
//		Lods Events Resulting In Reactivation After Retuning To Protection Twdm Channel
//			LODS events resulting in reactivation after retuning to protection TWDM channel: The count of
//			LODS events resolved through ONU reactivation after attempted protection switching, which turns
//			unsuccessful due to a handshake failure. (R) (mandatory) (4-bytes)
//
//		Lods Events Resulting In Reactivation After Retuning To Discretionary Twdm Channel
//			LODS events resulting in reactivation after retuning to discretionary TWDM channel: The count of
//			LODS events resolved through ONU reactivation after attempted retuning to a discretionary
//			channel, which turns unsuccessful due to a handshake failure. (R) (mandatory) (4-bytes)
//
type TwdmChannelPhyLodsPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	twdmchannelphylodsperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "TwdmChannelPhyLodsPerformanceMonitoringHistoryData",
		ClassID: 444,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			GetCurrentData,
			Set,
		),
		AllowedAttributeMask: 0xfffe,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", PointerAttributeType, 0x0000, 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, 0),
			1:  ByteField("IntervalEndTime", UnsignedIntegerAttributeType, 0x8000, 0, mapset.NewSetWith(Read), false, false, false, 1),
			2:  Uint16Field("ThresholdData12Id", UnsignedIntegerAttributeType, 0x4000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 2),
			3:  Uint64Field("TotalReceivedWordsProtectedByBitInterleavedParity32Bip32", CounterAttributeType, 0x2000, 0, mapset.NewSetWith(Read), false, false, false, 3),
			4:  Uint32Field("Bip32BitErrorCount", CounterAttributeType, 0x1000, 0, mapset.NewSetWith(Read), false, false, false, 4),
			5:  Uint32Field("CorrectedPsbdHecErrorCount", CounterAttributeType, 0x0800, 0, mapset.NewSetWith(Read), false, false, false, 5),
			6:  Uint32Field("UncorrectablePsbdHecErrorCount", CounterAttributeType, 0x0400, 0, mapset.NewSetWith(Read), false, false, false, 6),
			7:  Uint32Field("CorrectedDownstreamFsHeaderHecErrorCount", CounterAttributeType, 0x0200, 0, mapset.NewSetWith(Read), false, false, false, 7),
			8:  Uint32Field("UncorrectableDownstreamFsHeaderHecErrorCount", CounterAttributeType, 0x0100, 0, mapset.NewSetWith(Read), false, false, false, 8),
			9:  Uint32Field("TotalNumberOfLodsEvents", CounterAttributeType, 0x0080, 0, mapset.NewSetWith(Read), false, false, false, 9),
			10: Uint32Field("LodsEventsRestoredInOperatingTwdmChannel", CounterAttributeType, 0x0040, 0, mapset.NewSetWith(Read), false, false, false, 10),
			11: Uint32Field("LodsEventsRestoredInProtectionTwdmChannel", CounterAttributeType, 0x0020, 0, mapset.NewSetWith(Read), false, false, false, 11),
			12: Uint32Field("LodsEventsRestoredInDiscretionaryTwdmChannel", CounterAttributeType, 0x0010, 0, mapset.NewSetWith(Read), false, false, false, 12),
			13: Uint32Field("LodsEventsResultingInReactivation", CounterAttributeType, 0x0008, 0, mapset.NewSetWith(Read), false, false, false, 13),
			14: Uint32Field("LodsEventsResultingInReactivationAfterRetuningToProtectionTwdmChannel", CounterAttributeType, 0x0004, 0, mapset.NewSetWith(Read), false, false, false, 14),
			15: Uint32Field("LodsEventsResultingInReactivationAfterRetuningToDiscretionaryTwdmChannel", CounterAttributeType, 0x0002, 0, mapset.NewSetWith(Read), false, false, false, 15),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewTwdmChannelPhyLodsPerformanceMonitoringHistoryData (class ID 444) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewTwdmChannelPhyLodsPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*twdmchannelphylodsperformancemonitoringhistorydataBME, params...)
}
