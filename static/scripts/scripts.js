const portal = document.getElementById("portal");
function closeDialog() {
  portal.style.display = "none";
  document.body.classList.remove("body-no-scroll");
}

document.body.addEventListener("htmx:afterSwap", function (event) {
  if (event.detail.target.id === "portal") {
    document.body.classList.add("body-no-scroll");
    portal.style.display = "flex";

    portal.addEventListener("click", (event) => {
      const dialog = document.getElementById("dialog");
      const rect = dialog.getBoundingClientRect();
      if (
        event.clientX < rect.left ||
        event.clientX > rect.right ||
        event.clientY < rect.top ||
        event.clientY > rect.bottom
      ) {
        closeDialog();
      }
    });
  }
});
