const {Sketch, Page, Artboard} = require('sketch-constructor');
const fs = require('fs');

function buildSketch(designInfo) {
    const newSketch = new Sketch();

    for (const flow of designInfo.flows) {
        const page = new Page({
            name: flow.flowName
        });

        if (flow.interactions == null) {
            continue;
        }

        var index = 1;
        for (const interaction of flow.interactions) {
            let componentName = interaction.see.componentName.replace(/\"/g, "");
            let componentData = interaction.see.data.replace(/\"/g, "");
            const artboard = new Artboard({
                name: componentData + componentName,
                frame: {
                    width: '320px',
                    height: '480px',
                    x: 360 * index + 'px'
                },
            });

            index++;

            page.addArtboard(artboard);
        }

        newSketch.addPage(page);
    }
    newSketch.build('output.sketch');
}

function readJsonFile(path) {
    let rawdata = fs.readFileSync(path);
    return JSON.parse(rawdata);
}
//
// designInfo = readJsonFile('output.json');
// buildSketch(designInfo);

var data = "";
process.stdin.resume();
process.stdin.setEncoding('utf8');

process.stdin.on('data', function (chunk) {
    data += chunk;
});

process.stdin.on('end', function () {
    buildSketch(JSON.parse(data));
});
