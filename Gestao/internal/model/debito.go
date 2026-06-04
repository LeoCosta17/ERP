package model

type DebitoAvulsoCriar struct {
	FornecedorID    uint64  `json:"fornecedor_id"`
	Fornecedor      string  `json:"fornecedor"`
	Descricao       string  `json:"descricao"`
	NrDocumento     string  `json:"nr_documento"`
	NrNotaFiscal    string  `json:"nr_nota_fiscal"`
	Valor           float64 `json:"valor"`
	CategoriaDebito string  `json:"categoria_debito"`
	DtEntrada       string  `json:"dt_entrada"`
	DtVencimento    string  `json:"dt_vencimento"`
	NrParcela       uint64  `json:"nr_parcela"`
	NrTotalParcelas uint64  `json:"nr_total_parcelas"`
}
