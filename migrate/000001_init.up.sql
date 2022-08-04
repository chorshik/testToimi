CREATE TABLE ads
(
    id          serial         not null unique,
    title       varchar(200)   not null,
    price       numeric(10, 2) not null unique,
    description varchar(1000)  not null,
    photo       varchar(255)[3]
);