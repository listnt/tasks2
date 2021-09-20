CREATE OR REPLACE PROCEDURE public.insert_into_events(user_id_v integer, date_v text, event_v text, description_v text)
    LANGUAGE sql
AS $procedure$
INSERT INTO events(user_id,date ,event ,description )
VALUES (user_id_v,date_v ,event_v ,description_v);
$procedure$
;