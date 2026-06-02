create table if not exists tb_enderecos (
    id bigint primary key auto_increment,
    usuario_id bigint not null,
    cep varchar(20) not null,
    logradouro varchar(255) not null,
    numero varchar(20) not null,
    complemento varchar(255),
    bairro varchar(255) not null,
    cidade varchar(255) not null,
    uf varchar(2) not null,
    ibge varchar(20) not null,
    gia varchar(20),
    ddd varchar(20) not null,
    siafi varchar(20) not null,
    principal boolean not null default true,
    criado_em timestamp default current_timestamp,
    foreign key (usuario_id) references tb_usuarios(id) on delete cascade
);