// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package mock_nats is a generated GoMock package.
package mock_nats

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	nats "github.com/nats-io/nats.go"
)

// MockJetStreamContext is a mock of JetStreamContext interface.
type MockJetStreamContext struct {
	ctrl     *gomock.Controller
	recorder *MockJetStreamContextMockRecorder
}

// MockJetStreamContextMockRecorder is the mock recorder for MockJetStreamContext.
type MockJetStreamContextMockRecorder struct {
	mock *MockJetStreamContext
}

// NewMockJetStreamContext creates a new mock instance.
func NewMockJetStreamContext(ctrl *gomock.Controller) *MockJetStreamContext {
	mock := &MockJetStreamContext{ctrl: ctrl}
	mock.recorder = &MockJetStreamContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJetStreamContext) EXPECT() *MockJetStreamContextMockRecorder {
	return m.recorder
}

// AccountInfo mocks base method.
func (m *MockJetStreamContext) AccountInfo(opts ...nats.JSOpt) (*nats.AccountInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AccountInfo", varargs...)
	ret0, _ := ret[0].(*nats.AccountInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccountInfo indicates an expected call of AccountInfo.
func (mr *MockJetStreamContextMockRecorder) AccountInfo(opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountInfo", reflect.TypeOf((*MockJetStreamContext)(nil).AccountInfo), opts...)
}

// AddConsumer mocks base method.
func (m *MockJetStreamContext) AddConsumer(stream string, cfg *nats.ConsumerConfig, opts ...nats.JSOpt) (*nats.ConsumerInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{stream, cfg}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddConsumer", varargs...)
	ret0, _ := ret[0].(*nats.ConsumerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddConsumer indicates an expected call of AddConsumer.
func (mr *MockJetStreamContextMockRecorder) AddConsumer(stream, cfg interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{stream, cfg}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddConsumer", reflect.TypeOf((*MockJetStreamContext)(nil).AddConsumer), varargs...)
}

// AddStream mocks base method.
func (m *MockJetStreamContext) AddStream(cfg *nats.StreamConfig, opts ...nats.JSOpt) (*nats.StreamInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{cfg}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddStream", varargs...)
	ret0, _ := ret[0].(*nats.StreamInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddStream indicates an expected call of AddStream.
func (mr *MockJetStreamContextMockRecorder) AddStream(cfg interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{cfg}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStream", reflect.TypeOf((*MockJetStreamContext)(nil).AddStream), varargs...)
}

// ChanQueueSubscribe mocks base method.
func (m *MockJetStreamContext) ChanQueueSubscribe(subj, queue string, ch chan *nats.Msg, opts ...nats.SubOpt) (*nats.Subscription, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{subj, queue, ch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChanQueueSubscribe", varargs...)
	ret0, _ := ret[0].(*nats.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChanQueueSubscribe indicates an expected call of ChanQueueSubscribe.
func (mr *MockJetStreamContextMockRecorder) ChanQueueSubscribe(subj, queue, ch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{subj, queue, ch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChanQueueSubscribe", reflect.TypeOf((*MockJetStreamContext)(nil).ChanQueueSubscribe), varargs...)
}

// ChanSubscribe mocks base method.
func (m *MockJetStreamContext) ChanSubscribe(subj string, ch chan *nats.Msg, opts ...nats.SubOpt) (*nats.Subscription, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{subj, ch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChanSubscribe", varargs...)
	ret0, _ := ret[0].(*nats.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChanSubscribe indicates an expected call of ChanSubscribe.
func (mr *MockJetStreamContextMockRecorder) ChanSubscribe(subj, ch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{subj, ch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChanSubscribe", reflect.TypeOf((*MockJetStreamContext)(nil).ChanSubscribe), varargs...)
}

// ConsumerInfo mocks base method.
func (m *MockJetStreamContext) ConsumerInfo(stream, name string, opts ...nats.JSOpt) (*nats.ConsumerInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{stream, name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConsumerInfo", varargs...)
	ret0, _ := ret[0].(*nats.ConsumerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConsumerInfo indicates an expected call of ConsumerInfo.
func (mr *MockJetStreamContextMockRecorder) ConsumerInfo(stream, name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{stream, name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumerInfo", reflect.TypeOf((*MockJetStreamContext)(nil).ConsumerInfo), varargs...)
}

// ConsumerNames mocks base method.
func (m *MockJetStreamContext) ConsumerNames(stream string, opts ...nats.JSOpt) <-chan string {
	m.ctrl.T.Helper()
	varargs := []interface{}{stream}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConsumerNames", varargs...)
	ret0, _ := ret[0].(<-chan string)
	return ret0
}

// ConsumerNames indicates an expected call of ConsumerNames.
func (mr *MockJetStreamContextMockRecorder) ConsumerNames(stream interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{stream}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumerNames", reflect.TypeOf((*MockJetStreamContext)(nil).ConsumerNames), varargs...)
}

// ConsumersInfo mocks base method.
func (m *MockJetStreamContext) ConsumersInfo(stream string, opts ...nats.JSOpt) <-chan *nats.ConsumerInfo {
	m.ctrl.T.Helper()
	varargs := []interface{}{stream}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConsumersInfo", varargs...)
	ret0, _ := ret[0].(<-chan *nats.ConsumerInfo)
	return ret0
}

// ConsumersInfo indicates an expected call of ConsumersInfo.
func (mr *MockJetStreamContextMockRecorder) ConsumersInfo(stream interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{stream}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumersInfo", reflect.TypeOf((*MockJetStreamContext)(nil).ConsumersInfo), varargs...)
}

// CreateKeyValue mocks base method.
func (m *MockJetStreamContext) CreateKeyValue(cfg *nats.KeyValueConfig) (nats.KeyValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKeyValue", cfg)
	ret0, _ := ret[0].(nats.KeyValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateKeyValue indicates an expected call of CreateKeyValue.
func (mr *MockJetStreamContextMockRecorder) CreateKeyValue(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKeyValue", reflect.TypeOf((*MockJetStreamContext)(nil).CreateKeyValue), cfg)
}

// CreateObjectStore mocks base method.
func (m *MockJetStreamContext) CreateObjectStore(cfg *nats.ObjectStoreConfig) (nats.ObjectStore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateObjectStore", cfg)
	ret0, _ := ret[0].(nats.ObjectStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateObjectStore indicates an expected call of CreateObjectStore.
func (mr *MockJetStreamContextMockRecorder) CreateObjectStore(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateObjectStore", reflect.TypeOf((*MockJetStreamContext)(nil).CreateObjectStore), cfg)
}

// DeleteConsumer mocks base method.
func (m *MockJetStreamContext) DeleteConsumer(stream, consumer string, opts ...nats.JSOpt) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{stream, consumer}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteConsumer", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteConsumer indicates an expected call of DeleteConsumer.
func (mr *MockJetStreamContextMockRecorder) DeleteConsumer(stream, consumer interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{stream, consumer}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConsumer", reflect.TypeOf((*MockJetStreamContext)(nil).DeleteConsumer), varargs...)
}

// DeleteKeyValue mocks base method.
func (m *MockJetStreamContext) DeleteKeyValue(bucket string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKeyValue", bucket)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteKeyValue indicates an expected call of DeleteKeyValue.
func (mr *MockJetStreamContextMockRecorder) DeleteKeyValue(bucket interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKeyValue", reflect.TypeOf((*MockJetStreamContext)(nil).DeleteKeyValue), bucket)
}

// DeleteMsg mocks base method.
func (m *MockJetStreamContext) DeleteMsg(name string, seq uint64, opts ...nats.JSOpt) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, seq}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMsg", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMsg indicates an expected call of DeleteMsg.
func (mr *MockJetStreamContextMockRecorder) DeleteMsg(name, seq interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, seq}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMsg", reflect.TypeOf((*MockJetStreamContext)(nil).DeleteMsg), varargs...)
}

// DeleteObjectStore mocks base method.
func (m *MockJetStreamContext) DeleteObjectStore(bucket string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObjectStore", bucket)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObjectStore indicates an expected call of DeleteObjectStore.
func (mr *MockJetStreamContextMockRecorder) DeleteObjectStore(bucket interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjectStore", reflect.TypeOf((*MockJetStreamContext)(nil).DeleteObjectStore), bucket)
}

// DeleteStream mocks base method.
func (m *MockJetStreamContext) DeleteStream(name string, opts ...nats.JSOpt) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteStream", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStream indicates an expected call of DeleteStream.
func (mr *MockJetStreamContextMockRecorder) DeleteStream(name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStream", reflect.TypeOf((*MockJetStreamContext)(nil).DeleteStream), varargs...)
}

// GetMsg mocks base method.
func (m *MockJetStreamContext) GetMsg(name string, seq uint64, opts ...nats.JSOpt) (*nats.RawStreamMsg, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{name, seq}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMsg", varargs...)
	ret0, _ := ret[0].(*nats.RawStreamMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMsg indicates an expected call of GetMsg.
func (mr *MockJetStreamContextMockRecorder) GetMsg(name, seq interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name, seq}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMsg", reflect.TypeOf((*MockJetStreamContext)(nil).GetMsg), varargs...)
}

// KeyValue mocks base method.
func (m *MockJetStreamContext) KeyValue(bucket string) (nats.KeyValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyValue", bucket)
	ret0, _ := ret[0].(nats.KeyValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KeyValue indicates an expected call of KeyValue.
func (mr *MockJetStreamContextMockRecorder) KeyValue(bucket interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyValue", reflect.TypeOf((*MockJetStreamContext)(nil).KeyValue), bucket)
}

// ObjectStore mocks base method.
func (m *MockJetStreamContext) ObjectStore(bucket string) (nats.ObjectStore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectStore", bucket)
	ret0, _ := ret[0].(nats.ObjectStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectStore indicates an expected call of ObjectStore.
func (mr *MockJetStreamContextMockRecorder) ObjectStore(bucket interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStore", reflect.TypeOf((*MockJetStreamContext)(nil).ObjectStore), bucket)
}

// Publish mocks base method.
func (m *MockJetStreamContext) Publish(subj string, data []byte, opts ...nats.PubOpt) (*nats.PubAck, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{subj, data}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Publish", varargs...)
	ret0, _ := ret[0].(*nats.PubAck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish.
func (mr *MockJetStreamContextMockRecorder) Publish(subj, data interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{subj, data}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockJetStreamContext)(nil).Publish), varargs...)
}

// PublishAsync mocks base method.
func (m *MockJetStreamContext) PublishAsync(subj string, data []byte, opts ...nats.PubOpt) (nats.PubAckFuture, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{subj, data}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PublishAsync", varargs...)
	ret0, _ := ret[0].(nats.PubAckFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishAsync indicates an expected call of PublishAsync.
func (mr *MockJetStreamContextMockRecorder) PublishAsync(subj, data interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{subj, data}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishAsync", reflect.TypeOf((*MockJetStreamContext)(nil).PublishAsync), varargs...)
}

// PublishAsyncComplete mocks base method.
func (m *MockJetStreamContext) PublishAsyncComplete() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishAsyncComplete")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// PublishAsyncComplete indicates an expected call of PublishAsyncComplete.
func (mr *MockJetStreamContextMockRecorder) PublishAsyncComplete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishAsyncComplete", reflect.TypeOf((*MockJetStreamContext)(nil).PublishAsyncComplete))
}

// PublishAsyncPending mocks base method.
func (m *MockJetStreamContext) PublishAsyncPending() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishAsyncPending")
	ret0, _ := ret[0].(int)
	return ret0
}

// PublishAsyncPending indicates an expected call of PublishAsyncPending.
func (mr *MockJetStreamContextMockRecorder) PublishAsyncPending() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishAsyncPending", reflect.TypeOf((*MockJetStreamContext)(nil).PublishAsyncPending))
}

// PublishMsg mocks base method.
func (m_2 *MockJetStreamContext) PublishMsg(m *nats.Msg, opts ...nats.PubOpt) (*nats.PubAck, error) {
	m_2.ctrl.T.Helper()
	varargs := []interface{}{m}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m_2.ctrl.Call(m_2, "PublishMsg", varargs...)
	ret0, _ := ret[0].(*nats.PubAck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishMsg indicates an expected call of PublishMsg.
func (mr *MockJetStreamContextMockRecorder) PublishMsg(m interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{m}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishMsg", reflect.TypeOf((*MockJetStreamContext)(nil).PublishMsg), varargs...)
}

// PublishMsgAsync mocks base method.
func (m_2 *MockJetStreamContext) PublishMsgAsync(m *nats.Msg, opts ...nats.PubOpt) (nats.PubAckFuture, error) {
	m_2.ctrl.T.Helper()
	varargs := []interface{}{m}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m_2.ctrl.Call(m_2, "PublishMsgAsync", varargs...)
	ret0, _ := ret[0].(nats.PubAckFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishMsgAsync indicates an expected call of PublishMsgAsync.
func (mr *MockJetStreamContextMockRecorder) PublishMsgAsync(m interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{m}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishMsgAsync", reflect.TypeOf((*MockJetStreamContext)(nil).PublishMsgAsync), varargs...)
}

// PullSubscribe mocks base method.
func (m *MockJetStreamContext) PullSubscribe(subj, durable string, opts ...nats.SubOpt) (*nats.Subscription, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{subj, durable}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PullSubscribe", varargs...)
	ret0, _ := ret[0].(*nats.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PullSubscribe indicates an expected call of PullSubscribe.
func (mr *MockJetStreamContextMockRecorder) PullSubscribe(subj, durable interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{subj, durable}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullSubscribe", reflect.TypeOf((*MockJetStreamContext)(nil).PullSubscribe), varargs...)
}

// PurgeStream mocks base method.
func (m *MockJetStreamContext) PurgeStream(name string, opts ...nats.JSOpt) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PurgeStream", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PurgeStream indicates an expected call of PurgeStream.
func (mr *MockJetStreamContextMockRecorder) PurgeStream(name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurgeStream", reflect.TypeOf((*MockJetStreamContext)(nil).PurgeStream), varargs...)
}

// QueueSubscribe mocks base method.
func (m *MockJetStreamContext) QueueSubscribe(subj, queue string, cb nats.MsgHandler, opts ...nats.SubOpt) (*nats.Subscription, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{subj, queue, cb}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueueSubscribe", varargs...)
	ret0, _ := ret[0].(*nats.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueSubscribe indicates an expected call of QueueSubscribe.
func (mr *MockJetStreamContextMockRecorder) QueueSubscribe(subj, queue, cb interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{subj, queue, cb}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueSubscribe", reflect.TypeOf((*MockJetStreamContext)(nil).QueueSubscribe), varargs...)
}

// QueueSubscribeSync mocks base method.
func (m *MockJetStreamContext) QueueSubscribeSync(subj, queue string, opts ...nats.SubOpt) (*nats.Subscription, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{subj, queue}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueueSubscribeSync", varargs...)
	ret0, _ := ret[0].(*nats.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueSubscribeSync indicates an expected call of QueueSubscribeSync.
func (mr *MockJetStreamContextMockRecorder) QueueSubscribeSync(subj, queue interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{subj, queue}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueSubscribeSync", reflect.TypeOf((*MockJetStreamContext)(nil).QueueSubscribeSync), varargs...)
}

// StreamInfo mocks base method.
func (m *MockJetStreamContext) StreamInfo(stream string, opts ...nats.JSOpt) (*nats.StreamInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{stream}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StreamInfo", varargs...)
	ret0, _ := ret[0].(*nats.StreamInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamInfo indicates an expected call of StreamInfo.
func (mr *MockJetStreamContextMockRecorder) StreamInfo(stream interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{stream}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamInfo", reflect.TypeOf((*MockJetStreamContext)(nil).StreamInfo), varargs...)
}

// StreamNames mocks base method.
func (m *MockJetStreamContext) StreamNames(opts ...nats.JSOpt) <-chan string {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StreamNames", varargs...)
	ret0, _ := ret[0].(<-chan string)
	return ret0
}

// StreamNames indicates an expected call of StreamNames.
func (mr *MockJetStreamContextMockRecorder) StreamNames(opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamNames", reflect.TypeOf((*MockJetStreamContext)(nil).StreamNames), opts...)
}

// StreamsInfo mocks base method.
func (m *MockJetStreamContext) StreamsInfo(opts ...nats.JSOpt) <-chan *nats.StreamInfo {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StreamsInfo", varargs...)
	ret0, _ := ret[0].(<-chan *nats.StreamInfo)
	return ret0
}

// StreamsInfo indicates an expected call of StreamsInfo.
func (mr *MockJetStreamContextMockRecorder) StreamsInfo(opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamsInfo", reflect.TypeOf((*MockJetStreamContext)(nil).StreamsInfo), opts...)
}

// Subscribe mocks base method.
func (m *MockJetStreamContext) Subscribe(subj string, cb nats.MsgHandler, opts ...nats.SubOpt) (*nats.Subscription, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{subj, cb}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(*nats.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockJetStreamContextMockRecorder) Subscribe(subj, cb interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{subj, cb}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockJetStreamContext)(nil).Subscribe), varargs...)
}

// SubscribeSync mocks base method.
func (m *MockJetStreamContext) SubscribeSync(subj string, opts ...nats.SubOpt) (*nats.Subscription, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{subj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeSync", varargs...)
	ret0, _ := ret[0].(*nats.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeSync indicates an expected call of SubscribeSync.
func (mr *MockJetStreamContextMockRecorder) SubscribeSync(subj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{subj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeSync", reflect.TypeOf((*MockJetStreamContext)(nil).SubscribeSync), varargs...)
}

// UpdateConsumer mocks base method.
func (m *MockJetStreamContext) UpdateConsumer(stream string, cfg *nats.ConsumerConfig, opts ...nats.JSOpt) (*nats.ConsumerInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{stream, cfg}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateConsumer", varargs...)
	ret0, _ := ret[0].(*nats.ConsumerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateConsumer indicates an expected call of UpdateConsumer.
func (mr *MockJetStreamContextMockRecorder) UpdateConsumer(stream, cfg interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{stream, cfg}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConsumer", reflect.TypeOf((*MockJetStreamContext)(nil).UpdateConsumer), varargs...)
}

// UpdateStream mocks base method.
func (m *MockJetStreamContext) UpdateStream(cfg *nats.StreamConfig, opts ...nats.JSOpt) (*nats.StreamInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{cfg}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateStream", varargs...)
	ret0, _ := ret[0].(*nats.StreamInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStream indicates an expected call of UpdateStream.
func (mr *MockJetStreamContextMockRecorder) UpdateStream(cfg interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{cfg}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStream", reflect.TypeOf((*MockJetStreamContext)(nil).UpdateStream), varargs...)
}