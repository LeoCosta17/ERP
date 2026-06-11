import { getToken } from '../utils/auth.js';
import { carregarFornecedores, showError } from './listarFornecedores.js';

export function setupCriarFornecedor() {
    // 1. Busca o formulário na tela. Se não existir, interrompe a execução para evitar erros.
    const formNovo = document.getElementById('formNovoFornecedor');
    if (!formNovo) return;

    // 2. Intercepta o evento de envio (submit) do formulário
    formNovo.addEventListener('submit', async (e) => {
        e.preventDefault(); // Impede o comportamento padrão de recarregar a página

        // 3. Extrai o token de segurança e os dados preenchidos pelo usuário
        const token = getToken();
        const razao_social = document.getElementById('fornecedor_razao_social').value;
        const cnpj = document.getElementById('fornecedor_cnpj').value;
        const email = document.getElementById('fornecedor_email').value;

        try {
            // 4. Envia os dados para o backend via requisição assíncrona (AJAX / Fetch API)
            const res = await fetch('/api/fornecedores', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}` // Token de autenticação
                },
                body: JSON.stringify({ razao_social, cnpj, email })
            });

            // 5. Trata as respostas de erro da API (ex: campos inválidos ou duplicados)
            if (!res.ok) {
                const data = await res.json();
                showError(data.erro || "Erro ao cadastrar fornecedor.");
                return;
            }

            // 6. Sucesso: Busca e oculta o modal do Bootstrap atrelado ao formulário
            const modalEl = document.getElementById('modalFornecedor');
            const modal = bootstrap.Modal.getInstance(modalEl);
            if (modal) modal.hide();

            // 7. Limpa os campos do formulário para preparar o próximo uso
            formNovo.reset();
            
            // 8. Atualiza a tela de forma inteligente:
            // Verifica se a tabela de listagem está na página. Se sim, apenas recarrega os dados dela.
            // Se não, recarrega a página inteira.
            const tbody = document.getElementById('tabela_fornecedores_body');
            if (tbody) {
                carregarFornecedores(); // recarrega a lista sem atualizar a página
            } else {
                window.location.reload();
            }

        } catch (err) {
            // 9. Captura falhas inesperadas de rede ou quebras de script
            console.error(err);
            showError("Erro interno ao comunicar com o servidor.");
        }
    });
}
