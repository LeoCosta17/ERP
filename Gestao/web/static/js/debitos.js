document.addEventListener('DOMContentLoaded', () => {
    const token = localStorage.getItem('token');
    if (!token) {
        window.location.href = '/';
        return;
    }

    let debitosCarregados = [];

    const selectFornecedor = document.getElementById('id_fornecedor');
    const selectCategoria = document.getElementById('id_categoria');
    const formDebito = document.getElementById('formDebitoAvulso');
    const formFiltro = document.getElementById('formFiltroDebitos');
    const tabelaDebitos = document.getElementById('tabela_debitos_body');

    // Carregar dropdowns e débitos da tabela
    carregarDropdowns();
    carregarDebitos();

    formFiltro.addEventListener('submit', (e) => {
        e.preventDefault();
        carregarDebitos();
    });

    async function carregarDropdowns() {
        try {
            const editSelectFornecedor = document.getElementById('edit_id_fornecedor');
            const editSelectCategoria = document.getElementById('edit_id_categoria');

            // Fetch fornecedores
            const resF = await fetch('/api/fornecedores', {
                headers: { 'Authorization': `Bearer ${token}` }
            });
            if (resF.ok) {
                const fornecedores = await resF.json();
                selectFornecedor.innerHTML = '<option value="" selected disabled>Selecione o Fornecedor...</option>';
                if (editSelectFornecedor) editSelectFornecedor.innerHTML = '<option value="" selected disabled>Selecione o Fornecedor...</option>';

                fornecedores.forEach(f => {
                    const opt = document.createElement('option');
                    opt.value = f.id;
                    opt.textContent = `${f.razao_social} (CNPJ: ${f.cnpj})`;
                    selectFornecedor.appendChild(opt);

                    if (editSelectFornecedor) {
                        const opt2 = document.createElement('option');
                        opt2.value = f.id;
                        opt2.textContent = `${f.razao_social} (CNPJ: ${f.cnpj})`;
                        editSelectFornecedor.appendChild(opt2);
                    }
                });
            } else {
                selectFornecedor.innerHTML = '<option value="" disabled>Erro ao carregar</option>';
            }

            // Fetch categorias
            const resC = await fetch('/api/categorias', {
                headers: { 'Authorization': `Bearer ${token}` }
            });
            if (resC.ok) {
                const categorias = await resC.json();
                selectCategoria.innerHTML = '<option value="" selected>Sem Categoria (Opcional)</option>';
                if (editSelectCategoria) editSelectCategoria.innerHTML = '<option value="" selected>Sem Categoria (Opcional)</option>';

                categorias.forEach(c => {
                    const opt = document.createElement('option');
                    opt.value = c.id;
                    opt.textContent = c.nome;
                    selectCategoria.appendChild(opt);

                    if (editSelectCategoria) {
                        const opt2 = document.createElement('option');
                        opt2.value = c.id;
                        opt2.textContent = c.nome;
                        editSelectCategoria.appendChild(opt2);
                    }
                });
            } else {
                selectCategoria.innerHTML = '<option value="" selected>Erro ao carregar categorias</option>';
            }

        } catch (err) {
            console.error(err);
            selectFornecedor.innerHTML = '<option value="" disabled>Erro ao carregar</option>';
        }
    }

    // Handle form submit
    formDebito.addEventListener('submit', async (e) => {
        e.preventDefault();

        const id_fornecedor = parseInt(selectFornecedor.value, 10);
        let id_categoria = parseInt(selectCategoria.value, 10);
        if (isNaN(id_categoria)) id_categoria = null;

        const payload = {
            id_fornecedor: id_fornecedor,
            descricao: document.getElementById('descricao').value,
            nr_documento: document.getElementById('nr_documento').value,
            nr_nota_fiscal: document.getElementById('nr_nota_fiscal').value,
            valor: parseFloat(document.getElementById('valor').value),
            dt_entrada: document.getElementById('dt_entrada').value,
            dt_vencimento: document.getElementById('dt_vencimento').value,
            nr_parcela: parseInt(document.getElementById('nr_parcela').value, 10),
            nr_total_parcelas: parseInt(document.getElementById('nr_total_parcelas').value, 10)
        };

        if (id_categoria) {
            payload.id_categoria = id_categoria;
        }

        try {
            const res = await fetch('/api/debitos', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify(payload)
            });

            if (!res.ok) {
                const data = await res.json();
                alert(data.erro || "Erro ao cadastrar débito.");
                return;
            }

            alert("Pagamento avulso cadastrado com sucesso!");
            const modalEl = document.getElementById('modalDebitoAvulso');
            const modal = bootstrap.Modal.getInstance(modalEl);
            if (modal) modal.hide();

            formDebito.reset();
            carregarDebitos();
        } catch (err) {
            console.error(err);
            alert("Erro interno ao comunicar com o servidor.");
        }
    });

    async function carregarDebitos() {
        tabelaDebitos.innerHTML = '<tr><td colspan="7" class="text-muted py-5 text-center">Carregando...</td></tr>';
        
        const busca = document.getElementById('filtro_fornecedor').value;
        const vencimento = document.getElementById('filtro_vencimento').value;
        const status = document.getElementById('filtro_status').value;

        try {
            let url = '/api/debitos?';
            const params = new URLSearchParams();
            if (busca) params.append('busca', busca);
            if (vencimento) params.append('vencimento', vencimento);
            if (status) params.append('status', status);
            url += params.toString();

            const res = await fetch(url, {
                headers: { 'Authorization': `Bearer ${token}` }
            });

            if (res.status === 401) {
                window.location.href = '/';
                return;
            }

            if (!res.ok) {
                tabelaDebitos.innerHTML = '<tr><td colspan="7" class="text-danger py-5 text-center">Erro ao carregar débitos.</td></tr>';
                return;
            }

            const dados = await res.json();
            debitosCarregados = dados;
            renderTabelaDebitos(dados);
        } catch (err) {
            console.error(err);
            tabelaDebitos.innerHTML = '<tr><td colspan="7" class="text-danger py-5 text-center">Erro de comunicação.</td></tr>';
        }
    }

    function renderTabelaDebitos(debitos) {
        if (!debitos || debitos.length === 0) {
            tabelaDebitos.innerHTML = '<tr><td colspan="7" class="text-muted py-5 text-center">Nenhum débito encontrado.</td></tr>';
            return;
        }

        tabelaDebitos.innerHTML = '';
        debitos.forEach(d => {
            const dataVencimento = d.dt_vencimento ? d.dt_vencimento.substring(0, 10).split('-').reverse().join('/') : '-';
            const valorFormatado = new Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' }).format(d.valor);
            
            let badgeClass = 'bg-secondary';
            if (d.status === 'PENDENTE') badgeClass = 'bg-warning text-dark';
            if (d.status === 'PAGO') badgeClass = 'bg-success';
            if (d.status === 'CANCELADO') badgeClass = 'bg-danger';

            const tr = document.createElement('tr');
            tr.innerHTML = `
                <td>#${d.id}</td>
                <td class="text-start fw-bold">${d.fornecedor ? d.fornecedor.razao_social : '-'}</td>
                <td class="text-start text-muted">${d.descricao}</td>
                <td>${dataVencimento}</td>
                <td class="fw-bold text-danger">${valorFormatado}</td>
                <td><span class="badge ${badgeClass} px-3 py-2 rounded-pill">${d.status}</span></td>
                <td>
                    <button class="btn btn-sm btn-outline-primary" title="Visualizar"><i class="bi bi-eye"></i></button>
                    <button class="btn btn-sm btn-outline-warning" title="Editar" onclick="abrirModalEdicao(${d.id})"><i class="bi bi-pencil"></i></button>
                    ${d.status === 'PENDENTE' ? `<button class="btn btn-sm btn-outline-success" title="Dar Baixa" onclick="pagarDebito(${d.id})"><i class="bi bi-check2-circle"></i></button>` : ''}
                </td>
            `;
            tabelaDebitos.appendChild(tr);
        });
    }

    // Funcionalidades de Pagamento e Edição
    window.pagarDebito = async function(id) {
        if (!confirm("Tem certeza que deseja dar baixa (pagar) este débito?")) return;

        try {
            const res = await fetch(`/api/debitos/${id}/pagar`, {
                method: 'PUT',
                headers: { 'Authorization': `Bearer ${token}` }
            });

            if (!res.ok) {
                const data = await res.json();
                alert(data.erro || "Erro ao pagar débito.");
                return;
            }

            alert("Débito pago com sucesso!");
            carregarDebitos();
        } catch (err) {
            console.error(err);
            alert("Erro interno.");
        }
    };

    window.abrirModalEdicao = function(id) {
        const debito = debitosCarregados.find(d => d.id === id);
        if (!debito) return;

        document.getElementById('edit_id_debito').value = debito.id;
        document.getElementById('edit_id_fornecedor').value = debito.id_fornecedor;
        document.getElementById('edit_id_categoria').value = debito.id_categoria || '';
        document.getElementById('edit_descricao').value = debito.descricao;
        document.getElementById('edit_nr_documento').value = debito.nr_documento || '';
        document.getElementById('edit_nr_nota_fiscal').value = debito.nr_nota_fiscal || '';
        document.getElementById('edit_valor').value = debito.valor;
        document.getElementById('edit_dt_entrada').value = debito.dt_entrada.substring(0, 10);
        document.getElementById('edit_dt_vencimento').value = debito.dt_vencimento.substring(0, 10);
        document.getElementById('edit_nr_parcela').value = debito.nr_parcela;
        document.getElementById('edit_nr_total_parcelas').value = debito.nr_total_parcelas;

        const modalEl = document.getElementById('modalEditarDebito');
        const modal = new bootstrap.Modal(modalEl);
        modal.show();
    };

    const formEditarDebito = document.getElementById('formEditarDebito');
    if (formEditarDebito) {
        formEditarDebito.addEventListener('submit', async (e) => {
            e.preventDefault();

            const id = document.getElementById('edit_id_debito').value;
            const id_fornecedor = parseInt(document.getElementById('edit_id_fornecedor').value, 10);
            let id_categoria = parseInt(document.getElementById('edit_id_categoria').value, 10);
            if (isNaN(id_categoria)) id_categoria = null;

            const payload = {
                id_fornecedor: id_fornecedor,
                descricao: document.getElementById('edit_descricao').value,
                nr_documento: document.getElementById('edit_nr_documento').value,
                nr_nota_fiscal: document.getElementById('edit_nr_nota_fiscal').value,
                valor: parseFloat(document.getElementById('edit_valor').value),
                dt_entrada: document.getElementById('edit_dt_entrada').value,
                dt_vencimento: document.getElementById('edit_dt_vencimento').value,
                nr_parcela: parseInt(document.getElementById('edit_nr_parcela').value, 10),
                nr_total_parcelas: parseInt(document.getElementById('edit_nr_total_parcelas').value, 10)
            };

            if (id_categoria) {
                payload.id_categoria = id_categoria;
            }

            try {
                const res = await fetch(`/api/debitos/${id}`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`
                    },
                    body: JSON.stringify(payload)
                });

                if (!res.ok) {
                    const data = await res.json();
                    alert(data.erro || "Erro ao editar débito.");
                    return;
                }

                alert("Débito atualizado com sucesso!");
                const modalEl = document.getElementById('modalEditarDebito');
                const modal = bootstrap.Modal.getInstance(modalEl);
                if (modal) modal.hide();

                carregarDebitos();
            } catch (err) {
                console.error(err);
                alert("Erro interno ao comunicar com o servidor.");
            }
        });
    }
});
