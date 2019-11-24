const {Sketch, Page, Artboard, SharedStyle} = require('sketch-constructor');
const fs = require('fs');
const Swatch = require('./libs/swatch'); // custom component

function buildFlows(designInfo, newSketch) {
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
}

function buildLibraryItem(libraryPresets, sketch, artboard) {
    var colors = [];
    libraryPresets.forEach((preset, i) => {
        var color = {
            label: "",
            value: ""
        };
        preset.subProperties.forEach((property, j) => {
            if (property.key === "label") {
                color.label = property.value
            }
            if (property.key === "value") {
                color.value = property.value
            }
        });
        colors.push(color)
    });
    artboard.frame.width = colors.length * 220 + 20;
    artboard.frame.height = 190;

    colors.forEach((color, i) => {
        let swatch = new Swatch(Object.assign({}, color, {
            frame: {
                x: i * 220 + 20,
                y: 20,
                width: 200,
                height: 170
            }
        }));
        let layerStyle = SharedStyle.LayerStyle({
            name: color.label,
            fills: [{
                color: color.value
            }]
        });
        sketch.addLayerStyle(layerStyle);
        artboard.addLayer(swatch);
    });
}

function buildLibrary(designInfo, newSketch) {
    const libraryPage = new Page({
        name: "Library"
    });

    var offset = 20;
    var lastX = 20;
    for (const library of designInfo.libraries) {
        const artboard = new Artboard({
            name: library.libraryName
        });

        switch (library.libraryName) {
            case "Color":
                buildLibraryItem(library.libraryPresets, newSketch, artboard);
            default:
                break;
        }

        artboard.frame.x = lastX;
        lastX = lastX + artboard.frame.width + offset;
        libraryPage.addArtboard(artboard)
    }
    newSketch.addPage(libraryPage);
}

function buildSketch(designInfo) {
    const newSketch = new Sketch();

    buildFlows(designInfo, newSketch);
    buildLibrary(designInfo, newSketch);

    newSketch.build('output.sketch');
}

function readJsonFile(path) {
    let rawdata = fs.readFileSync(path);
    return JSON.parse(rawdata);
}

designInfo = readJsonFile('output.json');
buildSketch(designInfo);

// var data = "";
// process.stdin.resume();
// process.stdin.setEncoding('utf8');
//
// process.stdin.on('data', function (chunk) {
//     data += chunk;
// });
//
// process.stdin.on('end', function () {
//     buildSketch(JSON.parse(data));
// });
