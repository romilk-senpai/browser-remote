chrome.runtime.onInstalled.addListener(() => {
    chrome.contextMenus.create({
        id: "sendElement",
        title: "Remember this element"
    });
});

chrome.contextMenus.onClicked.addListener(async (info, tab) => {
    if (info.menuItemId === "sendElement" && tab?.id) {
        chrome.tabs.sendMessage(tab.id, { action: "sendElement", url: tab.url });
    }
});

