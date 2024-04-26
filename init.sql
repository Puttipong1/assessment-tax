CREATE TABLE IF NOT EXISTS deductions
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    type character varying(40) COLLATE pg_catalog."default" NOT NULL,
    amount numeric NOT NULL,
    CONSTRAINT deductions_pkey PRIMARY KEY (id),
    CONSTRAINT deductions_type UNIQUE (type)
);

INSERT INTO "deductions" ("type", "amount") VALUES ('personal', 50000.0)  ON conflict ("type") do nothing; 
INSERT INTO "deductions" ("type", "amount") VALUES ('k-receipt', 50000.0) ON conflict ("type") do nothing;
INSERT INTO "deductions" ("type", "amount") VALUES ('donation', 100000.0) ON conflict ("type") do nothing;