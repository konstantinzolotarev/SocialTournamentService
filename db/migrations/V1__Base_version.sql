--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.3
-- Dumped by pg_dump version 9.6.3

-- Started on 2017-06-25 00:53:59

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 2179 (class 1262 OID 16910)
-- Name: TournamenTGame; Type: DATABASE; Schema: -; Owner: GameRole
--

CREATE DATABASE "TournamenTGame" WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'Russian_Russia.1251' LC_CTYPE = 'Russian_Russia.1251';


ALTER DATABASE "TournamenTGame" OWNER TO "GameRole";

\connect "TournamenTGame"

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4 (class 2615 OID 16911)
-- Name: game; Type: SCHEMA; Schema: -; Owner: GameRole
--

CREATE SCHEMA game;


ALTER SCHEMA game OWNER TO "GameRole";

--
-- TOC entry 2180 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA game; Type: COMMENT; Schema: -; Owner: GameRole
--

COMMENT ON SCHEMA game IS 'standard public schema';


--
-- TOC entry 1 (class 3079 OID 12387)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2182 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = game, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 186 (class 1259 OID 16912)
-- Name: Backers; Type: TABLE; Schema: game; Owner: GameRole
--

CREATE TABLE "Backers" (
    id bigint NOT NULL,
    "playerId" bigint NOT NULL,
    "tournamentId" bigint NOT NULL,
    share numeric(10,2) NOT NULL
);


ALTER TABLE "Backers" OWNER TO "GameRole";

--
-- TOC entry 187 (class 1259 OID 16915)
-- Name: Backers_id_seq; Type: SEQUENCE; Schema: game; Owner: GameRole
--

CREATE SEQUENCE "Backers_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE "Backers_id_seq" OWNER TO "GameRole";

--
-- TOC entry 2183 (class 0 OID 0)
-- Dependencies: 187
-- Name: Backers_id_seq; Type: SEQUENCE OWNED BY; Schema: game; Owner: GameRole
--

ALTER SEQUENCE "Backers_id_seq" OWNED BY "Backers".id;


--
-- TOC entry 188 (class 1259 OID 16917)
-- Name: Players; Type: TABLE; Schema: game; Owner: GameRole
--

CREATE TABLE "Players" (
    id bigint NOT NULL,
    "playerName" text NOT NULL UNIQUE, 
    points numeric(10,2) NOT NULL
);


ALTER TABLE "Players" OWNER TO "GameRole";

--
-- TOC entry 189 (class 1259 OID 16923)
-- Name: Players_id_seq; Type: SEQUENCE; Schema: game; Owner: GameRole
--

CREATE SEQUENCE "Players_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE "Players_id_seq" OWNER TO "GameRole";

--
-- TOC entry 2184 (class 0 OID 0)
-- Dependencies: 189
-- Name: Players_id_seq; Type: SEQUENCE OWNED BY; Schema: game; Owner: GameRole
--

ALTER SEQUENCE "Players_id_seq" OWNED BY "Players".id;


--
-- TOC entry 190 (class 1259 OID 16925)
-- Name: Status; Type: TABLE; Schema: game; Owner: GameRole
--

CREATE TABLE "Status" (
    id bigint NOT NULL,
    name text NOT NULL
);


ALTER TABLE "Status" OWNER TO "GameRole";

--
-- TOC entry 191 (class 1259 OID 16931)
-- Name: Status_id_seq; Type: SEQUENCE; Schema: game; Owner: GameRole
--

CREATE SEQUENCE "Status_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE "Status_id_seq" OWNER TO "GameRole";

--
-- TOC entry 2185 (class 0 OID 0)
-- Dependencies: 191
-- Name: Status_id_seq; Type: SEQUENCE OWNED BY; Schema: game; Owner: GameRole
--

ALTER SEQUENCE "Status_id_seq" OWNED BY "Status".id;


--
-- TOC entry 192 (class 1259 OID 16933)
-- Name: TournamentPlayers; Type: TABLE; Schema: game; Owner: GameRole
--

CREATE TABLE "TournamentPlayers" (
    id bigint NOT NULL,
    "playerId" bigint NOT NULL,
    "tournamentId" bigint NOT NULL,
    share numeric(10,2) NOT NULL
);


ALTER TABLE "TournamentPlayers" OWNER TO "GameRole";

--
-- TOC entry 193 (class 1259 OID 16936)
-- Name: TournamentPlayers_id_seq; Type: SEQUENCE; Schema: game; Owner: GameRole
--

CREATE SEQUENCE "TournamentPlayers_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE "TournamentPlayers_id_seq" OWNER TO "GameRole";

--
-- TOC entry 2186 (class 0 OID 0)
-- Dependencies: 193
-- Name: TournamentPlayers_id_seq; Type: SEQUENCE OWNED BY; Schema: game; Owner: GameRole
--

ALTER SEQUENCE "TournamentPlayers_id_seq" OWNED BY "TournamentPlayers".id;


--
-- TOC entry 194 (class 1259 OID 16938)
-- Name: Tournaments; Type: TABLE; Schema: game; Owner: GameRole
--

CREATE TABLE "Tournaments" (
    id bigint NOT NULL,
    deposit numeric(10,2) NOT NULL,
    "statusId" bigint NOT NULL
);


ALTER TABLE "Tournaments" OWNER TO "GameRole";

--
-- TOC entry 195 (class 1259 OID 16941)
-- Name: Tournaments_id_seq; Type: SEQUENCE; Schema: game; Owner: GameRole
--

CREATE SEQUENCE "Tournaments_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE "Tournaments_id_seq" OWNER TO "GameRole";

--
-- TOC entry 2187 (class 0 OID 0)
-- Dependencies: 195
-- Name: Tournaments_id_seq; Type: SEQUENCE OWNED BY; Schema: game; Owner: GameRole
--

ALTER SEQUENCE "Tournaments_id_seq" OWNED BY "Tournaments".id;


--
-- TOC entry 2028 (class 2604 OID 16943)
-- Name: Backers id; Type: DEFAULT; Schema: game; Owner: GameRole
--

ALTER TABLE ONLY "Backers" ALTER COLUMN id SET DEFAULT nextval('"Backers_id_seq"'::regclass);


--
-- TOC entry 2029 (class 2604 OID 16944)
-- Name: Players id; Type: DEFAULT; Schema: game; Owner: GameRole
--

ALTER TABLE ONLY "Players" ALTER COLUMN id SET DEFAULT nextval('"Players_id_seq"'::regclass);


--
-- TOC entry 2030 (class 2604 OID 16945)
-- Name: Status id; Type: DEFAULT; Schema: game; Owner: GameRole
--

ALTER TABLE ONLY "Status" ALTER COLUMN id SET DEFAULT nextval('"Status_id_seq"'::regclass);


--
-- TOC entry 2031 (class 2604 OID 16946)
-- Name: TournamentPlayers id; Type: DEFAULT; Schema: game; Owner: GameRole
--

ALTER TABLE ONLY "TournamentPlayers" ALTER COLUMN id SET DEFAULT nextval('"TournamentPlayers_id_seq"'::regclass);


--
-- TOC entry 2032 (class 2604 OID 16947)
-- Name: Tournaments id; Type: DEFAULT; Schema: game; Owner: GameRole
--

ALTER TABLE ONLY "Tournaments" ALTER COLUMN id SET DEFAULT nextval('"Tournaments_id_seq"'::regclass);


--
-- TOC entry 2165 (class 0 OID 16912)
-- Dependencies: 186
-- Data for Name: Backers; Type: TABLE DATA; Schema: game; Owner: GameRole
--



--
-- TOC entry 2188 (class 0 OID 0)
-- Dependencies: 187
-- Name: Backers_id_seq; Type: SEQUENCE SET; Schema: game; Owner: GameRole
--

SELECT pg_catalog.setval('"Backers_id_seq"', 1, false);


--
-- TOC entry 2167 (class 0 OID 16917)
-- Dependencies: 188
-- Data for Name: Players; Type: TABLE DATA; Schema: game; Owner: GameRole
--

--
-- TOC entry 2189 (class 0 OID 0)
-- Dependencies: 189
-- Name: Players_id_seq; Type: SEQUENCE SET; Schema: game; Owner: GameRole
--

SELECT pg_catalog.setval('"Players_id_seq"', 2, true);


--
-- TOC entry 2169 (class 0 OID 16925)
-- Dependencies: 190
-- Data for Name: Status; Type: TABLE DATA; Schema: game; Owner: GameRole
--

INSERT INTO "Status" (id, name) VALUES (1, 'announcement');
INSERT INTO "Status" (id, name) VALUES (2, 'finished');


--
-- TOC entry 2190 (class 0 OID 0)
-- Dependencies: 191
-- Name: Status_id_seq; Type: SEQUENCE SET; Schema: game; Owner: GameRole
--

SELECT pg_catalog.setval('"Status_id_seq"', 2, true);


--
-- TOC entry 2171 (class 0 OID 16933)
-- Dependencies: 192
-- Data for Name: TournamentPlayers; Type: TABLE DATA; Schema: game; Owner: GameRole
--



--
-- TOC entry 2191 (class 0 OID 0)
-- Dependencies: 193
-- Name: TournamentPlayers_id_seq; Type: SEQUENCE SET; Schema: game; Owner: GameRole
--

SELECT pg_catalog.setval('"TournamentPlayers_id_seq"', 1, false);


--
-- TOC entry 2173 (class 0 OID 16938)
-- Dependencies: 194
-- Data for Name: Tournaments; Type: TABLE DATA; Schema: game; Owner: GameRole
--



--
-- TOC entry 2192 (class 0 OID 0)
-- Dependencies: 195
-- Name: Tournaments_id_seq; Type: SEQUENCE SET; Schema: game; Owner: GameRole
--

SELECT pg_catalog.setval('"Tournaments_id_seq"', 1, false);


--
-- TOC entry 2034 (class 2606 OID 16949)
-- Name: Backers Backers_pkey; Type: CONSTRAINT; Schema: game; Owner: GameRole
--

ALTER TABLE ONLY "Backers"
    ADD CONSTRAINT "Backers_pkey" PRIMARY KEY (id);


--
-- TOC entry 2036 (class 2606 OID 16951)
-- Name: Players Players_pkey; Type: CONSTRAINT; Schema: game; Owner: GameRole
--

ALTER TABLE ONLY "Players"
    ADD CONSTRAINT "Players_pkey" PRIMARY KEY (id);


--
-- TOC entry 2038 (class 2606 OID 16953)
-- Name: Status Status_pkey; Type: CONSTRAINT; Schema: game; Owner: GameRole
--

ALTER TABLE ONLY "Status"
    ADD CONSTRAINT "Status_pkey" PRIMARY KEY (id);


--
-- TOC entry 2040 (class 2606 OID 16955)
-- Name: TournamentPlayers TournamentPlayers_pkey; Type: CONSTRAINT; Schema: game; Owner: GameRole
--

ALTER TABLE ONLY "TournamentPlayers"
    ADD CONSTRAINT "TournamentPlayers_pkey" PRIMARY KEY (id);


--
-- TOC entry 2042 (class 2606 OID 16957)
-- Name: Tournaments Tournaments_pkey; Type: CONSTRAINT; Schema: game; Owner: GameRole
--

ALTER TABLE ONLY "Tournaments"
    ADD CONSTRAINT "Tournaments_pkey" PRIMARY KEY (id);


--
-- TOC entry 2043 (class 2606 OID 16958)
-- Name: Backers Backers_playerId_fkey; Type: FK CONSTRAINT; Schema: game; Owner: GameRole
--

ALTER TABLE ONLY "Backers"
    ADD CONSTRAINT "Backers_playerId_fkey" FOREIGN KEY ("playerId") REFERENCES "Players"(id);


--
-- TOC entry 2044 (class 2606 OID 16963)
-- Name: Backers Backers_tournamentId_fkey; Type: FK CONSTRAINT; Schema: game; Owner: GameRole
--

ALTER TABLE ONLY "Backers"
    ADD CONSTRAINT "Backers_tournamentId_fkey" FOREIGN KEY ("tournamentId") REFERENCES "Tournaments"(id);


--
-- TOC entry 2045 (class 2606 OID 16968)
-- Name: TournamentPlayers TournamentPlayers_playerId_fkey; Type: FK CONSTRAINT; Schema: game; Owner: GameRole
--

ALTER TABLE ONLY "TournamentPlayers"
    ADD CONSTRAINT "TournamentPlayers_playerId_fkey" FOREIGN KEY ("playerId") REFERENCES "Players"(id);


--
-- TOC entry 2046 (class 2606 OID 16973)
-- Name: TournamentPlayers TournamentPlayers_tournamentId_fkey; Type: FK CONSTRAINT; Schema: game; Owner: GameRole
--

ALTER TABLE ONLY "TournamentPlayers"
    ADD CONSTRAINT "TournamentPlayers_tournamentId_fkey" FOREIGN KEY ("tournamentId") REFERENCES "Tournaments"(id);


--
-- TOC entry 2047 (class 2606 OID 16978)
-- Name: Tournaments Tournaments_statusId_fkey; Type: FK CONSTRAINT; Schema: game; Owner: GameRole
--

ALTER TABLE ONLY "Tournaments"
    ADD CONSTRAINT "Tournaments_statusId_fkey" FOREIGN KEY ("statusId") REFERENCES "Status"(id);


-- Completed on 2017-06-25 00:54:02

--
-- PostgreSQL database dump complete
--

