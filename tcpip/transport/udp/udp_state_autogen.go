// automatically generated by stateify.

package udp

import (
	"gvisor.dev/gvisor/pkg/state"
	"gvisor.dev/gvisor/pkg/tcpip/buffer"
)

func (x *udpPacket) beforeSave() {}
func (x *udpPacket) save(m state.Map) {
	x.beforeSave()
	var data buffer.VectorisedView = x.saveData()
	m.SaveValue("data", data)
	m.Save("udpPacketEntry", &x.udpPacketEntry)
	m.Save("senderAddress", &x.senderAddress)
	m.Save("packetInfo", &x.packetInfo)
	m.Save("timestamp", &x.timestamp)
	m.Save("tos", &x.tos)
}

func (x *udpPacket) afterLoad() {}
func (x *udpPacket) load(m state.Map) {
	m.Load("udpPacketEntry", &x.udpPacketEntry)
	m.Load("senderAddress", &x.senderAddress)
	m.Load("packetInfo", &x.packetInfo)
	m.Load("timestamp", &x.timestamp)
	m.Load("tos", &x.tos)
	m.LoadValue("data", new(buffer.VectorisedView), func(y interface{}) { x.loadData(y.(buffer.VectorisedView)) })
}

func (x *endpoint) save(m state.Map) {
	x.beforeSave()
	var rcvBufSizeMax int = x.saveRcvBufSizeMax()
	m.SaveValue("rcvBufSizeMax", rcvBufSizeMax)
	m.Save("TransportEndpointInfo", &x.TransportEndpointInfo)
	m.Save("waiterQueue", &x.waiterQueue)
	m.Save("uniqueID", &x.uniqueID)
	m.Save("rcvReady", &x.rcvReady)
	m.Save("rcvList", &x.rcvList)
	m.Save("rcvBufSize", &x.rcvBufSize)
	m.Save("rcvClosed", &x.rcvClosed)
	m.Save("sndBufSize", &x.sndBufSize)
	m.Save("state", &x.state)
	m.Save("dstPort", &x.dstPort)
	m.Save("v6only", &x.v6only)
	m.Save("ttl", &x.ttl)
	m.Save("multicastTTL", &x.multicastTTL)
	m.Save("multicastAddr", &x.multicastAddr)
	m.Save("multicastNICID", &x.multicastNICID)
	m.Save("multicastLoop", &x.multicastLoop)
	m.Save("reusePort", &x.reusePort)
	m.Save("bindToDevice", &x.bindToDevice)
	m.Save("broadcast", &x.broadcast)
	m.Save("boundBindToDevice", &x.boundBindToDevice)
	m.Save("boundPortFlags", &x.boundPortFlags)
	m.Save("sendTOS", &x.sendTOS)
	m.Save("receiveTOS", &x.receiveTOS)
	m.Save("receiveIPPacketInfo", &x.receiveIPPacketInfo)
	m.Save("shutdownFlags", &x.shutdownFlags)
	m.Save("multicastMemberships", &x.multicastMemberships)
	m.Save("effectiveNetProtos", &x.effectiveNetProtos)
}

func (x *endpoint) load(m state.Map) {
	m.Load("TransportEndpointInfo", &x.TransportEndpointInfo)
	m.Load("waiterQueue", &x.waiterQueue)
	m.Load("uniqueID", &x.uniqueID)
	m.Load("rcvReady", &x.rcvReady)
	m.Load("rcvList", &x.rcvList)
	m.Load("rcvBufSize", &x.rcvBufSize)
	m.Load("rcvClosed", &x.rcvClosed)
	m.Load("sndBufSize", &x.sndBufSize)
	m.Load("state", &x.state)
	m.Load("dstPort", &x.dstPort)
	m.Load("v6only", &x.v6only)
	m.Load("ttl", &x.ttl)
	m.Load("multicastTTL", &x.multicastTTL)
	m.Load("multicastAddr", &x.multicastAddr)
	m.Load("multicastNICID", &x.multicastNICID)
	m.Load("multicastLoop", &x.multicastLoop)
	m.Load("reusePort", &x.reusePort)
	m.Load("bindToDevice", &x.bindToDevice)
	m.Load("broadcast", &x.broadcast)
	m.Load("boundBindToDevice", &x.boundBindToDevice)
	m.Load("boundPortFlags", &x.boundPortFlags)
	m.Load("sendTOS", &x.sendTOS)
	m.Load("receiveTOS", &x.receiveTOS)
	m.Load("receiveIPPacketInfo", &x.receiveIPPacketInfo)
	m.Load("shutdownFlags", &x.shutdownFlags)
	m.Load("multicastMemberships", &x.multicastMemberships)
	m.Load("effectiveNetProtos", &x.effectiveNetProtos)
	m.LoadValue("rcvBufSizeMax", new(int), func(y interface{}) { x.loadRcvBufSizeMax(y.(int)) })
	m.AfterLoad(x.afterLoad)
}

func (x *multicastMembership) beforeSave() {}
func (x *multicastMembership) save(m state.Map) {
	x.beforeSave()
	m.Save("nicID", &x.nicID)
	m.Save("multicastAddr", &x.multicastAddr)
}

func (x *multicastMembership) afterLoad() {}
func (x *multicastMembership) load(m state.Map) {
	m.Load("nicID", &x.nicID)
	m.Load("multicastAddr", &x.multicastAddr)
}

func (x *udpPacketList) beforeSave() {}
func (x *udpPacketList) save(m state.Map) {
	x.beforeSave()
	m.Save("head", &x.head)
	m.Save("tail", &x.tail)
}

func (x *udpPacketList) afterLoad() {}
func (x *udpPacketList) load(m state.Map) {
	m.Load("head", &x.head)
	m.Load("tail", &x.tail)
}

func (x *udpPacketEntry) beforeSave() {}
func (x *udpPacketEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("next", &x.next)
	m.Save("prev", &x.prev)
}

func (x *udpPacketEntry) afterLoad() {}
func (x *udpPacketEntry) load(m state.Map) {
	m.Load("next", &x.next)
	m.Load("prev", &x.prev)
}

func init() {
	state.Register("pkg/tcpip/transport/udp.udpPacket", (*udpPacket)(nil), state.Fns{Save: (*udpPacket).save, Load: (*udpPacket).load})
	state.Register("pkg/tcpip/transport/udp.endpoint", (*endpoint)(nil), state.Fns{Save: (*endpoint).save, Load: (*endpoint).load})
	state.Register("pkg/tcpip/transport/udp.multicastMembership", (*multicastMembership)(nil), state.Fns{Save: (*multicastMembership).save, Load: (*multicastMembership).load})
	state.Register("pkg/tcpip/transport/udp.udpPacketList", (*udpPacketList)(nil), state.Fns{Save: (*udpPacketList).save, Load: (*udpPacketList).load})
	state.Register("pkg/tcpip/transport/udp.udpPacketEntry", (*udpPacketEntry)(nil), state.Fns{Save: (*udpPacketEntry).save, Load: (*udpPacketEntry).load})
}
