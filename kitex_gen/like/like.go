// Code generated by thriftgo (0.3.18). DO NOT EDIT.

package like

import (
	"context"
	"fmt"
	"strings"
	"wiliwili/kitex_gen/model"
)

type LikeCommentReq struct {
	CommentId string `thrift:"commentId,1,required" frugal:"1,required,string" json:"commentId"`
	IsLike    bool   `thrift:"IsLike,2,required" frugal:"2,required,bool" json:"IsLike"`
}

func NewLikeCommentReq() *LikeCommentReq {
	return &LikeCommentReq{}
}

func (p *LikeCommentReq) InitDefault() {
}

func (p *LikeCommentReq) GetCommentId() (v string) {
	return p.CommentId
}

func (p *LikeCommentReq) GetIsLike() (v bool) {
	return p.IsLike
}
func (p *LikeCommentReq) SetCommentId(val string) {
	p.CommentId = val
}
func (p *LikeCommentReq) SetIsLike(val bool) {
	p.IsLike = val
}

func (p *LikeCommentReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LikeCommentReq(%+v)", *p)
}

func (p *LikeCommentReq) DeepEqual(ano *LikeCommentReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.CommentId) {
		return false
	}
	if !p.Field2DeepEqual(ano.IsLike) {
		return false
	}
	return true
}

func (p *LikeCommentReq) Field1DeepEqual(src string) bool {

	if strings.Compare(p.CommentId, src) != 0 {
		return false
	}
	return true
}
func (p *LikeCommentReq) Field2DeepEqual(src bool) bool {

	if p.IsLike != src {
		return false
	}
	return true
}

var fieldIDToName_LikeCommentReq = map[int16]string{
	1: "commentId",
	2: "IsLike",
}

type LikeCommentResp struct {
	BaseResp *model.BaseResp `thrift:"baseResp,1,required" frugal:"1,required,model.BaseResp" json:"baseResp"`
}

func NewLikeCommentResp() *LikeCommentResp {
	return &LikeCommentResp{}
}

func (p *LikeCommentResp) InitDefault() {
}

var LikeCommentResp_BaseResp_DEFAULT *model.BaseResp

func (p *LikeCommentResp) GetBaseResp() (v *model.BaseResp) {
	if !p.IsSetBaseResp() {
		return LikeCommentResp_BaseResp_DEFAULT
	}
	return p.BaseResp
}
func (p *LikeCommentResp) SetBaseResp(val *model.BaseResp) {
	p.BaseResp = val
}

func (p *LikeCommentResp) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *LikeCommentResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LikeCommentResp(%+v)", *p)
}

func (p *LikeCommentResp) DeepEqual(ano *LikeCommentResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.BaseResp) {
		return false
	}
	return true
}

func (p *LikeCommentResp) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_LikeCommentResp = map[int16]string{
	1: "baseResp",
}

type LikeVideoReq struct {
	VideoId string `thrift:"videoId,1,required" frugal:"1,required,string" json:"videoId"`
	IsLike  bool   `thrift:"IsLike,2,required" frugal:"2,required,bool" json:"IsLike"`
}

func NewLikeVideoReq() *LikeVideoReq {
	return &LikeVideoReq{}
}

func (p *LikeVideoReq) InitDefault() {
}

func (p *LikeVideoReq) GetVideoId() (v string) {
	return p.VideoId
}

func (p *LikeVideoReq) GetIsLike() (v bool) {
	return p.IsLike
}
func (p *LikeVideoReq) SetVideoId(val string) {
	p.VideoId = val
}
func (p *LikeVideoReq) SetIsLike(val bool) {
	p.IsLike = val
}

func (p *LikeVideoReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LikeVideoReq(%+v)", *p)
}

func (p *LikeVideoReq) DeepEqual(ano *LikeVideoReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.VideoId) {
		return false
	}
	if !p.Field2DeepEqual(ano.IsLike) {
		return false
	}
	return true
}

func (p *LikeVideoReq) Field1DeepEqual(src string) bool {

	if strings.Compare(p.VideoId, src) != 0 {
		return false
	}
	return true
}
func (p *LikeVideoReq) Field2DeepEqual(src bool) bool {

	if p.IsLike != src {
		return false
	}
	return true
}

var fieldIDToName_LikeVideoReq = map[int16]string{
	1: "videoId",
	2: "IsLike",
}

type LikeVideoResp struct {
	BaseResp *model.BaseResp `thrift:"baseResp,1,required" frugal:"1,required,model.BaseResp" json:"baseResp"`
}

func NewLikeVideoResp() *LikeVideoResp {
	return &LikeVideoResp{}
}

func (p *LikeVideoResp) InitDefault() {
}

var LikeVideoResp_BaseResp_DEFAULT *model.BaseResp

func (p *LikeVideoResp) GetBaseResp() (v *model.BaseResp) {
	if !p.IsSetBaseResp() {
		return LikeVideoResp_BaseResp_DEFAULT
	}
	return p.BaseResp
}
func (p *LikeVideoResp) SetBaseResp(val *model.BaseResp) {
	p.BaseResp = val
}

func (p *LikeVideoResp) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *LikeVideoResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LikeVideoResp(%+v)", *p)
}

func (p *LikeVideoResp) DeepEqual(ano *LikeVideoResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.BaseResp) {
		return false
	}
	return true
}

func (p *LikeVideoResp) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_LikeVideoResp = map[int16]string{
	1: "baseResp",
}

type CommentLikeNumReq struct {
	CommentId string `thrift:"commentId,1,required" frugal:"1,required,string" json:"commentId"`
}

func NewCommentLikeNumReq() *CommentLikeNumReq {
	return &CommentLikeNumReq{}
}

func (p *CommentLikeNumReq) InitDefault() {
}

func (p *CommentLikeNumReq) GetCommentId() (v string) {
	return p.CommentId
}
func (p *CommentLikeNumReq) SetCommentId(val string) {
	p.CommentId = val
}

func (p *CommentLikeNumReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CommentLikeNumReq(%+v)", *p)
}

func (p *CommentLikeNumReq) DeepEqual(ano *CommentLikeNumReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.CommentId) {
		return false
	}
	return true
}

func (p *CommentLikeNumReq) Field1DeepEqual(src string) bool {

	if strings.Compare(p.CommentId, src) != 0 {
		return false
	}
	return true
}

var fieldIDToName_CommentLikeNumReq = map[int16]string{
	1: "commentId",
}

type CommentLikeNumResp struct {
	BaseResp   *model.BaseResp `thrift:"baseResp,1,required" frugal:"1,required,model.BaseResp" json:"baseResp"`
	TotalCount int64           `thrift:"totalCount,2,required" frugal:"2,required,i64" json:"totalCount"`
}

func NewCommentLikeNumResp() *CommentLikeNumResp {
	return &CommentLikeNumResp{}
}

func (p *CommentLikeNumResp) InitDefault() {
}

var CommentLikeNumResp_BaseResp_DEFAULT *model.BaseResp

func (p *CommentLikeNumResp) GetBaseResp() (v *model.BaseResp) {
	if !p.IsSetBaseResp() {
		return CommentLikeNumResp_BaseResp_DEFAULT
	}
	return p.BaseResp
}

func (p *CommentLikeNumResp) GetTotalCount() (v int64) {
	return p.TotalCount
}
func (p *CommentLikeNumResp) SetBaseResp(val *model.BaseResp) {
	p.BaseResp = val
}
func (p *CommentLikeNumResp) SetTotalCount(val int64) {
	p.TotalCount = val
}

func (p *CommentLikeNumResp) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *CommentLikeNumResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CommentLikeNumResp(%+v)", *p)
}

func (p *CommentLikeNumResp) DeepEqual(ano *CommentLikeNumResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.BaseResp) {
		return false
	}
	if !p.Field2DeepEqual(ano.TotalCount) {
		return false
	}
	return true
}

func (p *CommentLikeNumResp) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}
func (p *CommentLikeNumResp) Field2DeepEqual(src int64) bool {

	if p.TotalCount != src {
		return false
	}
	return true
}

var fieldIDToName_CommentLikeNumResp = map[int16]string{
	1: "baseResp",
	2: "totalCount",
}

type VideoLikeNumReq struct {
	VideoId string `thrift:"videoId,1,required" frugal:"1,required,string" json:"videoId"`
}

func NewVideoLikeNumReq() *VideoLikeNumReq {
	return &VideoLikeNumReq{}
}

func (p *VideoLikeNumReq) InitDefault() {
}

func (p *VideoLikeNumReq) GetVideoId() (v string) {
	return p.VideoId
}
func (p *VideoLikeNumReq) SetVideoId(val string) {
	p.VideoId = val
}

func (p *VideoLikeNumReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VideoLikeNumReq(%+v)", *p)
}

func (p *VideoLikeNumReq) DeepEqual(ano *VideoLikeNumReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.VideoId) {
		return false
	}
	return true
}

func (p *VideoLikeNumReq) Field1DeepEqual(src string) bool {

	if strings.Compare(p.VideoId, src) != 0 {
		return false
	}
	return true
}

var fieldIDToName_VideoLikeNumReq = map[int16]string{
	1: "videoId",
}

type VideoLikeNumResp struct {
	BaseResp   *model.BaseResp `thrift:"baseResp,1,required" frugal:"1,required,model.BaseResp" json:"baseResp"`
	TotalCount int64           `thrift:"totalCount,2,required" frugal:"2,required,i64" json:"totalCount"`
}

func NewVideoLikeNumResp() *VideoLikeNumResp {
	return &VideoLikeNumResp{}
}

func (p *VideoLikeNumResp) InitDefault() {
}

var VideoLikeNumResp_BaseResp_DEFAULT *model.BaseResp

func (p *VideoLikeNumResp) GetBaseResp() (v *model.BaseResp) {
	if !p.IsSetBaseResp() {
		return VideoLikeNumResp_BaseResp_DEFAULT
	}
	return p.BaseResp
}

func (p *VideoLikeNumResp) GetTotalCount() (v int64) {
	return p.TotalCount
}
func (p *VideoLikeNumResp) SetBaseResp(val *model.BaseResp) {
	p.BaseResp = val
}
func (p *VideoLikeNumResp) SetTotalCount(val int64) {
	p.TotalCount = val
}

func (p *VideoLikeNumResp) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *VideoLikeNumResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VideoLikeNumResp(%+v)", *p)
}

func (p *VideoLikeNumResp) DeepEqual(ano *VideoLikeNumResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.BaseResp) {
		return false
	}
	if !p.Field2DeepEqual(ano.TotalCount) {
		return false
	}
	return true
}

func (p *VideoLikeNumResp) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}
func (p *VideoLikeNumResp) Field2DeepEqual(src int64) bool {

	if p.TotalCount != src {
		return false
	}
	return true
}

var fieldIDToName_VideoLikeNumResp = map[int16]string{
	1: "baseResp",
	2: "totalCount",
}

type LikeService interface {
	LikeComment(ctx context.Context, req *LikeCommentReq) (r *LikeCommentResp, err error)

	LikeVideo(ctx context.Context, req *LikeVideoReq) (r *LikeVideoResp, err error)

	CommentLikeNum(ctx context.Context, req *CommentLikeNumReq) (r *CommentLikeNumResp, err error)

	VideoLikeNum(ctx context.Context, req *VideoLikeNumReq) (r *VideoLikeNumResp, err error)
}

type LikeServiceLikeCommentArgs struct {
	Req *LikeCommentReq `thrift:"req,1" frugal:"1,default,LikeCommentReq" json:"req"`
}

func NewLikeServiceLikeCommentArgs() *LikeServiceLikeCommentArgs {
	return &LikeServiceLikeCommentArgs{}
}

func (p *LikeServiceLikeCommentArgs) InitDefault() {
}

var LikeServiceLikeCommentArgs_Req_DEFAULT *LikeCommentReq

func (p *LikeServiceLikeCommentArgs) GetReq() (v *LikeCommentReq) {
	if !p.IsSetReq() {
		return LikeServiceLikeCommentArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *LikeServiceLikeCommentArgs) SetReq(val *LikeCommentReq) {
	p.Req = val
}

func (p *LikeServiceLikeCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *LikeServiceLikeCommentArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LikeServiceLikeCommentArgs(%+v)", *p)
}

func (p *LikeServiceLikeCommentArgs) DeepEqual(ano *LikeServiceLikeCommentArgs) bool {
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

func (p *LikeServiceLikeCommentArgs) Field1DeepEqual(src *LikeCommentReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_LikeServiceLikeCommentArgs = map[int16]string{
	1: "req",
}

type LikeServiceLikeCommentResult struct {
	Success *LikeCommentResp `thrift:"success,0,optional" frugal:"0,optional,LikeCommentResp" json:"success,omitempty"`
}

func NewLikeServiceLikeCommentResult() *LikeServiceLikeCommentResult {
	return &LikeServiceLikeCommentResult{}
}

func (p *LikeServiceLikeCommentResult) InitDefault() {
}

var LikeServiceLikeCommentResult_Success_DEFAULT *LikeCommentResp

func (p *LikeServiceLikeCommentResult) GetSuccess() (v *LikeCommentResp) {
	if !p.IsSetSuccess() {
		return LikeServiceLikeCommentResult_Success_DEFAULT
	}
	return p.Success
}
func (p *LikeServiceLikeCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*LikeCommentResp)
}

func (p *LikeServiceLikeCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *LikeServiceLikeCommentResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LikeServiceLikeCommentResult(%+v)", *p)
}

func (p *LikeServiceLikeCommentResult) DeepEqual(ano *LikeServiceLikeCommentResult) bool {
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

func (p *LikeServiceLikeCommentResult) Field0DeepEqual(src *LikeCommentResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_LikeServiceLikeCommentResult = map[int16]string{
	0: "success",
}

type LikeServiceLikeVideoArgs struct {
	Req *LikeVideoReq `thrift:"req,1" frugal:"1,default,LikeVideoReq" json:"req"`
}

func NewLikeServiceLikeVideoArgs() *LikeServiceLikeVideoArgs {
	return &LikeServiceLikeVideoArgs{}
}

func (p *LikeServiceLikeVideoArgs) InitDefault() {
}

var LikeServiceLikeVideoArgs_Req_DEFAULT *LikeVideoReq

func (p *LikeServiceLikeVideoArgs) GetReq() (v *LikeVideoReq) {
	if !p.IsSetReq() {
		return LikeServiceLikeVideoArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *LikeServiceLikeVideoArgs) SetReq(val *LikeVideoReq) {
	p.Req = val
}

func (p *LikeServiceLikeVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *LikeServiceLikeVideoArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LikeServiceLikeVideoArgs(%+v)", *p)
}

func (p *LikeServiceLikeVideoArgs) DeepEqual(ano *LikeServiceLikeVideoArgs) bool {
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

func (p *LikeServiceLikeVideoArgs) Field1DeepEqual(src *LikeVideoReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_LikeServiceLikeVideoArgs = map[int16]string{
	1: "req",
}

type LikeServiceLikeVideoResult struct {
	Success *LikeVideoResp `thrift:"success,0,optional" frugal:"0,optional,LikeVideoResp" json:"success,omitempty"`
}

func NewLikeServiceLikeVideoResult() *LikeServiceLikeVideoResult {
	return &LikeServiceLikeVideoResult{}
}

func (p *LikeServiceLikeVideoResult) InitDefault() {
}

var LikeServiceLikeVideoResult_Success_DEFAULT *LikeVideoResp

func (p *LikeServiceLikeVideoResult) GetSuccess() (v *LikeVideoResp) {
	if !p.IsSetSuccess() {
		return LikeServiceLikeVideoResult_Success_DEFAULT
	}
	return p.Success
}
func (p *LikeServiceLikeVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*LikeVideoResp)
}

func (p *LikeServiceLikeVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *LikeServiceLikeVideoResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LikeServiceLikeVideoResult(%+v)", *p)
}

func (p *LikeServiceLikeVideoResult) DeepEqual(ano *LikeServiceLikeVideoResult) bool {
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

func (p *LikeServiceLikeVideoResult) Field0DeepEqual(src *LikeVideoResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_LikeServiceLikeVideoResult = map[int16]string{
	0: "success",
}

type LikeServiceCommentLikeNumArgs struct {
	Req *CommentLikeNumReq `thrift:"req,1" frugal:"1,default,CommentLikeNumReq" json:"req"`
}

func NewLikeServiceCommentLikeNumArgs() *LikeServiceCommentLikeNumArgs {
	return &LikeServiceCommentLikeNumArgs{}
}

func (p *LikeServiceCommentLikeNumArgs) InitDefault() {
}

var LikeServiceCommentLikeNumArgs_Req_DEFAULT *CommentLikeNumReq

func (p *LikeServiceCommentLikeNumArgs) GetReq() (v *CommentLikeNumReq) {
	if !p.IsSetReq() {
		return LikeServiceCommentLikeNumArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *LikeServiceCommentLikeNumArgs) SetReq(val *CommentLikeNumReq) {
	p.Req = val
}

func (p *LikeServiceCommentLikeNumArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *LikeServiceCommentLikeNumArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LikeServiceCommentLikeNumArgs(%+v)", *p)
}

func (p *LikeServiceCommentLikeNumArgs) DeepEqual(ano *LikeServiceCommentLikeNumArgs) bool {
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

func (p *LikeServiceCommentLikeNumArgs) Field1DeepEqual(src *CommentLikeNumReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_LikeServiceCommentLikeNumArgs = map[int16]string{
	1: "req",
}

type LikeServiceCommentLikeNumResult struct {
	Success *CommentLikeNumResp `thrift:"success,0,optional" frugal:"0,optional,CommentLikeNumResp" json:"success,omitempty"`
}

func NewLikeServiceCommentLikeNumResult() *LikeServiceCommentLikeNumResult {
	return &LikeServiceCommentLikeNumResult{}
}

func (p *LikeServiceCommentLikeNumResult) InitDefault() {
}

var LikeServiceCommentLikeNumResult_Success_DEFAULT *CommentLikeNumResp

func (p *LikeServiceCommentLikeNumResult) GetSuccess() (v *CommentLikeNumResp) {
	if !p.IsSetSuccess() {
		return LikeServiceCommentLikeNumResult_Success_DEFAULT
	}
	return p.Success
}
func (p *LikeServiceCommentLikeNumResult) SetSuccess(x interface{}) {
	p.Success = x.(*CommentLikeNumResp)
}

func (p *LikeServiceCommentLikeNumResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *LikeServiceCommentLikeNumResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LikeServiceCommentLikeNumResult(%+v)", *p)
}

func (p *LikeServiceCommentLikeNumResult) DeepEqual(ano *LikeServiceCommentLikeNumResult) bool {
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

func (p *LikeServiceCommentLikeNumResult) Field0DeepEqual(src *CommentLikeNumResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_LikeServiceCommentLikeNumResult = map[int16]string{
	0: "success",
}

type LikeServiceVideoLikeNumArgs struct {
	Req *VideoLikeNumReq `thrift:"req,1" frugal:"1,default,VideoLikeNumReq" json:"req"`
}

func NewLikeServiceVideoLikeNumArgs() *LikeServiceVideoLikeNumArgs {
	return &LikeServiceVideoLikeNumArgs{}
}

func (p *LikeServiceVideoLikeNumArgs) InitDefault() {
}

var LikeServiceVideoLikeNumArgs_Req_DEFAULT *VideoLikeNumReq

func (p *LikeServiceVideoLikeNumArgs) GetReq() (v *VideoLikeNumReq) {
	if !p.IsSetReq() {
		return LikeServiceVideoLikeNumArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *LikeServiceVideoLikeNumArgs) SetReq(val *VideoLikeNumReq) {
	p.Req = val
}

func (p *LikeServiceVideoLikeNumArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *LikeServiceVideoLikeNumArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LikeServiceVideoLikeNumArgs(%+v)", *p)
}

func (p *LikeServiceVideoLikeNumArgs) DeepEqual(ano *LikeServiceVideoLikeNumArgs) bool {
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

func (p *LikeServiceVideoLikeNumArgs) Field1DeepEqual(src *VideoLikeNumReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_LikeServiceVideoLikeNumArgs = map[int16]string{
	1: "req",
}

type LikeServiceVideoLikeNumResult struct {
	Success *VideoLikeNumResp `thrift:"success,0,optional" frugal:"0,optional,VideoLikeNumResp" json:"success,omitempty"`
}

func NewLikeServiceVideoLikeNumResult() *LikeServiceVideoLikeNumResult {
	return &LikeServiceVideoLikeNumResult{}
}

func (p *LikeServiceVideoLikeNumResult) InitDefault() {
}

var LikeServiceVideoLikeNumResult_Success_DEFAULT *VideoLikeNumResp

func (p *LikeServiceVideoLikeNumResult) GetSuccess() (v *VideoLikeNumResp) {
	if !p.IsSetSuccess() {
		return LikeServiceVideoLikeNumResult_Success_DEFAULT
	}
	return p.Success
}
func (p *LikeServiceVideoLikeNumResult) SetSuccess(x interface{}) {
	p.Success = x.(*VideoLikeNumResp)
}

func (p *LikeServiceVideoLikeNumResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *LikeServiceVideoLikeNumResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LikeServiceVideoLikeNumResult(%+v)", *p)
}

func (p *LikeServiceVideoLikeNumResult) DeepEqual(ano *LikeServiceVideoLikeNumResult) bool {
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

func (p *LikeServiceVideoLikeNumResult) Field0DeepEqual(src *VideoLikeNumResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_LikeServiceVideoLikeNumResult = map[int16]string{
	0: "success",
}
