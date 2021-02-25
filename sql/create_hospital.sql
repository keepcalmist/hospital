CREATE TABLE IF NOT EXISTS hospital (
    id serial  primary key,
    title text not null,
    created_on timestamp with time zone,
    description text,
);