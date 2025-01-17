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


var unknownVendorSpecificBME *ManagedEntityDefinition

type UnknownVendorSpecific struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	unknownVendorSpecificBME = &ManagedEntityDefinition{
		Name:    "UnknownVendorSpecificManagedEntity",
		ClassID: 0,
		MessageTypes: mapset.NewSetWith(
			MibUploadNext,
			AlarmNotification,
			AttributeValueChange,
		),
		AllowedAttributeMask: 0xffff,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0x0000, 0, mapset.NewSetWith(Read), false, false, false, 0),
			1: MultiByteField("UnknownAttr_1", OctetsAttributeType, 0x8000, 1, toOctets("AA=="), mapset.NewSetWith(Read), true, true, false, 1),
			2: MultiByteField("UnknownAttr_2", OctetsAttributeType, 0x4000, 1, toOctets("AA=="), mapset.NewSetWith(Read), true, true, false, 2),
			3: MultiByteField("UnknownAttr_3", OctetsAttributeType, 0x2000, 1, toOctets("AA=="), mapset.NewSetWith(Read), true, true, false, 3),
			4: MultiByteField("UnknownAttr_4", OctetsAttributeType, 0x1000, 1, toOctets("AA=="), mapset.NewSetWith(Read), true, true, false, 4),
			5: MultiByteField("UnknownAttr_5", OctetsAttributeType, 0x0800, 1, toOctets("AA=="), mapset.NewSetWith(Read), true, true, false, 5),
			6: MultiByteField("UnknownAttr_6", OctetsAttributeType, 0x0400, 1, toOctets("AA=="), mapset.NewSetWith(Read), true, true, false, 6),
			7: MultiByteField("UnknownAttr_7", OctetsAttributeType, 0x0200, 1, toOctets("AA=="), mapset.NewSetWith(Read), true, true, false, 7),
			8: MultiByteField("UnknownAttr_8", OctetsAttributeType, 0x0100, 1, toOctets("AA=="), mapset.NewSetWith(Read), true, true, false, 8),
			9: MultiByteField("UnknownAttr_9", OctetsAttributeType, 0x0080, 1, toOctets("AA=="), mapset.NewSetWith(Read), true, true, false, 9),
			10: MultiByteField("UnknownAttr_10", OctetsAttributeType, 0x0040, 1, toOctets("AA=="), mapset.NewSetWith(Read), true, true, false, 10),
			11: MultiByteField("UnknownAttr_11", OctetsAttributeType, 0x0020, 1, toOctets("AA=="), mapset.NewSetWith(Read), true, true, false, 11),
			12: MultiByteField("UnknownAttr_12", OctetsAttributeType, 0x0010, 1, toOctets("AA=="), mapset.NewSetWith(Read), true, true, false, 12),
			13: MultiByteField("UnknownAttr_13", OctetsAttributeType, 0x0008, 1, toOctets("AA=="), mapset.NewSetWith(Read), true, true, false, 13),
			14: MultiByteField("UnknownAttr_14", OctetsAttributeType, 0x0004, 1, toOctets("AA=="), mapset.NewSetWith(Read), true, true, false, 14),
			15: MultiByteField("UnknownAttr_15", OctetsAttributeType, 0x0002, 1, toOctets("AA=="), mapset.NewSetWith(Read), true, true, false, 15),
			16: MultiByteField("UnknownAttr_16", OctetsAttributeType, 0x0001, 1, toOctets("AA=="), mapset.NewSetWith(Read), true, true, false, 16),
		},
		Access:  UnknownAccess,
		Support: UnsupportedVendorSpecificManagedEntity,
	}
}

func NewUnknownVendorSpecificME(classID ClassID, params ...ParamData) (*ManagedEntity, OmciErrors) {
	return newUnknownManagedEntity(classID, *unknownVendorSpecificBME, params...)
}
