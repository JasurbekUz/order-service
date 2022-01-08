// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc.proto

package catalog

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xd1, 0x4a, 0x02, 0x41,
	0x14, 0xcd, 0x97, 0xcc, 0x9b, 0x6e, 0xeb, 0x48, 0x44, 0x46, 0xfb, 0x09, 0x83, 0x28, 0x18, 0x48,
	0x08, 0xad, 0xd5, 0x12, 0xf4, 0x54, 0xf4, 0x1c, 0x9b, 0x5e, 0x4c, 0x54, 0x66, 0xdb, 0x9d, 0x02,
	0xff, 0xa4, 0xa7, 0xbe, 0xa7, 0xc7, 0x3e, 0x21, 0xec, 0x47, 0x42, 0x67, 0xba, 0xa3, 0xbb, 0x93,
	0xb4, 0xbd, 0x39, 0x67, 0xce, 0x99, 0x7b, 0xce, 0x91, 0xbb, 0x50, 0x8a, 0xa3, 0x3e, 0x8f, 0x62,
	0x21, 0x05, 0x2b, 0xf6, 0x43, 0x19, 0x4e, 0xc4, 0xb0, 0x5e, 0x99, 0x62, 0x92, 0x84, 0x43, 0x54,
	0x78, 0xf3, 0xad, 0x04, 0x4e, 0x4f, 0x5d, 0xdd, 0x62, 0xfc, 0x32, 0xea, 0x23, 0x6b, 0x40, 0xb9,
	0x17, 0x63, 0x28, 0xf1, 0xec, 0x59, 0x3e, 0x8a, 0x98, 0xed, 0x71, 0xad, 0xe5, 0x0a, 0xa8, 0xa7,
	0x01, 0xc6, 0xa1, 0x14, 0xa0, 0xd4, 0x07, 0x97, 0x6e, 0xfd, 0xd9, 0xd5, 0xe0, 0x06, 0x9f, 0xb2,
	0xfc, 0x13, 0x00, 0xe2, 0x27, 0x2b, 0x82, 0xeb, 0x51, 0x22, 0x17, 0x82, 0x83, 0x94, 0x40, 0xe1,
	0x49, 0xb4, 0xb0, 0x76, 0x17, 0x0d, 0xf2, 0x58, 0x6b, 0x40, 0xf9, 0x1c, 0x27, 0x48, 0x8a, 0xac,
	0x3b, 0x87, 0x90, 0x8b, 0x69, 0x24, 0x67, 0xac, 0x0d, 0x8e, 0x8a, 0xdf, 0x0b, 0x25, 0x0e, 0x45,
	0x3c, 0x63, 0x55, 0x62, 0xfc, 0x40, 0xf5, 0x2c, 0xc4, 0x9a, 0xb0, 0x1b, 0xa0, 0xa4, 0x63, 0x76,
	0x90, 0x45, 0x73, 0x0a, 0x15, 0xa3, 0x19, 0xa1, 0xad, 0x8b, 0xc3, 0x8c, 0x8a, 0xda, 0x68, 0x83,
	0xa3, 0xda, 0xc8, 0xed, 0xd4, 0x51, 0x9d, 0x6c, 0x30, 0x9b, 0x6e, 0x85, 0x03, 0xa8, 0x56, 0x7c,
	0x21, 0xc6, 0xac, 0x62, 0xf8, 0x42, 0x8c, 0x57, 0x66, 0x04, 0x28, 0xef, 0x97, 0x0c, 0x0e, 0xc5,
	0x00, 0xe5, 0xf2, 0xe7, 0xa6, 0x26, 0x88, 0xdf, 0x82, 0x1d, 0xcd, 0xb7, 0x95, 0xb0, 0xbf, 0x36,
	0x8f, 0x0a, 0xe0, 0x00, 0xaa, 0x80, 0x3f, 0x9b, 0x02, 0x15, 0xfc, 0x17, 0x5f, 0xe9, 0xd0, 0x5d,
	0x60, 0x26, 0x34, 0x95, 0xb5, 0x6e, 0xc6, 0x52, 0x34, 0xcd, 0xbb, 0x84, 0x9a, 0x0e, 0xb5, 0xfa,
	0xdf, 0x59, 0xf2, 0x1d, 0x5b, 0x9f, 0xa4, 0x9c, 0x5d, 0x60, 0x26, 0xe7, 0x3f, 0x7c, 0x74, 0xc0,
	0x35, 0x39, 0xf4, 0x22, 0xd4, 0xd6, 0xd4, 0x7a, 0x7d, 0x2c, 0x5a, 0x1f, 0xaa, 0x3a, 0x83, 0xd9,
	0x45, 0x4b, 0x82, 0x23, 0xcb, 0x73, 0xe4, 0xbf, 0x03, 0xae, 0xf1, 0x9f, 0x6f, 0xbe, 0xef, 0xbe,
	0xcf, 0xbd, 0xc2, 0xc7, 0xdc, 0x2b, 0x7c, 0xce, 0xbd, 0xc2, 0xeb, 0x97, 0xb7, 0xf5, 0xb0, 0xbd,
	0xfc, 0x72, 0xb5, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x34, 0x39, 0x6f, 0xde, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CatalogServiceClient is the client API for CatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatalogServiceClient interface {
	CreateAuthor(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Author, error)
	GetAuthor(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Author, error)
	GetAuthors(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*AuthorListResp, error)
	UpdateAuthor(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Author, error)
	DeleteAuthor(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Empty, error)
	CreateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error)
	GetCategory(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Category, error)
	GetCategories(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*CategoryListResp, error)
	UpdateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error)
	DeleteCategory(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Empty, error)
	CreateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Get_Book, error)
	GetBook(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Get_Book, error)
	GetBooks(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*BookListResp, error)
	UpdateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Get_Book, error)
	DeleteBook(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Empty, error)
	CreateBookCategory(ctx context.Context, in *BookCategory, opts ...grpc.CallOption) (*Get_Book, error)
	GetBookCategoryList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*BookCategoryListResp, error)
	UpdateBookCategory(ctx context.Context, in *BookCategory, opts ...grpc.CallOption) (*Get_Book, error)
	CreateBookAuthor(ctx context.Context, in *BookAuthor, opts ...grpc.CallOption) (*Get_Book, error)
	GetBookAuthorList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*BookAuthorListResp, error)
	UpdateBookAuthor(ctx context.Context, in *BookAuthor, opts ...grpc.CallOption) (*Get_Book, error)
}

type catalogServiceClient struct {
	cc *grpc.ClientConn
}

func NewCatalogServiceClient(cc *grpc.ClientConn) CatalogServiceClient {
	return &catalogServiceClient{cc}
}

func (c *catalogServiceClient) CreateAuthor(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/CreateAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetAuthor(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/GetAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetAuthors(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*AuthorListResp, error) {
	out := new(AuthorListResp)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/GetAuthors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateAuthor(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/UpdateAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) DeleteAuthor(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/DeleteAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) CreateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetCategory(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/GetCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetCategories(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*CategoryListResp, error) {
	out := new(CategoryListResp)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/GetCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/UpdateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) DeleteCategory(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/DeleteCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) CreateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Get_Book, error) {
	out := new(Get_Book)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetBook(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Get_Book, error) {
	out := new(Get_Book)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetBooks(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*BookListResp, error) {
	out := new(BookListResp)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/GetBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Get_Book, error) {
	out := new(Get_Book)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) DeleteBook(ctx context.Context, in *ByIdReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) CreateBookCategory(ctx context.Context, in *BookCategory, opts ...grpc.CallOption) (*Get_Book, error) {
	out := new(Get_Book)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/CreateBookCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetBookCategoryList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*BookCategoryListResp, error) {
	out := new(BookCategoryListResp)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/GetBookCategoryList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateBookCategory(ctx context.Context, in *BookCategory, opts ...grpc.CallOption) (*Get_Book, error) {
	out := new(Get_Book)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/UpdateBookCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) CreateBookAuthor(ctx context.Context, in *BookAuthor, opts ...grpc.CallOption) (*Get_Book, error) {
	out := new(Get_Book)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/CreateBookAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetBookAuthorList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*BookAuthorListResp, error) {
	out := new(BookAuthorListResp)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/GetBookAuthorList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateBookAuthor(ctx context.Context, in *BookAuthor, opts ...grpc.CallOption) (*Get_Book, error) {
	out := new(Get_Book)
	err := c.cc.Invoke(ctx, "/catalog.CatalogService/UpdateBookAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServiceServer is the server API for CatalogService service.
type CatalogServiceServer interface {
	CreateAuthor(context.Context, *Author) (*Author, error)
	GetAuthor(context.Context, *ByIdReq) (*Author, error)
	GetAuthors(context.Context, *ListReq) (*AuthorListResp, error)
	UpdateAuthor(context.Context, *Author) (*Author, error)
	DeleteAuthor(context.Context, *ByIdReq) (*Empty, error)
	CreateCategory(context.Context, *Category) (*Category, error)
	GetCategory(context.Context, *ByIdReq) (*Category, error)
	GetCategories(context.Context, *ListReq) (*CategoryListResp, error)
	UpdateCategory(context.Context, *Category) (*Category, error)
	DeleteCategory(context.Context, *ByIdReq) (*Empty, error)
	CreateBook(context.Context, *Book) (*Get_Book, error)
	GetBook(context.Context, *ByIdReq) (*Get_Book, error)
	GetBooks(context.Context, *ListReq) (*BookListResp, error)
	UpdateBook(context.Context, *Book) (*Get_Book, error)
	DeleteBook(context.Context, *ByIdReq) (*Empty, error)
	CreateBookCategory(context.Context, *BookCategory) (*Get_Book, error)
	GetBookCategoryList(context.Context, *ListReq) (*BookCategoryListResp, error)
	UpdateBookCategory(context.Context, *BookCategory) (*Get_Book, error)
	CreateBookAuthor(context.Context, *BookAuthor) (*Get_Book, error)
	GetBookAuthorList(context.Context, *ListReq) (*BookAuthorListResp, error)
	UpdateBookAuthor(context.Context, *BookAuthor) (*Get_Book, error)
}

// UnimplementedCatalogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCatalogServiceServer struct {
}

func (*UnimplementedCatalogServiceServer) CreateAuthor(ctx context.Context, req *Author) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuthor not implemented")
}
func (*UnimplementedCatalogServiceServer) GetAuthor(ctx context.Context, req *ByIdReq) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthor not implemented")
}
func (*UnimplementedCatalogServiceServer) GetAuthors(ctx context.Context, req *ListReq) (*AuthorListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthors not implemented")
}
func (*UnimplementedCatalogServiceServer) UpdateAuthor(ctx context.Context, req *Author) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthor not implemented")
}
func (*UnimplementedCatalogServiceServer) DeleteAuthor(ctx context.Context, req *ByIdReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuthor not implemented")
}
func (*UnimplementedCatalogServiceServer) CreateCategory(ctx context.Context, req *Category) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (*UnimplementedCatalogServiceServer) GetCategory(ctx context.Context, req *ByIdReq) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (*UnimplementedCatalogServiceServer) GetCategories(ctx context.Context, req *ListReq) (*CategoryListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (*UnimplementedCatalogServiceServer) UpdateCategory(ctx context.Context, req *Category) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (*UnimplementedCatalogServiceServer) DeleteCategory(ctx context.Context, req *ByIdReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (*UnimplementedCatalogServiceServer) CreateBook(ctx context.Context, req *Book) (*Get_Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (*UnimplementedCatalogServiceServer) GetBook(ctx context.Context, req *ByIdReq) (*Get_Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (*UnimplementedCatalogServiceServer) GetBooks(ctx context.Context, req *ListReq) (*BookListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooks not implemented")
}
func (*UnimplementedCatalogServiceServer) UpdateBook(ctx context.Context, req *Book) (*Get_Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (*UnimplementedCatalogServiceServer) DeleteBook(ctx context.Context, req *ByIdReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (*UnimplementedCatalogServiceServer) CreateBookCategory(ctx context.Context, req *BookCategory) (*Get_Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBookCategory not implemented")
}
func (*UnimplementedCatalogServiceServer) GetBookCategoryList(ctx context.Context, req *ListReq) (*BookCategoryListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookCategoryList not implemented")
}
func (*UnimplementedCatalogServiceServer) UpdateBookCategory(ctx context.Context, req *BookCategory) (*Get_Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBookCategory not implemented")
}
func (*UnimplementedCatalogServiceServer) CreateBookAuthor(ctx context.Context, req *BookAuthor) (*Get_Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBookAuthor not implemented")
}
func (*UnimplementedCatalogServiceServer) GetBookAuthorList(ctx context.Context, req *ListReq) (*BookAuthorListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookAuthorList not implemented")
}
func (*UnimplementedCatalogServiceServer) UpdateBookAuthor(ctx context.Context, req *BookAuthor) (*Get_Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBookAuthor not implemented")
}

func RegisterCatalogServiceServer(s *grpc.Server, srv CatalogServiceServer) {
	s.RegisterService(&_CatalogService_serviceDesc, srv)
}

func _CatalogService_CreateAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Author)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/CreateAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateAuthor(ctx, req.(*Author))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/GetAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetAuthor(ctx, req.(*ByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetAuthors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetAuthors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/GetAuthors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetAuthors(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Author)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/UpdateAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateAuthor(ctx, req.(*Author))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_DeleteAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).DeleteAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/DeleteAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).DeleteAuthor(ctx, req.(*ByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateCategory(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/GetCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetCategory(ctx, req.(*ByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/GetCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetCategories(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/UpdateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateCategory(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/DeleteCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).DeleteCategory(ctx, req.(*ByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetBook(ctx, req.(*ByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/GetBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetBooks(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).DeleteBook(ctx, req.(*ByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_CreateBookCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateBookCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/CreateBookCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateBookCategory(ctx, req.(*BookCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetBookCategoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetBookCategoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/GetBookCategoryList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetBookCategoryList(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateBookCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateBookCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/UpdateBookCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateBookCategory(ctx, req.(*BookCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_CreateBookAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookAuthor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateBookAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/CreateBookAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateBookAuthor(ctx, req.(*BookAuthor))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetBookAuthorList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetBookAuthorList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/GetBookAuthorList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetBookAuthorList(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateBookAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookAuthor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateBookAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.CatalogService/UpdateBookAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateBookAuthor(ctx, req.(*BookAuthor))
	}
	return interceptor(ctx, in, info, handler)
}

var _CatalogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "catalog.CatalogService",
	HandlerType: (*CatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAuthor",
			Handler:    _CatalogService_CreateAuthor_Handler,
		},
		{
			MethodName: "GetAuthor",
			Handler:    _CatalogService_GetAuthor_Handler,
		},
		{
			MethodName: "GetAuthors",
			Handler:    _CatalogService_GetAuthors_Handler,
		},
		{
			MethodName: "UpdateAuthor",
			Handler:    _CatalogService_UpdateAuthor_Handler,
		},
		{
			MethodName: "DeleteAuthor",
			Handler:    _CatalogService_DeleteAuthor_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _CatalogService_CreateCategory_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _CatalogService_GetCategory_Handler,
		},
		{
			MethodName: "GetCategories",
			Handler:    _CatalogService_GetCategories_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _CatalogService_UpdateCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _CatalogService_DeleteCategory_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _CatalogService_CreateBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _CatalogService_GetBook_Handler,
		},
		{
			MethodName: "GetBooks",
			Handler:    _CatalogService_GetBooks_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _CatalogService_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _CatalogService_DeleteBook_Handler,
		},
		{
			MethodName: "CreateBookCategory",
			Handler:    _CatalogService_CreateBookCategory_Handler,
		},
		{
			MethodName: "GetBookCategoryList",
			Handler:    _CatalogService_GetBookCategoryList_Handler,
		},
		{
			MethodName: "UpdateBookCategory",
			Handler:    _CatalogService_UpdateBookCategory_Handler,
		},
		{
			MethodName: "CreateBookAuthor",
			Handler:    _CatalogService_CreateBookAuthor_Handler,
		},
		{
			MethodName: "GetBookAuthorList",
			Handler:    _CatalogService_GetBookAuthorList_Handler,
		},
		{
			MethodName: "UpdateBookAuthor",
			Handler:    _CatalogService_UpdateBookAuthor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}