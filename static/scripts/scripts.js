/*
 * Handle the display of the modal
 */
const portal = document.getElementById("portal");
function closeDialog() {
  portal.style.display = "none";
  document.body.classList.remove("body-no-scroll");
  // Remove the modal query param from the URL
  const url = new URL(window.location);
  url.searchParams.delete("modal");
  window.history.pushState({}, "", url);
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

/*
 * Handle the automatic loading of a modal from a query parameter
 */
function updateQueryParam(event) {
  console.log("ran updateQueryParam", event);
  const switchId = event.currentTarget.getAttribute("data-switch-id");
  console.log("ran updateQueryParam", switchId);
  const url = new URL(window.location);
  url.searchParams.set("modal", switchId);
  window.history.pushState({}, "", url);
  loadModalForSwitchId(switchId);
}

function getSwitchIdFromUrl() {
  const urlParams = new URLSearchParams(window.location.search);
  return urlParams.get("modal");
}

function loadModalForSwitchId(switchId) {
  const modalLoader = document.getElementById("modal-loader");
  modalLoader.setAttribute("hx-get", `/switches/${switchId}/modal`);
  modalLoader.setAttribute("hx-target", "#portal");
  modalLoader.setAttribute("hx-swap", "innerHTML");
  modalLoader.setAttribute("hx-trigger", "load");
  htmx.process(modalLoader);
}

document.addEventListener("DOMContentLoaded", function () {
  const switchId = getSwitchIdFromUrl();
  if (switchId) {
    // Delay the modal loading slightly to ensure everything is ready
    setTimeout(() => loadModalForSwitchId(switchId), 100);
  }
});

document.body.addEventListener("htmx:afterOnLoad", function (event) {
  if (event.detail.elt.id === "modal-loader") {
    // Modal content has been loaded
    // document.getElementById("portal").style.display = "block";
  }
});

// Update URL with query param when a switch card is clicked
document.body.addEventListener("htmx:beforeSend", function (event) {
  if (event.detail.elt.classList.contains("switch-card")) {
    const switchId = event.detail.elt.getAttribute("data-switch-id");
    if (switchId) {
      const url = new URL(window.location);
      url.searchParams.set("modal", switchId);
      window.history.pushState({}, "", url);
      loadModalForSwitchId(switchId);
    }
  }
});
