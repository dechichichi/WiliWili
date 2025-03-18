package pack
import (
	"wiliwili/kitex_gen/model"

	domainmodel "wiliwili/app/comment/domain/model"
)


func BuildCommentList(commentlist []*domainmodel.Comment)*[]*model.Comment {
	var commentList []*model.Comment
	for _, comment := range commentlist {
		commentList = append(commentList, BuildComment(comment))
	}
	return &commentList
}

func BuildComment(comment *domainmodel.Comment) *model.Comment {
	return &model.Comment{
		CommentId:      comment.CommentID,
		UserId:         comment.UserID,
		Content:        comment.CommentContent,
		CreateTime:     comment.CreatedAt.Unix(),
	}
}