CREATE TABLE users
(
    id         serial      not null primary key,
    created_at timestamp   not null default current_timestamp,
    login      varchar(50) not null,
    apikey     varchar(100)
);

CREATE FUNCTION generate_apikey() RETURNS trigger AS
$generate_apikey$
BEGIN
    UPDATE users SET apikey = '' WHERE id = NEW.id;
END;
$generate_apikey$ LANGUAGE plpgsql;

CREATE TRIGGER generate_apikey
    AFTER INSERT
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE generate_apikey();