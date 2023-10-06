package usecase

import (
	"context"
	"gosocial/internal/entity"
)

func (uc *ucImplement) CreatePost(ctx context.Context, req *entity.CreatePostRequest) (*entity.CreatePostResponse, error) {
	// TODO: upload photos to firebase storage

	// store image to bucket
	imgInfo, err := uc.imgBucket.SaveImage(ctx, req.Filename, req.File)
	if err != nil {
		return nil, err
	}
	imgInfo.URL = "/images/" + imgInfo.Filename
	// inline image info in Post
	req.Post.Image = imgInfo

	postDoc, err := uc.postStore.Save(req.Post)

	if err != nil {
		return nil, err
	}

	return &entity.CreatePostResponse{
		Postid: postDoc.DocId.Hex(),
	}, nil
}

func (uc *ucImplement) GetPostsByUser(ctx context.Context, req *entity.PostsByUserRequest) (*entity.PostsResponse, error) {

	postDocs, err := uc.postStore.GetManyByUser(req.Userid)

	if err != nil {
		return nil, err
	}

	// populate posts & profiles
	var posts []entity.PostRes
	for _, postDoc := range postDocs {
		post := postDoc.ToPost()
		profileDoc, _ := uc.profileStore.Get(post.Userid)
		posts = append(posts, entity.PostRes{
			Post:    post,
			Profile: profileDoc.ToProfile(),
		})
	}

	return &entity.PostsResponse{
		Posts: posts,
	}, nil
}

func (uc *ucImplement) GetPosts(ctx context.Context, req *entity.PostsRequest) (*entity.PostsResponse, error) {

	postDocs, err := uc.postStore.GetMany()

	if err != nil {
		return nil, err
	}

	// populate posts & profiles
	var posts []entity.PostRes
	for _, postDoc := range postDocs {
		post := postDoc.ToPost()
		profileDoc, _ := uc.profileStore.Get(post.Userid)
		posts = append(posts, entity.PostRes{
			Post:    post,
			Profile: profileDoc.ToProfile(),
		})
	}

	return &entity.PostsResponse{
		Posts: posts,
	}, nil
}

func (uc *ucImplement) DeletePost(ctx context.Context, req *entity.DeletePostRequest) (*entity.DeletePostResponse, error) {

	err := uc.postStore.DeleteOne(req.Postid)

	if err != nil {
		return nil, err
	}

	return &entity.DeletePostResponse{
		Postid: req.Postid,
	}, nil
}

func (uc *ucImplement) LikePost(ctx context.Context, req *entity.LikePostRequest) (*entity.LikePostResponse, error) {

	likesTotal, err := uc.postStore.LikePost(req.Postid, req.Userid)

	if err != nil {
		return nil, err
	}

	return &entity.LikePostResponse{
		LikesTotal: likesTotal,
	}, nil
}
