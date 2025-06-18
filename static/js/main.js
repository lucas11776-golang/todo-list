$(document).ready(function () {
  // Change navbar style on scroll
  function NavbarScroll() {
    const nav = document.querySelector(".navbar");

    if (nav == null) {
      return;
    }

    window.addEventListener("scroll", () => {
      if (window.scrollY > 50) {
        nav.classList.add("scrolled");
      } else {
        nav.classList.remove("scrolled");
      }
    });
  }

  NavbarScroll();
});


