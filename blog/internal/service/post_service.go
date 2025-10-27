package service

import (
	"golang_study/blog/internal/common"
	"golang_study/blog/internal/dto"
	"golang_study/blog/internal/dto/post_dto"
	"golang_study/blog/internal/model"
	"golang_study/blog/internal/repo"
)

type PostService struct {
	postRepository *repo.PostRepository
}

func NewPostService(postRepository *repo.PostRepository) *PostService {
	return &PostService{
		postRepository: postRepository,
	}

}

// 创建文章
func (postService *PostService) SavePost(title string, content string, userId uint) (bool, error) {

	post := model.Post{
		Title:   title,
		Content: content,
		UserID:  userId,
	}
	success, err := postService.postRepository.SavePost(&post)

	return success, err
}

// 删除文章
func (postService *PostService) DeletePost(deletePostDTO post_dto.DeletePostDTO) error {
	// 查询文章是否存在，报不存在异常
	post, err := postService.GetPostById(deletePostDTO.Id)

	if err != nil {
		return err
	}

	// 判断是否为本人更新， 报非法操作
	if post.UserID != deletePostDTO.UserId {
		return common.NewBusinessError(common.ErrUnAuthored)
	}

	return postService.postRepository.Delete(&model.Post{}, deletePostDTO.Id)
}

// 获取文章内容
func (postService *PostService) GetPostById(id uint) (*model.Post, error) {
	return postService.postRepository.GetById(id)

}

func (postService *PostService) PagePost(pageDTO dto.Pagination) ([]model.Post, int64, error) {
	return postService.postRepository.Page(pageDTO)
}

func (postService *PostService) UpdatePost(updatePostDTO post_dto.UpdatePostDTO) (bool, error) {
	// 查询文章是否存在，报不存在异常
	post, err := postService.GetPostById(updatePostDTO.Id)

	if err != nil {
		return false, err
	}

	// 判断是否为本人更新， 报非法操作
	if post.UserID != updatePostDTO.UserId {
		return false, common.NewBusinessError(common.ErrUnAuthored)
	}

	// 执行正常更新
	toUpdatePost := model.Post{
		ID:      post.ID,
		Title:   updatePostDTO.Title,
		Content: updatePostDTO.Content,
	}
	return postService.postRepository.Update(toUpdatePost)
}
