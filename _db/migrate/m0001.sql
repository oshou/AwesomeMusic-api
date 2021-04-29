-- +migrate Up
CREATE TABLE comment (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES user(id) ON DELETE CASCADE,
    post_id BIGINT NOT NULL REFERENCES post(id) ON DELETE CASCADE,
    comment TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

ALTER TABLE
    public.comment OWNER TO postgres;

COMMENT ON COLUMN public.comment.id IS 'コメントID';

COMMENT ON COLUMN public.comment.user_id IS 'ユーザーID';

COMMENT ON COLUMN public.comment.post_id IS '投稿ID';

COMMENT ON COLUMN public.comment.comment IS 'コメント本文';

COMMENT ON COLUMN public.comment.created_at IS '作成日時';

COMMENT ON COLUMN public.comment.updated_at IS '更新日時';

CREATE TABLE post (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES user(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

ALTER TABLE
    public.post OWNER TO postgres;

COMMENT ON COLUMN public.post.id IS '投稿ID';

COMMENT ON COLUMN public.post.user_id IS 'ユーザーID';

COMMENT ON COLUMN public.post.title IS '投稿タイトル';

COMMENT ON COLUMN public.post.url IS '投稿URL';

COMMENT ON COLUMN public.post.message IS '投稿メッセージ';

COMMENT ON COLUMN public.post.created_at IS '作成日時';

COMMENT ON COLUMN public.post.updated_at IS '更新日時';

CREATE TABLE public.post_tag (
    post_id BIGINT NOT NULL REFERENCES post(id) ON DELETE CASCADE,
    tag_id BIGINT NOT NULL REFERENCES tag(id) ON DELETE CASCADE
);

ALTER TABLE
    post_tag OWNER TO postgres;

COMMENT ON COLUMN public.post_tag.post_id IS '投稿ID';

COMMENT ON COLUMN public.post_tag.tag_id IS 'タグID';

CREATE TABLE tag (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(name)
);

ALTER TABLE
    public.tag OWNER TO postgres;

COMMENT ON COLUMN public.tag.id IS 'タグID';

COMMENT ON COLUMN public.tag.name IS 'タグ名';

COMMENT ON COLUMN public.tag.created_at IS '作成日時';

COMMENT ON COLUMN public.tag.updated_at IS '更新日時';

CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(name)
);

ALTER TABLE
    public."user" OWNER TO postgres;

COMMENT ON COLUMN public."user".id IS 'ユーザーID';

COMMENT ON COLUMN public."user".name IS 'ユーザー名';

COMMENT ON COLUMN public."user".password_hash IS 'パスワードハッシュ';

COMMENT ON COLUMN public."user".created_at IS '作成日時';

COMMENT ON COLUMN public."user".updated_at IS '更新日時';

-- +migrate Down
DROP TABLE comment;

DROP TABLE post;

DROP TABLE post_tag;

DROP TABLE tag;

DROP TABLE "user";
