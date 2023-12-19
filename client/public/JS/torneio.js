const API = "http://127.0.0.1:4000/api";

fetchTorneio()
.then((torn) => {
    mostrarTorneio(torn);
})
.catch((err) => {
    mostrarMensagemErro(err);
});

/**
 * Retorna todos os torneios da AFEB.
 * @returns {Array} - Torneios da AFEB.
 * @throws - Retorna um erro da API.
 */
async function fetchTorneio() {
    return await fetch(API + "/torneios")
    .then((res) => res.json())
    .then((res) => {
        if (res == null)
            return null;

        if (res.error) 
            return new Error(res.error);

        return res.torneios;
    })
    .catch((err) => {
        console.error(err);
        return new Error("Erro ao conectar com a API.");
    });
}

/**
 * Coloca os dados de torneios na tabela.
 * @param {Array} torneios - Torneios da AFEB.
 */
function mostrarTorneio(torneios) {
    if (!torneios)
        return;

    const torneioTabela = document.getElementById("tabela-torneio");

    for (const t of torneios) {
        const containerTorneio = document.createElement("a");
        //containerTorneio.classList.add("ranking-jogador-div");
        containerTorneio.href = `torneio.html?torneio=${t.codTorn}`;

        containerTorneio.innerHTML = `
            <div>
                <p class="text-sm"> ${t.titulo} </p>
                <p> ${t.modo} </p>
                <p> ${t.dataInicio} </p>
                <p> ${(!t.dataFim ? "" : t.dataFim)} </p>
            </div>
        `;

        torneioTabela.appendChild(containerTorneio);
    }
}

/**
 * Mostra uma mensagem de erro ao usuário.
 * @param {string} erro - Mensagem de erro.
 */
function mostrarMensagemErro(erro) {
    alert(erro);
}