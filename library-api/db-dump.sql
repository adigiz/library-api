--
-- PostgreSQL database dump
--

-- Dumped from database version 14.4 (Debian 14.4-1.pgdg110+1)
-- Dumped by pg_dump version 14.4 (Debian 14.4-1.pgdg110+1)

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
-- Name: authors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.authors (
                                id integer NOT NULL,
                                name character varying(50) NOT NULL,
                                created_at timestamp without time zone DEFAULT now(),
                                updated_at timestamp without time zone DEFAULT now(),
                                deleted_at timestamp without time zone
);


ALTER TABLE public.authors OWNER TO postgres;

--
-- Name: authors_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.authors_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.authors_id_seq OWNER TO postgres;

--
-- Name: authors_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.authors_id_seq OWNED BY public.authors.id;


--
-- Name: books; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.books (
                              id integer NOT NULL,
                              title character varying NOT NULL,
                              description character varying,
                              cover character varying,
                              created_at timestamp without time zone DEFAULT now(),
                              updated_at timestamp without time zone DEFAULT now(),
                              deleted_at timestamp without time zone,
                              author_id integer,
                              stock integer
);


ALTER TABLE public.books OWNER TO postgres;

--
-- Name: book_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.book_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.book_id_seq OWNER TO postgres;

--
-- Name: book_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.book_id_seq OWNED BY public.books.id;


--
-- Name: borrowing_records; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.borrowing_records (
                                          id integer NOT NULL,
                                          user_id integer NOT NULL,
                                          book_id integer NOT NULL,
                                          status character varying NOT NULL,
                                          borrowing_date timestamp without time zone,
                                          returning_date timestamp without time zone,
                                          created_at timestamp without time zone DEFAULT now(),
                                          updated_at timestamp without time zone DEFAULT now(),
                                          deleted_at timestamp without time zone
);


ALTER TABLE public.borrowing_records OWNER TO postgres;

--
-- Name: borrowing_records_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.borrowing_records_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.borrowing_records_id_seq OWNER TO postgres;

--
-- Name: borrowing_records_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.borrowing_records_id_seq OWNED BY public.borrowing_records.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
                              id integer NOT NULL,
                              name character varying NOT NULL,
                              phone character varying NOT NULL,
                              created_at timestamp without time zone DEFAULT now(),
                              updated_at timestamp without time zone DEFAULT now(),
                              deleted_at timestamp without time zone,
                              email character varying,
                              password character varying
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: authors id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors ALTER COLUMN id SET DEFAULT nextval('public.authors_id_seq'::regclass);


--
-- Name: books id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books ALTER COLUMN id SET DEFAULT nextval('public.book_id_seq'::regclass);


--
-- Name: borrowing_records id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.borrowing_records ALTER COLUMN id SET DEFAULT nextval('public.borrowing_records_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: authors; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.authors (id, name, created_at, updated_at, deleted_at) FROM stdin;
1	Dewa Awidiya	2022-08-01 13:24:56.401478	2022-08-01 13:24:56.419231	\N
\.


--
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.books (id, title, description, cover, created_at, updated_at, deleted_at, author_id, stock) FROM stdin;
24	Doraemon 3	Desc	https://mantap.com/mantap.png	2022-08-01 22:26:30.685187	2022-08-01 22:26:30.685187	\N	1	15
22	Doraemon 2	Desc	https://mantap.com/mantap.png	2022-08-01 22:17:42.912	2022-08-01 22:17:42.912	\N	1	12
18	Doraemon	Kucing sakti	https://google.com	2022-08-01 13:06:55.938551	2022-08-04 00:02:41.705355	\N	1	1
29	Clean Code	Desc	https://google.com/clean-code.png	2022-08-04 14:33:19.602567	2022-08-04 14:33:19.602567	\N	1	0
33	Clean Code 000	Desc	https://google.com/clean-code.png	2022-08-04 15:04:46.787894	2022-08-04 15:04:46.787894	\N	1	0
26	Doraemon 4	Desc	https://mantap.com/mantap.png	2022-08-01 23:09:47.879806	2022-08-01 23:09:47.879806	\N	1	13
35	Clean Code 11	Desc	https://google.com/clean-code.png	2022-08-07 19:42:29.090299	2022-08-07 19:42:29.090299	\N	1	0
37	Clean Code 122	Desc	https://google.com/clean-code.png	2022-08-07 19:42:38.59046	2022-08-07 19:42:38.59046	\N	1	0
39	Clean Code x2	Desc	https://google.com/clean-code.png	2022-08-07 19:42:47.060025	2022-08-07 19:42:47.060025	\N	1	0
41	Clean Code xx	Desc	https://google.com/clean-code.png	2022-08-07 19:50:34.581667	2022-08-07 19:50:34.581667	\N	1	0
\.


--
-- Data for Name: borrowing_records; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.borrowing_records (id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at) FROM stdin;
21	1	26	BORROWED	2022-08-04 00:22:29.578141	\N	2022-08-04 00:22:29.57925	2022-08-04 00:22:29.57925	\N
22	1	26	BORROWED	2022-08-04 00:22:30.068029	\N	2022-08-04 00:22:30.069378	2022-08-04 00:22:30.069378	\N
23	1	26	BORROWED	2022-08-04 00:22:30.57424	\N	2022-08-04 00:22:30.575104	2022-08-04 00:22:30.575104	\N
24	1	26	BORROWED	2022-08-04 00:22:30.975591	\N	2022-08-04 00:22:30.977449	2022-08-04 00:22:30.977449	\N
25	1	26	BORROWED	2022-08-04 00:23:43.310639	\N	2022-08-04 00:23:43.312395	2022-08-04 00:23:43.312395	\N
26	1	26	BORROWED	2022-08-04 00:24:28.431009	\N	2022-08-04 00:24:28.433619	2022-08-04 00:24:28.433619	\N
17	1	26	BORROWED	2022-08-04 00:14:57.801704	2022-08-04 00:40:31.24046	2022-08-04 00:14:57.808796	2022-08-04 00:40:31.24506	\N
18	1	26	BORROWED	2022-08-04 00:16:39.460406	2022-08-04 00:40:37.130275	2022-08-04 00:16:39.468246	2022-08-04 00:40:37.132725	\N
19	1	26	RETURNED	2022-08-04 00:21:38.386233	2022-08-04 00:47:50.124205	2022-08-04 00:21:38.393583	2022-08-04 00:47:50.128605	\N
27	1	26	BORROWED	2022-08-04 15:16:10.851212	\N	2022-08-04 15:16:10.853839	2022-08-04 15:16:10.853839	\N
28	1	26	BORROWED	2022-08-05 15:49:24.715984	\N	2022-08-05 15:49:24.718043	2022-08-05 15:49:24.718043	\N
20	1	26	RETURNED	2022-08-04 00:22:22.723335	2022-08-05 15:49:46.359005	2022-08-04 00:22:22.741981	2022-08-05 15:49:46.363931	\N
29	1	26	BORROWED	2022-08-05 15:49:49.957909	\N	2022-08-05 15:49:49.961759	2022-08-05 15:49:49.961759	\N
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, phone, created_at, updated_at, deleted_at, email, password) FROM stdin;
2	Tama	123456	2022-08-02 04:49:19.639516	2022-08-02 04:49:19.639516	\N	tama@gmail.com	P@ssw0rd
1	Dewa User	08599922991	2022-08-02 04:04:42.465784	2022-08-02 04:04:42.478492	\N	dewa@gmail.com	P@ssw0rd
\.


--
-- Name: authors_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.authors_id_seq', 1, true);


--
-- Name: book_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.book_id_seq', 43, true);


--
-- Name: borrowing_records_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.borrowing_records_id_seq', 29, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 2, true);


--
-- Name: authors authors_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.authors
    ADD CONSTRAINT authors_pk PRIMARY KEY (id);


--
-- Name: books book_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT book_pk PRIMARY KEY (id);


--
-- Name: borrowing_records borrowing_records_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.borrowing_records
    ADD CONSTRAINT borrowing_records_pk PRIMARY KEY (id);


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- Name: books_title_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX books_title_uindex ON public.books USING btree (title);


--
-- Name: books books_authors_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_authors_id_fk FOREIGN KEY (author_id) REFERENCES public.authors(id);


--
-- Name: borrowing_records borrowing_records_books_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.borrowing_records
    ADD CONSTRAINT borrowing_records_books_id_fk FOREIGN KEY (book_id) REFERENCES public.books(id);


--
-- Name: borrowing_records borrowing_records_users_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.borrowing_records
    ADD CONSTRAINT borrowing_records_users_id_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--
