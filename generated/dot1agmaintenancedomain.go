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

// Dot1AgMaintenanceDomainClassID is the 16-bit ID for the OMCI
// Managed entity Dot1ag maintenance domain
const Dot1AgMaintenanceDomainClassID ClassID = ClassID(299)

var dot1agmaintenancedomainBME *ManagedEntityDefinition

// Dot1AgMaintenanceDomain (class ID #299)
//	In [IEEE 802.1ag], a maintenance domain (MD) is a context within which configuration fault
//	management (CFM) connectivity verification can occur. Individual services (maintenance
//	associations, MAs) exist within an MD. An MD is created and deleted by the OLT. The MD ME is
//	specified by [IEEE 802.1ag] in such a way that the same provisioning can be used for all
//	associated systems in a network; the OMCI definition accordingly avoids ONU-specific information
//	such as pointers.
//
//	Relationships
//		Several MDs may be associated with a given bridge, at various MD levels, and a given MD may be
//		associated with any number of bridges.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies an instance of this ME. The values 0 and
//			0xFFFF are reserved. (R, setbycreate) (mandatory) (2-bytes)
//
//		Md Level
//			MD level:	This attribute ranges from 0..7 and specifies the maintenance level of this MD. Higher
//			numbers have wider geographic scope. (R,-W, setbycreate) (mandatory) (1-byte)
//
//		Md Name Format
//			MD name format: This attribute specifies one of several possible formats for the MD name
//			attribute. (R,-W, setbycreate) (mandatory) (1-byte)
//
//		Md Name 1 Md Name 2
//			MD name 1, MD name 2:These two attributes may be regarded as a 50-byte octet string whose value
//			is the left-justified maintenance domain name. The MD name may or may not be a printable
//			character string, so an octet string is the appropriate representation. If the MD name format
//			specifies a DNS-like name or a character string, the string is null-terminated; otherwise, its
//			length is determined by the MD name format. If the MD has no name (MD name format-=-0), this
//			attribute is undefined. Note that binary comparisons of the MD name are made in other CFM state
//			machines, so blanks, alphabetic case, etc., are significant. Also, note that the MD name and the
//			MA name must be packed (with additional bytes) into 48-byte CFM message headers. (R,-W)
//			(mandatory if MD name format is not 1) (25-bytes * 2 attributes)
//
//		Maintenance Domain Intermediate Point Half Function Mhf Creation
//			(R,-W, setbycreate) (mandatory) (1-byte)
//
//		Sender Id Permission
//			(R,-W, setbycreate) (mandatory) (1-byte)
//
type Dot1AgMaintenanceDomain struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	dot1agmaintenancedomainBME = &ManagedEntityDefinition{
		Name:    "Dot1AgMaintenanceDomain",
		ClassID: 299,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xf800,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0x0000, 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, 0),
			1: ByteField("MdLevel", UnsignedIntegerAttributeType, 0x8000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 1),
			2: ByteField("MdNameFormat", UnsignedIntegerAttributeType, 0x4000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 2),
			3: MultiByteField("MdName1MdName2", OctetsAttributeType, 0x2000, 25, toOctets("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=="), mapset.NewSetWith(Read, Write), false, false, false, 3),
			4: ByteField("MaintenanceDomainIntermediatePointHalfFunctionMhfCreation", UnsignedIntegerAttributeType, 0x1000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 4),
			5: ByteField("SenderIdPermission", UnsignedIntegerAttributeType, 0x0800, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 5),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewDot1AgMaintenanceDomain (class ID 299) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewDot1AgMaintenanceDomain(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*dot1agmaintenancedomainBME, params...)
}
