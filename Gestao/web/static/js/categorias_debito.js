document.addEventListener('DOMContentLoaded', () => {
    const token = localStorage.getItem('token');
    if (!token) {
        window.location.href = '/';
        return;
    }

    const tbody = document.getElementById('tabela_categorias_body');
    const formNovo = document.getElementById('formNovaCategoria');
    
    carregarCategorias();

    formNovo.addEventListener('submit', async (e) => {
        e.preventDefault();

        const nome = document.getElementById('categoria_nome').value;

        try {
            const res = await fetch('/api/categorias', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify({ nome })
            });

            if (!res.ok) {
                const data = await res.json();
                showError(data.erro || "Erro ao cadastrar categoria.");
                return;
            }

            const modalEl = document.getElementById('modalCategoria');
            const modal = bootstrap.Modal.getInstance(modalEl);
            if (modal) modal.hide();

            formNovo.reset();
            carregarCategorias();

        } catch (err) {
            console.error(err);
            showError("Erro interno ao comunicar com o servidor.");
        }
    });

    async function carregarCategorias() {
        tbody.innerHTML = `<tr><td colspan="3" class="text-muted py-5 text-center">Carregando...</td></tr>`;
        
        try {
            const res = await fetch('/api/categorias', {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });

            if (res.status === 401) {
                window.location.href = '/';
                return;
            }

            if (!res.ok) {
                showError("Erro ao carregar lista de categorias.");
                tbody.innerHTML = `<tr><td colspan="3" class="text-danger py-5 text-center">Erro ao carregar.</td></tr>`;
                return;
            }

            const dados = await res.json();
            renderTabela(dados);

        } catch (err) {
            console.error(err);
            tbody.innerHTML = `<tr><td colspan="3" class="text-danger py-5 text-center">Erro de comunicação.</td></tr>`;
        }
    }

    function renderTabela(categorias) {
        if (!categorias || categorias.length === 0) {
            tbody.innerHTML = `<tr><td colspan="3" class="text-muted py-5 text-center">Nenhuma categoria encontrada.</td></tr>`;
            return;
        }

        tbody.innerHTML = '';
        categorias.forEach(c => {
            const tr = document.createElement('tr');
            tr.innerHTML = `
                <td>#${c.id}</td>
                <td class="text-start fw-bold">${c.nome}</td>
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
