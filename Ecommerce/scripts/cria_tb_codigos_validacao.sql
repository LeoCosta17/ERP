CREATE TABLE IF NOT EXISTS tb_codigos_validacao (
    id bigint primary key auto_increment,
    usuario_id bigint NOT NULL,
    codigo VARCHAR(255) NOT NULL,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expira_em TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL 10 MINUTE),
    tipo VARCHAR(50) NOT NULL,
    usado BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (usuario_id) REFERENCES tb_usuarios(id) ON DELETE CASCADE
);