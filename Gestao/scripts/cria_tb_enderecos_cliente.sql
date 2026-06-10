create table if not exists tb_enderecos_clientes (
    id bigint auto_increment primary key,
    id_cliente bigint not null,
    cep varchar(9) not null,
    logradouro varchar(255) not null,
    numero varchar(20) not null,
    bairro varchar(100) not null,
    municipio varchar(100) not null,
    uf varchar(2) not null,
    codigo_municipio varchar(7) not null,
    is_principal bool default false,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    constraint fk_enderecos_clientes_clientes foreign key (id_cliente) references tb_clientes(id) on delete cascade on update cascade
);