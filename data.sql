create table bad_driver_statuses
(
    bad_driver_status_id serial  not null
        constraint bad_driver_statuses_pk
            primary key,
    title                varchar not null,
    description          text
);

create unique index bad_driver_statuses_title_uindex
    on bad_driver_statuses (title);

create table bad_drivers
(
    bad_driver_id    serial            not null
        constraint bad_drivers_pk
            primary key,
    name             varchar           not null,
    reason           varchar,
    status           integer           not null
        constraint bad_drivers_bad_driver_statuses_bad_driver_status_id_fk
            references bad_driver_statuses,
    accident_count   integer default 0 not null,
    ticket_count     integer default 0 not null,
    karens_irritated integer default 0 not null
);

INSERT INTO public.bad_driver_statuses (bad_driver_status_id, title, description) VALUES (1, 'Terrible', 'Universally regarded as a menace behind the wheel.');
INSERT INTO public.bad_driver_statuses (bad_driver_status_id, title, description) VALUES (2, 'Unsafe', 'Passengers frequently grab hand holds with legitimate fear.');
INSERT INTO public.bad_driver_statuses (bad_driver_status_id, title, description) VALUES (3, 'Catastrophe', 'It would be easier to count the drives that did not end in tragedy.');
INSERT INTO public.bad_driver_statuses (bad_driver_status_id, title, description) VALUES (4, 'Murderous', 'You will know they are coming by the screeching tires and the screaming pedestrians. Run. Do not walk, run.');
INSERT INTO public.bad_driver_statuses (bad_driver_status_id, title, description) VALUES (5, 'Saint', 'Angelic. Legends say as this driver passes a faint glow can be seen where the tires met the pavement.');
INSERT INTO public.bad_driver_statuses (bad_driver_status_id, title, description) VALUES (6, 'Normal', 'No tickets, no moving infractions, always goes the speed limit.');

INSERT INTO public.bad_drivers (bad_driver_id, name, reason, status, accident_count, ticket_count, karens_irritated) VALUES (1, 'Chris Goins', 'Likes to swerve between lanes, potential road rage.', 2, 3, 3, 10);
INSERT INTO public.bad_drivers (bad_driver_id, name, reason, status, accident_count, ticket_count, karens_irritated) VALUES (2, 'Tommy Matthews', 'There is a chance he could bring about the apocolypse.', 4, 900, 100, 1000);
INSERT INTO public.bad_drivers (bad_driver_id, name, reason, status, accident_count, ticket_count, karens_irritated) VALUES (3, 'Joe Mills', 'No one should be forced to witness how bad he drives.', 3, 100, 100, 500);
INSERT INTO public.bad_drivers (bad_driver_id, name, reason, status, accident_count, ticket_count, karens_irritated) VALUES (4, 'Kevin Kennedy', 'His driving could potentially end all wars as a secondary effect.', 5, -3, 0, 10000);
INSERT INTO public.bad_drivers (bad_driver_id, name, reason, status, accident_count, ticket_count, karens_irritated) VALUES (5, 'Kashif Mansoor', 'He is a hell raiser, really. A bad boy behind the wheel.', 1, 33, 20, 0);
INSERT INTO public.bad_drivers (bad_driver_id, name, reason, status, accident_count, ticket_count, karens_irritated) VALUES (6, 'Scott Weberg', 'The speed limit is his limit.', 6, 33, 20, 0);
INSERT INTO public.bad_drivers (bad_driver_id, name, reason, status, accident_count, ticket_count, karens_irritated) VALUES (7, 'Ananda SanKaran', 'Has never met a yellow light he didn''t see as a challenge.', 2, 33, 20, 0);