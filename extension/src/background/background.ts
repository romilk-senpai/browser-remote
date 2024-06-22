chrome.runtime.onInstalled.addListener(() => {
    chrome.contextMenus.create({
        id: "sendElement",
        title: "Remember this element"
    });
});

chrome.contextMenus.onClicked.addListener((info, tab) => {
    if (info.menuItemId === "sendElement" && tab?.id) {
        console.log(tab);
        chrome.tabs.sendMessage(tab.id, { action: "sendElement" });
    }
});

