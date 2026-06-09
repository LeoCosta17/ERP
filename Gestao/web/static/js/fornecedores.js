document.addEventListener('DOMContentLoaded', () => {
    const token = localStorage.getItem('token');
    if (!token) {
        window.location.href = '/';
        return;
    }

    const tbody = document.getElementById('tabela_fornecedores_body');
    const formFiltro = document.getElementById('formFiltroFornecedores');
    const inputBusca = document.getElementById('filtro_busca');
    const formNovo = document.getElementById('formNovoFornecedor');
    
    // Carregar fornecedores iniciais
    carregarFornecedores();

    // Filtro
    formFiltro.addEventListener('submit', (e) => {
        e.preventDefault();
        carregarFornecedores(inputBusca.value);
    });

    // Criação de Fornecedor
    formNovo.addEventListener('submit', async (e) => {
        e.preventDefault();

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
            carregarFornecedores(); // recarrega a lista

        } catch (err) {
            console.error(err);
            showError("Erro interno ao comunicar com o servidor.");
        }
    });

    async function carregarFornecedores(busca = "") {
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
                    <button class="btn btn-sm btn-outline-primary" title="Editar"><i class="bi bi-pencil"></i></button>
                </td>
            `;
            tbody.appendChild(tr);
        });
    }

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
});
