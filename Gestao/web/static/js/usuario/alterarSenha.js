import { getToken } from '/static/js/utils/auth.js';
import { showError } from '/static/js/utils/showError.js';
import { validaRespostaRequisicao } from '/static/js/utils/resposta.js';

async function alterarSenhaAPI(senhaAtual, novaSenha, senhaConfirmacao) {
    const token = getToken();
    const res = await fetch(`/api/usuario/alterar-senha`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify({
            senhaAtual: senha_atual,
            novaSenha: nova_senha,
            senhaConfirmacao: senha_confirmacao
        })
    });

    return await validaRespostaRequisicao(res);

}

export async function alterarSenha(){

    const formAlterarSenha = document.getElementById('formAlterarSenha');
    if(!formAlterarSenha) return;

    formAlterarSenha.addEventListener('submit', async (e) => {
        e.preventDefault();

        const senhaAtualInput = document.getElementById('senha_atual');
        const novaSenhaInput = document.getElementById('nova_senha');
        const senhaConfirmacaoInput = document.getElementById('senha_confirmacao');

        const senhaAtual = senhaAtualInput.value;
        const novaSenha = novaSenhaInput.value;
        const senhaConfirmacao = senhaConfirmacaoInput.value;

        try{
            const res = await alterarSenhaAPI(senhaAtual, novaSenha, senhaConfirmacao);
            senhaAtualInput.value = '';
            novaSenhaInput.value = '';
            senhaConfirmacaoInput.value = '';
            alert("Senha alterada com sucesso!");
        } catch(err){
            showError(err.message || "Erro ao alterar senha.");
        }
    });
}