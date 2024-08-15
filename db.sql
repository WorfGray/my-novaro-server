
--存储用户信息
CREATE TABLE users
(
    id                  VARCHAR(64) PRIMARY KEY,
    account             VARCHAR(42) NOT NULL UNIQUE,
    email               VARCHAR(64),
    did_name            VARCHAR(50),
    profile_picture_url TEXT,
    bio                 TEXT,
    created_at          TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    last_login          TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    user_level          INT            DEFAULT 1,
    user_score          DECIMAL(10, 2) DEFAULT 0,
    credit_score        DECIMAL(5, 2)  DEFAULT 100.00
);

COMMENT ON TABLE users IS '存储用户信息';
COMMENT ON COLUMN users.id IS 'id';
COMMENT ON COLUMN users.account IS '账户地址';
COMMENT ON COLUMN users.email IS '邮箱地址';
COMMENT ON COLUMN users.did_name IS '用户名';
COMMENT ON COLUMN users.profile_picture_url IS '头像url';
COMMENT ON COLUMN users.bio IS '个人简介';
COMMENT ON COLUMN users.created_at IS '创建时间';
COMMENT ON COLUMN users.last_login IS '最后登录时间';
COMMENT ON COLUMN users.user_level IS '用户等级';
COMMENT ON COLUMN users.user_score IS '用户积分';
COMMENT ON COLUMN users.credit_score IS '信用分';


--存储用户发布的帖子信息
CREATE TABLE posts
(
    id                 VARCHAR(64) PRIMARY KEY NOT NULL,
    user_id            VARCHAR(64)             NOT NULL,
    content            TEXT                    NOT NULL,
    comments_amount    INT,
    collections_amount INT,
    reposts_amount     INT,
    created_at         TIMESTAMP               NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE posts IS '存储用户发布的帖子信息';
COMMENT ON COLUMN posts.id IS '帖子ID，唯一标识帖子';
COMMENT ON COLUMN posts.user_id IS '用户ID，标识发布帖子的人';
COMMENT ON COLUMN posts.content IS '文章内容';
COMMENT ON COLUMN posts.comments_amount IS '评论数量';
COMMENT ON COLUMN posts.collections_amount IS '收藏数量';
COMMENT ON COLUMN posts.reposts_amount IS '转发数量';
COMMENT ON COLUMN posts.created_at IS '发布时间';

--存储用户点赞的帖子信息
CREATE TABLE tags
(
    id         VARCHAR(64) PRIMARY KEY NOT NULL,
    user_id    VARCHAR(64)             NOT NULL,
    post_id    VARCHAR(64)             NOT NULL,
    tag_type   VARCHAR(32)             NOT NULL,
    created_at TIMESTAMP               NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE tags IS '存储用户点赞的帖子信息';
COMMENT ON COLUMN tags.id IS '标签id';
COMMENT ON COLUMN tags.user_id IS '用户ID';
COMMENT ON COLUMN tags.post_id IS '帖子ID';
COMMENT ON COLUMN tags.tag_type IS '标签类型';
COMMENT ON COLUMN tags.created_at IS '创建时间';

--存储用户评论的帖子信息
CREATE TABLE comments
(
    id         VARCHAR(64) PRIMARY KEY NOT NULL,
    parent_id  VARCHAR(64),
    user_id    VARCHAR(64)             NOT NULL,
    post_id    VARCHAR(64)             NOT NULL,
    content    TEXT,
    created_at TIMESTAMP               NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE comments IS '存储用户评论的帖子信息';
COMMENT ON COLUMN comments.id IS '评论ID，唯一标识评论';
COMMENT ON COLUMN comments.parent_id IS '回复评论需要用到';
COMMENT ON COLUMN comments.user_id IS '用户ID，标识评论的人';
COMMENT ON COLUMN comments.post_id IS '帖子ID，标识被评论的帖子';
COMMENT ON COLUMN comments.content IS '评论内容';
COMMENT ON COLUMN comments.created_at IS '评论时间';


--存储用户收藏的帖子信息
CREATE TABLE collections
(
    id         VARCHAR(64) PRIMARY KEY NOT NULL,
    user_id    VARCHAR(64)             NOT NULL,
    post_id    VARCHAR(64)             NOT NULL,
    created_at TIMESTAMP               NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE collections IS '存储用户收藏的帖子信息';
COMMENT ON COLUMN collections.id IS '收藏ID，唯一标识收藏';
COMMENT ON COLUMN collections.user_id IS '用户ID，标识收藏的人';
COMMENT ON COLUMN collections.post_id IS '帖子ID，标识被收藏的帖子';
COMMENT ON COLUMN collections.created_at IS '收藏时间';

--存储推荐的帖子信息
CREATE TABLE reposts
(
    id         VARCHAR(64) PRIMARY KEY NOT NULL,
    user_id    VARCHAR(64)             NOT NULL,
    post_id    VARCHAR(64)             NOT NULL,
    status     INTEGER,
    created_at TIMESTAMP               NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE reposts IS '存储推荐的帖子信息';
COMMENT ON COLUMN reposts.id IS '推荐ID，唯一标识推荐';
COMMENT ON COLUMN reposts.user_id IS '用户ID';
COMMENT ON COLUMN reposts.post_id IS '帖子ID，标识被推荐的帖子';
COMMENT ON COLUMN reposts.status IS '推荐状态，1：Bullish，2:Bearish';
COMMENT ON COLUMN reposts.created_at IS '推荐时间';

--追加
alter table tags add column tag_color varchar(32) not null ;
comment on column tags.tag_color is '标签颜色';