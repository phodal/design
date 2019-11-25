const {Sketch, Page, Text, Artboard, SharedStyle} = require('sketch-constructor');
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

function buildColorLibrary(libraryPresets, sketch, artboard) {
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
    artboard.frame.width = 220;
    artboard.frame.height = colors.length * 190 + 20;

    colors.forEach((color, i) => {
        let swatch = new Swatch(Object.assign({}, color, {
            frame: {
                x: 10,
                y: i * 220 + 20,
                width: 200,
                height: 60
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

function buildFontSizeLibrary(libraryPresets, newSketch, artboard) {
    artboard.frame.width = 320;
    artboard.frame.height = libraryPresets.length * 50;
    libraryPresets.forEach((preset, i) => {
        var text = new Text({
            string: preset.key + "  FontSize  " + preset.value,
            name: preset.key,
            fontSize: parseInt(preset.value.replace("px", "")),
            color: '#000',
            frame: {
                x: 0,
                y: 40 * i,
            },
        });

        artboard.addLayer(text);
        text = null;
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
                buildColorLibrary(library.libraryPresets, newSketch, artboard);
                break;
            case "FontSize":
                buildFontSizeLibrary(library.libraryPresets, newSketch, artboard);
                break;
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
