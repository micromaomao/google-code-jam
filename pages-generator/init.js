#!/usr/bin/env node
const fs = require('fs')
const path = require('path')
const pug = require('pug')
const highlight = require('jstransformer')(require('jstransformer-highlight'))
const sass = require('node-sass')
const marked = require('marked')

if (process.argv.length !== 2) {
  process.stdout.write("Expected no arguments.\n")
  process.exit(1)
}

let projectRoot = path.resolve(__dirname, "..")
let outputDir = path.resolve(__dirname, '../dist')
try {
  fs.mkdirSync(outputDir)
} catch (e) {
  if (e.code !== 'EEXIST') throw e
}
try {
  fs.mkdirSync(path.resolve(outputDir, "commons"))
} catch (e) {
  if (e.code !== 'EEXIST') throw e
}
let dirs = fs.readdirSync(projectRoot).filter(name => /^20\d\d-(\d[A-Z]|qual)$/.test(name))

process.chdir(__dirname)
let codeTemplate = pug.compileFile(path.resolve(__dirname, 'templates', 'code.pug'))
let indexTemplate = pug.compileFile(path.resolve(__dirname, 'templates', 'index.pug'))
let hljsStyleContent = fs.readFileSync(path.resolve(__dirname, "node_modules/highlight.js/styles/github.css"), {encoding: 'utf8'})
let hljsStylePath = path.resolve(outputDir, "commons", "hljs.css")
process.stderr.write(`CP ${path.relative(outputDir, hljsStylePath)}\n`)
fs.writeFileSync(hljsStylePath, hljsStyleContent)

let mainStyleSassPath = path.resolve(__dirname, 'templates', 'style.sass')
let mainStyleSheetContent = sass.renderSync({
  file: mainStyleSassPath
}).css
let mainStyleSheetPath = path.resolve(outputDir, "commons", "style.css")
process.stderr.write(`SASS ${path.relative(outputDir, mainStyleSheetPath)}\n`)
fs.writeFileSync(mainStyleSheetPath, mainStyleSheetContent)

let generalJsContent = fs.readFileSync(path.resolve(__dirname, "templates", "general.js"), {encoding: 'utf8'})
let generalJsPath = path.resolve(outputDir, "commons", "general.js")
process.stderr.write(`CP ${path.relative(outputDir, generalJsPath)}\n`)
fs.writeFileSync(generalJsPath, generalJsContent)

process.chdir(outputDir)
let series = []
for (let seriesName of dirs) {
  let subdirs = fs.readdirSync(path.resolve(projectRoot, seriesName)).filter(subdirName => {
    let subdirPath = path.resolve(projectRoot, seriesName, subdirName)
    if (!fs.statSync(subdirPath).isDirectory()) return false
    if (!/^\d+$/.test(subdirName)) return false
    return true
  })
  let links = []
  for (let problemDir of subdirs) {
    let dirPath = path.resolve(projectRoot, seriesName, problemDir)
    try {
      fs.mkdirSync(path.resolve(outputDir, seriesName), {recursive: true})
    } catch (e) {
      if (e.code !== 'EEXIST') throw e
    }
    let outputFilePath = path.resolve(outputDir, seriesName, problemDir + '.html')
    let problemName = `${seriesName} Problem ${'ABCDEFGHIJKLMNOPQRSTUVWXYZ'[parseInt(problemDir)-1]}`
    links.push({
      href: path.relative(outputDir, outputFilePath),
      name: problemName
    })
    let filesObj = {}
    let dirs = []
    function rec(c) {
      let list = fs.readdirSync(c)
      for (let file of list) {
        let fp = path.resolve(c, file)
        let fileRelPath = path.relative(dirPath, fp)
        let stat = fs.statSync(fp)
        if (stat.isDirectory()) {
          dirs.push(fileRelPath)
          rec(fp)
        } else if (stat.isFile()) {
          let fileContent = fs.readFileSync(fp, {encoding: 'utf8'}).toString()
          let highlightRender = null
          if (file.endsWith('.go')) {
            let boilerPlateStartIndex = fileContent.indexOf('/*********Start boilerplate***********/')
            if (boilerPlateStartIndex >= 0) {
              fileContent = fileContent.substr(0, boilerPlateStartIndex).trimRight() + '\n\n// boilerplate omitted...'
            }
            fileContent = fileContent.replace(/^package main\n+import \([^\(\)]+\)\n*/m, "// package, import, etc...\n\n")
            const typicalStartFunc = `func start() {
	var t int
	mustReadLineOfInts(&t)
	for i := 0; i < t; i++ {
		fmt.Fprintf(stdout, "Case #%d: ", i+1)
		test()
	}
}`
            const typicalStartFunc2 = `func start() {
	var t int
	mustReadLineOfInts(&t)
	for i := 0; i < t; i++ {
		stdout.WriteString(fmt.Sprintf("Case #%d: ", i+1))
		test()
	}
}`
            const replaceWith = `func start() {
  // read T, repeat test() t times...
}`
            let idx = fileContent.indexOf(typicalStartFunc)
            if (idx >= 0) {
              fileContent = fileContent.substr(0, idx) + replaceWith + fileContent.substr(idx + typicalStartFunc.length)
            } else if ((idx = fileContent.indexOf(typicalStartFunc2)) >= 0) {
              fileContent = fileContent.substr(0, idx) + replaceWith + fileContent.substr(idx + typicalStartFunc2.length)
            }
            highlightRender = highlight.render(fileContent, {language: 'golang'}).body
          } else if (file.endsWith('.cc') || file.endsWith('.cpp') || file.endsWith('.c')) {
            highlightRender = highlight.render(fileContent, {language: 'c'}).body
          } else if (file.endsWith('.py')) {
            highlightRender = highlight.render(fileContent, {language: 'python'}).body
          } else if (file.endsWith('.md')) {
            highlightRender = marked.parse(fileContent, {
              sanitize: true
            })
          }
          filesObj[fileRelPath] = {
            path: fileRelPath,
            sourceUrl: 'https://github.com/micromaomao/google-code-jam/blob/master/' + path.relative(projectRoot, fp),
            downloadUrl: 'https://raw.githubusercontent.com/micromaomao/google-code-jam/master/' + path.relative(projectRoot, fp),
            content: fileContent,
            highlightRender
          }
        }
      }
    }
    rec(dirPath)
    let thisDirName = path.dirname(outputFilePath)
    let solutions = []
    if (filesObj['cmd.go']) {
      solutions.push({
        no: 1,
        cmd: 'cmd.go',
        correct: true
      })
    } else if (dirs.find(x => /^solution-\d+$/.test(x))) {
      let solutionNumbers = dirs.map(x => x.match(/^solution-(\d+)$/)).filter(x => x).map(x => parseInt(x[1]))
      for (let sln of solutionNumbers) {
        solutions.push({
          no: sln,
          cmd: `solution-${sln}/cmd.go`,
          correct: true
        })
      }
    }
    let wrongSolutions = dirs.filter(x => /^(wrong|incorrect)(-\d+)?$/.test(x))
    for (let sol of wrongSolutions) {
      solutions.push({
        no: sol,
        cmd: `${sol}/cmd.go`,
        correct: false
      })
    }
    if (dirs.includes('small')) {
      solutions.push({
        no: 'small',
        cmd: 'small/cmd.go',
        correct: 'small'
      })
    }
    let output = codeTemplate({files: filesObj, solutions, problemName, rootDir: path.relative(thisDirName, outputDir), hljsStyle: path.relative(thisDirName, hljsStylePath),
                                styleSheet: path.relative(thisDirName, mainStyleSheetPath), generalJs: path.relative(thisDirName, generalJsPath)})
    fs.writeFileSync(outputFilePath, output)
    process.stderr.write(`PUG ${path.relative(outputDir, outputFilePath)}\n`)
  }
  series.push({
    seriesName,
    links
  })
}

let indexHtml = indexTemplate({series, styleSheet: path.relative(outputDir, mainStyleSheetPath), generalJs: path.relative(outputDir, generalJsPath)})
fs.writeFileSync(path.resolve(outputDir, 'index.html'), indexHtml)
process.stderr.write(`PUG index.html\n`)
