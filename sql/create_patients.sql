CREATE TABLE IF NOT EXISTS patients
(
    id serial primary key unique,
    name TEXT NOT NULL,
    last_name text NOT NULL,
    hospital_id integer REFERENCES hospital(id) ON DELETE CASCADE,
    doctor_id integer REFERENCES doctors(id) ON DELETE SET NULL
);