function onDocReady () {
}

function seeIfReady () {
  document.removeEventListener("readystatechange", seeIfReady)
  if (document.readyState == "complete" || document.readyState == "interactive") {
    onDocReady()
  } else {
    document.addEventListener('readystatechange', seeIfReady)
  }
}

seeIfReady()
