CREATE TABLE tb_debitos (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    id_fornecedor BIGINT NOT NULL,
    id_categoria BIGINT,
    descricao VARCHAR(255) NOT NULL,
    nr_documento VARCHAR(255),
    nr_nota_fiscal VARCHAR(255),
    valor DECIMAL(15,2) NOT NULL,
    dt_entrada DATE NOT NULL,
    dt_vencimento DATE NOT NULL,
    nr_parcela INT NOT NULL,
    nr_total_parcelas INT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDENTE',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_fornecedor) REFERENCES tb_fornecedores(id),
    FOREIGN KEY (id_categoria) REFERENCES tb_categorias_debito(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
