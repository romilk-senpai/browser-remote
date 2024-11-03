import $ from "jquery";
import { API_HOST } from "../constants";

document.addEventListener('contextmenu', event => {
    window.lastRightClickedElement = event.target as HTMLElement;
});

chrome.runtime.onMessage.addListener(async (message, sender, sendResponse) => {
    if (message.action === "sendElement") {
        const elementText = $(window.lastRightClickedElement) || "No jquery available";
        console.log(elementText);
        const [tab] = await chrome.tabs.query({ active: true, lastFocusedWindow: true });
        await fetch(`${API_HOST}/elements/save`, {
            method: "POST",
            body: JSON.stringify({
                url: tab.url,
                name: "",
                query: elementText
            })
        });
    }
});
