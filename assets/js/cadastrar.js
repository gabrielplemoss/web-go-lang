async function fazPost(url ,body){
    await fetch(url,{
        method: "POST",
        mode: "same-origin",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(body)
    }).then(resposta => {
        if(resposta.redirected){
            window.location.href = resposta.url
        }
    })
}

document.getElementById("cadastrar-contato").addEventListener('submit', (e) => {
    e.preventDefault()

    const nome = document.getElementById("nome").value
    const apelido = document.getElementById("apelido").value
    const site = document.getElementById("site").value
    const email = document.getElementById("email").value
    const telefone = document.getElementById("telefone").value
    const endereco = document.getElementById("endereco").value

    const body = {
        "nome": nome,
        "apelido": apelido,
        "site": site,
        "email": email,
        "telefone": telefone,
        "endereco": endereco,
    }

    fazPost("/contato", body)   
})