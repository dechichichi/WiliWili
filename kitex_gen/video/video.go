// Code generated by thriftgo (0.3.18). DO NOT EDIT.

package video

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"wiliwili/kitex_gen/model"
)

type VideoSubmissionReq struct {
	VideoName     string `thrift:"video_name,1,required" frugal:"1,required,string" json:"video_name"`
	VideoDuration string `thrift:"video_duration,2,required" frugal:"2,required,string" json:"video_duration"`
	Video         []byte `thrift:"video,3,required" frugal:"3,required,binary" json:"video"`
}

func NewVideoSubmissionReq() *VideoSubmissionReq {
	return &VideoSubmissionReq{}
}

func (p *VideoSubmissionReq) InitDefault() {
}

func (p *VideoSubmissionReq) GetVideoName() (v string) {
	return p.VideoName
}

func (p *VideoSubmissionReq) GetVideoDuration() (v string) {
	return p.VideoDuration
}

func (p *VideoSubmissionReq) GetVideo() (v []byte) {
	return p.Video
}
func (p *VideoSubmissionReq) SetVideoName(val string) {
	p.VideoName = val
}
func (p *VideoSubmissionReq) SetVideoDuration(val string) {
	p.VideoDuration = val
}
func (p *VideoSubmissionReq) SetVideo(val []byte) {
	p.Video = val
}

func (p *VideoSubmissionReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VideoSubmissionReq(%+v)", *p)
}

func (p *VideoSubmissionReq) DeepEqual(ano *VideoSubmissionReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.VideoName) {
		return false
	}
	if !p.Field2DeepEqual(ano.VideoDuration) {
		return false
	}
	if !p.Field3DeepEqual(ano.Video) {
		return false
	}
	return true
}

func (p *VideoSubmissionReq) Field1DeepEqual(src string) bool {

	if strings.Compare(p.VideoName, src) != 0 {
		return false
	}
	return true
}
func (p *VideoSubmissionReq) Field2DeepEqual(src string) bool {

	if strings.Compare(p.VideoDuration, src) != 0 {
		return false
	}
	return true
}
func (p *VideoSubmissionReq) Field3DeepEqual(src []byte) bool {

	if bytes.Compare(p.Video, src) != 0 {
		return false
	}
	return true
}

var fieldIDToName_VideoSubmissionReq = map[int16]string{
	1: "video_name",
	2: "video_duration",
	3: "video",
}

type VideoSubmissionResp struct {
	Baseresp *model.BaseResp `thrift:"baseresp,1,required" frugal:"1,required,model.BaseResp" json:"baseresp"`
	VideoId  string          `thrift:"video_id,2,required" frugal:"2,required,string" json:"video_id"`
	VideoUrl string          `thrift:"video_url,3,required" frugal:"3,required,string" json:"video_url"`
}

func NewVideoSubmissionResp() *VideoSubmissionResp {
	return &VideoSubmissionResp{}
}

func (p *VideoSubmissionResp) InitDefault() {
}

var VideoSubmissionResp_Baseresp_DEFAULT *model.BaseResp

func (p *VideoSubmissionResp) GetBaseresp() (v *model.BaseResp) {
	if !p.IsSetBaseresp() {
		return VideoSubmissionResp_Baseresp_DEFAULT
	}
	return p.Baseresp
}

func (p *VideoSubmissionResp) GetVideoId() (v string) {
	return p.VideoId
}

func (p *VideoSubmissionResp) GetVideoUrl() (v string) {
	return p.VideoUrl
}
func (p *VideoSubmissionResp) SetBaseresp(val *model.BaseResp) {
	p.Baseresp = val
}
func (p *VideoSubmissionResp) SetVideoId(val string) {
	p.VideoId = val
}
func (p *VideoSubmissionResp) SetVideoUrl(val string) {
	p.VideoUrl = val
}

func (p *VideoSubmissionResp) IsSetBaseresp() bool {
	return p.Baseresp != nil
}

func (p *VideoSubmissionResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VideoSubmissionResp(%+v)", *p)
}

func (p *VideoSubmissionResp) DeepEqual(ano *VideoSubmissionResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Baseresp) {
		return false
	}
	if !p.Field2DeepEqual(ano.VideoId) {
		return false
	}
	if !p.Field3DeepEqual(ano.VideoUrl) {
		return false
	}
	return true
}

func (p *VideoSubmissionResp) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.Baseresp.DeepEqual(src) {
		return false
	}
	return true
}
func (p *VideoSubmissionResp) Field2DeepEqual(src string) bool {

	if strings.Compare(p.VideoId, src) != 0 {
		return false
	}
	return true
}
func (p *VideoSubmissionResp) Field3DeepEqual(src string) bool {

	if strings.Compare(p.VideoUrl, src) != 0 {
		return false
	}
	return true
}

var fieldIDToName_VideoSubmissionResp = map[int16]string{
	1: "baseresp",
	2: "video_id",
	3: "video_url",
}

type VideoGetReq struct {
	VideoId string `thrift:"video_id,1,required" frugal:"1,required,string" json:"video_id"`
}

func NewVideoGetReq() *VideoGetReq {
	return &VideoGetReq{}
}

func (p *VideoGetReq) InitDefault() {
}

func (p *VideoGetReq) GetVideoId() (v string) {
	return p.VideoId
}
func (p *VideoGetReq) SetVideoId(val string) {
	p.VideoId = val
}

func (p *VideoGetReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VideoGetReq(%+v)", *p)
}

func (p *VideoGetReq) DeepEqual(ano *VideoGetReq) bool {
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

func (p *VideoGetReq) Field1DeepEqual(src string) bool {

	if strings.Compare(p.VideoId, src) != 0 {
		return false
	}
	return true
}

var fieldIDToName_VideoGetReq = map[int16]string{
	1: "video_id",
}

type VideoGetResp struct {
	Baseresp *model.BaseResp `thrift:"baseresp,1,required" frugal:"1,required,model.BaseResp" json:"baseresp"`
	Video    *model.Video    `thrift:"video,2,required" frugal:"2,required,model.Video" json:"video"`
}

func NewVideoGetResp() *VideoGetResp {
	return &VideoGetResp{}
}

func (p *VideoGetResp) InitDefault() {
}

var VideoGetResp_Baseresp_DEFAULT *model.BaseResp

func (p *VideoGetResp) GetBaseresp() (v *model.BaseResp) {
	if !p.IsSetBaseresp() {
		return VideoGetResp_Baseresp_DEFAULT
	}
	return p.Baseresp
}

var VideoGetResp_Video_DEFAULT *model.Video

func (p *VideoGetResp) GetVideo() (v *model.Video) {
	if !p.IsSetVideo() {
		return VideoGetResp_Video_DEFAULT
	}
	return p.Video
}
func (p *VideoGetResp) SetBaseresp(val *model.BaseResp) {
	p.Baseresp = val
}
func (p *VideoGetResp) SetVideo(val *model.Video) {
	p.Video = val
}

func (p *VideoGetResp) IsSetBaseresp() bool {
	return p.Baseresp != nil
}

func (p *VideoGetResp) IsSetVideo() bool {
	return p.Video != nil
}

func (p *VideoGetResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VideoGetResp(%+v)", *p)
}

func (p *VideoGetResp) DeepEqual(ano *VideoGetResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Baseresp) {
		return false
	}
	if !p.Field2DeepEqual(ano.Video) {
		return false
	}
	return true
}

func (p *VideoGetResp) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.Baseresp.DeepEqual(src) {
		return false
	}
	return true
}
func (p *VideoGetResp) Field2DeepEqual(src *model.Video) bool {

	if !p.Video.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_VideoGetResp = map[int16]string{
	1: "baseresp",
	2: "video",
}

type VideoSearchReq struct {
	Keyword  string `thrift:"keyword,1,required" frugal:"1,required,string" json:"keyword"`
	PageNum  int64  `thrift:"page_num,2,required" frugal:"2,required,i64" json:"page_num"`
	PageSize int64  `thrift:"page_size,3,required" frugal:"3,required,i64" json:"page_size"`
}

func NewVideoSearchReq() *VideoSearchReq {
	return &VideoSearchReq{}
}

func (p *VideoSearchReq) InitDefault() {
}

func (p *VideoSearchReq) GetKeyword() (v string) {
	return p.Keyword
}

func (p *VideoSearchReq) GetPageNum() (v int64) {
	return p.PageNum
}

func (p *VideoSearchReq) GetPageSize() (v int64) {
	return p.PageSize
}
func (p *VideoSearchReq) SetKeyword(val string) {
	p.Keyword = val
}
func (p *VideoSearchReq) SetPageNum(val int64) {
	p.PageNum = val
}
func (p *VideoSearchReq) SetPageSize(val int64) {
	p.PageSize = val
}

func (p *VideoSearchReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VideoSearchReq(%+v)", *p)
}

func (p *VideoSearchReq) DeepEqual(ano *VideoSearchReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Keyword) {
		return false
	}
	if !p.Field2DeepEqual(ano.PageNum) {
		return false
	}
	if !p.Field3DeepEqual(ano.PageSize) {
		return false
	}
	return true
}

func (p *VideoSearchReq) Field1DeepEqual(src string) bool {

	if strings.Compare(p.Keyword, src) != 0 {
		return false
	}
	return true
}
func (p *VideoSearchReq) Field2DeepEqual(src int64) bool {

	if p.PageNum != src {
		return false
	}
	return true
}
func (p *VideoSearchReq) Field3DeepEqual(src int64) bool {

	if p.PageSize != src {
		return false
	}
	return true
}

var fieldIDToName_VideoSearchReq = map[int16]string{
	1: "keyword",
	2: "page_num",
	3: "page_size",
}

type VideoSearchResp struct {
	Baseresp *model.BaseResp `thrift:"baseresp,1,required" frugal:"1,required,model.BaseResp" json:"baseresp"`
	Videos   []*model.Video  `thrift:"videos,2,required" frugal:"2,required,list<model.Video>" json:"videos"`
}

func NewVideoSearchResp() *VideoSearchResp {
	return &VideoSearchResp{}
}

func (p *VideoSearchResp) InitDefault() {
}

var VideoSearchResp_Baseresp_DEFAULT *model.BaseResp

func (p *VideoSearchResp) GetBaseresp() (v *model.BaseResp) {
	if !p.IsSetBaseresp() {
		return VideoSearchResp_Baseresp_DEFAULT
	}
	return p.Baseresp
}

func (p *VideoSearchResp) GetVideos() (v []*model.Video) {
	return p.Videos
}
func (p *VideoSearchResp) SetBaseresp(val *model.BaseResp) {
	p.Baseresp = val
}
func (p *VideoSearchResp) SetVideos(val []*model.Video) {
	p.Videos = val
}

func (p *VideoSearchResp) IsSetBaseresp() bool {
	return p.Baseresp != nil
}

func (p *VideoSearchResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VideoSearchResp(%+v)", *p)
}

func (p *VideoSearchResp) DeepEqual(ano *VideoSearchResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Baseresp) {
		return false
	}
	if !p.Field2DeepEqual(ano.Videos) {
		return false
	}
	return true
}

func (p *VideoSearchResp) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.Baseresp.DeepEqual(src) {
		return false
	}
	return true
}
func (p *VideoSearchResp) Field2DeepEqual(src []*model.Video) bool {

	if len(p.Videos) != len(src) {
		return false
	}
	for i, v := range p.Videos {
		_src := src[i]
		if !v.DeepEqual(_src) {
			return false
		}
	}
	return true
}

var fieldIDToName_VideoSearchResp = map[int16]string{
	1: "baseresp",
	2: "videos",
}

type VideoTrendingReq struct {
	PageNum  int64 `thrift:"page_num,1,required" frugal:"1,required,i64" json:"page_num"`
	PageSize int64 `thrift:"page_size,2,required" frugal:"2,required,i64" json:"page_size"`
}

func NewVideoTrendingReq() *VideoTrendingReq {
	return &VideoTrendingReq{}
}

func (p *VideoTrendingReq) InitDefault() {
}

func (p *VideoTrendingReq) GetPageNum() (v int64) {
	return p.PageNum
}

func (p *VideoTrendingReq) GetPageSize() (v int64) {
	return p.PageSize
}
func (p *VideoTrendingReq) SetPageNum(val int64) {
	p.PageNum = val
}
func (p *VideoTrendingReq) SetPageSize(val int64) {
	p.PageSize = val
}

func (p *VideoTrendingReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VideoTrendingReq(%+v)", *p)
}

func (p *VideoTrendingReq) DeepEqual(ano *VideoTrendingReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.PageNum) {
		return false
	}
	if !p.Field2DeepEqual(ano.PageSize) {
		return false
	}
	return true
}

func (p *VideoTrendingReq) Field1DeepEqual(src int64) bool {

	if p.PageNum != src {
		return false
	}
	return true
}
func (p *VideoTrendingReq) Field2DeepEqual(src int64) bool {

	if p.PageSize != src {
		return false
	}
	return true
}

var fieldIDToName_VideoTrendingReq = map[int16]string{
	1: "page_num",
	2: "page_size",
}

type VideoTrendingResp struct {
	Baseresp *model.BaseResp `thrift:"baseresp,1,required" frugal:"1,required,model.BaseResp" json:"baseresp"`
	Videos   []*model.Video  `thrift:"videos,2,required" frugal:"2,required,list<model.Video>" json:"videos"`
}

func NewVideoTrendingResp() *VideoTrendingResp {
	return &VideoTrendingResp{}
}

func (p *VideoTrendingResp) InitDefault() {
}

var VideoTrendingResp_Baseresp_DEFAULT *model.BaseResp

func (p *VideoTrendingResp) GetBaseresp() (v *model.BaseResp) {
	if !p.IsSetBaseresp() {
		return VideoTrendingResp_Baseresp_DEFAULT
	}
	return p.Baseresp
}

func (p *VideoTrendingResp) GetVideos() (v []*model.Video) {
	return p.Videos
}
func (p *VideoTrendingResp) SetBaseresp(val *model.BaseResp) {
	p.Baseresp = val
}
func (p *VideoTrendingResp) SetVideos(val []*model.Video) {
	p.Videos = val
}

func (p *VideoTrendingResp) IsSetBaseresp() bool {
	return p.Baseresp != nil
}

func (p *VideoTrendingResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VideoTrendingResp(%+v)", *p)
}

func (p *VideoTrendingResp) DeepEqual(ano *VideoTrendingResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Baseresp) {
		return false
	}
	if !p.Field2DeepEqual(ano.Videos) {
		return false
	}
	return true
}

func (p *VideoTrendingResp) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.Baseresp.DeepEqual(src) {
		return false
	}
	return true
}
func (p *VideoTrendingResp) Field2DeepEqual(src []*model.Video) bool {

	if len(p.Videos) != len(src) {
		return false
	}
	for i, v := range p.Videos {
		_src := src[i]
		if !v.DeepEqual(_src) {
			return false
		}
	}
	return true
}

var fieldIDToName_VideoTrendingResp = map[int16]string{
	1: "baseresp",
	2: "videos",
}

type VideoService interface {
	VideoSubmission(ctx context.Context, req *VideoSubmissionReq) (r *VideoSubmissionResp, err error)

	VideoGet(ctx context.Context, req *VideoGetReq) (r *VideoGetResp, err error)

	VideoSearch(ctx context.Context, req *VideoSearchReq) (r *VideoSearchResp, err error)

	VideoTrending(ctx context.Context, req *VideoTrendingReq) (r *VideoTrendingResp, err error)
}

type VideoServiceVideoSubmissionArgs struct {
	Req *VideoSubmissionReq `thrift:"req,1" frugal:"1,default,VideoSubmissionReq" json:"req"`
}

func NewVideoServiceVideoSubmissionArgs() *VideoServiceVideoSubmissionArgs {
	return &VideoServiceVideoSubmissionArgs{}
}

func (p *VideoServiceVideoSubmissionArgs) InitDefault() {
}

var VideoServiceVideoSubmissionArgs_Req_DEFAULT *VideoSubmissionReq

func (p *VideoServiceVideoSubmissionArgs) GetReq() (v *VideoSubmissionReq) {
	if !p.IsSetReq() {
		return VideoServiceVideoSubmissionArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *VideoServiceVideoSubmissionArgs) SetReq(val *VideoSubmissionReq) {
	p.Req = val
}

func (p *VideoServiceVideoSubmissionArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *VideoServiceVideoSubmissionArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VideoServiceVideoSubmissionArgs(%+v)", *p)
}

func (p *VideoServiceVideoSubmissionArgs) DeepEqual(ano *VideoServiceVideoSubmissionArgs) bool {
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

func (p *VideoServiceVideoSubmissionArgs) Field1DeepEqual(src *VideoSubmissionReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_VideoServiceVideoSubmissionArgs = map[int16]string{
	1: "req",
}

type VideoServiceVideoSubmissionResult struct {
	Success *VideoSubmissionResp `thrift:"success,0,optional" frugal:"0,optional,VideoSubmissionResp" json:"success,omitempty"`
}

func NewVideoServiceVideoSubmissionResult() *VideoServiceVideoSubmissionResult {
	return &VideoServiceVideoSubmissionResult{}
}

func (p *VideoServiceVideoSubmissionResult) InitDefault() {
}

var VideoServiceVideoSubmissionResult_Success_DEFAULT *VideoSubmissionResp

func (p *VideoServiceVideoSubmissionResult) GetSuccess() (v *VideoSubmissionResp) {
	if !p.IsSetSuccess() {
		return VideoServiceVideoSubmissionResult_Success_DEFAULT
	}
	return p.Success
}
func (p *VideoServiceVideoSubmissionResult) SetSuccess(x interface{}) {
	p.Success = x.(*VideoSubmissionResp)
}

func (p *VideoServiceVideoSubmissionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *VideoServiceVideoSubmissionResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VideoServiceVideoSubmissionResult(%+v)", *p)
}

func (p *VideoServiceVideoSubmissionResult) DeepEqual(ano *VideoServiceVideoSubmissionResult) bool {
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

func (p *VideoServiceVideoSubmissionResult) Field0DeepEqual(src *VideoSubmissionResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_VideoServiceVideoSubmissionResult = map[int16]string{
	0: "success",
}

type VideoServiceVideoGetArgs struct {
	Req *VideoGetReq `thrift:"req,1" frugal:"1,default,VideoGetReq" json:"req"`
}

func NewVideoServiceVideoGetArgs() *VideoServiceVideoGetArgs {
	return &VideoServiceVideoGetArgs{}
}

func (p *VideoServiceVideoGetArgs) InitDefault() {
}

var VideoServiceVideoGetArgs_Req_DEFAULT *VideoGetReq

func (p *VideoServiceVideoGetArgs) GetReq() (v *VideoGetReq) {
	if !p.IsSetReq() {
		return VideoServiceVideoGetArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *VideoServiceVideoGetArgs) SetReq(val *VideoGetReq) {
	p.Req = val
}

func (p *VideoServiceVideoGetArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *VideoServiceVideoGetArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VideoServiceVideoGetArgs(%+v)", *p)
}

func (p *VideoServiceVideoGetArgs) DeepEqual(ano *VideoServiceVideoGetArgs) bool {
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

func (p *VideoServiceVideoGetArgs) Field1DeepEqual(src *VideoGetReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_VideoServiceVideoGetArgs = map[int16]string{
	1: "req",
}

type VideoServiceVideoGetResult struct {
	Success *VideoGetResp `thrift:"success,0,optional" frugal:"0,optional,VideoGetResp" json:"success,omitempty"`
}

func NewVideoServiceVideoGetResult() *VideoServiceVideoGetResult {
	return &VideoServiceVideoGetResult{}
}

func (p *VideoServiceVideoGetResult) InitDefault() {
}

var VideoServiceVideoGetResult_Success_DEFAULT *VideoGetResp

func (p *VideoServiceVideoGetResult) GetSuccess() (v *VideoGetResp) {
	if !p.IsSetSuccess() {
		return VideoServiceVideoGetResult_Success_DEFAULT
	}
	return p.Success
}
func (p *VideoServiceVideoGetResult) SetSuccess(x interface{}) {
	p.Success = x.(*VideoGetResp)
}

func (p *VideoServiceVideoGetResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *VideoServiceVideoGetResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VideoServiceVideoGetResult(%+v)", *p)
}

func (p *VideoServiceVideoGetResult) DeepEqual(ano *VideoServiceVideoGetResult) bool {
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

func (p *VideoServiceVideoGetResult) Field0DeepEqual(src *VideoGetResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_VideoServiceVideoGetResult = map[int16]string{
	0: "success",
}

type VideoServiceVideoSearchArgs struct {
	Req *VideoSearchReq `thrift:"req,1" frugal:"1,default,VideoSearchReq" json:"req"`
}

func NewVideoServiceVideoSearchArgs() *VideoServiceVideoSearchArgs {
	return &VideoServiceVideoSearchArgs{}
}

func (p *VideoServiceVideoSearchArgs) InitDefault() {
}

var VideoServiceVideoSearchArgs_Req_DEFAULT *VideoSearchReq

func (p *VideoServiceVideoSearchArgs) GetReq() (v *VideoSearchReq) {
	if !p.IsSetReq() {
		return VideoServiceVideoSearchArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *VideoServiceVideoSearchArgs) SetReq(val *VideoSearchReq) {
	p.Req = val
}

func (p *VideoServiceVideoSearchArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *VideoServiceVideoSearchArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VideoServiceVideoSearchArgs(%+v)", *p)
}

func (p *VideoServiceVideoSearchArgs) DeepEqual(ano *VideoServiceVideoSearchArgs) bool {
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

func (p *VideoServiceVideoSearchArgs) Field1DeepEqual(src *VideoSearchReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_VideoServiceVideoSearchArgs = map[int16]string{
	1: "req",
}

type VideoServiceVideoSearchResult struct {
	Success *VideoSearchResp `thrift:"success,0,optional" frugal:"0,optional,VideoSearchResp" json:"success,omitempty"`
}

func NewVideoServiceVideoSearchResult() *VideoServiceVideoSearchResult {
	return &VideoServiceVideoSearchResult{}
}

func (p *VideoServiceVideoSearchResult) InitDefault() {
}

var VideoServiceVideoSearchResult_Success_DEFAULT *VideoSearchResp

func (p *VideoServiceVideoSearchResult) GetSuccess() (v *VideoSearchResp) {
	if !p.IsSetSuccess() {
		return VideoServiceVideoSearchResult_Success_DEFAULT
	}
	return p.Success
}
func (p *VideoServiceVideoSearchResult) SetSuccess(x interface{}) {
	p.Success = x.(*VideoSearchResp)
}

func (p *VideoServiceVideoSearchResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *VideoServiceVideoSearchResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VideoServiceVideoSearchResult(%+v)", *p)
}

func (p *VideoServiceVideoSearchResult) DeepEqual(ano *VideoServiceVideoSearchResult) bool {
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

func (p *VideoServiceVideoSearchResult) Field0DeepEqual(src *VideoSearchResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_VideoServiceVideoSearchResult = map[int16]string{
	0: "success",
}

type VideoServiceVideoTrendingArgs struct {
	Req *VideoTrendingReq `thrift:"req,1" frugal:"1,default,VideoTrendingReq" json:"req"`
}

func NewVideoServiceVideoTrendingArgs() *VideoServiceVideoTrendingArgs {
	return &VideoServiceVideoTrendingArgs{}
}

func (p *VideoServiceVideoTrendingArgs) InitDefault() {
}

var VideoServiceVideoTrendingArgs_Req_DEFAULT *VideoTrendingReq

func (p *VideoServiceVideoTrendingArgs) GetReq() (v *VideoTrendingReq) {
	if !p.IsSetReq() {
		return VideoServiceVideoTrendingArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *VideoServiceVideoTrendingArgs) SetReq(val *VideoTrendingReq) {
	p.Req = val
}

func (p *VideoServiceVideoTrendingArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *VideoServiceVideoTrendingArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VideoServiceVideoTrendingArgs(%+v)", *p)
}

func (p *VideoServiceVideoTrendingArgs) DeepEqual(ano *VideoServiceVideoTrendingArgs) bool {
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

func (p *VideoServiceVideoTrendingArgs) Field1DeepEqual(src *VideoTrendingReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_VideoServiceVideoTrendingArgs = map[int16]string{
	1: "req",
}

type VideoServiceVideoTrendingResult struct {
	Success *VideoTrendingResp `thrift:"success,0,optional" frugal:"0,optional,VideoTrendingResp" json:"success,omitempty"`
}

func NewVideoServiceVideoTrendingResult() *VideoServiceVideoTrendingResult {
	return &VideoServiceVideoTrendingResult{}
}

func (p *VideoServiceVideoTrendingResult) InitDefault() {
}

var VideoServiceVideoTrendingResult_Success_DEFAULT *VideoTrendingResp

func (p *VideoServiceVideoTrendingResult) GetSuccess() (v *VideoTrendingResp) {
	if !p.IsSetSuccess() {
		return VideoServiceVideoTrendingResult_Success_DEFAULT
	}
	return p.Success
}
func (p *VideoServiceVideoTrendingResult) SetSuccess(x interface{}) {
	p.Success = x.(*VideoTrendingResp)
}

func (p *VideoServiceVideoTrendingResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *VideoServiceVideoTrendingResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VideoServiceVideoTrendingResult(%+v)", *p)
}

func (p *VideoServiceVideoTrendingResult) DeepEqual(ano *VideoServiceVideoTrendingResult) bool {
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

func (p *VideoServiceVideoTrendingResult) Field0DeepEqual(src *VideoTrendingResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

var fieldIDToName_VideoServiceVideoTrendingResult = map[int16]string{
	0: "success",
}
