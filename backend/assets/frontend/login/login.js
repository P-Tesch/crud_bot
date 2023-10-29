function login() {
    sessionStorage.setItem("username", document.getElementById("username").value)
    sessionStorage.setItem("password", document.getElementById("password").value)
    window.location.href = "../genres/genres.html";
}