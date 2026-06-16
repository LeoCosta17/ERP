import { checkAuth } from '../utils/auth.js';
import { carregarClientes } from './listarClientes.js';
import { setupCriarCliente } from './criarCliente.js';

document.addEventListener('DOMContentLoaded', () => {
    if (!checkAuth()) return;

    const tbody = document.getElementById('tabela_clientes_body');
    const formFiltro = document.getElementById('formFiltroClientes'); // Pode precisar confirmar o ID no HTML
    const inputBusca = document.getElementById('filtro_busca');

    // Carregar clientes iniciais
    if (tbody) {
        carregarClientes();
    }

    // Filtro
    if (formFiltro) {
        formFiltro.addEventListener('submit', (e) => {
            e.preventDefault();
            carregarClientes(inputBusca ? inputBusca.value : "");
        });
    }

    // Inicializa os modais
    setupCriarCliente();
});
