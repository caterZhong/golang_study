### blog 是一个使用golang + gin + gorm 开发的简易博客

#### 主要功能
- 登录（使用jwt）
- 发布文章
- 获取文章列表
- 修改文章
- 删除文章
- 发布文章评论
- 获取文章评论列表

#### 项目搭建
- 主要依赖 gin、 gorm、 mysql（库表初始化见blog/config/blog.sql）

#### 接口测试
- 提供curl测试文档（见blog/test/curl_test.txt）
