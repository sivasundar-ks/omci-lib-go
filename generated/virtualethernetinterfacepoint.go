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

// VirtualEthernetInterfacePointClassID is the 16-bit ID for the OMCI
// Managed entity Virtual Ethernet interface point
const VirtualEthernetInterfacePointClassID ClassID = ClassID(329)

var virtualethernetinterfacepointBME *ManagedEntityDefinition

// VirtualEthernetInterfacePoint (class ID #329)
//	This ME represents the data plane hand-off point in an ONU to a separate (non-OMCI) management
//	domain. The VEIP is managed by the OMCI, and is potentially known to the non-OMCI management
//	domain. One or more Ethernet traffic flows are present at this boundary.
//
//	Instances of this ME are automatically created and deleted by the ONU. This is necessary because
//	the required downstream priority queues are subject to physical implementation constraints. The
//	OLT may use one or more of the VEIPs created by the ONU.
//
//	It is expected that the ONU will create one VEIP for each non-OMCI management domain. At the
//	vendor's discretion, a VEIP may be created for each traffic class.
//
//	Relationships
//		An instance of this ME is associated with an instance of a virtual Ethernet interface between
//		OMCI and non-OMCI management domains.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. When used
//			independently of a cardholder and circuit pack, the ONU should assign IDs in the sequence 1, 2,
//			.... When used in conjunction with a cardholder and circuit pack, this 2 byte number indicates
//			the physical position of the VEIP. The first byte is the slot ID (defined in clause 9.1.5). The
//			second byte is the port ID, with the range 1..255. The values 0 and 0xFFFF are reserved. (R)
//			(mandatory) (2 bytes)
//
//		Administrative State
//			Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this
//			ME. Administrative state is further described in clause A.1.6. (R,-W) (mandatory) (1-byte)
//
//		Operational State
//			Operational state: This attribute indicates whether the ME is capable of performing its
//			function. Valid values are enabled (0) and disabled (1). (R) (optional) (1-byte)
//
//		Interdomain Name
//			Interdomain name: This attribute is a character string that provides an optional way to identify
//			the VEIP to a non-OMCI management domain. The interface may also be identified by its ME ID,
//			[b-IANA] assigned port and possibly other ways. If the vendor offers no information in this
//			attribute, it should be set to a sequence of null bytes. (R,-W) (optional) (25-bytes)
//
//		Tcp_Udp Pointer
//			TCP/UDP pointer: This attribute points to an instance of the TCP/UDP config data ME, which
//			provides for OMCI management of the non-OMCI management domain's IP connectivity. If no OMCI
//			management of the non-OMCI domain's IP connectivity is required, this attribute may be omitted
//			or set to its default, a null pointer. (R,-W) (optional) (2-bytes)
//
//		Iana Assigned Port
//			IANA assigned port: This attribute contains the TCP or UDP port value as assigned by  [b-IANA]
//			for the management protocol associated with this virtual Ethernet interface. This attribute is
//			to be regarded as a hint, not as a requirement that management communications use this port; the
//			actual port and protocol are specified in the associated TCP/UDP config data ME. If no port has
//			been assigned or if the management protocol is free to be chosen at run-time, this attribute
//			should be set to 0xFFFF. (R) (mandatory) (2-bytes)
//
type VirtualEthernetInterfacePoint struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	virtualethernetinterfacepointBME = &ManagedEntityDefinition{
		Name:    "VirtualEthernetInterfacePoint",
		ClassID: 329,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
		),
		AllowedAttributeMask: 0xf800,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0x0000, 0, mapset.NewSetWith(Read), false, false, false, 0),
			1: ByteField("AdministrativeState", EnumerationAttributeType, 0x8000, 0, mapset.NewSetWith(Read, Write), false, false, false, 1),
			2: ByteField("OperationalState", EnumerationAttributeType, 0x4000, 0, mapset.NewSetWith(Read), true, true, false, 2),
			3: MultiByteField("InterdomainName", StringAttributeType, 0x2000, 25, toOctets("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=="), mapset.NewSetWith(Read, Write), false, true, false, 3),
			4: Uint16Field("TcpUdpPointer", PointerAttributeType, 0x1000, 0, mapset.NewSetWith(Read, Write), false, true, false, 4),
			5: Uint16Field("IanaAssignedPort", UnsignedIntegerAttributeType, 0x0800, 65535, mapset.NewSetWith(Read), false, false, false, 5),
		},
		Access:  CreatedByOnu,
		Support: UnknownSupport,
	}
}

// NewVirtualEthernetInterfacePoint (class ID 329) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewVirtualEthernetInterfacePoint(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*virtualethernetinterfacepointBME, params...)
}
