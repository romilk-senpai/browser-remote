import $ from "jquery";

document.addEventListener('contextmenu', event => {
    window.lastRightClickedElement = event.target as HTMLElement;
});

chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
    if (message.action === "sendElement") {
        const elementText = $(window.lastRightClickedElement) || "No jquery available";
        console.log(elementText);
        sendResponse("ok");
    }
});