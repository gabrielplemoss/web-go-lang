async function fazPost(url ,body){
    await fetch(url,{
        method: "POST",
        mode: "cors",
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

document.getElementById('logar-usuario').addEventListener('submit', (e) => {
    e.preventDefault()

    const email = document.getElementById('login-user').value
    const password = document.getElementById('login-password').value
    const body = {
        "email": email,
        "senha": password
    }

    fazPost('login', body)    
})

document.getElementById("cadastrar-usuario").addEventListener('submit', (e) => {
    e.preventDefault()

    const user = document.getElementById("cadastro-user").value
    const email = document.getElementById("cadastro-email").value
    const password = document.getElementById("cadastro-password").value
    
    const body = {
        "usuario": user,
        "email": email,
        "senha": password
    }

    fazPost("cadastrar", body)   
})
