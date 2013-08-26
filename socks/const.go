package socks

const (
	VERSION = byte(5)

	METHOD_NOT_REQUIRED  = byte(0)
	METHOD_NO_ACCEPTABLE = byte(0xff)

	RESERVED = byte(0)

	ADDR_TYPE_IP     = byte(1)
	ADDR_TYPE_IPV6   = byte(4)
	ADDR_TYPE_DOMAIN = byte(3)

	CMD_CONNECT       = byte(1)
	CMD_BIND          = byte(2)
	CMD_UDP_ASSOCIATE = byte(3)

	REP_SUCCEED                    = byte(0)
	REP_SERVER_FAILURE             = byte(1)
	REP_CONNECTION_NOT_ALLOW       = byte(2)
	REP_NETWORK_UNREACHABLE        = byte(3)
	REP_HOST_UNREACHABLE           = byte(4)
	REP_CONNECTION_REFUSED         = byte(5)
	REP_TTL_EXPIRED                = byte(6)
	REP_COMMAND_NOT_SUPPORTED      = byte(7)
	REP_ADDRESS_TYPE_NOT_SUPPORTED = byte(8)
)
