create schema home_schema;

create table home_schema.house (
    id bigserial,
    street text not null,
    city text not null,
    state text not null,
    zip bigint not null,

    primary key (id)
);

create table home_schema.usr (
    id bigserial,
    first_name text not null,
    last_name text not null,
    email text not null,
    password_hash text not null,

    primary key (id),
    unique (email)
);

create table home_schema.house_to_usr (
    usr_id bigint, 
    house_id bigint,

    primary key (usr_id, house_id),
    foreign key (usr_id) references home_schema.usr (id),
    foreign key (house_id) references home_schema.house (id)
);

create table home_schema.financial_info (
    id bigserial,
    year int,
    mortgage real,
    insurance text,
    house_id bigint not null,

    primary key(id),
    unique (house_id),
    foreign key (house_id) references home_schema.house(id)
);

create table home_schema.rennovations (
    name text not null,
    description text,
    cost real,
    financial_info_id bigint not null,

    primary key (name, description, cost, financial_info_id),
    unique (name),
    foreign key (financial_info_id) references home_schema.financial_info (id)
);

create table home_schema.floors (
    id bigserial,
    area real,
    name text not null,
    description text,
    house_id bigint not null,

    primary key (id),
    unique (name),
    foreign key (house_id) references home_schema.house(id)
);

create table home_schema.entrance (
    name text not null,
    open boolean not null,
    floor_id bigint not null,

    primary key (name, open, floor_id),
    unique (name),
    foreign key (floor_id) references home_schema.floors(id)
);

create table home_schema.rooms (
    id bigserial,
    name text not null,
    description text,
    floor_id bigint not null,

    primary key (id),
    unique (name),
    foreign key (floor_id) references home_schema.floors(id)
);

create table home_schema.measurements (
    id bigserial,
    measurement_date timestamp not null,
    temperature real,
    humidity real,
    activity boolean,
    floor_id bigint not null,

    primary key (id),
    foreign key (floor_id) references home_schema.floors(id)
);
