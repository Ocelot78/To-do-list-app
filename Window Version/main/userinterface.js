document.addEventListener("DOMContentLoaded", function () {
    let nav = document.getElementById("navi");
    let buttonNew = document.getElementById("new");
    let windowHeight = window.innerHeight
    let navHeight = nav.offsetHeight
    let buttonHeight = buttonNew.offsetHeight
    let menu = document.getElementById("menu")
    let buttonAmount = menu.getElementsByTagName("button");
    let sum = windowHeight - navHeight - buttonAmount.length*buttonHeight
    buttonNew.style.marginTop = sum + "px"
    window.addEventListener("resize", function () {
        windowHeight = window.innerHeight
        buttonNew.style.marginTop = sum + "px"
    })
});