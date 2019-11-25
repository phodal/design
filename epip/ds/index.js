const ns = require('node-sketch');

ns.read('design.sketch').then(sketch => {
    console.log(sketch.pages)
    // sketch.document           // document data
    // sketch.meta               // meta data
    // sketch.user               // user data
    // sketch.pages              // array with all pages
    // sketch.symbolsPage        // the Symbols page
    // sketch.layerStyles        // array with the layer styles
    // sketch.textStyles         // array with the text styles
    // sketch.colors             // array containing the colors stored in the color palette
    // sketch.gradients          // array containing the gradients stored in the gradient palette
    // sketch.symbols            // array with all symbols stored in the document
    //
    // sketch.foreignSymbols     // array with the symbols loaded from external libraries
    // sketch.foreignLayerStyles // array with the layer styles loaded from external libraries
    // sketch.foreignTextStyles  // array with the text styles loaded from external libraries
});
