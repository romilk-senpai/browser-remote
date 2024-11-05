import $ from "jquery";
import { API_HOST } from "../constants";

document.addEventListener('contextmenu', event => {
    window.lastRightClickedElement = event.target as HTMLElement;
});

chrome.runtime.onMessage.addListener(async (message, sender, sendResponse) => {
    if (message.action === "sendElement") {
        const elementText = $(window.lastRightClickedElement)[0] || null; // todo: this is not text fix me
        if (!elementText) {
            return;
        }
        await fetch(`${API_HOST}/elements/save`, {
            method: "POST",
            body: JSON.stringify({
                url: message.url,
                name: "",
                query: elementText
            })
        });
    }
});
