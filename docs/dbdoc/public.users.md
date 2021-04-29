# public.users

## Description

## Columns

| Name | Type | Default | Nullable | Children | Parents | Comment |
| ---- | ---- | ------- | -------- | -------- | ------- | ------- |
| id | integer | nextval('users_id_seq'::regclass) | false | [public.posts](public.posts.md) [public.comments](public.comments.md) |  | ユーザーID |
| name | text |  | false |  |  | ユーザー名 |
| password_hash | text |  | false |  |  | パスワードハッシュ |
| created_at | timestamp with time zone | now() | false |  |  | 作成日時 |
| updated_at | timestamp with time zone | now() | false |  |  | 更新日時 |

## Constraints

| Name | Type | Definition |
| ---- | ---- | ---------- |
| users_pkey | PRIMARY KEY | PRIMARY KEY (id) |
| users_name_key | UNIQUE | UNIQUE (name) |

## Indexes

| Name | Definition |
| ---- | ---------- |
| users_pkey | CREATE UNIQUE INDEX users_pkey ON public.users USING btree (id) |
| users_name_key | CREATE UNIQUE INDEX users_name_key ON public.users USING btree (name) |

## Relations

![er](public.users.svg)

---

> Generated by [tbls](https://github.com/k1LoW/tbls)