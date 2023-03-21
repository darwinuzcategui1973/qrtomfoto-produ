-- ejecutar en cosnsola este scripts
-- mysql -u root -p < schema.sql
create database IF NOT EXISTS backenqrtomfoto;
use bdnube;


CREATE TABLE IF NOT EXISTS USERS (
    id int not null auto_increment,
    name varchar(255) not null,
    password varchar(300) not null,
    primary key (id)
);


CREATE TABLE IF NOT EXISTS ROLES(
    id int not null auto_increment,
    name varchar(255) not null,
    primary key (id)
);

CREATE TABLE IF NOT EXISTS USER_ROLES(
    id int not null auto_increment,
    user_id int not null,
    role_id int not null,
    primary key (id),
    foreign key (user_id) references USERS(id),
    foreign key (role_id) references ROLES(id)
);

CREATE TABLE IF NOT EXISTS EMPRESA (
    id int not null auto_increment,
    rif varchar(255) not null,
    status  varchar (20) not null DEFAULT "Licenciado",
    nombre varchar(255) not null,
    conformacion varchar(300) not null,
    fechaInstalacion date ,
    fechaVencimiento date ,
    sistema varchar(255),
    version varchar (8) not null DEFAULT "1.00.00",
    usuario  int,
    viewers   int,
    MACAddress varchar(20) not null DEFAULT "00:00:ff:00",
    fechaInicioSistema TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    primary key (id)
);

CREATE TABLE IF NOT EXISTS QRTOMFOTO (
    id int not null auto_increment,
    rif varchar(255) not null,
    ubicacion  varchar (255) not null DEFAULT "noHayUbicacion",
    nombre varchar(255) not null,
    archivo mediumblob not null,
    fechaInicioEnvio TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    primary key (id)
);

insert into ROLES (id, name) values (1, 'admin');
insert into ROLES (id, name) values (2, 'seller');
insert into ROLES (id, name) values (3, 'customer');

# las claves 123456 encriptadas
INSERT INTO backenqrtomfoto.USERS (name, password) VALUES('admon', 'jy7UwgD8v8eg2a6L6W15Mqwxq4VlG8hvkCZ+IcquD5pjlg');
INSERT INTO backenqrtomfoto.USERS (name, password) VALUES('darwin', 'YU3RyBam/ntBM+MSvs2AiowUWJkS2ZQ2ar5BWBVhPXxD4A');
INSERT INTO backenqrtomfoto.USERS (name, password) VALUES('tomas', 'Phcx4CoOGX6MFpNcxDoJ+xPQEkJBPsRJi0Sh13xmXnQGFQ');