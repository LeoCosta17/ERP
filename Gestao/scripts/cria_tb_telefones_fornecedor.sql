CREATE TABLE if not exists tb_telefones_fornecedores (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    id_fornecedor BIGINT NOT NULL,
    ddd CHAR(2) NOT NULL,
    numero VARCHAR(15) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_telefone_fornecedor FOREIGN KEY (id_fornecedor) REFERENCES tb_fornecedores(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;