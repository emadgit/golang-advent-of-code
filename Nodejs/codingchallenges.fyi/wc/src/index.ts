import { program }  from 'commander';
program
    .name('ccwc')
    .option('-c')
    .argument('<string>');

program.parse();

console.log(program.args);
