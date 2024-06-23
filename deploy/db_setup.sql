CREATE TABLE public.soundbox (
	id serial4 NOT NULL,
	"name" varchar NOT NULL,
	code varchar NOT NULL,
	capacity int2 DEFAULT 50 NOT NULL,
	members integer[],
	CONSTRAINT soundbox_pk PRIMARY KEY (id),
	CONSTRAINT soundbox_unique UNIQUE (code)
);

CREATE TABLE public."users" (
	id bigserial NOT NULL,
	authid varchar NOT NULL,
	nickname varchar NULL,
	picture varchar NULL,
	status varchar NULL,
	"role" _text NULL,
	CONSTRAINT user_pk PRIMARY KEY (id)
	CONSTRAINT users_unique UNIQUE (authid);

);

CREATE TABLE public.user_soundbox (
	user_authid varchar NOT NULL,
	soundbox_id int4 NOT NULL,
	CONSTRAINT user_soundbox_unique UNIQUE (user_authid)
);

ALTER TABLE public.user_soundbox ADD CONSTRAINT user_sb_users_fk FOREIGN KEY (user_authid) REFERENCES public.users(authid);
ALTER TABLE public.user_soundbox ADD CONSTRAINT user_sb_sb_fk FOREIGN KEY (soundbox_id) REFERENCES public.soundbox(id);

CREATE TABLE public.users_token (
	user_id bigserial NOT NULL,
	"token" varchar NULL,
	token_exp timestamp NULL,
	CONSTRAINT token_pk PRIMARY KEY (user_id)
);


-- public.users_token foreign keys

ALTER TABLE public.users_token ADD CONSTRAINT users_token_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id);