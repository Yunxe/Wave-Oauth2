const loginPage = document.querySelector("#login-page");
const authPage = document.querySelector("#auth-page");
const email = document.querySelector("input[name='email']");
const password = document.querySelector("input[name='password']");
const responseMsg = document.querySelector("#responseMsg");
const LoginButton = document.querySelector("#login");
const RegisterButton = document.querySelector("#register");
const authButton = document.querySelector("#authorization")

const checkToken = async (token) => {
    return await axios({
        method: "post",
        url: "http://localhost:8079/api/info",
        headers: {Authorization: "Bearer " + token},
        validateStatus: (status) => status,
    });
};

document.addEventListener("DOMContentLoaded", async () => {
    const token = localStorage.getItem("token")
    const res = await checkToken(token);
    const {code, message, data} = res.data;
    if (code === 0) {
        authPage.classList.remove("is-hidden");
        document.querySelector("#avatar").src = data.avatarUrl
        document.querySelector("#userName").innerHTML = data.userName
    } else {
        localStorage.clear()
        loginPage.classList.remove("is-hidden")
    }
});

LoginButton.addEventListener("click", async (e) => {
    e.preventDefault();
    LoginButton.classList.add("is-loading");
    const res = await axios({
        method: "post",
        url: "http://localhost:8079/api/login",
        data: {
            email: email.value,
            password: password.value,
        },
        validateStatus: (status) => status,
    });
    const {code, message, data} = res.data;
    responseMsg.classList.remove("is-hidden");
    if (code === 0) {
        responseMsg.classList.remove("is-danger");
        responseMsg.classList.add("is-success");
        responseMsg.innerHTML = "登入成功,正在跳转";
        setTimeout(async () => {
            localStorage.setItem("token", data.token);
            LoginButton.classList.remove("is-loading");
            loginPage.classList.add("is-hidden")
            authPage.classList.remove("is-hidden")
            const tokenInfo = await checkToken(data.token)
            localStorage.setItem("user", JSON.stringify(tokenInfo.data.data))
            document.querySelector("#userName").innerHTML = tokenInfo.data.data.userName
            document.querySelector("#avatar").setAttribute("src", tokenInfo.data.data.avatarUrl)
        }, 1000);
    } else {
        responseMsg.classList.add("is-danger");
        responseMsg.innerHTML = `登录失败,${message}`;
        setTimeout(() => {
            LoginButton.classList.remove("is-loading");
            responseMsg.classList.remove("is-danger");
            responseMsg.classList.add("is-hidden");
            responseMsg.innerHTML = "";
        }, 2000)
    }
});

RegisterButton.addEventListener("click", (e) => {
    e.preventDefault()
    RegisterButton.classList.add("is-loading");
    responseMsg.classList.remove("is-hidden")
    responseMsg.classList.add("is-info")
    responseMsg.innerHTML = "正在跳转注册页面";
    //TODO跳转注册页面
})

authButton.addEventListener("click", (e) => {
    e.preventDefault()
    const token = localStorage.getItem("token")
    const userJSON = localStorage.getItem("user")
    const user = JSON.parse(userJSON)
    window.location.href = `http://localhost:8078/oauth2/code${window.location.search}&uid=${user.uid}`
})


//进入页面->检查token->
//有token->checkToken()->success,更新user,跳转授权页面
//                     ->fail,删除user,跳转登录页面
//无token->跳转登录页面
//登录成功->存token->checkToken()
