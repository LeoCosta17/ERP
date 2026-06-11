CREATE TABLE IF NOT EXISTS tb_credenciais_empresa (
    id BIGINT PRIMARY KEY ,
    id_empresa BIGINT NOT NULL,
    tp_ambiente SMALLINT NOT NULL, -- 1 para Produção, 2 para Homologação
    certificado_digital BYTEA NOT NULL, -- BYTEA garante espaço seguro para o binário no Postgres
    senha_criptografada VARCHAR(255) NOT NULL,
    id_csc VARCHAR(6) NOT NULL, -- O identificador sequencial fornecido pela SEFAZ
    csc_nfe VARCHAR(255) NOT NULL, -- O token secreto
    FOREIGN KEY (id_empresa) REFERENCES tb_empresas(id),
    CONSTRAINT uq_empresa_ambiente UNIQUE (id_empresa, tp_ambiente) -- Garante uma config por ambiente
);