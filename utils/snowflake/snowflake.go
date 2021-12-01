package snowflake

import (
	"context"
	"net"
	"strconv"
	"strings"

	"github.com/bwmarrin/snowflake"
	"github.com/canbefree/magazine/utils"
	"github.com/canbefree/magazine/utils/log"
)

type SnowflakeIP struct {
	nodeInt int64
	node    *snowflake.Node
}

func NewSnowflakeIP() *SnowflakeIP {
	nodeInt := genNodeByIP()
	return &SnowflakeIP{
		nodeInt: nodeInt,
		node:    newNode(nodeInt),
	}
}

func (s *SnowflakeIP) GenerateID() uint64 {
	return uint64(s.node.Generate().Int64())
}

func (s *SnowflakeIP) GetNodeInt() int64 {
	return s.nodeInt
}

func newNode(nodeInt int64) *snowflake.Node {
	node, err := snowflake.NewNode(nodeInt)
	utils.PanicIfErr(err)
	return node
}

func genNodeByIP() (node int64) {
	// 理论上不会重复,建议对机器ip,生成的node进行上报
	// report node
	addrs, err := InterfaceAddrs()
	if err == nil {
		for _, address := range addrs {
			ipnet, ok := address.(*net.IPNet)
			if !ok {
				continue
			}

			log.Tracef(context.TODO(), "address:%v", address.String())
			log.Tracef(context.TODO(), "network:%v", address.Network())

			if !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				ip := ipnet.IP.To4().String()
				log.Tracef(context.TODO(), "ip:%v", ip)
				ns := strings.Split(ip, ".")
				var base int64 = 10
				for _, n := range ns[len(ns)-2:] {
					i, err := strconv.Atoi(n)
					utils.PanicIfErr(err)
					node = node + int64(i*int(base))
					base = base / 10
				}
				return node % 1024
			}
		}
	}
	// 获取不到node
	utils.PanicIfErr(err)
	return
}

// for gostub
var InterfaceAddrs = net.InterfaceAddrs
