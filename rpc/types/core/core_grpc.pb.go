// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: core.proto

package core

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CoreClient is the client API for Core service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoreClient interface {
	// init
	InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
	// user service
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*BaseResp, error)
	CreateOrUpdateUser(ctx context.Context, in *CreateOrUpdateUserReq, opts ...grpc.CallOption) (*BaseResp, error)
	GetUserById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*UserInfoResp, error)
	GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*UserListResp, error)
	DeleteUser(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileReq, opts ...grpc.CallOption) (*BaseResp, error)
	// menu service
	// menu management
	CreateOrUpdateMenu(ctx context.Context, in *CreateOrUpdateMenuReq, opts ...grpc.CallOption) (*BaseResp, error)
	DeleteMenu(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
	GetMenuListByRole(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*MenuInfoList, error)
	GetMenuByPage(ctx context.Context, in *PageInfoReq, opts ...grpc.CallOption) (*MenuInfoList, error)
	CreateOrUpdateMenuParam(ctx context.Context, in *CreateOrUpdateMenuParamReq, opts ...grpc.CallOption) (*BaseResp, error)
	DeleteMenuParam(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
	GeMenuParamListByMenuId(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*MenuParamListResp, error)
	// role service
	CreateOrUpdateRole(ctx context.Context, in *RoleInfo, opts ...grpc.CallOption) (*BaseResp, error)
	DeleteRole(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
	GetRoleById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*RoleInfo, error)
	GetRoleList(ctx context.Context, in *PageInfoReq, opts ...grpc.CallOption) (*RoleListResp, error)
	SetRoleStatus(ctx context.Context, in *SetStatusReq, opts ...grpc.CallOption) (*BaseResp, error)
	// api management service
	CreateOrUpdateApi(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*BaseResp, error)
	DeleteApi(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
	GetApiList(ctx context.Context, in *ApiPageReq, opts ...grpc.CallOption) (*ApiListResp, error)
	// authorization management service
	GetMenuAuthority(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*RoleMenuAuthorityResp, error)
	CreateOrUpdateMenuAuthority(ctx context.Context, in *RoleMenuAuthorityReq, opts ...grpc.CallOption) (*BaseResp, error)
	// example
	Hello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*BaseResp, error)
}

type coreClient struct {
	cc grpc.ClientConnInterface
}

func NewCoreClient(cc grpc.ClientConnInterface) CoreClient {
	return &coreClient{cc}
}

func (c *coreClient) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/core.core/initDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/core.core/login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/core.core/changePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) CreateOrUpdateUser(ctx context.Context, in *CreateOrUpdateUserReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/core.core/createOrUpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetUserById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	out := new(UserInfoResp)
	err := c.cc.Invoke(ctx, "/core.core/getUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	out := new(UserListResp)
	err := c.cc.Invoke(ctx, "/core.core/getUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) DeleteUser(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/core.core/deleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) UpdateProfile(ctx context.Context, in *UpdateProfileReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/core.core/updateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) CreateOrUpdateMenu(ctx context.Context, in *CreateOrUpdateMenuReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/core.core/createOrUpdateMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) DeleteMenu(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/core.core/deleteMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetMenuListByRole(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*MenuInfoList, error) {
	out := new(MenuInfoList)
	err := c.cc.Invoke(ctx, "/core.core/getMenuListByRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetMenuByPage(ctx context.Context, in *PageInfoReq, opts ...grpc.CallOption) (*MenuInfoList, error) {
	out := new(MenuInfoList)
	err := c.cc.Invoke(ctx, "/core.core/getMenuByPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) CreateOrUpdateMenuParam(ctx context.Context, in *CreateOrUpdateMenuParamReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/core.core/createOrUpdateMenuParam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) DeleteMenuParam(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/core.core/deleteMenuParam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GeMenuParamListByMenuId(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*MenuParamListResp, error) {
	out := new(MenuParamListResp)
	err := c.cc.Invoke(ctx, "/core.core/geMenuParamListByMenuId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) CreateOrUpdateRole(ctx context.Context, in *RoleInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/core.core/createOrUpdateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) DeleteRole(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/core.core/deleteRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetRoleById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*RoleInfo, error) {
	out := new(RoleInfo)
	err := c.cc.Invoke(ctx, "/core.core/getRoleById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetRoleList(ctx context.Context, in *PageInfoReq, opts ...grpc.CallOption) (*RoleListResp, error) {
	out := new(RoleListResp)
	err := c.cc.Invoke(ctx, "/core.core/getRoleList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) SetRoleStatus(ctx context.Context, in *SetStatusReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/core.core/setRoleStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) CreateOrUpdateApi(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/core.core/createOrUpdateApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) DeleteApi(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/core.core/deleteApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetApiList(ctx context.Context, in *ApiPageReq, opts ...grpc.CallOption) (*ApiListResp, error) {
	out := new(ApiListResp)
	err := c.cc.Invoke(ctx, "/core.core/getApiList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetMenuAuthority(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*RoleMenuAuthorityResp, error) {
	out := new(RoleMenuAuthorityResp)
	err := c.cc.Invoke(ctx, "/core.core/getMenuAuthority", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) CreateOrUpdateMenuAuthority(ctx context.Context, in *RoleMenuAuthorityReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/core.core/createOrUpdateMenuAuthority", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) Hello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/core.core/hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServer is the server API for Core service.
// All implementations must embed UnimplementedCoreServer
// for forward compatibility
type CoreServer interface {
	// init
	InitDatabase(context.Context, *Empty) (*BaseResp, error)
	// user service
	Login(context.Context, *LoginReq) (*LoginResp, error)
	ChangePassword(context.Context, *ChangePasswordReq) (*BaseResp, error)
	CreateOrUpdateUser(context.Context, *CreateOrUpdateUserReq) (*BaseResp, error)
	GetUserById(context.Context, *UUIDReq) (*UserInfoResp, error)
	GetUserList(context.Context, *GetUserListReq) (*UserListResp, error)
	DeleteUser(context.Context, *IDReq) (*BaseResp, error)
	UpdateProfile(context.Context, *UpdateProfileReq) (*BaseResp, error)
	// menu service
	// menu management
	CreateOrUpdateMenu(context.Context, *CreateOrUpdateMenuReq) (*BaseResp, error)
	DeleteMenu(context.Context, *IDReq) (*BaseResp, error)
	GetMenuListByRole(context.Context, *IDReq) (*MenuInfoList, error)
	GetMenuByPage(context.Context, *PageInfoReq) (*MenuInfoList, error)
	CreateOrUpdateMenuParam(context.Context, *CreateOrUpdateMenuParamReq) (*BaseResp, error)
	DeleteMenuParam(context.Context, *IDReq) (*BaseResp, error)
	GeMenuParamListByMenuId(context.Context, *IDReq) (*MenuParamListResp, error)
	// role service
	CreateOrUpdateRole(context.Context, *RoleInfo) (*BaseResp, error)
	DeleteRole(context.Context, *IDReq) (*BaseResp, error)
	GetRoleById(context.Context, *IDReq) (*RoleInfo, error)
	GetRoleList(context.Context, *PageInfoReq) (*RoleListResp, error)
	SetRoleStatus(context.Context, *SetStatusReq) (*BaseResp, error)
	// api management service
	CreateOrUpdateApi(context.Context, *ApiInfo) (*BaseResp, error)
	DeleteApi(context.Context, *IDReq) (*BaseResp, error)
	GetApiList(context.Context, *ApiPageReq) (*ApiListResp, error)
	// authorization management service
	GetMenuAuthority(context.Context, *IDReq) (*RoleMenuAuthorityResp, error)
	CreateOrUpdateMenuAuthority(context.Context, *RoleMenuAuthorityReq) (*BaseResp, error)
	// example
	Hello(context.Context, *HelloReq) (*BaseResp, error)
	mustEmbedUnimplementedCoreServer()
}

// UnimplementedCoreServer must be embedded to have forward compatible implementations.
type UnimplementedCoreServer struct {
}

func (UnimplementedCoreServer) InitDatabase(context.Context, *Empty) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitDatabase not implemented")
}
func (UnimplementedCoreServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedCoreServer) ChangePassword(context.Context, *ChangePasswordReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedCoreServer) CreateOrUpdateUser(context.Context, *CreateOrUpdateUserReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateUser not implemented")
}
func (UnimplementedCoreServer) GetUserById(context.Context, *UUIDReq) (*UserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedCoreServer) GetUserList(context.Context, *GetUserListReq) (*UserListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedCoreServer) DeleteUser(context.Context, *IDReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedCoreServer) UpdateProfile(context.Context, *UpdateProfileReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedCoreServer) CreateOrUpdateMenu(context.Context, *CreateOrUpdateMenuReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateMenu not implemented")
}
func (UnimplementedCoreServer) DeleteMenu(context.Context, *IDReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMenu not implemented")
}
func (UnimplementedCoreServer) GetMenuListByRole(context.Context, *IDReq) (*MenuInfoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenuListByRole not implemented")
}
func (UnimplementedCoreServer) GetMenuByPage(context.Context, *PageInfoReq) (*MenuInfoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenuByPage not implemented")
}
func (UnimplementedCoreServer) CreateOrUpdateMenuParam(context.Context, *CreateOrUpdateMenuParamReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateMenuParam not implemented")
}
func (UnimplementedCoreServer) DeleteMenuParam(context.Context, *IDReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMenuParam not implemented")
}
func (UnimplementedCoreServer) GeMenuParamListByMenuId(context.Context, *IDReq) (*MenuParamListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeMenuParamListByMenuId not implemented")
}
func (UnimplementedCoreServer) CreateOrUpdateRole(context.Context, *RoleInfo) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateRole not implemented")
}
func (UnimplementedCoreServer) DeleteRole(context.Context, *IDReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedCoreServer) GetRoleById(context.Context, *IDReq) (*RoleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleById not implemented")
}
func (UnimplementedCoreServer) GetRoleList(context.Context, *PageInfoReq) (*RoleListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleList not implemented")
}
func (UnimplementedCoreServer) SetRoleStatus(context.Context, *SetStatusReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRoleStatus not implemented")
}
func (UnimplementedCoreServer) CreateOrUpdateApi(context.Context, *ApiInfo) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateApi not implemented")
}
func (UnimplementedCoreServer) DeleteApi(context.Context, *IDReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApi not implemented")
}
func (UnimplementedCoreServer) GetApiList(context.Context, *ApiPageReq) (*ApiListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiList not implemented")
}
func (UnimplementedCoreServer) GetMenuAuthority(context.Context, *IDReq) (*RoleMenuAuthorityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenuAuthority not implemented")
}
func (UnimplementedCoreServer) CreateOrUpdateMenuAuthority(context.Context, *RoleMenuAuthorityReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateMenuAuthority not implemented")
}
func (UnimplementedCoreServer) Hello(context.Context, *HelloReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedCoreServer) mustEmbedUnimplementedCoreServer() {}

// UnsafeCoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoreServer will
// result in compilation errors.
type UnsafeCoreServer interface {
	mustEmbedUnimplementedCoreServer()
}

func RegisterCoreServer(s grpc.ServiceRegistrar, srv CoreServer) {
	s.RegisterService(&Core_ServiceDesc, srv)
}

func _Core_InitDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).InitDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/initDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).InitDatabase(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/changePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ChangePassword(ctx, req.(*ChangePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_CreateOrUpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).CreateOrUpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/createOrUpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).CreateOrUpdateUser(ctx, req.(*CreateOrUpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/getUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetUserById(ctx, req.(*UUIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/getUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetUserList(ctx, req.(*GetUserListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/deleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).DeleteUser(ctx, req.(*IDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/updateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).UpdateProfile(ctx, req.(*UpdateProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_CreateOrUpdateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateMenuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).CreateOrUpdateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/createOrUpdateMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).CreateOrUpdateMenu(ctx, req.(*CreateOrUpdateMenuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_DeleteMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).DeleteMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/deleteMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).DeleteMenu(ctx, req.(*IDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetMenuListByRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetMenuListByRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/getMenuListByRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetMenuListByRole(ctx, req.(*IDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetMenuByPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetMenuByPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/getMenuByPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetMenuByPage(ctx, req.(*PageInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_CreateOrUpdateMenuParam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateMenuParamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).CreateOrUpdateMenuParam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/createOrUpdateMenuParam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).CreateOrUpdateMenuParam(ctx, req.(*CreateOrUpdateMenuParamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_DeleteMenuParam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).DeleteMenuParam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/deleteMenuParam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).DeleteMenuParam(ctx, req.(*IDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GeMenuParamListByMenuId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GeMenuParamListByMenuId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/geMenuParamListByMenuId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GeMenuParamListByMenuId(ctx, req.(*IDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_CreateOrUpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).CreateOrUpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/createOrUpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).CreateOrUpdateRole(ctx, req.(*RoleInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/deleteRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).DeleteRole(ctx, req.(*IDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetRoleById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetRoleById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/getRoleById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetRoleById(ctx, req.(*IDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetRoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetRoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/getRoleList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetRoleList(ctx, req.(*PageInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_SetRoleStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).SetRoleStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/setRoleStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).SetRoleStatus(ctx, req.(*SetStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_CreateOrUpdateApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).CreateOrUpdateApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/createOrUpdateApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).CreateOrUpdateApi(ctx, req.(*ApiInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_DeleteApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).DeleteApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/deleteApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).DeleteApi(ctx, req.(*IDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetApiList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetApiList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/getApiList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetApiList(ctx, req.(*ApiPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetMenuAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetMenuAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/getMenuAuthority",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetMenuAuthority(ctx, req.(*IDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_CreateOrUpdateMenuAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleMenuAuthorityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).CreateOrUpdateMenuAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/createOrUpdateMenuAuthority",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).CreateOrUpdateMenuAuthority(ctx, req.(*RoleMenuAuthorityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.core/hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).Hello(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Core_ServiceDesc is the grpc.ServiceDesc for Core service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Core_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.core",
	HandlerType: (*CoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "initDatabase",
			Handler:    _Core_InitDatabase_Handler,
		},
		{
			MethodName: "login",
			Handler:    _Core_Login_Handler,
		},
		{
			MethodName: "changePassword",
			Handler:    _Core_ChangePassword_Handler,
		},
		{
			MethodName: "createOrUpdateUser",
			Handler:    _Core_CreateOrUpdateUser_Handler,
		},
		{
			MethodName: "getUserById",
			Handler:    _Core_GetUserById_Handler,
		},
		{
			MethodName: "getUserList",
			Handler:    _Core_GetUserList_Handler,
		},
		{
			MethodName: "deleteUser",
			Handler:    _Core_DeleteUser_Handler,
		},
		{
			MethodName: "updateProfile",
			Handler:    _Core_UpdateProfile_Handler,
		},
		{
			MethodName: "createOrUpdateMenu",
			Handler:    _Core_CreateOrUpdateMenu_Handler,
		},
		{
			MethodName: "deleteMenu",
			Handler:    _Core_DeleteMenu_Handler,
		},
		{
			MethodName: "getMenuListByRole",
			Handler:    _Core_GetMenuListByRole_Handler,
		},
		{
			MethodName: "getMenuByPage",
			Handler:    _Core_GetMenuByPage_Handler,
		},
		{
			MethodName: "createOrUpdateMenuParam",
			Handler:    _Core_CreateOrUpdateMenuParam_Handler,
		},
		{
			MethodName: "deleteMenuParam",
			Handler:    _Core_DeleteMenuParam_Handler,
		},
		{
			MethodName: "geMenuParamListByMenuId",
			Handler:    _Core_GeMenuParamListByMenuId_Handler,
		},
		{
			MethodName: "createOrUpdateRole",
			Handler:    _Core_CreateOrUpdateRole_Handler,
		},
		{
			MethodName: "deleteRole",
			Handler:    _Core_DeleteRole_Handler,
		},
		{
			MethodName: "getRoleById",
			Handler:    _Core_GetRoleById_Handler,
		},
		{
			MethodName: "getRoleList",
			Handler:    _Core_GetRoleList_Handler,
		},
		{
			MethodName: "setRoleStatus",
			Handler:    _Core_SetRoleStatus_Handler,
		},
		{
			MethodName: "createOrUpdateApi",
			Handler:    _Core_CreateOrUpdateApi_Handler,
		},
		{
			MethodName: "deleteApi",
			Handler:    _Core_DeleteApi_Handler,
		},
		{
			MethodName: "getApiList",
			Handler:    _Core_GetApiList_Handler,
		},
		{
			MethodName: "getMenuAuthority",
			Handler:    _Core_GetMenuAuthority_Handler,
		},
		{
			MethodName: "createOrUpdateMenuAuthority",
			Handler:    _Core_CreateOrUpdateMenuAuthority_Handler,
		},
		{
			MethodName: "hello",
			Handler:    _Core_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.proto",
}
