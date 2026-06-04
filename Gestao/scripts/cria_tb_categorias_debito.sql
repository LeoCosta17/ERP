create table if not exists tb_categorias_debito(
    id bigint auto_increment primary key,
    nome varchar(255) not null,
    id_empresa bigint not null,
    foreign key (id_empresa) references tb_empresas_gestao(id)
)