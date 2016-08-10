// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/monitoring/v3/metric_service.proto
// DO NOT EDIT!

package google_monitoring_v3

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_api5 "google.golang.org/genproto/googleapis/api/metric"
import google_api4 "google.golang.org/genproto/googleapis/api/monitoredres"
import google_protobuf4 "github.com/golang/protobuf/ptypes/empty"
import google_rpc "google.golang.org/genproto/googleapis/rpc/status"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Controls which fields are returned by `ListTimeSeries`.
type ListTimeSeriesRequest_TimeSeriesView int32

const (
	// Returns the identity of the metric(s), the time series,
	// and the time series data.
	ListTimeSeriesRequest_FULL ListTimeSeriesRequest_TimeSeriesView = 0
	// Returns the identity of the metric and the time series resource,
	// but not the time series data.
	ListTimeSeriesRequest_HEADERS ListTimeSeriesRequest_TimeSeriesView = 1
)

var ListTimeSeriesRequest_TimeSeriesView_name = map[int32]string{
	0: "FULL",
	1: "HEADERS",
}
var ListTimeSeriesRequest_TimeSeriesView_value = map[string]int32{
	"FULL":    0,
	"HEADERS": 1,
}

func (x ListTimeSeriesRequest_TimeSeriesView) String() string {
	return proto.EnumName(ListTimeSeriesRequest_TimeSeriesView_name, int32(x))
}
func (ListTimeSeriesRequest_TimeSeriesView) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor6, []int{8, 0}
}

// The `ListMonitoredResourceDescriptors` request.
type ListMonitoredResourceDescriptorsRequest struct {
	// The project on which to execute the request. The format is
	// `"projects/{project_id_or_number}"`.
	Name string `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	// An optional [filter](/monitoring/api/v3/filters) describing
	// the descriptors to be returned.  The filter can reference
	// the descriptor's type and labels. For example, the
	// following filter returns only Google Compute Engine descriptors
	// that have an `id` label:
	//
	//     resource.type = starts_with("gce_") AND resource.label:id
	Filter string `protobuf:"bytes,2,opt,name=filter" json:"filter,omitempty"`
	// A positive number that is the maximum number of results to return.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// If this field is not empty then it must contain the `nextPageToken` value
	// returned by a previous call to this method.  Using this field causes the
	// method to return additional results from the previous method call.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListMonitoredResourceDescriptorsRequest) Reset() {
	*m = ListMonitoredResourceDescriptorsRequest{}
}
func (m *ListMonitoredResourceDescriptorsRequest) String() string { return proto.CompactTextString(m) }
func (*ListMonitoredResourceDescriptorsRequest) ProtoMessage()    {}
func (*ListMonitoredResourceDescriptorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{0}
}

// The `ListMonitoredResourcDescriptors` response.
type ListMonitoredResourceDescriptorsResponse struct {
	// The monitored resource descriptors that are available to this project
	// and that match `filter`, if present.
	ResourceDescriptors []*google_api4.MonitoredResourceDescriptor `protobuf:"bytes,1,rep,name=resource_descriptors,json=resourceDescriptors" json:"resource_descriptors,omitempty"`
	// If there are more results than have been returned, then this field is set
	// to a non-empty value.  To see the additional results,
	// use that value as `pageToken` in the next call to this method.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListMonitoredResourceDescriptorsResponse) Reset() {
	*m = ListMonitoredResourceDescriptorsResponse{}
}
func (m *ListMonitoredResourceDescriptorsResponse) String() string { return proto.CompactTextString(m) }
func (*ListMonitoredResourceDescriptorsResponse) ProtoMessage()    {}
func (*ListMonitoredResourceDescriptorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{1}
}

func (m *ListMonitoredResourceDescriptorsResponse) GetResourceDescriptors() []*google_api4.MonitoredResourceDescriptor {
	if m != nil {
		return m.ResourceDescriptors
	}
	return nil
}

// The `GetMonitoredResourceDescriptor` request.
type GetMonitoredResourceDescriptorRequest struct {
	// The monitored resource descriptor to get.  The format is
	// `"projects/{project_id_or_number}/monitoredResourceDescriptors/{resource_type}"`.
	// The `{resource_type}` is a predefined type, such as
	// `cloudsql_database`.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *GetMonitoredResourceDescriptorRequest) Reset()         { *m = GetMonitoredResourceDescriptorRequest{} }
func (m *GetMonitoredResourceDescriptorRequest) String() string { return proto.CompactTextString(m) }
func (*GetMonitoredResourceDescriptorRequest) ProtoMessage()    {}
func (*GetMonitoredResourceDescriptorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{2}
}

// The `ListMetricDescriptors` request.
type ListMetricDescriptorsRequest struct {
	// The project on which to execute the request. The format is
	// `"projects/{project_id_or_number}"`.
	Name string `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	// If this field is empty, all custom and
	// system-defined metric descriptors are returned.
	// Otherwise, the [filter](/monitoring/api/v3/filters)
	// specifies which metric descriptors are to be
	// returned. For example, the following filter matches all
	// [custom metrics](/monitoring/custom-metrics):
	//
	//     metric.type = starts_with("custom.googleapis.com/")
	Filter string `protobuf:"bytes,2,opt,name=filter" json:"filter,omitempty"`
	// A positive number that is the maximum number of results to return.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// If this field is not empty then it must contain the `nextPageToken` value
	// returned by a previous call to this method.  Using this field causes the
	// method to return additional results from the previous method call.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListMetricDescriptorsRequest) Reset()                    { *m = ListMetricDescriptorsRequest{} }
func (m *ListMetricDescriptorsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListMetricDescriptorsRequest) ProtoMessage()               {}
func (*ListMetricDescriptorsRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

// The `ListMetricDescriptors` response.
type ListMetricDescriptorsResponse struct {
	// The metric descriptors that are available to the project
	// and that match the value of `filter`, if present.
	MetricDescriptors []*google_api5.MetricDescriptor `protobuf:"bytes,1,rep,name=metric_descriptors,json=metricDescriptors" json:"metric_descriptors,omitempty"`
	// If there are more results than have been returned, then this field is set
	// to a non-empty value.  To see the additional results,
	// use that value as `pageToken` in the next call to this method.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListMetricDescriptorsResponse) Reset()                    { *m = ListMetricDescriptorsResponse{} }
func (m *ListMetricDescriptorsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListMetricDescriptorsResponse) ProtoMessage()               {}
func (*ListMetricDescriptorsResponse) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

func (m *ListMetricDescriptorsResponse) GetMetricDescriptors() []*google_api5.MetricDescriptor {
	if m != nil {
		return m.MetricDescriptors
	}
	return nil
}

// The `GetMetricDescriptor` request.
type GetMetricDescriptorRequest struct {
	// The metric descriptor on which to execute the request. The format is
	// `"projects/{project_id_or_number}/metricDescriptors/{metric_id}"`.
	// An example value of `{metric_id}` is
	// `"compute.googleapis.com/instance/disk/read_bytes_count"`.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *GetMetricDescriptorRequest) Reset()                    { *m = GetMetricDescriptorRequest{} }
func (m *GetMetricDescriptorRequest) String() string            { return proto.CompactTextString(m) }
func (*GetMetricDescriptorRequest) ProtoMessage()               {}
func (*GetMetricDescriptorRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

// The `CreateMetricDescriptor` request.
type CreateMetricDescriptorRequest struct {
	// The project on which to execute the request. The format is
	// `"projects/{project_id_or_number}"`.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// The new [custom metric](/monitoring/custom-metrics)
	// descriptor.
	MetricDescriptor *google_api5.MetricDescriptor `protobuf:"bytes,2,opt,name=metric_descriptor,json=metricDescriptor" json:"metric_descriptor,omitempty"`
}

func (m *CreateMetricDescriptorRequest) Reset()                    { *m = CreateMetricDescriptorRequest{} }
func (m *CreateMetricDescriptorRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateMetricDescriptorRequest) ProtoMessage()               {}
func (*CreateMetricDescriptorRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{6} }

func (m *CreateMetricDescriptorRequest) GetMetricDescriptor() *google_api5.MetricDescriptor {
	if m != nil {
		return m.MetricDescriptor
	}
	return nil
}

// The `DeleteMetricDescriptor` request.
type DeleteMetricDescriptorRequest struct {
	// The metric descriptor on which to execute the request. The format is
	// `"projects/{project_id_or_number}/metricDescriptors/{metric_id}"`.
	// An example of `{metric_id}` is:
	// `"custom.googleapis.com/my_test_metric"`.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteMetricDescriptorRequest) Reset()                    { *m = DeleteMetricDescriptorRequest{} }
func (m *DeleteMetricDescriptorRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteMetricDescriptorRequest) ProtoMessage()               {}
func (*DeleteMetricDescriptorRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{7} }

// The `ListTimeSeries` request.
type ListTimeSeriesRequest struct {
	// The project on which to execute the request. The format is
	// "projects/{project_id_or_number}".
	Name string `protobuf:"bytes,10,opt,name=name" json:"name,omitempty"`
	// A [monitoring filter](/monitoring/api/v3/filters) that specifies which time
	// series should be returned.  The filter must specify a single metric type,
	// and can additionally specify metric labels and other information. For
	// example:
	//
	//     metric.type = "compute.googleapis.com/instance/cpu/usage_time" AND
	//         metric.label.instance_name = "my-instance-name"
	Filter string `protobuf:"bytes,2,opt,name=filter" json:"filter,omitempty"`
	// The time interval for which results should be returned. Only time series
	// that contain data points in the specified interval are included
	// in the response.
	Interval *TimeInterval `protobuf:"bytes,4,opt,name=interval" json:"interval,omitempty"`
	// By default, the raw time series data is returned.
	// Use this field to combine multiple time series for different
	// views of the data.
	Aggregation *Aggregation `protobuf:"bytes,5,opt,name=aggregation" json:"aggregation,omitempty"`
	// Specifies the order in which the points of the time series should
	// be returned.  By default, results are not ordered.  Currently,
	// this field must be left blank.
	OrderBy string `protobuf:"bytes,6,opt,name=order_by,json=orderBy" json:"order_by,omitempty"`
	// Specifies which information is returned about the time series.
	View ListTimeSeriesRequest_TimeSeriesView `protobuf:"varint,7,opt,name=view,enum=google.monitoring.v3.ListTimeSeriesRequest_TimeSeriesView" json:"view,omitempty"`
	// A positive number that is the maximum number of results to return.
	// When `view` field sets to `FULL`, it limits the number of `Points` server
	// will return; if `view` field is `HEADERS`, it limits the number of
	// `TimeSeries` server will return.
	PageSize int32 `protobuf:"varint,8,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// If this field is not empty then it must contain the `nextPageToken` value
	// returned by a previous call to this method.  Using this field causes the
	// method to return additional results from the previous method call.
	PageToken string `protobuf:"bytes,9,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListTimeSeriesRequest) Reset()                    { *m = ListTimeSeriesRequest{} }
func (m *ListTimeSeriesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTimeSeriesRequest) ProtoMessage()               {}
func (*ListTimeSeriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{8} }

func (m *ListTimeSeriesRequest) GetInterval() *TimeInterval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *ListTimeSeriesRequest) GetAggregation() *Aggregation {
	if m != nil {
		return m.Aggregation
	}
	return nil
}

// The `ListTimeSeries` response.
type ListTimeSeriesResponse struct {
	// One or more time series that match the filter included in the request.
	TimeSeries []*TimeSeries `protobuf:"bytes,1,rep,name=time_series,json=timeSeries" json:"time_series,omitempty"`
	// If there are more results than have been returned, then this field is set
	// to a non-empty value.  To see the additional results,
	// use that value as `pageToken` in the next call to this method.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListTimeSeriesResponse) Reset()                    { *m = ListTimeSeriesResponse{} }
func (m *ListTimeSeriesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTimeSeriesResponse) ProtoMessage()               {}
func (*ListTimeSeriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{9} }

func (m *ListTimeSeriesResponse) GetTimeSeries() []*TimeSeries {
	if m != nil {
		return m.TimeSeries
	}
	return nil
}

// The `CreateTimeSeries` request.
type CreateTimeSeriesRequest struct {
	// The project on which to execute the request. The format is
	// `"projects/{project_id_or_number}"`.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// The new data to be added to a list of time series.
	// Adds at most one data point to each of several time series.  The new data
	// point must be more recent than any other point in its time series.  Each
	// `TimeSeries` value must fully specify a unique time series by supplying
	// all label values for the metric and the monitored resource.
	TimeSeries []*TimeSeries `protobuf:"bytes,2,rep,name=time_series,json=timeSeries" json:"time_series,omitempty"`
}

func (m *CreateTimeSeriesRequest) Reset()                    { *m = CreateTimeSeriesRequest{} }
func (m *CreateTimeSeriesRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateTimeSeriesRequest) ProtoMessage()               {}
func (*CreateTimeSeriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{10} }

func (m *CreateTimeSeriesRequest) GetTimeSeries() []*TimeSeries {
	if m != nil {
		return m.TimeSeries
	}
	return nil
}

// Describes the result of a failed request to write data to a time series.
type CreateTimeSeriesError struct {
	// The time series, including the `Metric`, `MonitoredResource`,
	// and `Point`s (including timestamp and value) that resulted
	// in the error. This field provides all of the context that
	// would be needed to retry the operation.
	TimeSeries *TimeSeries `protobuf:"bytes,1,opt,name=time_series,json=timeSeries" json:"time_series,omitempty"`
	// The status of the requested write operation.
	Status *google_rpc.Status `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *CreateTimeSeriesError) Reset()                    { *m = CreateTimeSeriesError{} }
func (m *CreateTimeSeriesError) String() string            { return proto.CompactTextString(m) }
func (*CreateTimeSeriesError) ProtoMessage()               {}
func (*CreateTimeSeriesError) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{11} }

func (m *CreateTimeSeriesError) GetTimeSeries() *TimeSeries {
	if m != nil {
		return m.TimeSeries
	}
	return nil
}

func (m *CreateTimeSeriesError) GetStatus() *google_rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*ListMonitoredResourceDescriptorsRequest)(nil), "google.monitoring.v3.ListMonitoredResourceDescriptorsRequest")
	proto.RegisterType((*ListMonitoredResourceDescriptorsResponse)(nil), "google.monitoring.v3.ListMonitoredResourceDescriptorsResponse")
	proto.RegisterType((*GetMonitoredResourceDescriptorRequest)(nil), "google.monitoring.v3.GetMonitoredResourceDescriptorRequest")
	proto.RegisterType((*ListMetricDescriptorsRequest)(nil), "google.monitoring.v3.ListMetricDescriptorsRequest")
	proto.RegisterType((*ListMetricDescriptorsResponse)(nil), "google.monitoring.v3.ListMetricDescriptorsResponse")
	proto.RegisterType((*GetMetricDescriptorRequest)(nil), "google.monitoring.v3.GetMetricDescriptorRequest")
	proto.RegisterType((*CreateMetricDescriptorRequest)(nil), "google.monitoring.v3.CreateMetricDescriptorRequest")
	proto.RegisterType((*DeleteMetricDescriptorRequest)(nil), "google.monitoring.v3.DeleteMetricDescriptorRequest")
	proto.RegisterType((*ListTimeSeriesRequest)(nil), "google.monitoring.v3.ListTimeSeriesRequest")
	proto.RegisterType((*ListTimeSeriesResponse)(nil), "google.monitoring.v3.ListTimeSeriesResponse")
	proto.RegisterType((*CreateTimeSeriesRequest)(nil), "google.monitoring.v3.CreateTimeSeriesRequest")
	proto.RegisterType((*CreateTimeSeriesError)(nil), "google.monitoring.v3.CreateTimeSeriesError")
	proto.RegisterEnum("google.monitoring.v3.ListTimeSeriesRequest_TimeSeriesView", ListTimeSeriesRequest_TimeSeriesView_name, ListTimeSeriesRequest_TimeSeriesView_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for MetricService service

type MetricServiceClient interface {
	// Lists monitored resource descriptors that match a filter.
	ListMonitoredResourceDescriptors(ctx context.Context, in *ListMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (*ListMonitoredResourceDescriptorsResponse, error)
	// Gets a single monitored resource descriptor.
	GetMonitoredResourceDescriptor(ctx context.Context, in *GetMonitoredResourceDescriptorRequest, opts ...grpc.CallOption) (*google_api4.MonitoredResourceDescriptor, error)
	// Lists metric descriptors that match a filter.
	ListMetricDescriptors(ctx context.Context, in *ListMetricDescriptorsRequest, opts ...grpc.CallOption) (*ListMetricDescriptorsResponse, error)
	// Gets a single metric descriptor.
	GetMetricDescriptor(ctx context.Context, in *GetMetricDescriptorRequest, opts ...grpc.CallOption) (*google_api5.MetricDescriptor, error)
	// Creates a new metric descriptor.
	// User-created metric descriptors define
	// [custom metrics](/monitoring/custom-metrics).
	CreateMetricDescriptor(ctx context.Context, in *CreateMetricDescriptorRequest, opts ...grpc.CallOption) (*google_api5.MetricDescriptor, error)
	// Deletes a metric descriptor. Only user-created
	// [custom metrics](/monitoring/custom-metrics) can be deleted.
	DeleteMetricDescriptor(ctx context.Context, in *DeleteMetricDescriptorRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error)
	// Lists time series that match a filter.
	ListTimeSeries(ctx context.Context, in *ListTimeSeriesRequest, opts ...grpc.CallOption) (*ListTimeSeriesResponse, error)
	// Creates or adds data to one or more time series.
	// The response is empty if all time series in the request were written.
	// If any time series could not be written, a corresponding failure message is
	// included in the error response.
	CreateTimeSeries(ctx context.Context, in *CreateTimeSeriesRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error)
}

type metricServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetricServiceClient(cc *grpc.ClientConn) MetricServiceClient {
	return &metricServiceClient{cc}
}

func (c *metricServiceClient) ListMonitoredResourceDescriptors(ctx context.Context, in *ListMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (*ListMonitoredResourceDescriptorsResponse, error) {
	out := new(ListMonitoredResourceDescriptorsResponse)
	err := grpc.Invoke(ctx, "/google.monitoring.v3.MetricService/ListMonitoredResourceDescriptors", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricServiceClient) GetMonitoredResourceDescriptor(ctx context.Context, in *GetMonitoredResourceDescriptorRequest, opts ...grpc.CallOption) (*google_api4.MonitoredResourceDescriptor, error) {
	out := new(google_api4.MonitoredResourceDescriptor)
	err := grpc.Invoke(ctx, "/google.monitoring.v3.MetricService/GetMonitoredResourceDescriptor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricServiceClient) ListMetricDescriptors(ctx context.Context, in *ListMetricDescriptorsRequest, opts ...grpc.CallOption) (*ListMetricDescriptorsResponse, error) {
	out := new(ListMetricDescriptorsResponse)
	err := grpc.Invoke(ctx, "/google.monitoring.v3.MetricService/ListMetricDescriptors", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricServiceClient) GetMetricDescriptor(ctx context.Context, in *GetMetricDescriptorRequest, opts ...grpc.CallOption) (*google_api5.MetricDescriptor, error) {
	out := new(google_api5.MetricDescriptor)
	err := grpc.Invoke(ctx, "/google.monitoring.v3.MetricService/GetMetricDescriptor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricServiceClient) CreateMetricDescriptor(ctx context.Context, in *CreateMetricDescriptorRequest, opts ...grpc.CallOption) (*google_api5.MetricDescriptor, error) {
	out := new(google_api5.MetricDescriptor)
	err := grpc.Invoke(ctx, "/google.monitoring.v3.MetricService/CreateMetricDescriptor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricServiceClient) DeleteMetricDescriptor(ctx context.Context, in *DeleteMetricDescriptorRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error) {
	out := new(google_protobuf4.Empty)
	err := grpc.Invoke(ctx, "/google.monitoring.v3.MetricService/DeleteMetricDescriptor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricServiceClient) ListTimeSeries(ctx context.Context, in *ListTimeSeriesRequest, opts ...grpc.CallOption) (*ListTimeSeriesResponse, error) {
	out := new(ListTimeSeriesResponse)
	err := grpc.Invoke(ctx, "/google.monitoring.v3.MetricService/ListTimeSeries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricServiceClient) CreateTimeSeries(ctx context.Context, in *CreateTimeSeriesRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error) {
	out := new(google_protobuf4.Empty)
	err := grpc.Invoke(ctx, "/google.monitoring.v3.MetricService/CreateTimeSeries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MetricService service

type MetricServiceServer interface {
	// Lists monitored resource descriptors that match a filter.
	ListMonitoredResourceDescriptors(context.Context, *ListMonitoredResourceDescriptorsRequest) (*ListMonitoredResourceDescriptorsResponse, error)
	// Gets a single monitored resource descriptor.
	GetMonitoredResourceDescriptor(context.Context, *GetMonitoredResourceDescriptorRequest) (*google_api4.MonitoredResourceDescriptor, error)
	// Lists metric descriptors that match a filter.
	ListMetricDescriptors(context.Context, *ListMetricDescriptorsRequest) (*ListMetricDescriptorsResponse, error)
	// Gets a single metric descriptor.
	GetMetricDescriptor(context.Context, *GetMetricDescriptorRequest) (*google_api5.MetricDescriptor, error)
	// Creates a new metric descriptor.
	// User-created metric descriptors define
	// [custom metrics](/monitoring/custom-metrics).
	CreateMetricDescriptor(context.Context, *CreateMetricDescriptorRequest) (*google_api5.MetricDescriptor, error)
	// Deletes a metric descriptor. Only user-created
	// [custom metrics](/monitoring/custom-metrics) can be deleted.
	DeleteMetricDescriptor(context.Context, *DeleteMetricDescriptorRequest) (*google_protobuf4.Empty, error)
	// Lists time series that match a filter.
	ListTimeSeries(context.Context, *ListTimeSeriesRequest) (*ListTimeSeriesResponse, error)
	// Creates or adds data to one or more time series.
	// The response is empty if all time series in the request were written.
	// If any time series could not be written, a corresponding failure message is
	// included in the error response.
	CreateTimeSeries(context.Context, *CreateTimeSeriesRequest) (*google_protobuf4.Empty, error)
}

func RegisterMetricServiceServer(s *grpc.Server, srv MetricServiceServer) {
	s.RegisterService(&_MetricService_serviceDesc, srv)
}

func _MetricService_ListMonitoredResourceDescriptors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMonitoredResourceDescriptorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServiceServer).ListMonitoredResourceDescriptors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.MetricService/ListMonitoredResourceDescriptors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServiceServer).ListMonitoredResourceDescriptors(ctx, req.(*ListMonitoredResourceDescriptorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricService_GetMonitoredResourceDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMonitoredResourceDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServiceServer).GetMonitoredResourceDescriptor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.MetricService/GetMonitoredResourceDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServiceServer).GetMonitoredResourceDescriptor(ctx, req.(*GetMonitoredResourceDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricService_ListMetricDescriptors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMetricDescriptorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServiceServer).ListMetricDescriptors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.MetricService/ListMetricDescriptors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServiceServer).ListMetricDescriptors(ctx, req.(*ListMetricDescriptorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricService_GetMetricDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServiceServer).GetMetricDescriptor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.MetricService/GetMetricDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServiceServer).GetMetricDescriptor(ctx, req.(*GetMetricDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricService_CreateMetricDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMetricDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServiceServer).CreateMetricDescriptor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.MetricService/CreateMetricDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServiceServer).CreateMetricDescriptor(ctx, req.(*CreateMetricDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricService_DeleteMetricDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMetricDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServiceServer).DeleteMetricDescriptor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.MetricService/DeleteMetricDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServiceServer).DeleteMetricDescriptor(ctx, req.(*DeleteMetricDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricService_ListTimeSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTimeSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServiceServer).ListTimeSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.MetricService/ListTimeSeries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServiceServer).ListTimeSeries(ctx, req.(*ListTimeSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricService_CreateTimeSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTimeSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServiceServer).CreateTimeSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.MetricService/CreateTimeSeries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServiceServer).CreateTimeSeries(ctx, req.(*CreateTimeSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.monitoring.v3.MetricService",
	HandlerType: (*MetricServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMonitoredResourceDescriptors",
			Handler:    _MetricService_ListMonitoredResourceDescriptors_Handler,
		},
		{
			MethodName: "GetMonitoredResourceDescriptor",
			Handler:    _MetricService_GetMonitoredResourceDescriptor_Handler,
		},
		{
			MethodName: "ListMetricDescriptors",
			Handler:    _MetricService_ListMetricDescriptors_Handler,
		},
		{
			MethodName: "GetMetricDescriptor",
			Handler:    _MetricService_GetMetricDescriptor_Handler,
		},
		{
			MethodName: "CreateMetricDescriptor",
			Handler:    _MetricService_CreateMetricDescriptor_Handler,
		},
		{
			MethodName: "DeleteMetricDescriptor",
			Handler:    _MetricService_DeleteMetricDescriptor_Handler,
		},
		{
			MethodName: "ListTimeSeries",
			Handler:    _MetricService_ListTimeSeries_Handler,
		},
		{
			MethodName: "CreateTimeSeries",
			Handler:    _MetricService_CreateTimeSeries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor6,
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/monitoring/v3/metric_service.proto", fileDescriptor6)
}

var fileDescriptor6 = []byte{
	// 1009 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x56, 0xdf, 0x72, 0xdb, 0xc4,
	0x17, 0xfe, 0x6d, 0x92, 0xe6, 0xcf, 0xf1, 0x34, 0xbf, 0x74, 0xdb, 0x1a, 0x23, 0x1a, 0xc6, 0x68,
	0xa6, 0xc4, 0xb8, 0x45, 0xca, 0xd8, 0x1d, 0x2e, 0xd2, 0xa6, 0x43, 0xd2, 0x04, 0x9a, 0x21, 0x40,
	0x46, 0x2e, 0x5c, 0x70, 0xe3, 0x51, 0x94, 0x13, 0xb1, 0x60, 0x69, 0xc5, 0x6a, 0xed, 0x92, 0x32,
	0xe1, 0x02, 0x66, 0x7a, 0xcf, 0x00, 0x33, 0xf0, 0x12, 0x30, 0xbc, 0x03, 0x6f, 0xc0, 0x35, 0x77,
	0xbc, 0x02, 0xf7, 0x8c, 0x56, 0x72, 0x1c, 0xcb, 0x92, 0x2c, 0x97, 0x0b, 0x6e, 0x6c, 0x69, 0x77,
	0xcf, 0x39, 0xdf, 0xf9, 0xce, 0xd9, 0xf3, 0x09, 0x1e, 0xbb, 0x9c, 0xbb, 0x3d, 0x34, 0x5c, 0xde,
	0xb3, 0x7d, 0xd7, 0xe0, 0xc2, 0x35, 0x5d, 0xf4, 0x03, 0xc1, 0x25, 0x37, 0xe3, 0x2d, 0x3b, 0x60,
	0xa1, 0xe9, 0x71, 0x9f, 0x49, 0x2e, 0x98, 0xef, 0x9a, 0x83, 0xb6, 0xe9, 0xa1, 0x14, 0xcc, 0xe9,
	0x86, 0x28, 0x06, 0xcc, 0x41, 0x43, 0x9d, 0xa6, 0x37, 0x12, 0x4f, 0xa3, 0xa3, 0xc6, 0xa0, 0xad,
	0x1d, 0x94, 0xf3, 0x6f, 0x07, 0xcc, 0x4c, 0xdc, 0x39, 0xdc, 0x3f, 0x65, 0xae, 0x69, 0xfb, 0x3e,
	0x97, 0xb6, 0x64, 0xdc, 0x0f, 0xe3, 0x00, 0xda, 0x76, 0x79, 0x57, 0x31, 0xc0, 0xe4, 0x2f, 0x31,
	0xff, 0x70, 0x06, 0xf3, 0x38, 0x05, 0x3c, 0x11, 0x18, 0x8e, 0x5e, 0xba, 0x02, 0x43, 0xde, 0x17,
	0xc3, 0x84, 0xb5, 0xb7, 0x5f, 0x84, 0x3a, 0x87, 0x7b, 0x1e, 0xf7, 0xff, 0x8d, 0x87, 0xb1, 0xa4,
	0xda, 0x2e, 0x93, 0x9f, 0xf6, 0x8f, 0x0d, 0x87, 0x7b, 0x66, 0xec, 0xc5, 0x54, 0x1b, 0xc7, 0xfd,
	0x53, 0x33, 0x90, 0x67, 0x01, 0x86, 0x26, 0x7a, 0x81, 0x3c, 0x8b, 0x7f, 0x67, 0x23, 0x52, 0x04,
	0x8e, 0x19, 0x4a, 0x5b, 0xf6, 0xc3, 0xe4, 0x2f, 0x36, 0xd7, 0xbf, 0x23, 0xb0, 0x71, 0xc8, 0x42,
	0xf9, 0xfe, 0x90, 0x18, 0x2b, 0xe1, 0x65, 0x0f, 0x43, 0x47, 0xb0, 0x40, 0x72, 0x11, 0x5a, 0xf8,
	0x45, 0x1f, 0x43, 0x49, 0x29, 0x2c, 0xf8, 0xb6, 0x87, 0xb5, 0x2b, 0x75, 0xd2, 0x58, 0xb1, 0xd4,
	0x33, 0xad, 0xc2, 0xe2, 0x29, 0xeb, 0x49, 0x14, 0xb5, 0x39, 0xb5, 0x9a, 0xbc, 0xd1, 0x57, 0x60,
	0x25, 0xb0, 0x5d, 0xec, 0x86, 0xec, 0x19, 0xd6, 0xe6, 0xeb, 0xa4, 0x71, 0xc5, 0x5a, 0x8e, 0x16,
	0x3a, 0xec, 0x19, 0xd2, 0x75, 0x00, 0xb5, 0x29, 0xf9, 0xe7, 0xe8, 0xd7, 0x16, 0x94, 0xa1, 0x3a,
	0xfe, 0x24, 0x5a, 0xd0, 0x7f, 0x21, 0xd0, 0x98, 0x8e, 0x29, 0x0c, 0xb8, 0x1f, 0x22, 0xfd, 0x04,
	0x6e, 0x0c, 0x4b, 0xd9, 0x3d, 0x19, 0xed, 0xd7, 0x48, 0x7d, 0xbe, 0x51, 0x69, 0x6d, 0x18, 0x09,
	0x3d, 0x76, 0xc0, 0x8c, 0x02, 0x7f, 0xd6, 0x75, 0x31, 0x19, 0x83, 0xbe, 0x0e, 0xff, 0xf7, 0xf1,
	0x4b, 0xd9, 0xbd, 0x04, 0x36, 0xce, 0xf2, 0x6a, 0xb4, 0x7c, 0x74, 0x01, 0xf8, 0x3e, 0xdc, 0x7e,
	0x17, 0x8b, 0xe0, 0xa6, 0x19, 0x9c, 0x1f, 0x31, 0xa8, 0x3f, 0x27, 0x70, 0x4b, 0x65, 0xab, 0x5a,
	0xe1, 0x3f, 0xa4, 0xfd, 0x07, 0x02, 0xeb, 0x39, 0x40, 0x12, 0xae, 0xdf, 0x03, 0x9a, 0x4c, 0x8b,
	0x49, 0xa6, 0x6f, 0x8d, 0x31, 0x9d, 0x72, 0x61, 0x5d, 0xf3, 0xd2, 0x4e, 0x4b, 0x93, 0xbb, 0x09,
	0x5a, 0x44, 0x6e, 0xda, 0x63, 0x01, 0xa3, 0x5f, 0xc3, 0xfa, 0x23, 0x81, 0xb6, 0xc4, 0x19, 0x8c,
	0xe8, 0x01, 0x5c, 0x9b, 0xc8, 0x4d, 0x01, 0x9a, 0x96, 0xda, 0x5a, 0x3a, 0x35, 0xbd, 0x0d, 0xeb,
	0x7b, 0xd8, 0xc3, 0x99, 0xe2, 0xeb, 0x3f, 0xcd, 0xc3, 0xcd, 0x88, 0xfd, 0x27, 0xcc, 0xc3, 0x0e,
	0x0a, 0x86, 0x13, 0xf5, 0x87, 0x12, 0xf5, 0x7f, 0x08, 0xcb, 0xcc, 0x97, 0x28, 0x06, 0x76, 0x4f,
	0x15, 0xb8, 0xd2, 0xd2, 0x8d, 0xac, 0x51, 0x6e, 0x44, 0x61, 0x0e, 0x92, 0x93, 0xd6, 0x85, 0x0d,
	0x7d, 0x04, 0x15, 0xdb, 0x75, 0x05, 0xba, 0x6a, 0x58, 0xab, 0x96, 0xab, 0xb4, 0x5e, 0xcb, 0x76,
	0xb1, 0x33, 0x3a, 0x68, 0x5d, 0xb6, 0xa2, 0x2f, 0xc3, 0x32, 0x17, 0x27, 0x28, 0xba, 0xc7, 0x67,
	0xb5, 0x45, 0x05, 0x6f, 0x49, 0xbd, 0xef, 0x9e, 0xd1, 0x0f, 0x60, 0x61, 0xc0, 0xf0, 0x69, 0x6d,
	0xa9, 0x4e, 0x1a, 0xab, 0xad, 0xad, 0x6c, 0xc7, 0x99, 0x34, 0x18, 0xa3, 0x95, 0x8f, 0x19, 0x3e,
	0xb5, 0x94, 0x9f, 0xf1, 0x7e, 0x5f, 0x2e, 0xec, 0xf7, 0x95, 0x74, 0xbf, 0x6f, 0xc0, 0xea, 0xb8,
	0x4f, 0xba, 0x0c, 0x0b, 0xef, 0x7c, 0x74, 0x78, 0xb8, 0xf6, 0x3f, 0x5a, 0x81, 0xa5, 0xc7, 0xfb,
	0x3b, 0x7b, 0xfb, 0x56, 0x67, 0x8d, 0xe8, 0xdf, 0x12, 0xa8, 0xa6, 0x31, 0x25, 0x37, 0x62, 0x07,
	0x2a, 0x92, 0x79, 0x18, 0xa9, 0x27, 0xc3, 0xe1, 0x55, 0xa8, 0xe7, 0x53, 0x9e, 0x98, 0x83, 0xbc,
	0x78, 0x2e, 0x7d, 0x0f, 0x02, 0x78, 0x29, 0xee, 0xea, 0xfc, 0x0e, 0xb9, 0xdc, 0xcf, 0x29, 0x64,
	0x73, 0xb3, 0x23, 0x8b, 0x26, 0xd3, 0xcd, 0x74, 0xc8, 0x7d, 0x21, 0xb8, 0x98, 0x4c, 0x9b, 0xcc,
	0x9c, 0x76, 0x13, 0x16, 0x63, 0x21, 0x4a, 0x2e, 0x19, 0x1d, 0x5a, 0x8b, 0xc0, 0x31, 0x3a, 0x6a,
	0xc7, 0x4a, 0x4e, 0xb4, 0xfe, 0x06, 0xb8, 0x1a, 0xdf, 0xa5, 0x4e, 0xfc, 0x59, 0x41, 0xff, 0x24,
	0x50, 0x9f, 0x26, 0x11, 0x74, 0x3b, 0xbf, 0xbd, 0x4a, 0xc8, 0x9d, 0xf6, 0xf0, 0x45, 0xcd, 0xe3,
	0xde, 0xd0, 0xb7, 0xbe, 0xf9, 0xe3, 0xaf, 0xef, 0xe7, 0xee, 0xd1, 0x56, 0x24, 0xf4, 0x5f, 0x45,
	0x45, 0xd9, 0x0e, 0x04, 0xff, 0x0c, 0x1d, 0x19, 0x9a, 0xcd, 0xf3, 0xd1, 0xa7, 0x48, 0x16, 0xf4,
	0xdf, 0x09, 0xbc, 0x5a, 0x2c, 0x29, 0xf4, 0x7e, 0x36, 0xbc, 0x52, 0x42, 0xa4, 0x95, 0xd5, 0x45,
	0xfd, 0x81, 0x4a, 0xe2, 0x2d, 0x7a, 0x2f, 0x2b, 0x89, 0xc2, 0x1c, 0xcc, 0xe6, 0x39, 0xfd, 0x8d,
	0xc4, 0x43, 0x6d, 0x42, 0x52, 0x68, 0xab, 0x80, 0xdc, 0x1c, 0x21, 0xd4, 0xda, 0x33, 0xd9, 0x24,
	0x55, 0x30, 0x55, 0x02, 0x6f, 0xd0, 0x8d, 0x9c, 0x2a, 0x4c, 0x20, 0xfb, 0x99, 0xc0, 0xf5, 0x0c,
	0xc1, 0xa1, 0x9b, 0xf9, 0x7c, 0x67, 0x8f, 0x79, 0xad, 0x50, 0x37, 0xf4, 0x96, 0x02, 0x76, 0x97,
	0x36, 0xb3, 0x99, 0x4d, 0xe3, 0x32, 0x9b, 0xcd, 0x73, 0xfa, 0x2b, 0x81, 0x6a, 0xb6, 0xb4, 0xd1,
	0x1c, 0x72, 0x0a, 0x85, 0x70, 0x0a, 0xc2, 0x5d, 0x85, 0xf0, 0x81, 0x5e, 0x96, 0xba, 0xad, 0x49,
	0x05, 0x8d, 0xd8, 0xac, 0x66, 0x8b, 0x61, 0x1e, 0xe2, 0x42, 0xe9, 0xd4, 0xaa, 0x43, 0xa3, 0xe1,
	0x97, 0xb1, 0xb1, 0x1f, 0x7d, 0x0c, 0x0f, 0xd9, 0x6c, 0xce, 0xc2, 0xe6, 0x8f, 0x04, 0x56, 0xc7,
	0xe7, 0x3a, 0xbd, 0x33, 0x83, 0x22, 0x69, 0x77, 0xcb, 0x1d, 0x4e, 0x1a, 0xb1, 0xa1, 0x10, 0xea,
	0xb4, 0x9e, 0xcd, 0xe6, 0xa5, 0xd1, 0xf8, 0x9c, 0xc0, 0x5a, 0x7a, 0xee, 0xd2, 0x37, 0x8b, 0xea,
	0x3b, 0x89, 0x2d, 0x8f, 0xa7, 0x3b, 0x0a, 0xc5, 0x6d, 0x7d, 0x2a, 0x8a, 0x2d, 0xd2, 0xdc, 0xdd,
	0x84, 0x9a, 0xc3, 0xbd, 0xcc, 0xc0, 0xbb, 0x74, 0x6c, 0x20, 0x1f, 0x45, 0x51, 0x8e, 0xc8, 0xf1,
	0xa2, 0x0a, 0xd7, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x3a, 0xd0, 0x46, 0x8a, 0x0e, 0x00,
	0x00,
}
