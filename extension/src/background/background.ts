chrome.runtime.onInstalled.addListener(() => {
    chrome.contextMenus.create({
        id: "sendElement",
        title: "Remember this element",
        contexts: ["all"]
    });
});

chrome.contextMenus.onClicked.addListener(async (info, tab) => {
    if (info.menuItemId === "sendElement" && tab?.id) {
        if (!tab.url) {
            return;
        }
        try {
            let url = new URL(tab.url);
            chrome.tabs.sendMessage(tab.id, { action: "sendElement", url: url });
        }
        catch (error) {
            console.log(error);
        }
    }
});

