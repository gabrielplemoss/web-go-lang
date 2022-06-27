let editar = false

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


document.querySelector(".editar").addEventListener("click",(e)=>{
    e.preventDefault()
    e.target.innerText = "salvar"
    const inputs = document.querySelectorAll("input")
    if(!editar){
        editar = true
        for(const input of inputs){
            input.disabled = false
        }
    } else if(editar){
        editar = false
        const nome = document.getElementById("nome").value
        const apelido = document.getElementById("apelido").value
        const site = document.getElementById("site").value
        const email = document.getElementById("email").value
        const telefone = document.getElementById("telefone").value
        const endereco = document.getElementById("endereco").value

        const url = window.location.pathname
        const id = url.replace("/contato/", "")
    
        const body = {
            "nome": nome,
            "apelido": apelido,
            "site": site,
            "email": email,
            "telefone": telefone,
            "endereco": endereco,
        }
    
       fazPost(`/editar/contato/${id}`, body) 
    }
})
