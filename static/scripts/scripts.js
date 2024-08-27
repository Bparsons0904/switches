// const dialog = document.getElementsByClassName("dialog");
// for (let i = 0; i < dialog.length; i++) {
//   portal[i].style.display = "none";
// }
//
const portal = document.getElementById("portal");
function closeDialog() {
  portal.style.display = "none";
  //   dialog.close();
}

document.body.addEventListener("htmx:afterSwap", function (event) {
  if (event.detail.target.id === "portal") {
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
