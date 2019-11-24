const {Sketch, Page, Artboard} = require('sketch-constructor');

// Without a path it creates an empty sketch class to work with
const newSketch = new Sketch();

// Create a new Page instance
const page = new Page({ });

// Add an artboard to the page
const artboard = new Artboard({});
page.addArtboard(artboard);

// Add the page to the Sketch instance
newSketch.addPage( page );

// You can also add a page by giving it an object, the arguments
// are the same as that of the Page constructor
newSketch.addPage({
    name: 'My Page'
});

newSketch.build('output.sketch');
