--
-- PostgreSQL database dump
--

-- Dumped from database version 10.5
-- Dumped by pg_dump version 11.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: json_web_tokens; Type: TABLE; Schema: public; Owner: yannickchiasson
--

CREATE TABLE public.json_web_tokens (
    id uuid NOT NULL,
    resource_id integer DEFAULT 0 NOT NULL,
    resource_type character varying(255) NOT NULL,
    token text NOT NULL,
    expires_at timestamp without time zone,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.json_web_tokens OWNER TO yannickchiasson;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: yannickchiasson
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO yannickchiasson;

--
-- Name: users; Type: TABLE; Schema: public; Owner: yannickchiasson
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    email character varying(255) DEFAULT ''::character varying NOT NULL,
    password_hash text NOT NULL,
    name character varying(255) DEFAULT ''::character varying NOT NULL,
    avatar_url character varying(255) DEFAULT ''::character varying NOT NULL,
    deleted_at timestamp without time zone,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO yannickchiasson;

--
-- Name: json_web_tokens_i01; Type: INDEX; Schema: public; Owner: yannickchiasson
--

CREATE INDEX json_web_tokens_i01 ON public.json_web_tokens USING btree (resource_id, resource_type, expires_at);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: yannickchiasson
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: users_i01; Type: INDEX; Schema: public; Owner: yannickchiasson
--

CREATE UNIQUE INDEX users_i01 ON public.users USING btree (email);


--
-- PostgreSQL database dump complete
--

