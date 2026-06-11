CREATE TABLE IF NOT EXISTS tb_config_fiscais_empresa (
    id BIGINT PRIMARY KEY ,
    id_empresa BIGINT NOT NULL UNIQUE,
    cd_regime_tributario INT NOT NULL,
    FOREIGN KEY (id_empresa) REFERENCES tb_empresas(id)
);