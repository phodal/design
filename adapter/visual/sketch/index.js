const {Sketch, Page, Artboard} = require('sketch-constructor');
const fs = require('fs');

function buildSketch(designInfo) {
    const newSketch = new Sketch();
    const page = new Page({
        name: 'HomePage'
    });
    var index = 1;
    for (const key in designInfo.components) {
        const artboard = new Artboard({
            name: key,
            frame: {
                width: '320px',
                height: '480px',
                x: 360 * index + 'px'
            },
        });

        page.addArtboard(artboard);
        index++;
    }
    newSketch.addPage(page);
    // const page = new Page({});
    // const artboard = new Artboard({});
    // page.addArtboard(artboard);
    //
    // newSketch.addPage(page);

    newSketch.build('output.sketch');
}

var data = "";
process.stdin.resume();
process.stdin.setEncoding('utf8');

process.stdin.on('data', function (chunk) {
    data += chunk;
});

process.stdin.on('end', function () {
    buildSketch(JSON.parse(data));
});
