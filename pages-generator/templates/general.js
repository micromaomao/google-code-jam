function onDocReady () {
  let aDownloads = document.querySelectorAll('a.download')
  function downloadLinkClickHandler (evt) {
    evt.preventDefault()
    if (this.classList.contains('loading')) {
      return
    }
    let url = this.href
    let saveName = this.download
    let origText = this.innerText
    fetch(url, {credentials: 'omit'}).then(res => res.blob()).then(blob => {
      this.classList.remove('loading')
      this.innerText = origText
      let blobUrl = URL.createObjectURL(blob)
      this.href = blobUrl
      this.innerText = 'click again'
      this.removeEventListener('click', downloadLinkClickHandler)
      this.addEventListener('click', evt => {
        this.innerText = origText
      })
    }, err => {
      this.classList.remove('loading')
      this.innerText = err.toString()
      this.classList.add('error')
    })
    this.classList.add('loading')
    this.innerText = 'loading...'
  }
  for (let ele of aDownloads) {
    ele.addEventListener('click', downloadLinkClickHandler)
  }
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
