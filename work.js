    const { spawn } = require('child_process');
    const pyProg = spawn('C:/Users/aryan/AppData/Local/Programs/Python/Python310/python.exe', ['c:/Users/aryan/Desktop/Amigo_Ass/back.py']);

    pyProg.stdout.on('data', function(data) {

        console.log(data.toString());
    });