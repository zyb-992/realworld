package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, loginReq *v1.LoginRequest) (*v1.UserReply, error) {
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "boom",
		},
	}, nil
}

func (s *RealWorldService) Register(context.Context, *v1.RegisterRequest) (*v1.UserReply, error) {
	return &v1.UserReply{}, nil
}

func (s *RealWorldService) GetCurrentUser(context.Context, *v1.GetCurrentUserRequest) (*v1.UserReply, error) {
	return &v1.UserReply{}, nil
}

func (s *RealWorldService) UpdateUser(context.Context, *v1.UpdateUserRequest) (*v1.UserReply, error) {
	return &v1.UserReply{}, nil
}

func (s *RealWorldService) GetProfile(context.Context, *v1.GetProfileRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil

}

func (s *RealWorldService) GetFollowUser(context.Context, *v1.FollowUserRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil

}

func (s *RealWorldService) GetUnFollowUser(context.Context, *v1.UnFollowUserRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil

}

func (s *RealWorldService) GetListArticles(context.Context, *v1.ListArticlesRequest) (*v1.MultipleArticlesReply, error) {
	return &v1.MultipleArticlesReply{}, nil

}

func (s *RealWorldService) FeedArticles(context.Context, *v1.FeedArticlesRequest) (*v1.MultipleArticlesReply, error) {
	return &v1.MultipleArticlesReply{}, nil

}

func (s *RealWorldService) GetArticle(context.Context, *v1.GetArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{}, nil

}

func (s *RealWorldService) CreateArticle(context.Context, *v1.CreateArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{}, nil

}

func (s *RealWorldService) UpdateArticle(context.Context, *v1.UpdateArticleRequest) (*v1.DeleteArticleReply, error) {
	return &v1.DeleteArticleReply{}, nil

}

func (s *RealWorldService) DeleteArticle(context.Context, *v1.DeleteArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{}, nil

}

func (s *RealWorldService) AddCommentsToAnArticle(context.Context, *v1.AddCommentsRequest) (*v1.SingleCommentReply, error) {
	return &v1.SingleCommentReply{}, nil

}

func (s *RealWorldService) GetCommentsFromAnArticle(context.Context, *v1.GetCommentsRequest) (*v1.MultipleCommentsReply, error) {
	return &v1.MultipleCommentsReply{}, nil

}

func (s *RealWorldService) DeleteComment(context.Context, *v1.DeleteCommentRequest) (*v1.DeleteCommentReply, error) {
	return &v1.DeleteCommentReply{}, nil

}

func (s *RealWorldService) FavoriteArticle(context.Context, *v1.FavoriteArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{}, nil

}
func (s *RealWorldService) UnFavoriteArticle(context.Context, *v1.UnFavoriteArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{}, nil

}

func (s *RealWorldService) GetTags(context.Context, *v1.GetTagsRequest) (*v1.TagListReply, error) {
	return &v1.TagListReply{}, nil
}
