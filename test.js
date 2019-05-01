'use strict';

const
    { spawn } = require( 'child_process' ),
    ls = spawn( './main', [ '-token', '123', '-url', 'https://github.com/Microsoft/BotFramework-Emulator/releases/download/v4.3.3/BotFramework-Emulator-4.3.3-linux-i386.AppImage', '-saved_path', 'data.bump' ] );

ls.stdout.on( 'data', data => {
    console.log( `stdout: ${data}` );
});

ls.stderr.on( 'data', data => {
    console.log( `stderr: ${parseInt(data)}` );
});

ls.on( 'close', code => {
    console.log( `child process exited with code ${code}` );
});
