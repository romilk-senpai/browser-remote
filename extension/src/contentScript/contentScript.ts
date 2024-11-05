import $ from "jquery";
import { API_HOST } from "../constants";

document.addEventListener('contextmenu', event => {
    window.lastRightClickedElement = event.target as HTMLElement;
});

chrome.runtime.onMessage.addListener(async (message, _sender, _sendResponse) => {
    if (message.action === "sendElement") {
        const element = $(window.lastRightClickedElement)[0] || null; // todo: this is not text fix me
        if (!element) {
            return;
        }
        let selector = generateSelector(element);
        console.log(selector);
        await fetch(`${API_HOST}/elements/save`, {
            method: "POST",
            body: JSON.stringify({
                url: message.url,
                name: "",
                query: selector
            })
        });
    }
});

function generateSelector(element: HTMLElement): string {
    const path: string[] = [];
    while (element.nodeType === Node.ELEMENT_NODE) {
        let selector = element.nodeName.toLowerCase();
        if (element.id) {
            selector += `#${element.id}`;
            path.unshift(selector);
            break;
        } else {
            let sibling: HTMLElement | null = element;
            let nth = 1;
            while ((sibling = sibling.previousElementSibling as HTMLElement)) {
                if (sibling.nodeName.toLowerCase() === selector) nth++;
            }
            selector += `:nth-of-type(${nth})`;
        }
        path.unshift(selector);
        element = element.parentElement as HTMLElement;
    }
    return path.join(" > ");
}
