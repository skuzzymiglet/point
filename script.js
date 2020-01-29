var len
var slides
var index = 0
document.addEventListener('keydown', function (e) {
  slides = document.getElementsByClassName('slide')
  len = slides.length
  if (e.keyCode === 37) {
    index--
  } else if (e.keyCode === 39) {
    index++
  }
  index %= len
  if (index < 0) { index = 0 }
  Array.from(slides).forEach(function (e, i) {
    if (i === index) {
      e.style.display = 'block'
    } else {
      e.style.display = 'none'
    }
  })
})
