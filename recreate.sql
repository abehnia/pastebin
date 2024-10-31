--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4 (Ubuntu 16.4-0ubuntu0.24.04.2)
-- Dumped by pg_dump version 16.4 (Ubuntu 16.4-0ubuntu0.24.04.2)

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
-- Name: pastebin; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pastebin (
    id bigint NOT NULL,
    "timestamp" bigint NOT NULL,
    title text NOT NULL,
    content text NOT NULL
);


ALTER TABLE public.pastebin OWNER TO postgres;

--
-- Name: pastebin_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.pastebin_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.pastebin_id_seq OWNER TO postgres;

--
-- Name: pastebin_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.pastebin_id_seq OWNED BY public.pastebin.id;


--
-- Name: pastebin id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pastebin ALTER COLUMN id SET DEFAULT nextval('public.pastebin_id_seq'::regclass);


--
-- Name: pastebin pastebin_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pastebin
    ADD CONSTRAINT pastebin_pkey PRIMARY KEY (id);


--
-- Name: TABLE pastebin; Type: ACL; Schema: public; Owner: postgres
--

GRANT ALL ON TABLE public.pastebin TO ubuntu;


--
-- Name: SEQUENCE pastebin_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.pastebin_id_seq TO ubuntu;


--
-- PostgreSQL database dump complete
--

