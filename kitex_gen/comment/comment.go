// Code generated by thriftgo (0.3.18). DO NOT EDIT.

package comment

import (
	"context"
	"fmt"
	"strings"
	"wiliwili/kitex_gen/model"
)

type CommentVideoReq struct {
	VideoId string `thrift:"videoId,1,required" frugal:"1,required,string" json:"videoId"`
	Content string `thrift:"content,2,required" frugal:"2,required,string" json:"content"`
}

func NewCommentVideoReq() *CommentVideoReq {
	return &CommentVideoReq{}
}

func (p *CommentVideoReq) InitDefault() {
}

func (p *CommentVideoReq) GetVideoId() (v string) {
	return p.VideoId
}

func (p *CommentVideoReq) GetContent() (v string) {
	return p.Content
}
func (p *CommentVideoReq) SetVideoId(val string) {
	p.VideoId = val
}
func (p *CommentVideoReq) SetContent(val string) {
	p.Content = val
}

func (p *CommentVideoReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CommentVideoReq(%+v)", *p)
}

func (p *CommentVideoReq) DeepEqual(ano *CommentVideoReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.VideoId) {
		return false
	}
	if !p.Field2DeepEqual(ano.Content) {
		return false
	}
	return true
}

func (p *CommentVideoReq) Field1DeepEqual(src string) bool {

	if strings.Compare(p.VideoId, src) != 0 {
		return false
	}
	return true
}
func (p *CommentVideoReq) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Content, src) != 0 {
		return false
	}
	return true
}

var fieldIDToName_CommentVideoReq = map[int16]string{
	1: "videoId",
	2: "content",
}

type CommentVideoResp struct {
	BaseResp  *model.BaseResp `thrift:"baseResp,1,required" frugal:"1,required,model.BaseResp" json:"baseResp"`
	CommentId string          `thrift:"commentId,2,required" frugal:"2,required,string" json:"commentId"`
}

func NewCommentVideoResp() *CommentVideoResp {
	return &CommentVideoResp{}
}

func (p *CommentVideoResp) InitDefault() {
}

var CommentVideoResp_BaseResp_DEFAULT *model.BaseResp

func (p *CommentVideoResp) GetBaseResp() (v *model.BaseResp) {
	if !p.IsSetBaseResp() {
		return CommentVideoResp_BaseResp_DEFAULT
	}
	return p.BaseResp
}

func (p *CommentVideoResp) GetCommentId() (v string) {
	return p.CommentId
}
func (p *CommentVideoResp) SetBaseResp(val *model.BaseResp) {
	p.BaseResp = val
}
func (p *CommentVideoResp) SetCommentId(val string) {
	p.CommentId = val
}

func (p *CommentVideoResp) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *CommentVideoResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CommentVideoResp(%+v)", *p)
}

func (p *CommentVideoResp) DeepEqual(ano *CommentVideoResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.BaseResp) {
		return false
	}
	if !p.Field2DeepEqual(ano.CommentId) {
		return false
	}
	return true
}

func (p *CommentVideoResp) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}
func (p *CommentVideoResp) Field2DeepEqual(src string) bool {

	if strings.Compare(p.CommentId, src) != 0 {
		return false
	}
	return true
}

var fieldIDToName_CommentVideoResp = map[int16]string{
	1: "baseResp",
	2: "commentId",
}

type ReplyCommentReq struct {
	CommentId string `thrift:"commentId,1,required" frugal:"1,required,string" json:"commentId"`
	Content   string `thrift:"content,2,required" frugal:"2,required,string" json:"content"`
}

func NewReplyCommentReq() *ReplyCommentReq {
	return &ReplyCommentReq{}
}

func (p *ReplyCommentReq) InitDefault() {
}

func (p *ReplyCommentReq) GetCommentId() (v string) {
	return p.CommentId
}

func (p *ReplyCommentReq) GetContent() (v string) {
	return p.Content
}
func (p *ReplyCommentReq) SetCommentId(val string) {
	p.CommentId = val
}
func (p *ReplyCommentReq) SetContent(val string) {
	p.Content = val
}

func (p *ReplyCommentReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ReplyCommentReq(%+v)", *p)
}

func (p *ReplyCommentReq) DeepEqual(ano *ReplyCommentReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.CommentId) {
		return false
	}
	if !p.Field2DeepEqual(ano.Content) {
		return false
	}
	return true
}

func (p *ReplyCommentReq) Field1DeepEqual(src string) bool {

	if strings.Compare(p.CommentId, src) != 0 {
		return false
	}
	return true
}
func (p *ReplyCommentReq) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Content, src) != 0 {
		return false
	}
	return true
}

var fieldIDToName_ReplyCommentReq = map[int16]string{
	1: "commentId",
	2: "content",
}

type ReplyCommentResp struct {
	BaseResp  *model.BaseResp `thrift:"baseResp,1,required" frugal:"1,required,model.BaseResp" json:"baseResp"`
	CommentId string          `thrift:"commentId,2,required" frugal:"2,required,string" json:"commentId"`
}

func NewReplyCommentResp() *ReplyCommentResp {
	return &ReplyCommentResp{}
}

func (p *ReplyCommentResp) InitDefault() {
}

var ReplyCommentResp_BaseResp_DEFAULT *model.BaseResp

func (p *ReplyCommentResp) GetBaseResp() (v *model.BaseResp) {
	if !p.IsSetBaseResp() {
		return ReplyCommentResp_BaseResp_DEFAULT
	}
	return p.BaseResp
}

func (p *ReplyCommentResp) GetCommentId() (v string) {
	return p.CommentId
}
func (p *ReplyCommentResp) SetBaseResp(val *model.BaseResp) {
	p.BaseResp = val
}
func (p *ReplyCommentResp) SetCommentId(val string) {
	p.CommentId = val
}

func (p *ReplyCommentResp) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *ReplyCommentResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ReplyCommentResp(%+v)", *p)
}

func (p *ReplyCommentResp) DeepEqual(ano *ReplyCommentResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.BaseResp) {
		return false
	}
	if !p.Field2DeepEqual(ano.CommentId) {
		return false
	}
	return true
}

func (p *ReplyCommentResp) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}
func (p *ReplyCommentResp) Field2DeepEqual(src string) bool {

	if strings.Compare(p.CommentId, src) != 0 {
		return false
	}
	return true
}

var fieldIDToName_ReplyCommentResp = map[int16]string{
	1: "baseResp",
	2: "commentId",
}

type GetCommentListReq struct {
	Id          string `thrift:"Id,1,required" frugal:"1,required,string" json:"Id"`
	Page        int64  `thrift:"page,2,required" frugal:"2,required,i64" json:"page"`
	PageSize    int64  `thrift:"pageSize,3,required" frugal:"3,required,i64" json:"pageSize"`
	CommentTpye int64  `thrift:"CommentTpye,4,required" frugal:"4,required,i64" json:"CommentTpye"`
}

func NewGetCommentListReq() *GetCommentListReq {
	return &GetCommentListReq{}
}

func (p *GetCommentListReq) InitDefault() {
}

func (p *GetCommentListReq) GetId() (v string) {
	return p.Id
}

func (p *GetCommentListReq) GetPage() (v int64) {
	return p.Page
}

func (p *GetCommentListReq) GetPageSize() (v int64) {
	return p.PageSize
}

func (p *GetCommentListReq) GetCommentTpye() (v int64) {
	return p.CommentTpye
}
func (p *GetCommentListReq) SetId(val string) {
	p.Id = val
}
func (p *GetCommentListReq) SetPage(val int64) {
	p.Page = val
}
func (p *GetCommentListReq) SetPageSize(val int64) {
	p.PageSize = val
}
func (p *GetCommentListReq) SetCommentTpye(val int64) {
	p.CommentTpye = val
}

func (p *GetCommentListReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetCommentListReq(%+v)", *p)
}

func (p *GetCommentListReq) DeepEqual(ano *GetCommentListReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Id) {
		return false
	}
	if !p.Field2DeepEqual(ano.Page) {
		return false
	}
	if !p.Field3DeepEqual(ano.PageSize) {
		return false
	}
	if !p.Field4DeepEqual(ano.CommentTpye) {
		return false
	}
	return true
}

func (p *GetCommentListReq) Field1DeepEqual(src string) bool {

	if strings.Compare(p.Id, src) != 0 {
		return false
	}
	return true
}
func (p *GetCommentListReq) Field2DeepEqual(src int64) bool {

	if p.Page != src {
		return false
	}
	return true
}
func (p *GetCommentListReq) Field3DeepEqual(src int64) bool {

	if p.PageSize != src {
		return false
	}
	return true
}
func (p *GetCommentListReq) Field4DeepEqual(src int64) bool {

	if p.CommentTpye != src {
		return false
	}
	return true
}

var fieldIDToName_GetCommentListReq = map[int16]string{
	1: "Id",
	2: "page",
	3: "pageSize",
	4: "CommentTpye",
}

type GetCommentListResp struct {
	BaseResp    *model.BaseResp  `thrift:"baseResp,1,required" frugal:"1,required,model.BaseResp" json:"baseResp"`
	CommentList []*model.Comment `thrift:"commentList,2,required" frugal:"2,required,list<model.Comment>" json:"commentList"`
}

func NewGetCommentListResp() *GetCommentListResp {
	return &GetCommentListResp{}
}

func (p *GetCommentListResp) InitDefault() {
}

var GetCommentListResp_BaseResp_DEFAULT *model.BaseResp

func (p *GetCommentListResp) GetBaseResp() (v *model.BaseResp) {
	if !p.IsSetBaseResp() {
		return GetCommentListResp_BaseResp_DEFAULT
	}
	return p.BaseResp
}

func (p *GetCommentListResp) GetCommentList() (v []*model.Comment) {
	return p.CommentList
}
func (p *GetCommentListResp) SetBaseResp(val *model.BaseResp) {
	p.BaseResp = val
}
func (p *GetCommentListResp) SetCommentList(val []*model.Comment) {
	p.CommentList = val
}

func (p *GetCommentListResp) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *GetCommentListResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetCommentListResp(%+v)", *p)
}

func (p *GetCommentListResp) DeepEqual(ano *GetCommentListResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.BaseResp) {
		return false
	}
	if !p.Field2DeepEqual(ano.CommentList) {
		return false
	}
	return true
}

func (p *GetCommentListResp) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}
func (p *GetCommentListResp) Field2DeepEqual(src []*model.Comment) bool {

	if len(p.CommentList) != len(src) {
		return false
	}
	for i, v := range p.CommentList {
		_src := src[i]
		if !v.DeepEqual(_src) {
			return false
		}
	}
	return true
}

var fieldIDToName_GetCommentListResp = map[int16]string{
	1: "baseResp",
	2: "commentList",
}

type DeleteCommentReq struct {
	CommentId string `thrift:"commentId,1,required" frugal:"1,required,string" json:"commentId"`
}

func NewDeleteCommentReq() *DeleteCommentReq {
	return &DeleteCommentReq{}
}

func (p *DeleteCommentReq) InitDefault() {
}

func (p *DeleteCommentReq) GetCommentId() (v string) {
	return p.CommentId
}
func (p *DeleteCommentReq) SetCommentId(val string) {
	p.CommentId = val
}

func (p *DeleteCommentReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DeleteCommentReq(%+v)", *p)
}

func (p *DeleteCommentReq) DeepEqual(ano *DeleteCommentReq) bool {
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

func (p *DeleteCommentReq) Field1DeepEqual(src string) bool {

	if strings.Compare(p.CommentId, src) != 0 {
		return false
	}
	return true
}

var fieldIDToName_DeleteCommentReq = map[int16]string{
	1: "commentId",
}

type DeleteCommentResp struct {
	BaseResp *model.BaseResp `thrift:"baseResp,1,required" frugal:"1,required,model.BaseResp" json:"baseResp"`
}

func NewDeleteCommentResp() *DeleteCommentResp {
	return &DeleteCommentResp{}
}

func (p *DeleteCommentResp) InitDefault() {
}

var DeleteCommentResp_BaseResp_DEFAULT *model.BaseResp

func (p *DeleteCommentResp) GetBaseResp() (v *model.BaseResp) {
	if !p.IsSetBaseResp() {
		return DeleteCommentResp_BaseResp_DEFAULT
	}
	return p.BaseResp
}
func (p *DeleteCommentResp) SetBaseResp(val *model.BaseResp) {
	p.BaseResp = val
}

func (p *DeleteCommentResp) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *DeleteCommentResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DeleteCommentResp(%+v)", *p)
}

func (p *DeleteCommentResp) DeepEqual(ano *DeleteCommentResp) bool {
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

func (p *DeleteCommentResp) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_DeleteCommentResp = map[int16]string{
	1: "baseResp",
}

type CommentService interface {
	CommentVideo(ctx context.Context, req *CommentVideoReq) (r *CommentVideoResp, err error)

	ReplyComment(ctx context.Context, req *ReplyCommentReq) (r *ReplyCommentResp, err error)

	GetCommentList(ctx context.Context, req *GetCommentListReq) (r *GetCommentListResp, err error)

	DeleteComment(ctx context.Context, req *DeleteCommentReq) (r *DeleteCommentResp, err error)
}

type CommentServiceCommentVideoArgs struct {
	Req *CommentVideoReq `thrift:"req,1" frugal:"1,default,CommentVideoReq" json:"req"`
}

func NewCommentServiceCommentVideoArgs() *CommentServiceCommentVideoArgs {
	return &CommentServiceCommentVideoArgs{}
}

func (p *CommentServiceCommentVideoArgs) InitDefault() {
}

var CommentServiceCommentVideoArgs_Req_DEFAULT *CommentVideoReq

func (p *CommentServiceCommentVideoArgs) GetReq() (v *CommentVideoReq) {
	if !p.IsSetReq() {
		return CommentServiceCommentVideoArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *CommentServiceCommentVideoArgs) SetReq(val *CommentVideoReq) {
	p.Req = val
}

func (p *CommentServiceCommentVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CommentServiceCommentVideoArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CommentServiceCommentVideoArgs(%+v)", *p)
}

func (p *CommentServiceCommentVideoArgs) DeepEqual(ano *CommentServiceCommentVideoArgs) bool {
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

func (p *CommentServiceCommentVideoArgs) Field1DeepEqual(src *CommentVideoReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_CommentServiceCommentVideoArgs = map[int16]string{
	1: "req",
}

type CommentServiceCommentVideoResult struct {
	Success *CommentVideoResp `thrift:"success,0,optional" frugal:"0,optional,CommentVideoResp" json:"success,omitempty"`
}

func NewCommentServiceCommentVideoResult() *CommentServiceCommentVideoResult {
	return &CommentServiceCommentVideoResult{}
}

func (p *CommentServiceCommentVideoResult) InitDefault() {
}

var CommentServiceCommentVideoResult_Success_DEFAULT *CommentVideoResp

func (p *CommentServiceCommentVideoResult) GetSuccess() (v *CommentVideoResp) {
	if !p.IsSetSuccess() {
		return CommentServiceCommentVideoResult_Success_DEFAULT
	}
	return p.Success
}
func (p *CommentServiceCommentVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*CommentVideoResp)
}

func (p *CommentServiceCommentVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CommentServiceCommentVideoResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CommentServiceCommentVideoResult(%+v)", *p)
}

func (p *CommentServiceCommentVideoResult) DeepEqual(ano *CommentServiceCommentVideoResult) bool {
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

func (p *CommentServiceCommentVideoResult) Field0DeepEqual(src *CommentVideoResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_CommentServiceCommentVideoResult = map[int16]string{
	0: "success",
}

type CommentServiceReplyCommentArgs struct {
	Req *ReplyCommentReq `thrift:"req,1" frugal:"1,default,ReplyCommentReq" json:"req"`
}

func NewCommentServiceReplyCommentArgs() *CommentServiceReplyCommentArgs {
	return &CommentServiceReplyCommentArgs{}
}

func (p *CommentServiceReplyCommentArgs) InitDefault() {
}

var CommentServiceReplyCommentArgs_Req_DEFAULT *ReplyCommentReq

func (p *CommentServiceReplyCommentArgs) GetReq() (v *ReplyCommentReq) {
	if !p.IsSetReq() {
		return CommentServiceReplyCommentArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *CommentServiceReplyCommentArgs) SetReq(val *ReplyCommentReq) {
	p.Req = val
}

func (p *CommentServiceReplyCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CommentServiceReplyCommentArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CommentServiceReplyCommentArgs(%+v)", *p)
}

func (p *CommentServiceReplyCommentArgs) DeepEqual(ano *CommentServiceReplyCommentArgs) bool {
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

func (p *CommentServiceReplyCommentArgs) Field1DeepEqual(src *ReplyCommentReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_CommentServiceReplyCommentArgs = map[int16]string{
	1: "req",
}

type CommentServiceReplyCommentResult struct {
	Success *ReplyCommentResp `thrift:"success,0,optional" frugal:"0,optional,ReplyCommentResp" json:"success,omitempty"`
}

func NewCommentServiceReplyCommentResult() *CommentServiceReplyCommentResult {
	return &CommentServiceReplyCommentResult{}
}

func (p *CommentServiceReplyCommentResult) InitDefault() {
}

var CommentServiceReplyCommentResult_Success_DEFAULT *ReplyCommentResp

func (p *CommentServiceReplyCommentResult) GetSuccess() (v *ReplyCommentResp) {
	if !p.IsSetSuccess() {
		return CommentServiceReplyCommentResult_Success_DEFAULT
	}
	return p.Success
}
func (p *CommentServiceReplyCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*ReplyCommentResp)
}

func (p *CommentServiceReplyCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CommentServiceReplyCommentResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CommentServiceReplyCommentResult(%+v)", *p)
}

func (p *CommentServiceReplyCommentResult) DeepEqual(ano *CommentServiceReplyCommentResult) bool {
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

func (p *CommentServiceReplyCommentResult) Field0DeepEqual(src *ReplyCommentResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_CommentServiceReplyCommentResult = map[int16]string{
	0: "success",
}

type CommentServiceGetCommentListArgs struct {
	Req *GetCommentListReq `thrift:"req,1" frugal:"1,default,GetCommentListReq" json:"req"`
}

func NewCommentServiceGetCommentListArgs() *CommentServiceGetCommentListArgs {
	return &CommentServiceGetCommentListArgs{}
}

func (p *CommentServiceGetCommentListArgs) InitDefault() {
}

var CommentServiceGetCommentListArgs_Req_DEFAULT *GetCommentListReq

func (p *CommentServiceGetCommentListArgs) GetReq() (v *GetCommentListReq) {
	if !p.IsSetReq() {
		return CommentServiceGetCommentListArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *CommentServiceGetCommentListArgs) SetReq(val *GetCommentListReq) {
	p.Req = val
}

func (p *CommentServiceGetCommentListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CommentServiceGetCommentListArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CommentServiceGetCommentListArgs(%+v)", *p)
}

func (p *CommentServiceGetCommentListArgs) DeepEqual(ano *CommentServiceGetCommentListArgs) bool {
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

func (p *CommentServiceGetCommentListArgs) Field1DeepEqual(src *GetCommentListReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_CommentServiceGetCommentListArgs = map[int16]string{
	1: "req",
}

type CommentServiceGetCommentListResult struct {
	Success *GetCommentListResp `thrift:"success,0,optional" frugal:"0,optional,GetCommentListResp" json:"success,omitempty"`
}

func NewCommentServiceGetCommentListResult() *CommentServiceGetCommentListResult {
	return &CommentServiceGetCommentListResult{}
}

func (p *CommentServiceGetCommentListResult) InitDefault() {
}

var CommentServiceGetCommentListResult_Success_DEFAULT *GetCommentListResp

func (p *CommentServiceGetCommentListResult) GetSuccess() (v *GetCommentListResp) {
	if !p.IsSetSuccess() {
		return CommentServiceGetCommentListResult_Success_DEFAULT
	}
	return p.Success
}
func (p *CommentServiceGetCommentListResult) SetSuccess(x interface{}) {
	p.Success = x.(*GetCommentListResp)
}

func (p *CommentServiceGetCommentListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CommentServiceGetCommentListResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CommentServiceGetCommentListResult(%+v)", *p)
}

func (p *CommentServiceGetCommentListResult) DeepEqual(ano *CommentServiceGetCommentListResult) bool {
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

func (p *CommentServiceGetCommentListResult) Field0DeepEqual(src *GetCommentListResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_CommentServiceGetCommentListResult = map[int16]string{
	0: "success",
}

type CommentServiceDeleteCommentArgs struct {
	Req *DeleteCommentReq `thrift:"req,1" frugal:"1,default,DeleteCommentReq" json:"req"`
}

func NewCommentServiceDeleteCommentArgs() *CommentServiceDeleteCommentArgs {
	return &CommentServiceDeleteCommentArgs{}
}

func (p *CommentServiceDeleteCommentArgs) InitDefault() {
}

var CommentServiceDeleteCommentArgs_Req_DEFAULT *DeleteCommentReq

func (p *CommentServiceDeleteCommentArgs) GetReq() (v *DeleteCommentReq) {
	if !p.IsSetReq() {
		return CommentServiceDeleteCommentArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *CommentServiceDeleteCommentArgs) SetReq(val *DeleteCommentReq) {
	p.Req = val
}

func (p *CommentServiceDeleteCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CommentServiceDeleteCommentArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CommentServiceDeleteCommentArgs(%+v)", *p)
}

func (p *CommentServiceDeleteCommentArgs) DeepEqual(ano *CommentServiceDeleteCommentArgs) bool {
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

func (p *CommentServiceDeleteCommentArgs) Field1DeepEqual(src *DeleteCommentReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_CommentServiceDeleteCommentArgs = map[int16]string{
	1: "req",
}

type CommentServiceDeleteCommentResult struct {
	Success *DeleteCommentResp `thrift:"success,0,optional" frugal:"0,optional,DeleteCommentResp" json:"success,omitempty"`
}

func NewCommentServiceDeleteCommentResult() *CommentServiceDeleteCommentResult {
	return &CommentServiceDeleteCommentResult{}
}

func (p *CommentServiceDeleteCommentResult) InitDefault() {
}

var CommentServiceDeleteCommentResult_Success_DEFAULT *DeleteCommentResp

func (p *CommentServiceDeleteCommentResult) GetSuccess() (v *DeleteCommentResp) {
	if !p.IsSetSuccess() {
		return CommentServiceDeleteCommentResult_Success_DEFAULT
	}
	return p.Success
}
func (p *CommentServiceDeleteCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*DeleteCommentResp)
}

func (p *CommentServiceDeleteCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CommentServiceDeleteCommentResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CommentServiceDeleteCommentResult(%+v)", *p)
}

func (p *CommentServiceDeleteCommentResult) DeepEqual(ano *CommentServiceDeleteCommentResult) bool {
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

func (p *CommentServiceDeleteCommentResult) Field0DeepEqual(src *DeleteCommentResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_CommentServiceDeleteCommentResult = map[int16]string{
	0: "success",
}
