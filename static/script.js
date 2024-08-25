const keymap = [
    { id: 'btn-volumeup', command: 'volumeup' },
    { id: 'btn-volumedown', command: 'volumedown' },
    { id: 'btn-back', command: 'back' },
    { id: 'btn-forward', command: 'forward' },
    { id: 'btn-playpause', command: 'playpause' },
];

keymap.forEach((mappedKey) => {
    document.querySelector(`#${mappedKey.id}`).addEventListener('click', () => {
        sendCommand(mappedKey.command);
    });
});

const trackpad = document.querySelector('#trackpad');
const trackpadComputed = {
    width: parseFloat(
        window.getComputedStyle(trackpad).width.replace('px', '')
    ),
    height: parseFloat(
        window.getComputedStyle(trackpad).height.replace('px', '')
    ),
};

// Prevent scrolling
trackpad.addEventListener('touchmove', (event) => event.preventDefault());

trackpad.addEventListener('touchstart', (event) => {
    touchStart(event.touches[0].clientX, event.touches[0].clientY);
});
trackpad.addEventListener('touchend', (event) => {
    touchEnd(event.changedTouches[0].clientX, event.changedTouches[0].clientY);
});
trackpad.addEventListener('mousedown', (event) => {
    touchStart(event.clientX, event.clientY);
});
trackpad.addEventListener('mouseup', (event) => {
    touchEnd(event.clientX, event.clientY);
});

var trackpadStart = { x: 0, y: 0 };
function touchStart(x, y) {
    trackpadStart = { x, y };
}
function touchEnd(x, y) {
    const deltaX = x - trackpadStart.x;
    const deltaY = y - trackpadStart.y;

    const scaledDeltaX = (deltaX / trackpadComputed.width) * 100;
    const scaledDeltaY = (deltaY / trackpadComputed.height) * 100;

    sendTrackpadMove(scaledDeltaX.toFixed(), scaledDeltaY.toFixed());
}

const keyboardInput = document.querySelector('#keyboard-input');
keyboardInput.addEventListener('input', (event) => {
    console.log(event);
    switch (event.inputType) {
        case 'insertText':
            sendKeyboardInput(event.data);
            setTimeout(clearKeyboardInput, 800);
            break;
        case 'deleteContentBackward':
            sendCommand('backspace');
            clearKeyboardInput();
        default:
            break;
    }
});
keyboardInput.addEventListener('keypress', (event) => {
    if (event.key === 'Enter') {
        sendCommand('enter');
    }
});

function clearKeyboardInput() {
    //Leave a character so you can backspace
    keyboardInput.value = ' ';
}

function sendCommand(command) {
    fetch('/btn', {
        method: 'POST',
        body: command,
    });
}
function sendTrackpadMove(deltaX, deltaY) {
    fetch('/trackpad', {
        method: 'POST',
        body: `${deltaX},${deltaY}`,
    });
}
function sendKeyboardInput(text) {
    fetch('/kbd', {
        method: 'POST',
        body: text,
    });
}
