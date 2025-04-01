// Code generated by thriftgo (0.3.18). DO NOT EDIT.

package user

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"wiliwili/kitex_gen/model"
)

type UserRegisterReq struct {
	Username  string `thrift:"username,1,required" frugal:"1,required,string" json:"username"`
	Password  string `thrift:"password,2,required" frugal:"2,required,string" json:"password"`
	Email     string `thrift:"email,3,required" frugal:"3,required,string" json:"email"`
	Gender    string `thrift:"gender,4,required" frugal:"4,required,string" json:"gender"`
	Signature string `thrift:"signature,5,required" frugal:"5,required,string" json:"signature"`
}

func NewUserRegisterReq() *UserRegisterReq {
	return &UserRegisterReq{}
}

func (p *UserRegisterReq) InitDefault() {
}

func (p *UserRegisterReq) GetUsername() (v string) {
	return p.Username
}

func (p *UserRegisterReq) GetPassword() (v string) {
	return p.Password
}

func (p *UserRegisterReq) GetEmail() (v string) {
	return p.Email
}

func (p *UserRegisterReq) GetGender() (v string) {
	return p.Gender
}

func (p *UserRegisterReq) GetSignature() (v string) {
	return p.Signature
}
func (p *UserRegisterReq) SetUsername(val string) {
	p.Username = val
}
func (p *UserRegisterReq) SetPassword(val string) {
	p.Password = val
}
func (p *UserRegisterReq) SetEmail(val string) {
	p.Email = val
}
func (p *UserRegisterReq) SetGender(val string) {
	p.Gender = val
}
func (p *UserRegisterReq) SetSignature(val string) {
	p.Signature = val
}

func (p *UserRegisterReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserRegisterReq(%+v)", *p)
}

func (p *UserRegisterReq) DeepEqual(ano *UserRegisterReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Username) {
		return false
	}
	if !p.Field2DeepEqual(ano.Password) {
		return false
	}
	if !p.Field3DeepEqual(ano.Email) {
		return false
	}
	if !p.Field4DeepEqual(ano.Gender) {
		return false
	}
	if !p.Field5DeepEqual(ano.Signature) {
		return false
	}
	return true
}

func (p *UserRegisterReq) Field1DeepEqual(src string) bool {

	if strings.Compare(p.Username, src) != 0 {
		return false
	}
	return true
}
func (p *UserRegisterReq) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Password, src) != 0 {
		return false
	}
	return true
}
func (p *UserRegisterReq) Field3DeepEqual(src string) bool {

	if strings.Compare(p.Email, src) != 0 {
		return false
	}
	return true
}
func (p *UserRegisterReq) Field4DeepEqual(src string) bool {

	if strings.Compare(p.Gender, src) != 0 {
		return false
	}
	return true
}
func (p *UserRegisterReq) Field5DeepEqual(src string) bool {

	if strings.Compare(p.Signature, src) != 0 {
		return false
	}
	return true
}

var fieldIDToName_UserRegisterReq = map[int16]string{
	1: "username",
	2: "password",
	3: "email",
	4: "gender",
	5: "signature",
}

type UserRegisterResp struct {
	BaseResp *model.BaseResp `thrift:"baseResp,1,required" frugal:"1,required,model.BaseResp" json:"baseResp"`
	Uid      int64           `thrift:"Uid,2,required" frugal:"2,required,i64" json:"Uid"`
}

func NewUserRegisterResp() *UserRegisterResp {
	return &UserRegisterResp{}
}

func (p *UserRegisterResp) InitDefault() {
}

var UserRegisterResp_BaseResp_DEFAULT *model.BaseResp

func (p *UserRegisterResp) GetBaseResp() (v *model.BaseResp) {
	if !p.IsSetBaseResp() {
		return UserRegisterResp_BaseResp_DEFAULT
	}
	return p.BaseResp
}

func (p *UserRegisterResp) GetUid() (v int64) {
	return p.Uid
}
func (p *UserRegisterResp) SetBaseResp(val *model.BaseResp) {
	p.BaseResp = val
}
func (p *UserRegisterResp) SetUid(val int64) {
	p.Uid = val
}

func (p *UserRegisterResp) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *UserRegisterResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserRegisterResp(%+v)", *p)
}

func (p *UserRegisterResp) DeepEqual(ano *UserRegisterResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.BaseResp) {
		return false
	}
	if !p.Field2DeepEqual(ano.Uid) {
		return false
	}
	return true
}

func (p *UserRegisterResp) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}
func (p *UserRegisterResp) Field2DeepEqual(src int64) bool {

	if p.Uid != src {
		return false
	}
	return true
}

var fieldIDToName_UserRegisterResp = map[int16]string{
	1: "baseResp",
	2: "Uid",
}

type UserLoginReq struct {
	Uid      int64  `thrift:"Uid,1,required" frugal:"1,required,i64" json:"Uid"`
	Password string `thrift:"password,2,required" frugal:"2,required,string" json:"password"`
}

func NewUserLoginReq() *UserLoginReq {
	return &UserLoginReq{}
}

func (p *UserLoginReq) InitDefault() {
}

func (p *UserLoginReq) GetUid() (v int64) {
	return p.Uid
}

func (p *UserLoginReq) GetPassword() (v string) {
	return p.Password
}
func (p *UserLoginReq) SetUid(val int64) {
	p.Uid = val
}
func (p *UserLoginReq) SetPassword(val string) {
	p.Password = val
}

func (p *UserLoginReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserLoginReq(%+v)", *p)
}

func (p *UserLoginReq) DeepEqual(ano *UserLoginReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Uid) {
		return false
	}
	if !p.Field2DeepEqual(ano.Password) {
		return false
	}
	return true
}

func (p *UserLoginReq) Field1DeepEqual(src int64) bool {

	if p.Uid != src {
		return false
	}
	return true
}
func (p *UserLoginReq) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Password, src) != 0 {
		return false
	}
	return true
}

var fieldIDToName_UserLoginReq = map[int16]string{
	1: "Uid",
	2: "password",
}

type UserLoginResp struct {
	BaseResp *model.BaseResp `thrift:"baseResp,1" frugal:"1,default,model.BaseResp" json:"baseResp"`
	UserInfo *model.UserInfo `thrift:"userInfo,2" frugal:"2,default,model.UserInfo" json:"userInfo"`
}

func NewUserLoginResp() *UserLoginResp {
	return &UserLoginResp{}
}

func (p *UserLoginResp) InitDefault() {
}

var UserLoginResp_BaseResp_DEFAULT *model.BaseResp

func (p *UserLoginResp) GetBaseResp() (v *model.BaseResp) {
	if !p.IsSetBaseResp() {
		return UserLoginResp_BaseResp_DEFAULT
	}
	return p.BaseResp
}

var UserLoginResp_UserInfo_DEFAULT *model.UserInfo

func (p *UserLoginResp) GetUserInfo() (v *model.UserInfo) {
	if !p.IsSetUserInfo() {
		return UserLoginResp_UserInfo_DEFAULT
	}
	return p.UserInfo
}
func (p *UserLoginResp) SetBaseResp(val *model.BaseResp) {
	p.BaseResp = val
}
func (p *UserLoginResp) SetUserInfo(val *model.UserInfo) {
	p.UserInfo = val
}

func (p *UserLoginResp) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *UserLoginResp) IsSetUserInfo() bool {
	return p.UserInfo != nil
}

func (p *UserLoginResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserLoginResp(%+v)", *p)
}

func (p *UserLoginResp) DeepEqual(ano *UserLoginResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.BaseResp) {
		return false
	}
	if !p.Field2DeepEqual(ano.UserInfo) {
		return false
	}
	return true
}

func (p *UserLoginResp) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}
func (p *UserLoginResp) Field2DeepEqual(src *model.UserInfo) bool {

	if !p.UserInfo.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_UserLoginResp = map[int16]string{
	1: "baseResp",
	2: "userInfo",
}

type UserProfileReq struct {
	Uid int64 `thrift:"Uid,1,required" frugal:"1,required,i64" json:"Uid"`
}

func NewUserProfileReq() *UserProfileReq {
	return &UserProfileReq{}
}

func (p *UserProfileReq) InitDefault() {
}

func (p *UserProfileReq) GetUid() (v int64) {
	return p.Uid
}
func (p *UserProfileReq) SetUid(val int64) {
	p.Uid = val
}

func (p *UserProfileReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserProfileReq(%+v)", *p)
}

func (p *UserProfileReq) DeepEqual(ano *UserProfileReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Uid) {
		return false
	}
	return true
}

func (p *UserProfileReq) Field1DeepEqual(src int64) bool {

	if p.Uid != src {
		return false
	}
	return true
}

var fieldIDToName_UserProfileReq = map[int16]string{
	1: "Uid",
}

type UserProfileResp struct {
	BaseResp    *model.BaseResp    `thrift:"baseResp,1" frugal:"1,default,model.BaseResp" json:"baseResp"`
	UserProfile *model.UserProfile `thrift:"userProfile,2" frugal:"2,default,model.UserProfile" json:"userProfile"`
}

func NewUserProfileResp() *UserProfileResp {
	return &UserProfileResp{}
}

func (p *UserProfileResp) InitDefault() {
}

var UserProfileResp_BaseResp_DEFAULT *model.BaseResp

func (p *UserProfileResp) GetBaseResp() (v *model.BaseResp) {
	if !p.IsSetBaseResp() {
		return UserProfileResp_BaseResp_DEFAULT
	}
	return p.BaseResp
}

var UserProfileResp_UserProfile_DEFAULT *model.UserProfile

func (p *UserProfileResp) GetUserProfile() (v *model.UserProfile) {
	if !p.IsSetUserProfile() {
		return UserProfileResp_UserProfile_DEFAULT
	}
	return p.UserProfile
}
func (p *UserProfileResp) SetBaseResp(val *model.BaseResp) {
	p.BaseResp = val
}
func (p *UserProfileResp) SetUserProfile(val *model.UserProfile) {
	p.UserProfile = val
}

func (p *UserProfileResp) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *UserProfileResp) IsSetUserProfile() bool {
	return p.UserProfile != nil
}

func (p *UserProfileResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserProfileResp(%+v)", *p)
}

func (p *UserProfileResp) DeepEqual(ano *UserProfileResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.BaseResp) {
		return false
	}
	if !p.Field2DeepEqual(ano.UserProfile) {
		return false
	}
	return true
}

func (p *UserProfileResp) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}
func (p *UserProfileResp) Field2DeepEqual(src *model.UserProfile) bool {

	if !p.UserProfile.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_UserProfileResp = map[int16]string{
	1: "baseResp",
	2: "userProfile",
}

type UserAvatarUploadReq struct {
	Avatar []byte `thrift:"avatar,1,required" frugal:"1,required,binary" json:"avatar"`
	Uid    int64  `thrift:"Uid,2,required" frugal:"2,required,i64" json:"Uid"`
}

func NewUserAvatarUploadReq() *UserAvatarUploadReq {
	return &UserAvatarUploadReq{}
}

func (p *UserAvatarUploadReq) InitDefault() {
}

func (p *UserAvatarUploadReq) GetAvatar() (v []byte) {
	return p.Avatar
}

func (p *UserAvatarUploadReq) GetUid() (v int64) {
	return p.Uid
}
func (p *UserAvatarUploadReq) SetAvatar(val []byte) {
	p.Avatar = val
}
func (p *UserAvatarUploadReq) SetUid(val int64) {
	p.Uid = val
}

func (p *UserAvatarUploadReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserAvatarUploadReq(%+v)", *p)
}

func (p *UserAvatarUploadReq) DeepEqual(ano *UserAvatarUploadReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Avatar) {
		return false
	}
	if !p.Field2DeepEqual(ano.Uid) {
		return false
	}
	return true
}

func (p *UserAvatarUploadReq) Field1DeepEqual(src []byte) bool {

	if bytes.Compare(p.Avatar, src) != 0 {
		return false
	}
	return true
}
func (p *UserAvatarUploadReq) Field2DeepEqual(src int64) bool {

	if p.Uid != src {
		return false
	}
	return true
}

var fieldIDToName_UserAvatarUploadReq = map[int16]string{
	1: "avatar",
	2: "Uid",
}

type UserAvatarUploadResp struct {
	BaseResp *model.BaseResp `thrift:"baseResp,1" frugal:"1,default,model.BaseResp" json:"baseResp"`
	Image    *model.Image    `thrift:"image,2" frugal:"2,default,model.Image" json:"image"`
}

func NewUserAvatarUploadResp() *UserAvatarUploadResp {
	return &UserAvatarUploadResp{}
}

func (p *UserAvatarUploadResp) InitDefault() {
}

var UserAvatarUploadResp_BaseResp_DEFAULT *model.BaseResp

func (p *UserAvatarUploadResp) GetBaseResp() (v *model.BaseResp) {
	if !p.IsSetBaseResp() {
		return UserAvatarUploadResp_BaseResp_DEFAULT
	}
	return p.BaseResp
}

var UserAvatarUploadResp_Image_DEFAULT *model.Image

func (p *UserAvatarUploadResp) GetImage() (v *model.Image) {
	if !p.IsSetImage() {
		return UserAvatarUploadResp_Image_DEFAULT
	}
	return p.Image
}
func (p *UserAvatarUploadResp) SetBaseResp(val *model.BaseResp) {
	p.BaseResp = val
}
func (p *UserAvatarUploadResp) SetImage(val *model.Image) {
	p.Image = val
}

func (p *UserAvatarUploadResp) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *UserAvatarUploadResp) IsSetImage() bool {
	return p.Image != nil
}

func (p *UserAvatarUploadResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserAvatarUploadResp(%+v)", *p)
}

func (p *UserAvatarUploadResp) DeepEqual(ano *UserAvatarUploadResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.BaseResp) {
		return false
	}
	if !p.Field2DeepEqual(ano.Image) {
		return false
	}
	return true
}

func (p *UserAvatarUploadResp) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}
func (p *UserAvatarUploadResp) Field2DeepEqual(src *model.Image) bool {

	if !p.Image.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_UserAvatarUploadResp = map[int16]string{
	1: "baseResp",
	2: "image",
}

type UserService interface {
	UserRegister(ctx context.Context, req *UserRegisterReq) (r *UserRegisterResp, err error)

	UserLogin(ctx context.Context, req *UserLoginReq) (r *UserLoginResp, err error)

	UserProfile(ctx context.Context, req *UserProfileReq) (r *UserProfileResp, err error)

	UserAvatarUpload(ctx context.Context, req *UserAvatarUploadReq) (r *UserAvatarUploadResp, err error)
}

type UserServiceUserRegisterArgs struct {
	Req *UserRegisterReq `thrift:"req,1" frugal:"1,default,UserRegisterReq" json:"req"`
}

func NewUserServiceUserRegisterArgs() *UserServiceUserRegisterArgs {
	return &UserServiceUserRegisterArgs{}
}

func (p *UserServiceUserRegisterArgs) InitDefault() {
}

var UserServiceUserRegisterArgs_Req_DEFAULT *UserRegisterReq

func (p *UserServiceUserRegisterArgs) GetReq() (v *UserRegisterReq) {
	if !p.IsSetReq() {
		return UserServiceUserRegisterArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *UserServiceUserRegisterArgs) SetReq(val *UserRegisterReq) {
	p.Req = val
}

func (p *UserServiceUserRegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UserServiceUserRegisterArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserServiceUserRegisterArgs(%+v)", *p)
}

func (p *UserServiceUserRegisterArgs) DeepEqual(ano *UserServiceUserRegisterArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *UserServiceUserRegisterArgs) Field1DeepEqual(src *UserRegisterReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_UserServiceUserRegisterArgs = map[int16]string{
	1: "req",
}

type UserServiceUserRegisterResult struct {
	Success *UserRegisterResp `thrift:"success,0,optional" frugal:"0,optional,UserRegisterResp" json:"success,omitempty"`
}

func NewUserServiceUserRegisterResult() *UserServiceUserRegisterResult {
	return &UserServiceUserRegisterResult{}
}

func (p *UserServiceUserRegisterResult) InitDefault() {
}

var UserServiceUserRegisterResult_Success_DEFAULT *UserRegisterResp

func (p *UserServiceUserRegisterResult) GetSuccess() (v *UserRegisterResp) {
	if !p.IsSetSuccess() {
		return UserServiceUserRegisterResult_Success_DEFAULT
	}
	return p.Success
}
func (p *UserServiceUserRegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*UserRegisterResp)
}

func (p *UserServiceUserRegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UserServiceUserRegisterResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserServiceUserRegisterResult(%+v)", *p)
}

func (p *UserServiceUserRegisterResult) DeepEqual(ano *UserServiceUserRegisterResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *UserServiceUserRegisterResult) Field0DeepEqual(src *UserRegisterResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_UserServiceUserRegisterResult = map[int16]string{
	0: "success",
}

type UserServiceUserLoginArgs struct {
	Req *UserLoginReq `thrift:"req,1" frugal:"1,default,UserLoginReq" json:"req"`
}

func NewUserServiceUserLoginArgs() *UserServiceUserLoginArgs {
	return &UserServiceUserLoginArgs{}
}

func (p *UserServiceUserLoginArgs) InitDefault() {
}

var UserServiceUserLoginArgs_Req_DEFAULT *UserLoginReq

func (p *UserServiceUserLoginArgs) GetReq() (v *UserLoginReq) {
	if !p.IsSetReq() {
		return UserServiceUserLoginArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *UserServiceUserLoginArgs) SetReq(val *UserLoginReq) {
	p.Req = val
}

func (p *UserServiceUserLoginArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UserServiceUserLoginArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserServiceUserLoginArgs(%+v)", *p)
}

func (p *UserServiceUserLoginArgs) DeepEqual(ano *UserServiceUserLoginArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *UserServiceUserLoginArgs) Field1DeepEqual(src *UserLoginReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_UserServiceUserLoginArgs = map[int16]string{
	1: "req",
}

type UserServiceUserLoginResult struct {
	Success *UserLoginResp `thrift:"success,0,optional" frugal:"0,optional,UserLoginResp" json:"success,omitempty"`
}

func NewUserServiceUserLoginResult() *UserServiceUserLoginResult {
	return &UserServiceUserLoginResult{}
}

func (p *UserServiceUserLoginResult) InitDefault() {
}

var UserServiceUserLoginResult_Success_DEFAULT *UserLoginResp

func (p *UserServiceUserLoginResult) GetSuccess() (v *UserLoginResp) {
	if !p.IsSetSuccess() {
		return UserServiceUserLoginResult_Success_DEFAULT
	}
	return p.Success
}
func (p *UserServiceUserLoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*UserLoginResp)
}

func (p *UserServiceUserLoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UserServiceUserLoginResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserServiceUserLoginResult(%+v)", *p)
}

func (p *UserServiceUserLoginResult) DeepEqual(ano *UserServiceUserLoginResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *UserServiceUserLoginResult) Field0DeepEqual(src *UserLoginResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_UserServiceUserLoginResult = map[int16]string{
	0: "success",
}

type UserServiceUserProfileArgs struct {
	Req *UserProfileReq `thrift:"req,1" frugal:"1,default,UserProfileReq" json:"req"`
}

func NewUserServiceUserProfileArgs() *UserServiceUserProfileArgs {
	return &UserServiceUserProfileArgs{}
}

func (p *UserServiceUserProfileArgs) InitDefault() {
}

var UserServiceUserProfileArgs_Req_DEFAULT *UserProfileReq

func (p *UserServiceUserProfileArgs) GetReq() (v *UserProfileReq) {
	if !p.IsSetReq() {
		return UserServiceUserProfileArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *UserServiceUserProfileArgs) SetReq(val *UserProfileReq) {
	p.Req = val
}

func (p *UserServiceUserProfileArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UserServiceUserProfileArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserServiceUserProfileArgs(%+v)", *p)
}

func (p *UserServiceUserProfileArgs) DeepEqual(ano *UserServiceUserProfileArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *UserServiceUserProfileArgs) Field1DeepEqual(src *UserProfileReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_UserServiceUserProfileArgs = map[int16]string{
	1: "req",
}

type UserServiceUserProfileResult struct {
	Success *UserProfileResp `thrift:"success,0,optional" frugal:"0,optional,UserProfileResp" json:"success,omitempty"`
}

func NewUserServiceUserProfileResult() *UserServiceUserProfileResult {
	return &UserServiceUserProfileResult{}
}

func (p *UserServiceUserProfileResult) InitDefault() {
}

var UserServiceUserProfileResult_Success_DEFAULT *UserProfileResp

func (p *UserServiceUserProfileResult) GetSuccess() (v *UserProfileResp) {
	if !p.IsSetSuccess() {
		return UserServiceUserProfileResult_Success_DEFAULT
	}
	return p.Success
}
func (p *UserServiceUserProfileResult) SetSuccess(x interface{}) {
	p.Success = x.(*UserProfileResp)
}

func (p *UserServiceUserProfileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UserServiceUserProfileResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserServiceUserProfileResult(%+v)", *p)
}

func (p *UserServiceUserProfileResult) DeepEqual(ano *UserServiceUserProfileResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *UserServiceUserProfileResult) Field0DeepEqual(src *UserProfileResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_UserServiceUserProfileResult = map[int16]string{
	0: "success",
}

type UserServiceUserAvatarUploadArgs struct {
	Req *UserAvatarUploadReq `thrift:"req,1" frugal:"1,default,UserAvatarUploadReq" json:"req"`
}

func NewUserServiceUserAvatarUploadArgs() *UserServiceUserAvatarUploadArgs {
	return &UserServiceUserAvatarUploadArgs{}
}

func (p *UserServiceUserAvatarUploadArgs) InitDefault() {
}

var UserServiceUserAvatarUploadArgs_Req_DEFAULT *UserAvatarUploadReq

func (p *UserServiceUserAvatarUploadArgs) GetReq() (v *UserAvatarUploadReq) {
	if !p.IsSetReq() {
		return UserServiceUserAvatarUploadArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *UserServiceUserAvatarUploadArgs) SetReq(val *UserAvatarUploadReq) {
	p.Req = val
}

func (p *UserServiceUserAvatarUploadArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UserServiceUserAvatarUploadArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserServiceUserAvatarUploadArgs(%+v)", *p)
}

func (p *UserServiceUserAvatarUploadArgs) DeepEqual(ano *UserServiceUserAvatarUploadArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *UserServiceUserAvatarUploadArgs) Field1DeepEqual(src *UserAvatarUploadReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_UserServiceUserAvatarUploadArgs = map[int16]string{
	1: "req",
}

type UserServiceUserAvatarUploadResult struct {
	Success *UserAvatarUploadResp `thrift:"success,0,optional" frugal:"0,optional,UserAvatarUploadResp" json:"success,omitempty"`
}

func NewUserServiceUserAvatarUploadResult() *UserServiceUserAvatarUploadResult {
	return &UserServiceUserAvatarUploadResult{}
}

func (p *UserServiceUserAvatarUploadResult) InitDefault() {
}

var UserServiceUserAvatarUploadResult_Success_DEFAULT *UserAvatarUploadResp

func (p *UserServiceUserAvatarUploadResult) GetSuccess() (v *UserAvatarUploadResp) {
	if !p.IsSetSuccess() {
		return UserServiceUserAvatarUploadResult_Success_DEFAULT
	}
	return p.Success
}
func (p *UserServiceUserAvatarUploadResult) SetSuccess(x interface{}) {
	p.Success = x.(*UserAvatarUploadResp)
}

func (p *UserServiceUserAvatarUploadResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UserServiceUserAvatarUploadResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserServiceUserAvatarUploadResult(%+v)", *p)
}

func (p *UserServiceUserAvatarUploadResult) DeepEqual(ano *UserServiceUserAvatarUploadResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *UserServiceUserAvatarUploadResult) Field0DeepEqual(src *UserAvatarUploadResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_UserServiceUserAvatarUploadResult = map[int16]string{
	0: "success",
}
