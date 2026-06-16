import { getToken } from "../utils/auth.js";
import { showError } from "../fornecedores/listarFornecedores.js";

export async function carregarClientes(busca = "") {
    const tbody = document.getElementById('tabela_clientes_body');
    if (!tbody) return;

    const token = getToken();
    tbody.innerHTML = `<tr><td colspan="5" class="text-muted py-5 text-center">Carregando...</td></tr>`

    try {
        let url = '/api/clientes';
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
            return
        }

        if (!res.ok) {
            showError("Erro ao carregar lista de clientes.")
            tbody.innerHTML = `<tr><td colspan="5" class="text-danger py-5 text-center">Erro ao carregar.</td></tr>`
            return;
        }

        const dados = await res.json();
        renderTabela(dados)
    } catch (err) {
        console.error(err);
        tbody.innerHTML = `<tr><td colspan="5" class="text-danger py-5 text-center">Erro de comunicação.</td></tr>`;
    }
}

function renderTabela(clientes) {
    const tbody = document.getElementById('tabela_clientes_body');
    if (!tbody) return;

    if (!clientes || clientes.length === 0) {
        tbody.innerHTML = `<tr><td colspan="5" class="text-muted py-5 text-center">Nenhum cliente encontrado.</td></tr>`;
        return;
    }

    tbody.innerHTML = '';
    clientes.forEach(c => {
        const tr = document.createElement('tr');
        // Usamos c.cpf ou c.cnpj dependendo de qual estiver preenchido
        const documento = c.cpf ? c.cpf : (c.cnpj ? c.cnpj : '-');
        
        tr.innerHTML = `
            <td>#${c.id}</td>
            <td class="text-start fw-bold">${c.nome}</td>
            <td>${c.email || '-'}</td>
            <td>${c.telefone || '-'}</td>
            <td><span class="badge bg-secondary">${c.tipo || '-'}</span></td>
            <td>${documento}</td>
            <td>
                <button class="btn btn-sm btn-outline-primary" title="Editar" onclick="window.abrirModalEditarCliente(${c.id})"><i class="bi bi-pencil"></i></button>
            </td>
        `;
        tbody.appendChild(tr);
    });
}