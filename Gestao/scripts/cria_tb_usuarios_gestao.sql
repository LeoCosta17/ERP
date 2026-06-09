create table if not exists tb_usuarios_gestao(
    id bigint primary key auto_increment,
    nome varchar(255) not null,
    cpf varchar(14) unique,
    telefone varchar(20) unique,
    email varchar(255) not null unique,
    senha varchar(255) not null,
    criado_em timestamp default current_timestamp,
    atualizado_em timestamp default current_timestamp on update current_timestamp,
    ativo boolean default false
);