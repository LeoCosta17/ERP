import { checkAuth } from '../utils/auth.js';
import { carregarDadosUsuario } from './carregarDadoUsuario.js';

document.addEventListener('DOMContentLoaded', () => {
    if (!checkAuth()) return;

    carregarDadosUsuario();
});
