'use strict';

var gulp = require('gulp');
var path = require('path');
var jsonlint = require('gulp-jsonlint');
// var watch = require('gulp-watch');
var gulpFilter = require('gulp-filter');
var debug = require('gulp-debug');

var ramllint = require('./lib/gulp-ramllint');
var deref = require('./lib/deref-raml-schema.js');
var ramlStruct = require('./lib/json-to-raml-struct.js');
var fixRamlOutput = require('./lib/fix-raml-output.js');
var raml2html = require('./lib/gulp-raml2html.js');
var validateExamples = require('./lib/validate-examples.js');

var API_SPEC = '*.raml';
var API_DEST = 'public';
var RAML2HTML_OPTIONS = {
  // Task-specific options go here.
  mainTemplate: 'template.nunjucks',
  templatesPath: 'templates'
};

function handleError(err) {
  console.error(err.toString());
  this.emit('end');
}

gulp.task('apidoc', function () {
  var rename = require('gulp-rename'),
    schemaFolder = path.resolve(process.cwd(), 'schema');

  return gulp.src(API_SPEC)
    // .pipe(watch(['*.raml', 'raml/**/*.raml', 'common-raml/**/*.raml', 'examples/**/*.json', 'schema/**/*.json', 'templates/*.nunjucks']))
    .pipe(debug())
    .pipe(gulpFilter(API_SPEC))
    //.pipe(ramllint())
    //.pipe(debug())
    //.pipe(ramllint.reporter())
    //.pipe(deref(schemaFolder))
    //.pipe(debug())
    //.pipe(ramlStruct())
    //.pipe(debug())
    //.pipe(validateExamples())
    //.pipe(validateExamples.reporter())
    //.pipe(fixRamlOutput())
    .pipe(raml2html(RAML2HTML_OPTIONS))
    .on('error', handleError)
    .pipe(rename({
      extname: '.html'
    }))
    .pipe(gulp.dest(API_DEST));
});

gulp.task('jsonlint', function () {
  gulp.src(['./examples/**/*.json', './schema/**/*.json'])
  // .pipe(watch(['./examples/**/*.json', './schema/**/*.json']))
  .pipe(jsonlint())
    .pipe(jsonlint.reporter());
});


gulp.task('default', ['jsonlint', 'apidoc']);
