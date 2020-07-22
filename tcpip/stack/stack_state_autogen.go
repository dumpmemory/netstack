// automatically generated by stateify.

package stack

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *tuple) StateTypeName() string {
	return "pkg/tcpip/stack.tuple"
}

func (x *tuple) StateFields() []string {
	return []string{
		"tupleEntry",
		"tupleID",
		"conn",
		"direction",
	}
}

func (x *tuple) beforeSave() {}

func (x *tuple) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.tupleEntry)
	m.Save(1, &x.tupleID)
	m.Save(2, &x.conn)
	m.Save(3, &x.direction)
}

func (x *tuple) afterLoad() {}

func (x *tuple) StateLoad(m state.Source) {
	m.Load(0, &x.tupleEntry)
	m.Load(1, &x.tupleID)
	m.Load(2, &x.conn)
	m.Load(3, &x.direction)
}

func (x *tupleID) StateTypeName() string {
	return "pkg/tcpip/stack.tupleID"
}

func (x *tupleID) StateFields() []string {
	return []string{
		"srcAddr",
		"srcPort",
		"dstAddr",
		"dstPort",
		"transProto",
		"netProto",
	}
}

func (x *tupleID) beforeSave() {}

func (x *tupleID) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.srcAddr)
	m.Save(1, &x.srcPort)
	m.Save(2, &x.dstAddr)
	m.Save(3, &x.dstPort)
	m.Save(4, &x.transProto)
	m.Save(5, &x.netProto)
}

func (x *tupleID) afterLoad() {}

func (x *tupleID) StateLoad(m state.Source) {
	m.Load(0, &x.srcAddr)
	m.Load(1, &x.srcPort)
	m.Load(2, &x.dstAddr)
	m.Load(3, &x.dstPort)
	m.Load(4, &x.transProto)
	m.Load(5, &x.netProto)
}

func (x *conn) StateTypeName() string {
	return "pkg/tcpip/stack.conn"
}

func (x *conn) StateFields() []string {
	return []string{
		"original",
		"reply",
		"manip",
		"tcbHook",
		"tcb",
		"lastUsed",
	}
}

func (x *conn) beforeSave() {}

func (x *conn) StateSave(m state.Sink) {
	x.beforeSave()
	var lastUsed unixTime = x.saveLastUsed()
	m.SaveValue(5, lastUsed)
	m.Save(0, &x.original)
	m.Save(1, &x.reply)
	m.Save(2, &x.manip)
	m.Save(3, &x.tcbHook)
	m.Save(4, &x.tcb)
}

func (x *conn) afterLoad() {}

func (x *conn) StateLoad(m state.Source) {
	m.Load(0, &x.original)
	m.Load(1, &x.reply)
	m.Load(2, &x.manip)
	m.Load(3, &x.tcbHook)
	m.Load(4, &x.tcb)
	m.LoadValue(5, new(unixTime), func(y interface{}) { x.loadLastUsed(y.(unixTime)) })
}

func (x *ConnTrack) StateTypeName() string {
	return "pkg/tcpip/stack.ConnTrack"
}

func (x *ConnTrack) StateFields() []string {
	return []string{
		"seed",
		"buckets",
	}
}

func (x *ConnTrack) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.seed)
	m.Save(1, &x.buckets)
}

func (x *ConnTrack) afterLoad() {}

func (x *ConnTrack) StateLoad(m state.Source) {
	m.Load(0, &x.seed)
	m.Load(1, &x.buckets)
}

func (x *bucket) StateTypeName() string {
	return "pkg/tcpip/stack.bucket"
}

func (x *bucket) StateFields() []string {
	return []string{
		"tuples",
	}
}

func (x *bucket) beforeSave() {}

func (x *bucket) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.tuples)
}

func (x *bucket) afterLoad() {}

func (x *bucket) StateLoad(m state.Source) {
	m.Load(0, &x.tuples)
}

func (x *unixTime) StateTypeName() string {
	return "pkg/tcpip/stack.unixTime"
}

func (x *unixTime) StateFields() []string {
	return []string{
		"second",
		"nano",
	}
}

func (x *unixTime) beforeSave() {}

func (x *unixTime) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.second)
	m.Save(1, &x.nano)
}

func (x *unixTime) afterLoad() {}

func (x *unixTime) StateLoad(m state.Source) {
	m.Load(0, &x.second)
	m.Load(1, &x.nano)
}

func (x *IPTables) StateTypeName() string {
	return "pkg/tcpip/stack.IPTables"
}

func (x *IPTables) StateFields() []string {
	return []string{
		"mu",
		"tables",
		"priorities",
		"modified",
		"connections",
		"reaperDone",
	}
}

func (x *IPTables) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.mu)
	m.Save(1, &x.tables)
	m.Save(2, &x.priorities)
	m.Save(3, &x.modified)
	m.Save(4, &x.connections)
	m.Save(5, &x.reaperDone)
}

func (x *IPTables) StateLoad(m state.Source) {
	m.Load(0, &x.mu)
	m.Load(1, &x.tables)
	m.Load(2, &x.priorities)
	m.Load(3, &x.modified)
	m.Load(4, &x.connections)
	m.Load(5, &x.reaperDone)
	m.AfterLoad(x.afterLoad)
}

func (x *Table) StateTypeName() string {
	return "pkg/tcpip/stack.Table"
}

func (x *Table) StateFields() []string {
	return []string{
		"Rules",
		"BuiltinChains",
		"Underflows",
	}
}

func (x *Table) beforeSave() {}

func (x *Table) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Rules)
	m.Save(1, &x.BuiltinChains)
	m.Save(2, &x.Underflows)
}

func (x *Table) afterLoad() {}

func (x *Table) StateLoad(m state.Source) {
	m.Load(0, &x.Rules)
	m.Load(1, &x.BuiltinChains)
	m.Load(2, &x.Underflows)
}

func (x *Rule) StateTypeName() string {
	return "pkg/tcpip/stack.Rule"
}

func (x *Rule) StateFields() []string {
	return []string{
		"Filter",
		"Matchers",
		"Target",
	}
}

func (x *Rule) beforeSave() {}

func (x *Rule) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Filter)
	m.Save(1, &x.Matchers)
	m.Save(2, &x.Target)
}

func (x *Rule) afterLoad() {}

func (x *Rule) StateLoad(m state.Source) {
	m.Load(0, &x.Filter)
	m.Load(1, &x.Matchers)
	m.Load(2, &x.Target)
}

func (x *IPHeaderFilter) StateTypeName() string {
	return "pkg/tcpip/stack.IPHeaderFilter"
}

func (x *IPHeaderFilter) StateFields() []string {
	return []string{
		"Protocol",
		"Dst",
		"DstMask",
		"DstInvert",
		"Src",
		"SrcMask",
		"SrcInvert",
		"OutputInterface",
		"OutputInterfaceMask",
		"OutputInterfaceInvert",
	}
}

func (x *IPHeaderFilter) beforeSave() {}

func (x *IPHeaderFilter) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Protocol)
	m.Save(1, &x.Dst)
	m.Save(2, &x.DstMask)
	m.Save(3, &x.DstInvert)
	m.Save(4, &x.Src)
	m.Save(5, &x.SrcMask)
	m.Save(6, &x.SrcInvert)
	m.Save(7, &x.OutputInterface)
	m.Save(8, &x.OutputInterfaceMask)
	m.Save(9, &x.OutputInterfaceInvert)
}

func (x *IPHeaderFilter) afterLoad() {}

func (x *IPHeaderFilter) StateLoad(m state.Source) {
	m.Load(0, &x.Protocol)
	m.Load(1, &x.Dst)
	m.Load(2, &x.DstMask)
	m.Load(3, &x.DstInvert)
	m.Load(4, &x.Src)
	m.Load(5, &x.SrcMask)
	m.Load(6, &x.SrcInvert)
	m.Load(7, &x.OutputInterface)
	m.Load(8, &x.OutputInterfaceMask)
	m.Load(9, &x.OutputInterfaceInvert)
}

func (x *linkAddrEntryList) StateTypeName() string {
	return "pkg/tcpip/stack.linkAddrEntryList"
}

func (x *linkAddrEntryList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (x *linkAddrEntryList) beforeSave() {}

func (x *linkAddrEntryList) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.head)
	m.Save(1, &x.tail)
}

func (x *linkAddrEntryList) afterLoad() {}

func (x *linkAddrEntryList) StateLoad(m state.Source) {
	m.Load(0, &x.head)
	m.Load(1, &x.tail)
}

func (x *linkAddrEntryEntry) StateTypeName() string {
	return "pkg/tcpip/stack.linkAddrEntryEntry"
}

func (x *linkAddrEntryEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (x *linkAddrEntryEntry) beforeSave() {}

func (x *linkAddrEntryEntry) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.next)
	m.Save(1, &x.prev)
}

func (x *linkAddrEntryEntry) afterLoad() {}

func (x *linkAddrEntryEntry) StateLoad(m state.Source) {
	m.Load(0, &x.next)
	m.Load(1, &x.prev)
}

func (x *PacketBufferList) StateTypeName() string {
	return "pkg/tcpip/stack.PacketBufferList"
}

func (x *PacketBufferList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (x *PacketBufferList) beforeSave() {}

func (x *PacketBufferList) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.head)
	m.Save(1, &x.tail)
}

func (x *PacketBufferList) afterLoad() {}

func (x *PacketBufferList) StateLoad(m state.Source) {
	m.Load(0, &x.head)
	m.Load(1, &x.tail)
}

func (x *PacketBufferEntry) StateTypeName() string {
	return "pkg/tcpip/stack.PacketBufferEntry"
}

func (x *PacketBufferEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (x *PacketBufferEntry) beforeSave() {}

func (x *PacketBufferEntry) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.next)
	m.Save(1, &x.prev)
}

func (x *PacketBufferEntry) afterLoad() {}

func (x *PacketBufferEntry) StateLoad(m state.Source) {
	m.Load(0, &x.next)
	m.Load(1, &x.prev)
}

func (x *TransportEndpointID) StateTypeName() string {
	return "pkg/tcpip/stack.TransportEndpointID"
}

func (x *TransportEndpointID) StateFields() []string {
	return []string{
		"LocalPort",
		"LocalAddress",
		"RemotePort",
		"RemoteAddress",
	}
}

func (x *TransportEndpointID) beforeSave() {}

func (x *TransportEndpointID) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.LocalPort)
	m.Save(1, &x.LocalAddress)
	m.Save(2, &x.RemotePort)
	m.Save(3, &x.RemoteAddress)
}

func (x *TransportEndpointID) afterLoad() {}

func (x *TransportEndpointID) StateLoad(m state.Source) {
	m.Load(0, &x.LocalPort)
	m.Load(1, &x.LocalAddress)
	m.Load(2, &x.RemotePort)
	m.Load(3, &x.RemoteAddress)
}

func (x *GSOType) StateTypeName() string {
	return "pkg/tcpip/stack.GSOType"
}

func (x *GSOType) StateFields() []string {
	return nil
}

func (x *GSO) StateTypeName() string {
	return "pkg/tcpip/stack.GSO"
}

func (x *GSO) StateFields() []string {
	return []string{
		"Type",
		"NeedsCsum",
		"CsumOffset",
		"MSS",
		"L3HdrLen",
		"MaxSize",
	}
}

func (x *GSO) beforeSave() {}

func (x *GSO) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Type)
	m.Save(1, &x.NeedsCsum)
	m.Save(2, &x.CsumOffset)
	m.Save(3, &x.MSS)
	m.Save(4, &x.L3HdrLen)
	m.Save(5, &x.MaxSize)
}

func (x *GSO) afterLoad() {}

func (x *GSO) StateLoad(m state.Source) {
	m.Load(0, &x.Type)
	m.Load(1, &x.NeedsCsum)
	m.Load(2, &x.CsumOffset)
	m.Load(3, &x.MSS)
	m.Load(4, &x.L3HdrLen)
	m.Load(5, &x.MaxSize)
}

func (x *TransportEndpointInfo) StateTypeName() string {
	return "pkg/tcpip/stack.TransportEndpointInfo"
}

func (x *TransportEndpointInfo) StateFields() []string {
	return []string{
		"NetProto",
		"TransProto",
		"ID",
		"BindNICID",
		"BindAddr",
		"RegisterNICID",
	}
}

func (x *TransportEndpointInfo) beforeSave() {}

func (x *TransportEndpointInfo) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.NetProto)
	m.Save(1, &x.TransProto)
	m.Save(2, &x.ID)
	m.Save(3, &x.BindNICID)
	m.Save(4, &x.BindAddr)
	m.Save(5, &x.RegisterNICID)
}

func (x *TransportEndpointInfo) afterLoad() {}

func (x *TransportEndpointInfo) StateLoad(m state.Source) {
	m.Load(0, &x.NetProto)
	m.Load(1, &x.TransProto)
	m.Load(2, &x.ID)
	m.Load(3, &x.BindNICID)
	m.Load(4, &x.BindAddr)
	m.Load(5, &x.RegisterNICID)
}

func (x *multiPortEndpoint) StateTypeName() string {
	return "pkg/tcpip/stack.multiPortEndpoint"
}

func (x *multiPortEndpoint) StateFields() []string {
	return []string{
		"demux",
		"netProto",
		"transProto",
		"endpoints",
		"flags",
	}
}

func (x *multiPortEndpoint) beforeSave() {}

func (x *multiPortEndpoint) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.demux)
	m.Save(1, &x.netProto)
	m.Save(2, &x.transProto)
	m.Save(3, &x.endpoints)
	m.Save(4, &x.flags)
}

func (x *multiPortEndpoint) afterLoad() {}

func (x *multiPortEndpoint) StateLoad(m state.Source) {
	m.Load(0, &x.demux)
	m.Load(1, &x.netProto)
	m.Load(2, &x.transProto)
	m.Load(3, &x.endpoints)
	m.Load(4, &x.flags)
}

func (x *tupleList) StateTypeName() string {
	return "pkg/tcpip/stack.tupleList"
}

func (x *tupleList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (x *tupleList) beforeSave() {}

func (x *tupleList) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.head)
	m.Save(1, &x.tail)
}

func (x *tupleList) afterLoad() {}

func (x *tupleList) StateLoad(m state.Source) {
	m.Load(0, &x.head)
	m.Load(1, &x.tail)
}

func (x *tupleEntry) StateTypeName() string {
	return "pkg/tcpip/stack.tupleEntry"
}

func (x *tupleEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (x *tupleEntry) beforeSave() {}

func (x *tupleEntry) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.next)
	m.Save(1, &x.prev)
}

func (x *tupleEntry) afterLoad() {}

func (x *tupleEntry) StateLoad(m state.Source) {
	m.Load(0, &x.next)
	m.Load(1, &x.prev)
}

func init() {
	state.Register((*tuple)(nil))
	state.Register((*tupleID)(nil))
	state.Register((*conn)(nil))
	state.Register((*ConnTrack)(nil))
	state.Register((*bucket)(nil))
	state.Register((*unixTime)(nil))
	state.Register((*IPTables)(nil))
	state.Register((*Table)(nil))
	state.Register((*Rule)(nil))
	state.Register((*IPHeaderFilter)(nil))
	state.Register((*linkAddrEntryList)(nil))
	state.Register((*linkAddrEntryEntry)(nil))
	state.Register((*PacketBufferList)(nil))
	state.Register((*PacketBufferEntry)(nil))
	state.Register((*TransportEndpointID)(nil))
	state.Register((*GSOType)(nil))
	state.Register((*GSO)(nil))
	state.Register((*TransportEndpointInfo)(nil))
	state.Register((*multiPortEndpoint)(nil))
	state.Register((*tupleList)(nil))
	state.Register((*tupleEntry)(nil))
}
