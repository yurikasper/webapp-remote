var keymap = [
    {id: 'btn-volumeup', command: 'volumeup'},
    {id: 'btn-volumedown', command: 'volumedown'},
    {id: 'btn-back', command: 'back'},
    {id: 'btn-forward', command: 'forward'},
    {id: 'btn-playpause', command: 'playpause'}
];

keymap.forEach(mappedKey => {
    document.querySelector(`#${mappedKey.id}`).addEventListener('click', () => {
        sendCommand(mappedKey.command);
    });
});

function sendCommand(command) {
    fetch('/btn', {
    method: 'POST',
    body: command
    });
}
