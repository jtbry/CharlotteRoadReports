document.addEventListener("DOMContentLoaded", () => {
    const pathname = window.location.pathname
    const navLinks = document.getElementsByClassName('nav-link')
    for(let i = 0; i < navLinks.length; ++i) {
        if(navLinks[i].getAttribute('href') == pathname) navLinks[i].classList.add('active')
        else if(navLinks[i].classList.contains('active')) navLinks[i].classList.remove('active')
    }
})