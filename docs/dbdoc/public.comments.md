# public.comments

## Description

## Columns

| Name | Type | Default | Nullable | Children | Parents | Comment |
| ---- | ---- | ------- | -------- | -------- | ------- | ------- |
| id | integer | nextval('comments_id_seq'::regclass) | false |  |  | コメントID |
| user_id | bigint |  | false |  | [public.users](public.users.md) | ユーザーID |
| post_id | bigint |  | false |  | [public.posts](public.posts.md) | 投稿ID |
| comment | text |  | false |  |  | コメント本文 |
| created_at | timestamp with time zone | now() | false |  |  | 作成日時 |
| updated_at | timestamp with time zone | now() | false |  |  | 更新日時 |

## Constraints

| Name | Type | Definition |
| ---- | ---- | ---------- |
| comments_user_id_fkey | FOREIGN KEY | FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE |
| comments_post_id_fkey | FOREIGN KEY | FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE |
| comments_pkey | PRIMARY KEY | PRIMARY KEY (id) |

## Indexes

| Name | Definition |
| ---- | ---------- |
| comments_pkey | CREATE UNIQUE INDEX comments_pkey ON public.comments USING btree (id) |

## Relations

![er](public.comments.svg)

---

> Generated by [tbls](https://github.com/k1LoW/tbls)
