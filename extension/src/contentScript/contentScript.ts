import $ from "jquery";
import { API_HOST } from "../constants";

document.addEventListener('contextmenu', event => {
    window.lastRightClickedElement = event.target as HTMLElement;
});

chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
    if (message.action === "sendElement") {
        const elementText = $(window.lastRightClickedElement) || "No jquery available";
        console.log(elementText);
        chrome.tabs.query({ active: true, lastFocusedWindow: true })
            .then(tabs => {
                fetch(`${API_HOST}/elements/save`, {
                    method: "POST",
                    body: JSON.stringify({
                        url: tabs[0].url,
                        name: "",
                        query: elementText
                    })
                });
            });
    }
});
