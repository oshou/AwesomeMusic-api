--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3
-- Dumped by pg_dump version 12.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: comment; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.comment (
    id integer NOT NULL,
    user_id bigint NOT NULL,
    post_id bigint NOT NULL,
    comment text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.comment OWNER TO root;

--
-- Name: COLUMN comment.id; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.comment.id IS 'コメントID';


--
-- Name: COLUMN comment.user_id; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.comment.user_id IS 'ユーザーID';


--
-- Name: COLUMN comment.post_id; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.comment.post_id IS '投稿ID';


--
-- Name: COLUMN comment.comment; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.comment.comment IS 'コメント本文';


--
-- Name: COLUMN comment.created_at; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.comment.created_at IS '作成日時';


--
-- Name: COLUMN comment.updated_at; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.comment.updated_at IS '更新日時';


--
-- Name: comment_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.comment_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.comment_id_seq OWNER TO root;

--
-- Name: comment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.comment_id_seq OWNED BY public.comment.id;


--
-- Name: migrations; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.migrations (
    id text NOT NULL,
    applied_at timestamp with time zone
);


ALTER TABLE public.migrations OWNER TO root;

--
-- Name: post; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.post (
    id integer NOT NULL,
    user_id bigint NOT NULL,
    title text NOT NULL,
    url text NOT NULL,
    message text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.post OWNER TO root;

--
-- Name: COLUMN post.id; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.post.id IS '投稿ID';


--
-- Name: COLUMN post.user_id; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.post.user_id IS 'ユーザーID';


--
-- Name: COLUMN post.title; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.post.title IS '投稿タイトル';


--
-- Name: COLUMN post.url; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.post.url IS '投稿URL';


--
-- Name: COLUMN post.message; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.post.message IS '投稿メッセージ';


--
-- Name: COLUMN post.created_at; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.post.created_at IS '作成日時';


--
-- Name: COLUMN post.updated_at; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.post.updated_at IS '更新日時';


--
-- Name: post_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.post_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_id_seq OWNER TO root;

--
-- Name: post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.post_id_seq OWNED BY public.post.id;


--
-- Name: post_tag; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.post_tag (
    post_id bigint NOT NULL,
    tag_id bigint NOT NULL
);


ALTER TABLE public.post_tag OWNER TO root;

--
-- Name: COLUMN post_tag.post_id; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.post_tag.post_id IS '投稿ID';


--
-- Name: COLUMN post_tag.tag_id; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.post_tag.tag_id IS 'タグID';


--
-- Name: tag; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.tag (
    id integer NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.tag OWNER TO root;

--
-- Name: COLUMN tag.id; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.tag.id IS 'タグID';


--
-- Name: COLUMN tag.name; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.tag.name IS 'タグ名';


--
-- Name: COLUMN tag.created_at; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.tag.created_at IS '作成日時';


--
-- Name: COLUMN tag.updated_at; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.tag.updated_at IS '更新日時';


--
-- Name: tag_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.tag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tag_id_seq OWNER TO root;

--
-- Name: tag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.tag_id_seq OWNED BY public.tag.id;


--
-- Name: user; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public."user" (
    id integer NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public."user" OWNER TO root;

--
-- Name: COLUMN "user".id; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public."user".id IS 'ユーザーID';


--
-- Name: COLUMN "user".name; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public."user".name IS 'ユーザー名';


--
-- Name: COLUMN "user".created_at; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public."user".created_at IS '作成日時';


--
-- Name: COLUMN "user".updated_at; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public."user".updated_at IS '更新日時';


--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_id_seq OWNER TO root;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.user_id_seq OWNED BY public."user".id;


--
-- Name: comment id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.comment ALTER COLUMN id SET DEFAULT nextval('public.comment_id_seq'::regclass);


--
-- Name: post id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.post ALTER COLUMN id SET DEFAULT nextval('public.post_id_seq'::regclass);


--
-- Name: tag id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.tag ALTER COLUMN id SET DEFAULT nextval('public.tag_id_seq'::regclass);


--
-- Name: user id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- Name: comment comment_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.comment
    ADD CONSTRAINT comment_pkey PRIMARY KEY (id);


--
-- Name: migrations migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.migrations
    ADD CONSTRAINT migrations_pkey PRIMARY KEY (id);


--
-- Name: post post_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.post
    ADD CONSTRAINT post_pkey PRIMARY KEY (id);


--
-- Name: tag tag_name_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.tag
    ADD CONSTRAINT tag_name_key UNIQUE (name);


--
-- Name: tag tag_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.tag
    ADD CONSTRAINT tag_pkey PRIMARY KEY (id);


--
-- Name: user user_name_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_name_key UNIQUE (name);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

