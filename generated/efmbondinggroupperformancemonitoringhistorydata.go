/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 * Copyright 2020-present Open Networking Foundation

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

// EfmBondingGroupPerformanceMonitoringHistoryDataClassID is the 16-bit ID for the OMCI
// Managed entity EFM bonding group performance monitoring history data
const EfmBondingGroupPerformanceMonitoringHistoryDataClassID ClassID = ClassID(421)

var efmbondinggroupperformancemonitoringhistorydataBME *ManagedEntityDefinition

// EfmBondingGroupPerformanceMonitoringHistoryData (class ID #421)
//	This ME collects PM data as seen at the xTU-C. Instances of this ME are created and deleted by
//	the OLT.
//
//	Relationships
//		An instance of this ME is associated with an xDSL UNI.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the EFM bonding group. (R,
//			setbycreate) (mandatory) (2-bytes)
//
//		Interval End Time
//			Interval end time: This attribute identifies the most recently finished 15-min interval. (R)
//			(mandatory) (1-byte)
//
//		Threshold Data 1_2 Id
//			Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 and 2 MEs
//			that contain PM threshold values. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Rx Bad Fragments
//			Rx bad fragments: Clause 45.2.3.33 of [IEEE 802.3]. (R) (mandatory) (4-bytes)
//
//		Rx Lost Fragments
//			Rx lost fragments: Clause 45.2.3.34 of [IEEE 802.3]. (R) (mandatory) (4-bytes)
//
//		Rx Lost Starts
//			Rx lost starts: Clause 45.2.3.35 of [IEEE 802.3]. (R) (mandatory) (4-bytes)
//
//		Rx Lost Ends
//			Rx lost ends: Clause 45.2.3.36 of [IEEE 802.3]. (R) (mandatory) (4-bytes)
//
//		Rx Frames
//			Rx frames: Number of Ethernet frames received over this group. (R) (mandatory) (4-bytes)
//
//		Tx Frames
//			Tx frames: Number of Ethernet frames transmitted over this group. (R) (mandatory) (4-bytes)
//
//		Rx Bytes
//			Rx bytes: Number of bytes contained in the Ethernet frames received over this group. (R)
//			(mandatory) (8-bytes)
//
//		Tx Bytes
//			Tx bytes: Number of bytes contained in the Ethernet frames transmitted over this group. (R)
//			(mandatory) (8-bytes)
//
//		Tx Discarded Frames
//			Tx discarded frames: Number of Ethernet frames discarded by the group transmit function. (R)
//			(mandatory) (4-bytes)
//
//		Tx Discarded Bytes
//			Tx discarded bytes: Number of bytes contained in the Ethernet frames discarded by the group
//			transmit function. (R) (mandatory) (4-bytes)
//
type EfmBondingGroupPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	efmbondinggroupperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "EfmBondingGroupPerformanceMonitoringHistoryData",
		ClassID: 421,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
			GetCurrentData,
		),
		AllowedAttributeMask: 0xfff0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", PointerAttributeType, 0x0000, 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, 0),
			1:  ByteField("IntervalEndTime", UnsignedIntegerAttributeType, 0x8000, 0, mapset.NewSetWith(Read), false, false, false, 1),
			2:  Uint16Field("ThresholdData12Id", UnsignedIntegerAttributeType, 0x4000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 2),
			3:  Uint32Field("RxBadFragments", CounterAttributeType, 0x2000, 0, mapset.NewSetWith(Read), false, false, false, 3),
			4:  Uint32Field("RxLostFragments", CounterAttributeType, 0x1000, 0, mapset.NewSetWith(Read), false, false, false, 4),
			5:  Uint32Field("RxLostStarts", CounterAttributeType, 0x0800, 0, mapset.NewSetWith(Read), false, false, false, 5),
			6:  Uint32Field("RxLostEnds", CounterAttributeType, 0x0400, 0, mapset.NewSetWith(Read), false, false, false, 6),
			7:  Uint32Field("RxFrames", CounterAttributeType, 0x0200, 0, mapset.NewSetWith(Read), false, false, false, 7),
			8:  Uint32Field("TxFrames", CounterAttributeType, 0x0100, 0, mapset.NewSetWith(Read), false, false, false, 8),
			9:  Uint64Field("RxBytes", CounterAttributeType, 0x0080, 0, mapset.NewSetWith(Read), false, false, false, 9),
			10: Uint64Field("TxBytes", CounterAttributeType, 0x0040, 0, mapset.NewSetWith(Read), false, false, false, 10),
			11: Uint32Field("TxDiscardedFrames", CounterAttributeType, 0x0020, 0, mapset.NewSetWith(Read), false, false, false, 11),
			12: Uint32Field("TxDiscardedBytes", CounterAttributeType, 0x0010, 0, mapset.NewSetWith(Read), false, false, false, 12),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
		Alarms: AlarmMap{
			0: "Rx bad fragments",
			1: "Rx lost fragments",
			2: "Rx lost starts",
			3: "Rx lost ends",
		},
	}
}

// NewEfmBondingGroupPerformanceMonitoringHistoryData (class ID 421) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewEfmBondingGroupPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*efmbondinggroupperformancemonitoringhistorydataBME, params...)
}
