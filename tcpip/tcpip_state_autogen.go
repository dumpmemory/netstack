// automatically generated by stateify.

package tcpip

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (l *sockErrorList) StateTypeName() string {
	return "pkg/tcpip.sockErrorList"
}

func (l *sockErrorList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *sockErrorList) beforeSave() {}

func (l *sockErrorList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *sockErrorList) afterLoad() {}

func (l *sockErrorList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *sockErrorEntry) StateTypeName() string {
	return "pkg/tcpip.sockErrorEntry"
}

func (e *sockErrorEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *sockErrorEntry) beforeSave() {}

func (e *sockErrorEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *sockErrorEntry) afterLoad() {}

func (e *sockErrorEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (so *SocketOptions) StateTypeName() string {
	return "pkg/tcpip.SocketOptions"
}

func (so *SocketOptions) StateFields() []string {
	return []string{
		"handler",
		"broadcastEnabled",
		"passCredEnabled",
		"noChecksumEnabled",
		"reuseAddressEnabled",
		"reusePortEnabled",
		"keepAliveEnabled",
		"multicastLoopEnabled",
		"receiveTOSEnabled",
		"receiveTClassEnabled",
		"receivePacketInfoEnabled",
		"hdrIncludedEnabled",
		"v6OnlyEnabled",
		"quickAckEnabled",
		"delayOptionEnabled",
		"corkOptionEnabled",
		"receiveOriginalDstAddress",
		"recvErrEnabled",
		"errQueue",
		"bindToDevice",
		"sendBufferSize",
		"linger",
	}
}

func (so *SocketOptions) beforeSave() {}

func (so *SocketOptions) StateSave(stateSinkObject state.Sink) {
	so.beforeSave()
	stateSinkObject.Save(0, &so.handler)
	stateSinkObject.Save(1, &so.broadcastEnabled)
	stateSinkObject.Save(2, &so.passCredEnabled)
	stateSinkObject.Save(3, &so.noChecksumEnabled)
	stateSinkObject.Save(4, &so.reuseAddressEnabled)
	stateSinkObject.Save(5, &so.reusePortEnabled)
	stateSinkObject.Save(6, &so.keepAliveEnabled)
	stateSinkObject.Save(7, &so.multicastLoopEnabled)
	stateSinkObject.Save(8, &so.receiveTOSEnabled)
	stateSinkObject.Save(9, &so.receiveTClassEnabled)
	stateSinkObject.Save(10, &so.receivePacketInfoEnabled)
	stateSinkObject.Save(11, &so.hdrIncludedEnabled)
	stateSinkObject.Save(12, &so.v6OnlyEnabled)
	stateSinkObject.Save(13, &so.quickAckEnabled)
	stateSinkObject.Save(14, &so.delayOptionEnabled)
	stateSinkObject.Save(15, &so.corkOptionEnabled)
	stateSinkObject.Save(16, &so.receiveOriginalDstAddress)
	stateSinkObject.Save(17, &so.recvErrEnabled)
	stateSinkObject.Save(18, &so.errQueue)
	stateSinkObject.Save(19, &so.bindToDevice)
	stateSinkObject.Save(20, &so.sendBufferSize)
	stateSinkObject.Save(21, &so.linger)
}

func (so *SocketOptions) afterLoad() {}

func (so *SocketOptions) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &so.handler)
	stateSourceObject.Load(1, &so.broadcastEnabled)
	stateSourceObject.Load(2, &so.passCredEnabled)
	stateSourceObject.Load(3, &so.noChecksumEnabled)
	stateSourceObject.Load(4, &so.reuseAddressEnabled)
	stateSourceObject.Load(5, &so.reusePortEnabled)
	stateSourceObject.Load(6, &so.keepAliveEnabled)
	stateSourceObject.Load(7, &so.multicastLoopEnabled)
	stateSourceObject.Load(8, &so.receiveTOSEnabled)
	stateSourceObject.Load(9, &so.receiveTClassEnabled)
	stateSourceObject.Load(10, &so.receivePacketInfoEnabled)
	stateSourceObject.Load(11, &so.hdrIncludedEnabled)
	stateSourceObject.Load(12, &so.v6OnlyEnabled)
	stateSourceObject.Load(13, &so.quickAckEnabled)
	stateSourceObject.Load(14, &so.delayOptionEnabled)
	stateSourceObject.Load(15, &so.corkOptionEnabled)
	stateSourceObject.Load(16, &so.receiveOriginalDstAddress)
	stateSourceObject.Load(17, &so.recvErrEnabled)
	stateSourceObject.Load(18, &so.errQueue)
	stateSourceObject.Load(19, &so.bindToDevice)
	stateSourceObject.Load(20, &so.sendBufferSize)
	stateSourceObject.Load(21, &so.linger)
}

func (s *SockError) StateTypeName() string {
	return "pkg/tcpip.SockError"
}

func (s *SockError) StateFields() []string {
	return []string{
		"sockErrorEntry",
		"Err",
		"ErrOrigin",
		"ErrType",
		"ErrCode",
		"ErrInfo",
		"Payload",
		"Dst",
		"Offender",
		"NetProto",
	}
}

func (s *SockError) beforeSave() {}

func (s *SockError) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.sockErrorEntry)
	stateSinkObject.Save(1, &s.Err)
	stateSinkObject.Save(2, &s.ErrOrigin)
	stateSinkObject.Save(3, &s.ErrType)
	stateSinkObject.Save(4, &s.ErrCode)
	stateSinkObject.Save(5, &s.ErrInfo)
	stateSinkObject.Save(6, &s.Payload)
	stateSinkObject.Save(7, &s.Dst)
	stateSinkObject.Save(8, &s.Offender)
	stateSinkObject.Save(9, &s.NetProto)
}

func (s *SockError) afterLoad() {}

func (s *SockError) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.sockErrorEntry)
	stateSourceObject.Load(1, &s.Err)
	stateSourceObject.Load(2, &s.ErrOrigin)
	stateSourceObject.Load(3, &s.ErrType)
	stateSourceObject.Load(4, &s.ErrCode)
	stateSourceObject.Load(5, &s.ErrInfo)
	stateSourceObject.Load(6, &s.Payload)
	stateSourceObject.Load(7, &s.Dst)
	stateSourceObject.Load(8, &s.Offender)
	stateSourceObject.Load(9, &s.NetProto)
}

func (e *Error) StateTypeName() string {
	return "pkg/tcpip.Error"
}

func (e *Error) StateFields() []string {
	return []string{
		"msg",
		"ignoreStats",
	}
}

func (e *Error) beforeSave() {}

func (e *Error) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.msg)
	stateSinkObject.Save(1, &e.ignoreStats)
}

func (e *Error) afterLoad() {}

func (e *Error) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.msg)
	stateSourceObject.Load(1, &e.ignoreStats)
}

func (f *FullAddress) StateTypeName() string {
	return "pkg/tcpip.FullAddress"
}

func (f *FullAddress) StateFields() []string {
	return []string{
		"NIC",
		"Addr",
		"Port",
	}
}

func (f *FullAddress) beforeSave() {}

func (f *FullAddress) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.NIC)
	stateSinkObject.Save(1, &f.Addr)
	stateSinkObject.Save(2, &f.Port)
}

func (f *FullAddress) afterLoad() {}

func (f *FullAddress) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.NIC)
	stateSourceObject.Load(1, &f.Addr)
	stateSourceObject.Load(2, &f.Port)
}

func (c *ControlMessages) StateTypeName() string {
	return "pkg/tcpip.ControlMessages"
}

func (c *ControlMessages) StateFields() []string {
	return []string{
		"HasTimestamp",
		"Timestamp",
		"HasInq",
		"Inq",
		"HasTOS",
		"TOS",
		"HasTClass",
		"TClass",
		"HasIPPacketInfo",
		"PacketInfo",
		"HasOriginalDstAddress",
		"OriginalDstAddress",
		"SockErr",
	}
}

func (c *ControlMessages) beforeSave() {}

func (c *ControlMessages) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.HasTimestamp)
	stateSinkObject.Save(1, &c.Timestamp)
	stateSinkObject.Save(2, &c.HasInq)
	stateSinkObject.Save(3, &c.Inq)
	stateSinkObject.Save(4, &c.HasTOS)
	stateSinkObject.Save(5, &c.TOS)
	stateSinkObject.Save(6, &c.HasTClass)
	stateSinkObject.Save(7, &c.TClass)
	stateSinkObject.Save(8, &c.HasIPPacketInfo)
	stateSinkObject.Save(9, &c.PacketInfo)
	stateSinkObject.Save(10, &c.HasOriginalDstAddress)
	stateSinkObject.Save(11, &c.OriginalDstAddress)
	stateSinkObject.Save(12, &c.SockErr)
}

func (c *ControlMessages) afterLoad() {}

func (c *ControlMessages) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.HasTimestamp)
	stateSourceObject.Load(1, &c.Timestamp)
	stateSourceObject.Load(2, &c.HasInq)
	stateSourceObject.Load(3, &c.Inq)
	stateSourceObject.Load(4, &c.HasTOS)
	stateSourceObject.Load(5, &c.TOS)
	stateSourceObject.Load(6, &c.HasTClass)
	stateSourceObject.Load(7, &c.TClass)
	stateSourceObject.Load(8, &c.HasIPPacketInfo)
	stateSourceObject.Load(9, &c.PacketInfo)
	stateSourceObject.Load(10, &c.HasOriginalDstAddress)
	stateSourceObject.Load(11, &c.OriginalDstAddress)
	stateSourceObject.Load(12, &c.SockErr)
}

func (l *LinkPacketInfo) StateTypeName() string {
	return "pkg/tcpip.LinkPacketInfo"
}

func (l *LinkPacketInfo) StateFields() []string {
	return []string{
		"Protocol",
		"PktType",
	}
}

func (l *LinkPacketInfo) beforeSave() {}

func (l *LinkPacketInfo) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.Protocol)
	stateSinkObject.Save(1, &l.PktType)
}

func (l *LinkPacketInfo) afterLoad() {}

func (l *LinkPacketInfo) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.Protocol)
	stateSourceObject.Load(1, &l.PktType)
}

func (l *LingerOption) StateTypeName() string {
	return "pkg/tcpip.LingerOption"
}

func (l *LingerOption) StateFields() []string {
	return []string{
		"Enabled",
		"Timeout",
	}
}

func (l *LingerOption) beforeSave() {}

func (l *LingerOption) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.Enabled)
	stateSinkObject.Save(1, &l.Timeout)
}

func (l *LingerOption) afterLoad() {}

func (l *LingerOption) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.Enabled)
	stateSourceObject.Load(1, &l.Timeout)
}

func (i *IPPacketInfo) StateTypeName() string {
	return "pkg/tcpip.IPPacketInfo"
}

func (i *IPPacketInfo) StateFields() []string {
	return []string{
		"NIC",
		"LocalAddr",
		"DestinationAddr",
	}
}

func (i *IPPacketInfo) beforeSave() {}

func (i *IPPacketInfo) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.NIC)
	stateSinkObject.Save(1, &i.LocalAddr)
	stateSinkObject.Save(2, &i.DestinationAddr)
}

func (i *IPPacketInfo) afterLoad() {}

func (i *IPPacketInfo) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.NIC)
	stateSourceObject.Load(1, &i.LocalAddr)
	stateSourceObject.Load(2, &i.DestinationAddr)
}

func init() {
	state.Register((*sockErrorList)(nil))
	state.Register((*sockErrorEntry)(nil))
	state.Register((*SocketOptions)(nil))
	state.Register((*SockError)(nil))
	state.Register((*Error)(nil))
	state.Register((*FullAddress)(nil))
	state.Register((*ControlMessages)(nil))
	state.Register((*LinkPacketInfo)(nil))
	state.Register((*LingerOption)(nil))
	state.Register((*IPPacketInfo)(nil))
}
