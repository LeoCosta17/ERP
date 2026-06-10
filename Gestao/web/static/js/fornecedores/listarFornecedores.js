import { getToken } from '../utils/auth.js';

function showError(message) {
    const modalBody = document.getElementById('errorModalBody');
    if (modalBody) {
        modalBody.textContent = message;
        const errorModalElement = document.getElementById('errorModal');
        const errorModal = bootstrap.Modal.getOrCreateInstance(errorModalElement);
        errorModal.show();
    } else {
        alert(message);
    }
}

export async function carregarFornecedores(busca = "") {
    const tbody = document.getElementById('tabela_fornecedores_body');
    if (!tbody) return;
    
    const token = getToken();
    tbody.innerHTML = `<tr><td colspan="5" class="text-muted py-5 text-center">Carregando...</td></tr>`;
    
    try {
        let url = '/api/fornecedores';
        if (busca) {
            url += `?busca=${encodeURIComponent(busca)}`;
        }

        const res = await fetch(url, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });

        if (res.status === 401) {
            window.location.href = '/';
            return;
        }

        if (!res.ok) {
            showError("Erro ao carregar lista de fornecedores.");
            tbody.innerHTML = `<tr><td colspan="5" class="text-danger py-5 text-center">Erro ao carregar.</td></tr>`;
            return;
        }

        const dados = await res.json();
        renderTabela(dados);

    } catch (err) {
        console.error(err);
        tbody.innerHTML = `<tr><td colspan="5" class="text-danger py-5 text-center">Erro de comunicação.</td></tr>`;
    }
}

function renderTabela(fornecedores) {
    const tbody = document.getElementById('tabela_fornecedores_body');
    if (!tbody) return;

    if (!fornecedores || fornecedores.length === 0) {
        tbody.innerHTML = `<tr><td colspan="5" class="text-muted py-5 text-center">Nenhum fornecedor encontrado.</td></tr>`;
        return;
    }

    tbody.innerHTML = '';
    fornecedores.forEach(f => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>#${f.id}</td>
            <td class="text-start fw-bold">${f.razao_social}</td>
            <td>${f.cnpj}</td>
            <td>${f.email || '-'}</td>
            <td>
                <button class="btn btn-sm btn-outline-primary" title="Editar" onclick="window.abrirModalEditarFornecedor(${f.id})"><i class="bi bi-pencil"></i></button>
            </td>
        `;
        tbody.appendChild(tr);
    });
}

export { showError };
