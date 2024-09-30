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

async function shareSwitchDetail(event) {
  event.preventDefault();
  event.stopPropagation();
  const switchName = event.currentTarget.getAttribute("data-switch-name");
  const switchID = event.currentTarget.getAttribute("data-switch-id");
  console.log(switchID);
  const domain = window.location.origin;
  const url = `${domain}/switches/${switchID}`;

  if (navigator.share && navigator.canShare(shareData)) {
    const shareData = {
      title: "Check out this awesome switch!",
      text: switchName,
      url: url,
    };
    await navigator.share(shareData);
  } else {
    await navigator.clipboard.writeText(url);
    alert("Switch address copied to clipboard!");
  }
}

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

document.addEventListener("DOMContentLoaded", function () {
  const menuButton = document.getElementById("menu-mobile");
  menuButton.addEventListener("click", function () {
    openMobileMenu();
  });
});

function openMobileMenu() {
  const sidenav = document.getElementById("sidenav");
  sidenav.classList.add("open");

  const closeMenuButton = document.getElementById("sidenav");
  closeMenuButton.addEventListener("click", function () {
    sidenav.classList.remove("open");
  });
}
