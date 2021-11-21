CREATE TABLE pickup
(
    --- идентификатор
    id             serial       not null primary key ,
    --- дата создания
    created_at     timestamp default current_timestamp,
    --- дата обновления
    updated_at     timestamp    not null,
    --- активность
    active         boolean,
    --- код транспортной компании
    provider_id    varchar(100) not null,
    --- название транспортной компании
    provider_name  varchar(100) not null,
    --- тип точки
    type           varchar(100) not null,
    --- тип операции в точке
    type_operation varchar(100) not null,
    --- индекс
    zip            integer      not null,
    --- широта
    latitude       varchar(100) not null,
    --- долгота
    longitude      varchar(100) not null,
    --- название
    name           varchar(100) not null,
    --- код страны
    country_code   varchar(20),
    --- регион
    region         varchar(100),
    --- тип региона
    region_type    varchar(50),
    --- населенный пункт
    city           varchar(100) not null,
    --- тип населенного пункта
    city_type      varchar(50),
    --- код населенного пункта ФИАС
    city_fias      varchar(50),
    --- район
    area           varchar(50),
    --- улица
    street         varchar(100),
    --- тип улицы
    street_type    varchar(10),
    --- дом
    house          varchar(10),
    --- корпус
    block          varchar(10),
    --- офис
    office         varchar(10),
    --- полный адрес
    full_address   varchar(255) not null,
    --- старница
    url            varchar(50),
    --- почта
    email          varchar(50),
    --- телефон
    phone          varchar(50),
    --- время работы
    worktime       varchar(100),
    --- примерочная
    fitting_room   boolean,
    --- наложенный платеж
    cod            boolean,
    --- оплата наличными
    payment_cash   boolean,
    --- оплата картой
    payment_card   boolean,
    --- описание
    description    text,
    --- фотографии
    photo          text,
    --- максимальный размер посылки А
    max_size_a     integer,
    --- максимальный размер посылки Б
    max_size_b     integer,
    --- максимальный размер посылки Ц
    max_size_c     integer,
    --- максимульная сумма размеров посылки
    max_size_sum   integer,
    --- минимальный вес
    weight_min     integer,
    --- максимальный вес
    weight_max     integer,
    --- максимальная сумма наложенного платежа
    cod_max        integer,
    --- максимальный объем посылки в м3
    bulk_max       integer
);