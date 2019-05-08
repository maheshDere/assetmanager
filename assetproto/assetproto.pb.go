// Code generated by protoc-gen-go.
// source: assetproto.proto
// DO NOT EDIT!

/*
Package assetproto is a generated protocol buffer package.

It is generated from these files:
	assetproto.proto

It has these top-level messages:
	CreateAssetReq
	AssetReq
	Asset
	ListAssetResp
*/
package assetproto

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type CreateAssetReq struct {
	Caption     string   `protobuf:"bytes,1,opt,name=caption" json:"caption,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	AssetType   string   `protobuf:"bytes,3,opt,name=asset_type" json:"asset_type,omitempty"`
	Url         string   `protobuf:"bytes,4,opt,name=url" json:"url,omitempty"`
	Source      string   `protobuf:"bytes,5,opt,name=source" json:"source,omitempty"`
	Owner       string   `protobuf:"bytes,6,opt,name=owner" json:"owner,omitempty"`
	Private     bool     `protobuf:"varint,7,opt,name=private" json:"private,omitempty"`
	Tags        []string `protobuf:"bytes,8,rep,name=tags" json:"tags,omitempty"`
}

func (m *CreateAssetReq) Reset()         { *m = CreateAssetReq{} }
func (m *CreateAssetReq) String() string { return proto.CompactTextString(m) }
func (*CreateAssetReq) ProtoMessage()    {}

type AssetReq struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Tag     string `protobuf:"bytes,2,opt,name=tag" json:"tag,omitempty"`
	Owner   string `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
	Private bool   `protobuf:"varint,4,opt,name=private" json:"private,omitempty"`
	Video   bool   `protobuf:"varint,5,opt,name=video" json:"video,omitempty"`
	Audio   bool   `protobuf:"varint,6,opt,name=audio" json:"audio,omitempty"`
	Image   bool   `protobuf:"varint,7,opt,name=image" json:"image,omitempty"`
	Limit   int32  `protobuf:"varint,8,opt,name=limit" json:"limit,omitempty"`
	Offset  int32  `protobuf:"varint,9,opt,name=offset" json:"offset,omitempty"`
}

func (m *AssetReq) Reset()         { *m = AssetReq{} }
func (m *AssetReq) String() string { return proto.CompactTextString(m) }
func (*AssetReq) ProtoMessage()    {}

type Asset struct {
	Url         string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	AssetType   string   `protobuf:"bytes,2,opt,name=asset_type" json:"asset_type,omitempty"`
	Id          string   `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Tags        []string `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	Caption     string   `protobuf:"bytes,5,opt,name=caption" json:"caption,omitempty"`
	Description string   `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}

type ListAssetResp struct {
	Count  int32    `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Assets []*Asset `protobuf:"bytes,2,rep,name=assets" json:"assets,omitempty"`
}

func (m *ListAssetResp) Reset()         { *m = ListAssetResp{} }
func (m *ListAssetResp) String() string { return proto.CompactTextString(m) }
func (*ListAssetResp) ProtoMessage()    {}

func (m *ListAssetResp) GetAssets() []*Asset {
	if m != nil {
		return m.Assets
	}
	return nil
}

func init() {
}

// Client API for AssetService service

type AssetServiceClient interface {
	CreateAsset(ctx context.Context, in *CreateAssetReq, opts ...grpc.CallOption) (*ListAssetResp, error)
	GetAsset(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*ListAssetResp, error)
	GetMediaLibrary(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*ListAssetResp, error)
}

type assetServiceClient struct {
	cc *grpc.ClientConn
}

func NewAssetServiceClient(cc *grpc.ClientConn) AssetServiceClient {
	return &assetServiceClient{cc}
}

func (c *assetServiceClient) CreateAsset(ctx context.Context, in *CreateAssetReq, opts ...grpc.CallOption) (*ListAssetResp, error) {
	out := new(ListAssetResp)
	err := grpc.Invoke(ctx, "/assetproto.AssetService/CreateAsset", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) GetAsset(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*ListAssetResp, error) {
	out := new(ListAssetResp)
	err := grpc.Invoke(ctx, "/assetproto.AssetService/GetAsset", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) GetMediaLibrary(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*ListAssetResp, error) {
	out := new(ListAssetResp)
	err := grpc.Invoke(ctx, "/assetproto.AssetService/GetMediaLibrary", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AssetService service

type AssetServiceServer interface {
	CreateAsset(context.Context, *CreateAssetReq) (*ListAssetResp, error)
	GetAsset(context.Context, *AssetReq) (*ListAssetResp, error)
	GetMediaLibrary(context.Context, *AssetReq) (*ListAssetResp, error)
}

func RegisterAssetServiceServer(s *grpc.Server, srv AssetServiceServer) {
	s.RegisterService(&_AssetService_serviceDesc, srv)
}

func _AssetService_CreateAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAssetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AssetServiceServer).CreateAsset(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _AssetService_GetAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AssetServiceServer).GetAsset(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _AssetService_GetMediaLibrary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AssetServiceServer).GetMediaLibrary(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _AssetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "assetproto.AssetService",
	HandlerType: (*AssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAsset",
			Handler:    _AssetService_CreateAsset_Handler,
		},
		{
			MethodName: "GetAsset",
			Handler:    _AssetService_GetAsset_Handler,
		},
		{
			MethodName: "GetMediaLibrary",
			Handler:    _AssetService_GetMediaLibrary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}