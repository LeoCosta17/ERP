create table if not exists tb_clientes(
    id bigint auto_increment primary key,
    nome varchar(150) not null unique,
    tipo enum('PF', 'PJ') not null,
    email varchar(200) unique default null,
    telefone varchar(20) default null,
    cpf varchar(14) unique default null,
    cnpj varchar(18) unique default null,
    contribuinte enum('1','2','9') default null,
    is_consumidor_final bool default false,
    ie varchar(14) unique default null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
);