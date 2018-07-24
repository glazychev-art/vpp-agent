// Code generated by govpp binapi-generator DO NOT EDIT.
// Package sr represents the VPP binary API of the 'sr' VPP module.
// Generated from '/usr/share/vpp/api/sr.api.json'
package sr

import "git.fd.io/govpp.git/api"

// Srv6Sid represents the VPP binary API data type 'srv6_sid'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 577:
//
//            "srv6_sid",
//            [
//                "u8",
//                "addr",
//                16
//            ],
//            {
//                "crc": "0x6ee67284"
//            }
//
type Srv6Sid struct {
	Addr []byte `struc:"[16]byte"`
}

func (*Srv6Sid) GetTypeName() string {
	return "srv6_sid"
}
func (*Srv6Sid) GetCrcString() string {
	return "6ee67284"
}

// Srv6SidList represents the VPP binary API data type 'srv6_sid_list'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 588:
//
//            "srv6_sid_list",
//            [
//                "u8",
//                "num_sids"
//            ],
//            [
//                "u32",
//                "weight"
//            ],
//            [
//                "vl_api_srv6_sid_t",
//                "sids",
//                16
//            ],
//            {
//                "crc": "0x4066af74"
//            }
//
type Srv6SidList struct {
	NumSids uint8 `struc:"sizeof=Sids"`
	Weight  uint32
	Sids    []Srv6Sid
}

func (*Srv6SidList) GetTypeName() string {
	return "srv6_sid_list"
}
func (*Srv6SidList) GetCrcString() string {
	return "4066af74"
}

// SrIP6Address represents the VPP binary API data type 'sr_ip6_address'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 607:
//
//            "sr_ip6_address",
//            [
//                "u8",
//                "data",
//                16
//            ],
//            {
//                "crc": "0xbea0c5e6"
//            }
//
type SrIP6Address struct {
	Data []byte `struc:"[16]byte"`
}

func (*SrIP6Address) GetTypeName() string {
	return "sr_ip6_address"
}
func (*SrIP6Address) GetCrcString() string {
	return "bea0c5e6"
}

// SrLocalsidAddDel represents the VPP binary API message 'sr_localsid_add_del'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 4:
//
//            "sr_localsid_add_del",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "is_del"
//            ],
//            [
//                "vl_api_srv6_sid_t",
//                "localsid"
//            ],
//            [
//                "u8",
//                "end_psp"
//            ],
//            [
//                "u8",
//                "behavior"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            [
//                "u32",
//                "vlan_index"
//            ],
//            [
//                "u32",
//                "fib_table"
//            ],
//            [
//                "u8",
//                "nh_addr6",
//                16
//            ],
//            [
//                "u8",
//                "nh_addr4",
//                4
//            ],
//            {
//                "crc": "0x20d478a0"
//            }
//
type SrLocalsidAddDel struct {
	IsDel     uint8
	Localsid  Srv6Sid
	EndPsp    uint8
	Behavior  uint8
	SwIfIndex uint32
	VlanIndex uint32
	FibTable  uint32
	NhAddr6   []byte `struc:"[16]byte"`
	NhAddr4   []byte `struc:"[4]byte"`
}

func (*SrLocalsidAddDel) GetMessageName() string {
	return "sr_localsid_add_del"
}
func (*SrLocalsidAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrLocalsidAddDel) GetCrcString() string {
	return "20d478a0"
}
func NewSrLocalsidAddDel() api.Message {
	return &SrLocalsidAddDel{}
}

// SrLocalsidAddDelReply represents the VPP binary API message 'sr_localsid_add_del_reply'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 60:
//
//            "sr_localsid_add_del_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type SrLocalsidAddDelReply struct {
	Retval int32
}

func (*SrLocalsidAddDelReply) GetMessageName() string {
	return "sr_localsid_add_del_reply"
}
func (*SrLocalsidAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrLocalsidAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func NewSrLocalsidAddDelReply() api.Message {
	return &SrLocalsidAddDelReply{}
}

// SrPolicyAdd represents the VPP binary API message 'sr_policy_add'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 78:
//
//            "sr_policy_add",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "bsid_addr",
//                16
//            ],
//            [
//                "u32",
//                "weight"
//            ],
//            [
//                "u8",
//                "is_encap"
//            ],
//            [
//                "u8",
//                "type"
//            ],
//            [
//                "u32",
//                "fib_table"
//            ],
//            [
//                "vl_api_srv6_sid_list_t",
//                "sids"
//            ],
//            {
//                "crc": "0xa1676c1f"
//            }
//
type SrPolicyAdd struct {
	BsidAddr []byte `struc:"[16]byte"`
	Weight   uint32
	IsEncap  uint8
	Type     uint8
	FibTable uint32
	Sids     Srv6SidList
}

func (*SrPolicyAdd) GetMessageName() string {
	return "sr_policy_add"
}
func (*SrPolicyAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrPolicyAdd) GetCrcString() string {
	return "a1676c1f"
}
func NewSrPolicyAdd() api.Message {
	return &SrPolicyAdd{}
}

// SrPolicyAddReply represents the VPP binary API message 'sr_policy_add_reply'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 121:
//
//            "sr_policy_add_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type SrPolicyAddReply struct {
	Retval int32
}

func (*SrPolicyAddReply) GetMessageName() string {
	return "sr_policy_add_reply"
}
func (*SrPolicyAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrPolicyAddReply) GetCrcString() string {
	return "e8d4e804"
}
func NewSrPolicyAddReply() api.Message {
	return &SrPolicyAddReply{}
}

// SrPolicyMod represents the VPP binary API message 'sr_policy_mod'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 139:
//
//            "sr_policy_mod",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "bsid_addr",
//                16
//            ],
//            [
//                "u32",
//                "sr_policy_index"
//            ],
//            [
//                "u32",
//                "fib_table"
//            ],
//            [
//                "u8",
//                "operation"
//            ],
//            [
//                "u32",
//                "sl_index"
//            ],
//            [
//                "u32",
//                "weight"
//            ],
//            [
//                "vl_api_srv6_sid_list_t",
//                "sids"
//            ],
//            {
//                "crc": "0x51252136"
//            }
//
type SrPolicyMod struct {
	BsidAddr      []byte `struc:"[16]byte"`
	SrPolicyIndex uint32
	FibTable      uint32
	Operation     uint8
	SlIndex       uint32
	Weight        uint32
	Sids          Srv6SidList
}

func (*SrPolicyMod) GetMessageName() string {
	return "sr_policy_mod"
}
func (*SrPolicyMod) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrPolicyMod) GetCrcString() string {
	return "51252136"
}
func NewSrPolicyMod() api.Message {
	return &SrPolicyMod{}
}

// SrPolicyModReply represents the VPP binary API message 'sr_policy_mod_reply'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 186:
//
//            "sr_policy_mod_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type SrPolicyModReply struct {
	Retval int32
}

func (*SrPolicyModReply) GetMessageName() string {
	return "sr_policy_mod_reply"
}
func (*SrPolicyModReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrPolicyModReply) GetCrcString() string {
	return "e8d4e804"
}
func NewSrPolicyModReply() api.Message {
	return &SrPolicyModReply{}
}

// SrPolicyDel represents the VPP binary API message 'sr_policy_del'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 204:
//
//            "sr_policy_del",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "vl_api_srv6_sid_t",
//                "bsid_addr"
//            ],
//            [
//                "u32",
//                "sr_policy_index"
//            ],
//            {
//                "crc": "0x168e1a98"
//            }
//
type SrPolicyDel struct {
	BsidAddr      Srv6Sid
	SrPolicyIndex uint32
}

func (*SrPolicyDel) GetMessageName() string {
	return "sr_policy_del"
}
func (*SrPolicyDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrPolicyDel) GetCrcString() string {
	return "168e1a98"
}
func NewSrPolicyDel() api.Message {
	return &SrPolicyDel{}
}

// SrPolicyDelReply represents the VPP binary API message 'sr_policy_del_reply'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 230:
//
//            "sr_policy_del_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type SrPolicyDelReply struct {
	Retval int32
}

func (*SrPolicyDelReply) GetMessageName() string {
	return "sr_policy_del_reply"
}
func (*SrPolicyDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrPolicyDelReply) GetCrcString() string {
	return "e8d4e804"
}
func NewSrPolicyDelReply() api.Message {
	return &SrPolicyDelReply{}
}

// SrSetEncapSource represents the VPP binary API message 'sr_set_encap_source'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 248:
//
//            "sr_set_encap_source",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "encaps_source",
//                16
//            ],
//            {
//                "crc": "0xd05bb4de"
//            }
//
type SrSetEncapSource struct {
	EncapsSource []byte `struc:"[16]byte"`
}

func (*SrSetEncapSource) GetMessageName() string {
	return "sr_set_encap_source"
}
func (*SrSetEncapSource) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrSetEncapSource) GetCrcString() string {
	return "d05bb4de"
}
func NewSrSetEncapSource() api.Message {
	return &SrSetEncapSource{}
}

// SrSetEncapSourceReply represents the VPP binary API message 'sr_set_encap_source_reply'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 271:
//
//            "sr_set_encap_source_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type SrSetEncapSourceReply struct {
	Retval int32
}

func (*SrSetEncapSourceReply) GetMessageName() string {
	return "sr_set_encap_source_reply"
}
func (*SrSetEncapSourceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrSetEncapSourceReply) GetCrcString() string {
	return "e8d4e804"
}
func NewSrSetEncapSourceReply() api.Message {
	return &SrSetEncapSourceReply{}
}

// SrSteeringAddDel represents the VPP binary API message 'sr_steering_add_del'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 289:
//
//            "sr_steering_add_del",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "is_del"
//            ],
//            [
//                "u8",
//                "bsid_addr",
//                16
//            ],
//            [
//                "u32",
//                "sr_policy_index"
//            ],
//            [
//                "u32",
//                "table_id"
//            ],
//            [
//                "u8",
//                "prefix_addr",
//                16
//            ],
//            [
//                "u32",
//                "mask_width"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            [
//                "u8",
//                "traffic_type"
//            ],
//            {
//                "crc": "0x28b5dcab"
//            }
//
type SrSteeringAddDel struct {
	IsDel         uint8
	BsidAddr      []byte `struc:"[16]byte"`
	SrPolicyIndex uint32
	TableID       uint32
	PrefixAddr    []byte `struc:"[16]byte"`
	MaskWidth     uint32
	SwIfIndex     uint32
	TrafficType   uint8
}

func (*SrSteeringAddDel) GetMessageName() string {
	return "sr_steering_add_del"
}
func (*SrSteeringAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrSteeringAddDel) GetCrcString() string {
	return "28b5dcab"
}
func NewSrSteeringAddDel() api.Message {
	return &SrSteeringAddDel{}
}

// SrSteeringAddDelReply represents the VPP binary API message 'sr_steering_add_del_reply'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 341:
//
//            "sr_steering_add_del_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type SrSteeringAddDelReply struct {
	Retval int32
}

func (*SrSteeringAddDelReply) GetMessageName() string {
	return "sr_steering_add_del_reply"
}
func (*SrSteeringAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrSteeringAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func NewSrSteeringAddDelReply() api.Message {
	return &SrSteeringAddDelReply{}
}

// SrLocalsidsDump represents the VPP binary API message 'sr_localsids_dump'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 359:
//
//            "sr_localsids_dump",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            {
//                "crc": "0x51077d14"
//            }
//
type SrLocalsidsDump struct {
}

func (*SrLocalsidsDump) GetMessageName() string {
	return "sr_localsids_dump"
}
func (*SrLocalsidsDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrLocalsidsDump) GetCrcString() string {
	return "51077d14"
}
func NewSrLocalsidsDump() api.Message {
	return &SrLocalsidsDump{}
}

// SrLocalsidsDetails represents the VPP binary API message 'sr_localsids_details'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 377:
//
//            "sr_localsids_details",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "vl_api_srv6_sid_t",
//                "addr"
//            ],
//            [
//                "u8",
//                "end_psp"
//            ],
//            [
//                "u16",
//                "behavior"
//            ],
//            [
//                "u32",
//                "fib_table"
//            ],
//            [
//                "u32",
//                "vlan_index"
//            ],
//            [
//                "u8",
//                "xconnect_nh_addr6",
//                16
//            ],
//            [
//                "u8",
//                "xconnect_nh_addr4",
//                4
//            ],
//            [
//                "u32",
//                "xconnect_iface_or_vrf_table"
//            ],
//            {
//                "crc": "0x7ff35765"
//            }
//
type SrLocalsidsDetails struct {
	Addr                    Srv6Sid
	EndPsp                  uint8
	Behavior                uint16
	FibTable                uint32
	VlanIndex               uint32
	XconnectNhAddr6         []byte `struc:"[16]byte"`
	XconnectNhAddr4         []byte `struc:"[4]byte"`
	XconnectIfaceOrVrfTable uint32
}

func (*SrLocalsidsDetails) GetMessageName() string {
	return "sr_localsids_details"
}
func (*SrLocalsidsDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrLocalsidsDetails) GetCrcString() string {
	return "7ff35765"
}
func NewSrLocalsidsDetails() api.Message {
	return &SrLocalsidsDetails{}
}

// SrPoliciesDump represents the VPP binary API message 'sr_policies_dump'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 425:
//
//            "sr_policies_dump",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            {
//                "crc": "0x51077d14"
//            }
//
type SrPoliciesDump struct {
}

func (*SrPoliciesDump) GetMessageName() string {
	return "sr_policies_dump"
}
func (*SrPoliciesDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrPoliciesDump) GetCrcString() string {
	return "51077d14"
}
func NewSrPoliciesDump() api.Message {
	return &SrPoliciesDump{}
}

// SrPoliciesDetails represents the VPP binary API message 'sr_policies_details'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 443:
//
//            "sr_policies_details",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "vl_api_srv6_sid_t",
//                "bsid"
//            ],
//            [
//                "u8",
//                "type"
//            ],
//            [
//                "u8",
//                "is_encap"
//            ],
//            [
//                "u32",
//                "fib_table"
//            ],
//            [
//                "u8",
//                "num_sid_lists"
//            ],
//            [
//                "vl_api_srv6_sid_list_t",
//                "sid_lists",
//                0,
//                "num_sid_lists"
//            ],
//            {
//                "crc": "0xae838a76"
//            }
//
type SrPoliciesDetails struct {
	Bsid        Srv6Sid
	Type        uint8
	IsEncap     uint8
	FibTable    uint32
	NumSidLists uint8 `struc:"sizeof=SidLists"`
	SidLists    []Srv6SidList
}

func (*SrPoliciesDetails) GetMessageName() string {
	return "sr_policies_details"
}
func (*SrPoliciesDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrPoliciesDetails) GetCrcString() string {
	return "ae838a76"
}
func NewSrPoliciesDetails() api.Message {
	return &SrPoliciesDetails{}
}

// SrSteeringPolDump represents the VPP binary API message 'sr_steering_pol_dump'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 483:
//
//            "sr_steering_pol_dump",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            {
//                "crc": "0x51077d14"
//            }
//
type SrSteeringPolDump struct {
}

func (*SrSteeringPolDump) GetMessageName() string {
	return "sr_steering_pol_dump"
}
func (*SrSteeringPolDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SrSteeringPolDump) GetCrcString() string {
	return "51077d14"
}
func NewSrSteeringPolDump() api.Message {
	return &SrSteeringPolDump{}
}

// SrSteeringPolDetails represents the VPP binary API message 'sr_steering_pol_details'.
// Generated from '/usr/share/vpp/api/sr.api.json', line 501:
//
//            "sr_steering_pol_details",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "traffic_type"
//            ],
//            [
//                "u32",
//                "fib_table"
//            ],
//            [
//                "u8",
//                "prefix_addr",
//                16
//            ],
//            [
//                "u32",
//                "mask_width"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            [
//                "vl_api_srv6_sid_t",
//                "bsid"
//            ],
//            {
//                "crc": "0x1c756f85"
//            }
//
type SrSteeringPolDetails struct {
	TrafficType uint8
	FibTable    uint32
	PrefixAddr  []byte `struc:"[16]byte"`
	MaskWidth   uint32
	SwIfIndex   uint32
	Bsid        Srv6Sid
}

func (*SrSteeringPolDetails) GetMessageName() string {
	return "sr_steering_pol_details"
}
func (*SrSteeringPolDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SrSteeringPolDetails) GetCrcString() string {
	return "1c756f85"
}
func NewSrSteeringPolDetails() api.Message {
	return &SrSteeringPolDetails{}
}
