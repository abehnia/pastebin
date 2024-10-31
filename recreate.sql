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
-- Name: pastebin; Type: TABLE; Schema: public; Owner: ubuntu
--

DROP TABLE IF EXISTS public.pastebin;

CREATE TABLE public.pastebin (
    id uuid NOT NULL,
    "timestamp" bigint NOT NULL,
    title text NOT NULL,
    content TEXT NOT NULL,
    seen_counter INTEGER NOT NULL DEFAULT 0,
    star_counter INTEGER NOT NULL DEFAULT 0
);


ALTER TABLE public.pastebin OWNER TO ubuntu;

--
-- Name: pastebin pastebin_pkey; Type: CONSTRAINT; Schema: public; Owner: ubuntu
--

ALTER TABLE ONLY public.pastebin
    ADD CONSTRAINT pastebin_pkey PRIMARY KEY (id);


--
-- Name: TABLE pastebin; Type: ACL; Schema: public; Owner: ubuntu
--

GRANT ALL ON TABLE public.pastebin TO ubuntu;

--
-- PostgreSQL database dump complete
--

