DROP DATABASE IF EXISTS agenda;
CREATE DATABASE IF NOT EXISTS agenda;
USE agenda;

DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS contatos;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    usuario varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null
) ENGINE=INNODB;

CREATE TABLE contatos(
    id int auto_increment primary key,
    usuario_dono int not null ,
    nome varchar(50) not null,
    apelido varchar(50) not null,
    site varchar(50) not null,
    email varchar(50) not null,
    telefone varchar(100) not null,
    endereco varchar(100) not null,
    FOREIGN KEY (usuario_dono) REFERENCES usuarios(id)
) ENGINE=INNODB;