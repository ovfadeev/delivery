CREATE TABLE users
(
    id         serial      not null primary key,
    created_at timestamp   not null default current_timestamp,
    login      varchar(50) not null unique,
    apikey     varchar(100)
);

CREATE FUNCTION generate_apikey() RETURNS trigger AS
$generate_apikey$
BEGIN
    UPDATE users SET apikey = uuid_generate_v5(uuid_ns_oid(), concat(NEW.created_at, NEW.login)) WHERE id = NEW.id;
    RETURN NEW;
END;
$generate_apikey$ LANGUAGE plpgsql;

CREATE TRIGGER generate_apikey
    AFTER INSERT
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE generate_apikey();