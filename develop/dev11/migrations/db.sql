-- public.items definition

-- Drop table

-- DROP TABLE public.items;

CREATE TABLE Events (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	user_id int4 NULL,
	date text NULL,
    event text NULL,
	description text NULL,
	CONSTRAINT items_pk PRIMARY KEY (id)
);
