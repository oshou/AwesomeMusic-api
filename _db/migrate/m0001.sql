-- +migrate Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(name)
);

ALTER TABLE
    public.users OWNER TO postgres;

COMMENT ON COLUMN public.users.id IS 'ユーザーID';

COMMENT ON COLUMN public.users.name IS 'ユーザー名';

COMMENT ON COLUMN public.users.password_hash IS 'パスワードハッシュ';

COMMENT ON COLUMN public.users.created_at IS '作成日時';

COMMENT ON COLUMN public.users.updated_at IS '更新日時';

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

ALTER TABLE
    public.posts OWNER TO postgres;

COMMENT ON COLUMN public.posts.id IS '投稿ID';

COMMENT ON COLUMN public.posts.user_id IS 'ユーザーID';

COMMENT ON COLUMN public.posts.title IS '投稿タイトル';

COMMENT ON COLUMN public.posts.url IS '投稿URL';

COMMENT ON COLUMN public.posts.message IS '投稿メッセージ';

COMMENT ON COLUMN public.posts.created_at IS '作成日時';

COMMENT ON COLUMN public.posts.updated_at IS '更新日時';

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    post_id BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    comment TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

ALTER TABLE
    public.comments OWNER TO postgres;

COMMENT ON COLUMN public.comments.id IS 'コメントID';

COMMENT ON COLUMN public.comments.user_id IS 'ユーザーID';

COMMENT ON COLUMN public.comments.post_id IS '投稿ID';

COMMENT ON COLUMN public.comments.comment IS 'コメント本文';

COMMENT ON COLUMN public.comments.created_at IS '作成日時';

COMMENT ON COLUMN public.comments.updated_at IS '更新日時';

CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(name)
);

ALTER TABLE
    public.tags OWNER TO postgres;

COMMENT ON COLUMN public.tags.id IS 'タグID';

COMMENT ON COLUMN public.tags.name IS 'タグ名';

COMMENT ON COLUMN public.tags.created_at IS '作成日時';

COMMENT ON COLUMN public.tags.updated_at IS '更新日時';

CREATE TABLE public.post_tag (
    post_id BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    tag_id BIGINT NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY(post_id, tag_id)
);

ALTER TABLE
    post_tag OWNER TO postgres;

COMMENT ON COLUMN public.post_tag.post_id IS '投稿ID';

COMMENT ON COLUMN public.post_tag.tag_id IS 'タグID';

-- +migrate Down
DROP TABLE comments;

DROP TABLE posts;

DROP TABLE post_tag;

DROP TABLE tags;

DROP TABLE users;
