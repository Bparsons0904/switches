function closeDialog() {
  const portal = document.getElementById("portal");
  portal.style.display = "none";
  document.body.classList.remove("body-no-scroll");
  const url = new URL(window.location);
  url.searchParams.delete("modal");
  window.history.pushState({}, "", url);
  const modalLoader = document.getElementById("modal-loader");
  modalLoader.innerHTML = "";
}

document.body.addEventListener("htmx:afterSwap", function (event) {
  const portal = document.getElementById("portal");
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

async function shareSwitch(event) {
  if (navigator.share && navigator.canShare(shareData)) {
    const switchName = event.currentTarget.getAttribute("data-switch-name");
    const shareData = {
      title: "Check out this awesome switch!",
      text: switchName,
      url: window.location.href,
    };
    await navigator.share(shareData);
  } else {
    await navigator.clipboard.writeText(window.location.href);
    alert("Switch address copied to clipboard!");
  }
}

// function updateQueryParam(switchId) {
//   // const switchId = event.currentTarget.getAttribute("data-switch-id");
//   const url = new URL(window.location);
//   url.searchParams.set("modal", switchId);
//   window.history.pushState({}, "", url);
//   loadModalForSwitchId(switchId);
// }

function getSwitchIdFromUrl() {
  const urlParams = new URLSearchParams(window.location.search);
  return urlParams.get("modal");
}

function loadModalForSwitchId(switchId) {
  const modalLoader = document.getElementById("modal-loader");
  modalLoader.innerHTML = `
    <div hx-get="/switches/${switchId}/modal"
         hx-trigger="load"
         hx-target="#portal"
         hx-swap="innerHTML"></div>
  `;
  htmx.process(modalLoader);
}

document.addEventListener("DOMContentLoaded", function () {
  const switchId = getSwitchIdFromUrl();
  if (switchId) {
    setTimeout(() => loadModalForSwitchId(switchId), 100);
  }
});

// document.body.addEventListener("htmx:beforeSend", function (event) {
//   if (event.detail.elt.classList.contains("switch-card")) {
//     const switchId = event.detail.elt.getAttribute("data-switch-id");
//     if (switchId) {
//       const url = new URL(window.location);
//       url.searchParams.set("modal", switchId);
//       window.history.pushState({}, "", url);
//       loadModalForSwitchId(switchId);
//     }
//   }
//
//   if (event.detail.elt.classList.contains("switch-card")) {
//     closeDialog();
//   }
// });
