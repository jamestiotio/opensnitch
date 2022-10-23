package exprs

// keywords used in the configuration to define rules.
const (
	// https://wiki.nftables.org/wiki-nftables/index.php/Netfilter_hooks#Priority_within_hook
	NFT_CHAIN_MANGLE    = "mangle"
	NFT_CHAIN_FILTER    = "filter"
	NFT_CHAIN_RAW       = "raw"
	NFT_CHAIN_SECURITY  = "security"
	NFT_CHAIN_NATDEST   = "natdest"
	NFT_CHAIN_NATSOURCE = "natsource"
	NFT_CHAIN_CONNTRACK = "conntrack"
	NFT_CHAIN_SELINUX   = "selinux"

	NFT_HOOK_INPUT       = "input"
	NFT_HOOK_OUTPUT      = "output"
	NFT_HOOK_PREROUTING  = "prerouting"
	NFT_HOOK_POSTROUTING = "postrouting"
	NFT_HOOK_INGRESS     = "ingress"
	NFT_HOOK_EGRESS      = "egress"
	NFT_HOOK_FORWARD     = "forward"

	NFT_TABLE_INET = "inet"
	NFT_TABLE_NAT  = "nat"
	// TODO
	NFT_TABLE_ARP    = "arp"
	NFT_TABLE_BRIDGE = "bridge"
	NFT_TABLE_NETDEV = "netdev"

	NFT_FAMILY_IP     = "ip"
	NFT_FAMILY_IP6    = "ip6"
	NFT_FAMILY_INET   = "inet"
	NFT_FAMILY_BRIDGE = "bridge"
	NFT_FAMILY_ARP    = "arp"
	NFT_FAMILY_NETDEV = "netdev"

	VERDICT_ACCEPT = "accept"
	VERDICT_DROP   = "drop"
	VERDICT_REJECT = "reject"
	VERDICT_RETURN = "return"
	VERDICT_QUEUE  = "queue"

	VERDICT_JUMP = "jump"
	// TODO
	VERDICT_GOTO       = "goto"
	VERDICT_STOP       = "stop"
	VERDICT_STOLEN     = "stolen"
	VERDICT_CONTINUE   = "continue"
	VERDICT_MASQUERADE = "masquerade"
	VERDICT_DNAT       = "dnat"
	VERDICT_SNAT       = "snat"
	VERDICT_REDIRECT   = "redirect"
	VERDICT_TPROXY     = "tproxy"

	NFT_PARM_TO = "to"

	NFT_QUEUE_NUM     = "num"
	NFT_QUEUE_BY_PASS = "queue-bypass"

	NFT_MASQ_RANDOM       = "random"
	NFT_MASQ_FULLY_RANDOM = "fully-random"
	NFT_MASQ_PERSISTENT   = "persistent"

	NFT_PROTOCOL  = "protocol"
	NFT_SPORT     = "sport"
	NFT_DPORT     = "dport"
	NFT_SADDR     = "saddr"
	NFT_DADDR     = "daddr"
	NFT_ICMP_CODE = "code"
	NFT_ICMP_TYPE = "type"

	NFT_ETHER = "ether"

	NFT_IIFNAME = "iifname"
	NFT_OIFNAME = "oifname"

	NFT_LOG        = "log"
	NFT_LOG_PREFIX = "prefix"
	// TODO
	NFT_LOG_LEVEL        = "level"
	NFT_LOG_LEVEL_EMERG  = "emerg"
	NFT_LOG_LEVEL_ALERT  = "alert"
	NFT_LOG_LEVEL_CRIT   = "crit"
	NFT_LOG_LEVEL_ERR    = "err"
	NFT_LOG_LEVEL_WARN   = "warn"
	NFT_LOG_LEVEL_NOTICE = "notice"
	NFT_LOG_LEVEL_INFO   = "info"
	NFT_LOG_LEVEL_DEBUG  = "debug"
	NFT_LOG_LEVEL_AUDIT  = "audit"
	NFT_LOG_FLAGS        = "flags"

	NFT_CT               = "ct"
	NFT_CT_STATE         = "state"
	NFT_CT_SET_MARK      = "set"
	NFT_CT_MARK          = "mark"
	CT_STATE_NEW         = "new"
	CT_STATE_ESTABLISHED = "established"
	CT_STATE_RELATED     = "related"
	CT_STATE_INVALID     = "invalid"

	NFT_NOTRACK = "notrack"

	NFT_QUOTA            = "quota"
	NFT_QUOTA_UNTIL      = "until"
	NFT_QUOTA_OVER       = "over"
	NFT_QUOTA_USED       = "used"
	NFT_QUOTA_UNIT_BYTES = "bytes"
	NFT_QUOTA_UNIT_KB    = "kbytes"
	NFT_QUOTA_UNIT_MB    = "mbytes"
	NFT_QUOTA_UNIT_GB    = "gbytes"

	NFT_COUNTER         = "counter"
	NFT_COUNTER_NAME    = "name"
	NFT_COUNTER_PACKETS = "packets"
	NFT_COUNTER_BYTES   = "bytes"

	NFT_LIMIT             = "limit"
	NFT_LIMIT_OVER        = "over"
	NFT_LIMIT_BURST       = "burst"
	NFT_LIMIT_UNITS_RATE  = "rate-units"
	NFT_LIMIT_UNITS_TIME  = "time-units"
	NFT_LIMIT_UNITS       = "units"
	NFT_LIMIT_UNIT_SECOND = "second"
	NFT_LIMIT_UNIT_MINUTE = "minute"
	NFT_LIMIT_UNIT_HOUR   = "hour"
	NFT_LIMIT_UNIT_DAY    = "day"
	NFT_LIMIT_UNIT_KBYTES = "kbytes"
	NFT_LIMIT_UNIT_MBYTES = "mbytes"

	NFT_META          = "meta"
	NFT_META_MARK     = "mark"
	NFT_META_SET_MARK = "set"
	NFT_META_PRIORITY = "priority"
	NFT_META_NFTRACE  = "nftrace"
	NFT_META_SET      = "set"
	NFT_META_SKUID    = "skuid"
	NFT_META_SKGID    = "skgid"
	NFT_META_L4PROTO  = "l4proto"
	NFT_META_PROTOCOL = "protocol"

	NFT_PROTO_UDP      = "udp"
	NFT_PROTO_UDPLITE  = "udplite"
	NFT_PROTO_TCP      = "tcp"
	NFT_PROTO_SCTP     = "sctp"
	NFT_PROTO_DCCP     = "dccp"
	NFT_PROTO_ICMP     = "icmp"
	NFT_PROTO_ICMPX    = "icmpx"
	NFT_PROTO_ICMPv6   = "icmpv6"
	NFT_PROTO_AH       = "ah"
	NFT_PROTO_ETHERNET = "ethernet"
	NFT_PROTO_GRE      = "gre"
	NFT_PROTO_IP       = "ip"
	NFT_PROTO_IPIP     = "ipip"
	NFT_PROTO_L2TP     = "l2tp"
	NFT_PROTO_COMP     = "comp"
	NFT_PROTO_IGMP     = "igmp"
	NFT_PROTO_ESP      = "esp"
	NFT_PROTO_RAW      = "raw"
	NFT_PROTO_ENCAP    = "encap"

	ICMP_NO_ROUTE           = "no-route"
	ICMP_PROT_UNREACHABLE   = "prot-unreachable"
	ICMP_PORT_UNREACHABLE   = "port-unreachable"
	ICMP_NET_UNREACHABLE    = "net-unreachable"
	ICMP_ADDR_UNREACHABLE   = "addr-unreachable"
	ICMP_HOST_UNREACHABLE   = "host-unreachable"
	ICMP_NET_PROHIBITED     = "net-prohibited"
	ICMP_HOST_PROHIBITED    = "host-prohibited"
	ICMP_ADMIN_PROHIBITED   = "admin-prohibited"
	ICMP_REJECT_ROUTE       = "reject-route"
	ICMP_REJECT_POLICY_FAIL = "policy-fail"

	ICMP_ECHO_REPLY           = "echo-reply"
	ICMP_ECHO_REQUEST         = "echo-request"
	ICMP_SOURCE_QUENCH        = "source-quench"
	ICMP_DEST_UNREACHABLE     = "destination-unreachable"
	ICMP_REDIRECT             = "redirect"
	ICMP_TIME_EXCEEDED        = "time-exceeded"
	ICMP_INFO_REQUEST         = "info-request"
	ICMP_INFO_REPLY           = "info-reply"
	ICMP_PARAMETER_PROBLEM    = "parameter-problem"
	ICMP_TIMESTAMP_REQUEST    = "timestamp-request"
	ICMP_TIMESTAMP_REPLY      = "timestamp-reply"
	ICMP_ROUTER_ADVERTISEMENT = "router-advertisement"
	ICMP_ROUTER_SOLICITATION  = "router-solicitation"
	ICMP_ADDRESS_MASK_REQUEST = "address-mask-request"
	ICMP_ADDRESS_MASK_REPLY   = "address-mask-reply"

	ICMP_PACKET_TOO_BIG          = "packet-too-big"
	ICMP_NEIGHBOUR_SOLICITATION  = "neighbour-solicitation"
	ICMP_NEIGHBOUR_ADVERTISEMENT = "neighbour-advertisement"
)
