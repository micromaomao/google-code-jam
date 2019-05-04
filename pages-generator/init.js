#!/usr/bin/env node
const fs = require('fs')
const path = require('path')
const pug = require('pug')
const highlight = require('jstransformer')(require('jstransformer-highlight'))
const sass = require('node-sass')
const marked = require('marked')
const ReplitClient = require('repl.it-api')
let replit = new ReplitClient()

if (process.argv.length !== 2) {
  process.stdout.write("Expected no arguments.\n")
  process.exit(1)
}

function delay (t) {
  return new Promise((resolve, reject) => {
    setTimeout(() => resolve(), t)
  })
}

let uploadToReplit = false

async function main () {

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

if (fs.readFileSync(path.resolve(projectRoot, '.git', 'HEAD'), {encoding: 'utf8'}).indexOf('master') >= 0 || process.env.TRAVIS_BRANCH === "master") {
  uploadToReplit = true
  await replit.login(process.env.REPLIT_SID)
} else {
  replit = null
}

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
    let thisDirName = path.dirname(outputFilePath)
    let problemNo = parseInt(problemDir)
    let problemName = `${seriesName} Problem ${'ABCDEFGHIJKLMNOPQRSTUVWXYZ'[problemNo-1]}`
    let filesObj = {}
    let dirs = []
    let replName = 'gcj-' + problemName.replace(/ /g, '-')
    if (uploadToReplit) {
      const tryAmount = 10
      for (let i = 1; i <= tryAmount; i ++) {
        try {
          process.stderr.write(`REPLIT-CREATE ${replName}\n`)
          await replit.create('bash')
          await replit.connect()
          break
        } catch (e) {
          if (i == tryAmount) {
            throw e
          }
          process.stderr.write(`... failed (${e}), retrying in ${i*5}s...\n`)
          await delay(i * 5000)
          continue
        }
      }
    }
    async function rec(c) {
      let list = fs.readdirSync(c)
      for (let file of list) {
        let fp = path.resolve(c, file)
        let fileRelPath = path.relative(dirPath, fp)
        let stat = fs.statSync(fp)
        if (stat.isDirectory()) {
          dirs.push(fileRelPath)
          await rec(fp)
        } else if (stat.isFile()) {
          let fileOrigContent = fs.readFileSync(fp, {encoding: null})
          let fileContent = fileOrigContent.toString('utf8')
          if (uploadToReplit) {
            process.stderr.write(`REPLIT-WRITE ${replName} ${fileRelPath}\n`)
            await replit.write(fileRelPath, fileContent)
          }
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
          } else if (file.endsWith('.diff')) {
            highlightRender = highlight.render(fileContent, {language: 'diff'}).body
          } else if (file.endsWith('.md')) {
            highlightRender = marked.parse(fileContent, {
              sanitize: true
            })
          }
          let emitPath = path.resolve(thisDirName, "files", fileRelPath)
          let emitDir = path.dirname(emitPath)
          try {
            fs.mkdirSync(emitDir)
          } catch (e) {
            if (e.code !== 'EEXIST') throw e
          }
          process.stderr.write(`CP ${path.relative(outputDir, emitPath)}\n`)
          fs.writeFileSync(emitPath, fileOrigContent)
          filesObj[fileRelPath] = {
            path: fileRelPath,
            githubUrl: 'https://github.com/micromaomao/google-code-jam/blob/master/' + path.relative(projectRoot, fp),
            downloadUrl: path.relative(thisDirName, emitPath),
            content: fileContent,
            highlightRender
          }
        }
      }
    }
    await rec(dirPath)
    let solutions = []
    function findInitFile(prefix) {
      let names = ["cmd.go", "why.diff"]
      for (let n of names) {
        let path = prefix + '/' + n
        if (prefix == '') {
          path = n
        }
        if (filesObj[path]) {
          return path
        }
      }
      return null
    }
    if (filesObj['cmd.go']) {
      solutions.push({
        no: 1,
        cmd: findInitFile(''),
        correct: true
      })
    } else if (dirs.find(x => /^solution-\d+$/.test(x))) {
      let solutionNumbers = dirs.map(x => x.match(/^solution-(\d+)$/)).filter(x => x).map(x => parseInt(x[1]))
      for (let sln of solutionNumbers) {
        solutions.push({
          no: sln,
          cmd: findInitFile(`solution-${sln}`),
          correct: true
        })
      }
    }
    let wrongSolutions = dirs.filter(x => /^(wrong|incorrect)(-\d+)?$/.test(x))
    for (let sol of wrongSolutions) {
      solutions.push({
        no: sol,
        cmd: findInitFile(sol),
        correct: false
      })
    }
    let smallSolutions = dirs.filter(x => /^small(-\d+)?$/.test(x))
    for (let sol of smallSolutions) {
      solutions.push({
        no: sol,
        cmd: findInitFile(sol),
        correct: 'small'
      })
    }
    let replitinfo = null
    if (uploadToReplit) {
      let mainSh = `go build && ./runner${filesObj['sample.in'] ? ' < sample.in' : ''}\n# Check out files from the sidebar.`
      await replit.writeMain(mainSh)
      replitinfo = replit.getInfo()
    }
    let problemSolved = false
    if (solutions.find(x => x.correct === true)) {
      problemSolved = true
    } else if (solutions.find(x => x.correct === 'small')) {
      problemSolved = 'small'
    }
    links.push({
      href: path.relative(outputDir, outputFilePath),
      name: problemName,
      no: problemNo,
      correct: problemSolved
    })
    solutions = solutions.sort((a, b) => {
      if (a.no < b.no) {
        return -1
      } else if (a.no > b.no) {
        return 1
      } else {
        return 0
      }
    })
    let output = codeTemplate({files: filesObj, solutions, problemName, replitUrl: replitinfo ? replitinfo.url : null,
                                rootDir: path.relative(thisDirName, outputDir), hljsStyle: path.relative(thisDirName, hljsStylePath),
                                styleSheet: path.relative(thisDirName, mainStyleSheetPath), generalJs: path.relative(thisDirName, generalJsPath)})
    fs.writeFileSync(outputFilePath, output)
    process.stderr.write(`PUG ${path.relative(outputDir, outputFilePath)}\n`)
  }
  series.push({
    seriesName,
    links: links.sort((a, b) => Math.sign(a.no - b.no))
  })
}

series.sort((a, b) => {
  let [aYear, aRound] = a.seriesName.split('-')
  let [bYear, bRound] = b.seriesName.split('-')
  aYear = parseInt(aYear)
  bYear = parseInt(bYear)
  let yearCompare = Math.sign(aYear - bYear)
  if (yearCompare !== 0) {
    return yearCompare
  }
  if (aRound == bRound) {
    return 0
  }
  if (aRound == 'qual') {
    return -1
  } else if (bRound == 'qual') {
    return 1
  }
  return aRound<bRound ? -1 : 1
})

let readmeMdContent = fs.readFileSync(path.resolve(projectRoot, "README.md"), {encoding: 'utf8'})

let indexHtml = indexTemplate({series, readme: marked.parse(readmeMdContent, {
  sanitize: false
}), styleSheet: path.relative(outputDir, mainStyleSheetPath), generalJs: path.relative(outputDir, generalJsPath)})
fs.writeFileSync(path.resolve(outputDir, 'index.html'), indexHtml)
process.stderr.write(`PUG index.html\n`)

}

main().then(() => {
  process.exit(0)
}).catch(err => {
  process.stderr.write(err.toString() + '\n' + err.stack + '\n')
  process.exit(1)
})
