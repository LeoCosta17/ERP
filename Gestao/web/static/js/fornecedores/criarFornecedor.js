import { getToken } from '../utils/auth.js';
import { carregarFornecedores, showError } from './listarFornecedores.js';

export function setupCriarFornecedor() {
    const formNovo = document.getElementById('formNovoFornecedor');
    if (!formNovo) return;

    formNovo.addEventListener('submit', async (e) => {
        e.preventDefault();

        const token = getToken();
        const razao_social = document.getElementById('fornecedor_razao_social').value;
        const cnpj = document.getElementById('fornecedor_cnpj').value;
        const email = document.getElementById('fornecedor_email').value;

        try {
            const res = await fetch('/api/fornecedores', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify({ razao_social, cnpj, email })
            });

            if (!res.ok) {
                const data = await res.json();
                showError(data.erro || "Erro ao cadastrar fornecedor.");
                return;
            }

            // Sucesso
            const modalEl = document.getElementById('modalFornecedor');
            const modal = bootstrap.Modal.getInstance(modalEl);
            if (modal) modal.hide();

            formNovo.reset();
            const tbody = document.getElementById('tabela_fornecedores_body');
            if (tbody) {
                carregarFornecedores(); // recarrega a lista
            } else {
                window.location.reload();
            }

        } catch (err) {
            console.error(err);
            showError("Erro interno ao comunicar com o servidor.");
        }
    });
}
