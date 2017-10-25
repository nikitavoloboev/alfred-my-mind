#!/usr/bin/env node
const path = require('path');
const fs = require('fs');
const walkSync = require('fs-walk').walkSync;

const { getText, getURL } = require('./regex');

// These two arguments must be directories.
const input = process.argv[2];
const output = process.argv[3];

// Using for converting URLs with IDs to full URLs.
const mapsLookup = {};

if (input === undefined || output === undefined) {
  // eslint-disable-next-line no-console
  console.log(
    'No files were parsed due to insufficient arguments \nPlease run the parser with the following command: npm run parse "path/to/mindmap/json/folder" "path/to/output/folder"'
  );
  process.exit();
}

/*
 * Equivalent to mkdir -p dirname.
 */
const mkdirs = dirname => {
  const parentDir = path.dirname(dirname);

  if (!fs.existsSync(parentDir)) {
    mkdirs(parentDir);
  }

  fs.mkdirSync(dirname);
};

/*
 * Recursively walk a directory and call a function on all its files.
 * Imported file and absolute path are the parameters passed to the callback function.
 */
const walkDir = (dirname, fn) => {
  walkSync(dirname, (basedir, filename, stat) => {
    const absPath = path.resolve('./', basedir, filename);

    if (stat.isDirectory()) {
      return walkDir(absPath, fn);
    }

    if (typeof fn === 'function') {
      // eslint-disable-next-line import/no-dynamic-require, global-require
      fn(require(absPath), absPath);
    }

    return null;
  });
};

const parseNode = node => {
  if (!node.title) {
    return [];
  }

  const text = getText(node.title.text);
  const parsedNode = {
    uid: text,
    title: text,
    autocomplete: text,
    subtitle: node.note ? getText(node.note.text) : undefined,
    arg: getURL(node.title.text)
  };
  const nodes = [parsedNode];

  parseSubnodes(node.nodes).forEach(subnode => nodes.push(subnode));
  return nodes;
};

/*
 * Parse an array of subnodes recursively. Uses parseNode, so the structure of
 * subnodes will be the same. The only difference is a color attribute on
 * subnodes, which has a string with a valid CSS color format.
 */
const parseSubnodes = subnodes => {
  const flatSubs = [];

  subnodes.forEach(subnode => {
    parseNode(subnode).forEach(el => flatSubs.push(el));
  });

  return flatSubs;
};

walkDir(input, (map, filename) => {
  const parsedMap = {};

  // Parse all nodes and populate the lookup table, which will be used for
  // converting IDs to node title on connections.
  parsedMap.items = [];

  map.nodes.forEach(node =>
    parseNode(node).forEach(parsedNode => parsedMap.items.push(parsedNode))
  );

  parsedMap.items = parsedMap.items.filter(item => !!item.arg);

  // Find out the path for the output file.
  const inputBasePath = path.resolve('./', input);
  const outputFile = path.join(output, filename.replace(inputBasePath, ''));
  const outputPath = path.dirname(outputFile);

  // Create folder if it doesn't exist.
  if (!fs.existsSync(outputPath)) {
    mkdirs(outputPath);
  }

  // Write parsed map to new location.
  fs.writeFileSync(outputFile, JSON.stringify(parsedMap, null, 2));
});
