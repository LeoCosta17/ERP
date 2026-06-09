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
    if (tbody) {
        carregarFornecedores();
    }

    // Filtro
    if (formFiltro) {
        formFiltro.addEventListener('submit', (e) => {
            e.preventDefault();
            carregarFornecedores(inputBusca.value);
        });
    }

    // Criação de Fornecedor
    if (formNovo) {
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

    async function carregarFornecedores(busca = "") {
        if (!tbody) return;
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
                    <button class="btn btn-sm btn-outline-primary" title="Editar" onclick="abrirModalEditarFornecedor(${f.id})"><i class="bi bi-pencil"></i></button>
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

    // Funções para Editar Fornecedor
    window.abrirModalEditarFornecedor = async function(id) {
        try {
            const res = await fetch(`/api/fornecedores/${id}`, {
                headers: { 'Authorization': `Bearer ${token}` }
            });
            if (!res.ok) {
                showError("Erro ao buscar fornecedor.");
                return;
            }
            const fornecedor = await res.json();
            
            document.getElementById('edit_fornecedor_id').value = fornecedor.id;
            document.getElementById('edit_fornecedor_razao_social').value = fornecedor.razao_social || '';
            document.getElementById('edit_fornecedor_cnpj').value = fornecedor.cnpj || '';
            document.getElementById('edit_fornecedor_inscricao_estadual').value = fornecedor.inscricao_estadual || '';
            document.getElementById('edit_fornecedor_email').value = fornecedor.email || '';

            const containerEnderecos = document.getElementById('containerEnderecos');
            containerEnderecos.innerHTML = '';
            if (fornecedor.enderecos) {
                fornecedor.enderecos.forEach(end => adicionarEndereco(end));
            }

            const containerTelefones = document.getElementById('containerTelefones');
            containerTelefones.innerHTML = '';
            if (fornecedor.telefones) {
                fornecedor.telefones.forEach(tel => adicionarTelefone(tel));
            }

            const modalEditarEl = document.getElementById('modalEditarFornecedor');
            const modalEditar = bootstrap.Modal.getOrCreateInstance(modalEditarEl);
            modalEditar.show();
        } catch(err) {
            console.error(err);
            showError("Erro de comunicação ao buscar fornecedor.");
        }
    };

    function adicionarEndereco(end = {}) {
        const container = document.getElementById('containerEnderecos');
        const div = document.createElement('div');
        div.className = 'card mb-3 p-3 endereco-item';
        div.innerHTML = `
            <div class="row g-2">
                <div class="col-md-3">
                    <label class="form-label form-label-sm">CEP</label>
                    <input type="text" class="form-control form-control-sm addr-cep" value="${end.cep || ''}" required>
                </div>
                <div class="col-md-6">
                    <label class="form-label form-label-sm">Logradouro</label>
                    <input type="text" class="form-control form-control-sm addr-logradouro" value="${end.logradouro || ''}" required>
                </div>
                <div class="col-md-3">
                    <label class="form-label form-label-sm">Número</label>
                    <input type="text" class="form-control form-control-sm addr-numero" value="${end.numero || ''}" required>
                </div>
                <div class="col-md-4">
                    <label class="form-label form-label-sm">Bairro</label>
                    <input type="text" class="form-control form-control-sm addr-bairro" value="${end.bairro || ''}" required>
                </div>
                <div class="col-md-5">
                    <label class="form-label form-label-sm">Município</label>
                    <input type="text" class="form-control form-control-sm addr-municipio" value="${end.municipio || ''}" required>
                </div>
                <div class="col-md-2">
                    <label class="form-label form-label-sm">UF</label>
                    <input type="text" class="form-control form-control-sm addr-uf" value="${end.uf || ''}" required>
                </div>
                <div class="col-md-1 d-flex align-items-end">
                    <button type="button" class="btn btn-outline-danger btn-sm" onclick="this.closest('.endereco-item').remove()"><i class="bi bi-trash"></i></button>
                </div>
            </div>
        `;
        container.appendChild(div);
    }

    function adicionarTelefone(tel = {}) {
        const container = document.getElementById('containerTelefones');
        const div = document.createElement('div');
        div.className = 'row g-2 mb-2 telefone-item align-items-end';
        div.innerHTML = `
            <div class="col-md-2">
                <label class="form-label form-label-sm">DDD</label>
                <input type="text" class="form-control form-control-sm tel-ddd" value="${tel.ddd || ''}" required>
            </div>
            <div class="col-md-4">
                <label class="form-label form-label-sm">Número</label>
                <input type="text" class="form-control form-control-sm tel-numero" value="${tel.numero || ''}" required>
            </div>
            <div class="col-md-2">
                <button type="button" class="btn btn-outline-danger btn-sm w-100" onclick="this.closest('.telefone-item').remove()"><i class="bi bi-trash"></i> Remover</button>
            </div>
        `;
        container.appendChild(div);
    }

    const btnAddEndereco = document.getElementById('btnAddEndereco');
    if (btnAddEndereco) btnAddEndereco.addEventListener('click', () => adicionarEndereco());

    const btnAddTelefone = document.getElementById('btnAddTelefone');
    if (btnAddTelefone) btnAddTelefone.addEventListener('click', () => adicionarTelefone());

    const formEditar = document.getElementById('formEditarFornecedor');
    if (formEditar) {
        formEditar.addEventListener('submit', async (e) => {
            e.preventDefault();
            const id = document.getElementById('edit_fornecedor_id').value;
            
            const payload = {
                razao_social: document.getElementById('edit_fornecedor_razao_social').value,
                cnpj: document.getElementById('edit_fornecedor_cnpj').value,
                inscricao_estadual: document.getElementById('edit_fornecedor_inscricao_estadual').value,
                email: document.getElementById('edit_fornecedor_email').value,
                enderecos: [],
                telefones: []
            };

            document.querySelectorAll('.endereco-item').forEach(item => {
                payload.enderecos.push({
                    cep: item.querySelector('.addr-cep').value,
                    logradouro: item.querySelector('.addr-logradouro').value,
                    numero: item.querySelector('.addr-numero').value,
                    bairro: item.querySelector('.addr-bairro').value,
                    municipio: item.querySelector('.addr-municipio').value,
                    uf: item.querySelector('.addr-uf').value,
                    codigo_municipio: "",
                    is_principal: false
                });
            });

            document.querySelectorAll('.telefone-item').forEach(item => {
                payload.telefones.push({
                    ddd: item.querySelector('.tel-ddd').value,
                    numero: item.querySelector('.tel-numero').value
                });
            });

            try {
                const res = await fetch(`/api/fornecedores/${id}`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`
                    },
                    body: JSON.stringify(payload)
                });

                if (!res.ok) {
                    const data = await res.json();
                    showError(data.erro || "Erro ao atualizar fornecedor.");
                    return;
                }

                const modalEditarEl = document.getElementById('modalEditarFornecedor');
                const modalEditar = bootstrap.Modal.getInstance(modalEditarEl);
                if (modalEditar) modalEditar.hide();

                carregarFornecedores();
            } catch(err) {
                console.error(err);
                showError("Erro interno ao comunicar com o servidor.");
            }
        });
    }
});
