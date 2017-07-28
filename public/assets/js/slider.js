//side menu//

$(document).scroll(function() {

y = $(window).height()+$(document).scrollTop();
  if (y >= $(document).height() -50) {
    $('.menu-wrapper').addClass('active');
     $('.arrow-img').addClass('active');
  } else {
    $('.menu-wrapper').removeClass('active');
     $('.arrow-img').removeClass('active');
   }
});

//Image Slider Code//
var sliderIndex = 0;
var refreshIntervalId;
var $images = $(".slide-group").children();

$(function(){
  var $bar = $('.slide-buttons');
  addButton("Previous", prev, $bar);
  addButton("Next", next, $bar);
})

function next(){
  sliderIndex = (sliderIndex<$images.length-1)?sliderIndex+1:0;
  slideNext();
}

function prev(){
  sliderIndex = (sliderIndex>0)?sliderIndex-1:$images.length-1;
  slidePrev();
}
function slideNext(){
  $images.hide("slide", { direction: "left" }, 1000);
  $($images[sliderIndex]).show("slide", { direction: "right" }, 1000);
}

function slidePrev(){
  $images.hide("slide", { direction: "right" }, 1000);
  $($images[sliderIndex]).show("slide", { direction: "left" }, 1000);
}
function validate() {
  if (document.getElementById('remember').checked) {
    refreshIntervalId = window.setInterval(next, 5000);
  } else {
    window.clearInterval(refreshIntervalId)
    refreshIntervalId = null;
  }
}

function addButton(name, click, $element){
  var $button = $("<button>");
  // $button.text(name);
  $button.addClass(name);
  $button.on('click', click);
  $element.append($button);
}

$(document).ready(function(){
  validate();
});


//This function Pauses Slider Good for old people...
$('.slider').on({
   mouseenter: function() {
       window.clearInterval(refreshIntervalId)
    refreshIntervalId = null;
    },
    mouseleave: function() {
        refreshIntervalId = window.setInterval(next, 5000);
        
    }
})

//PDF SLIDER//

var pdfINDEX = 0;
var slides = $('#pdf-slider').children();
// Gah dynamic typing JAVASCRIPT SUCKS, this should all be remade in typescript...
var forcedNext;
console.log(slides + "slide numb");

function  startNext(){
  pdfINDEX = (pdfINDEX < slides.length-1)?pdfINDEX+1:0;
  pdfNext()
};

function pdfNext(){
  slides.hide("slide", { direction: "left"}, 1000);
  $(slides[pdfINDEX]).show("slide", {direction: "right"}, 1000);
};

function forceNext(){
 forcedNext = window.setInterval(startNext, 5000)
};

$(document).ready(function(){
  forceNext();
});

$('.pdf-slider-container').on({
   mouseenter: function() {
       window.clearInterval(forceNext)
    forcedNext = undefined;
    },
    mouseleave: function() {
        forcedNext = window.setInterval(startNext, 5000);
    }
});

//All of this is horrible!! bad way to make links
$('.main-bottle-container div:nth-child(1)').on('click', function(){
  window.location = "http://easycarewater.com/product/algatec";
});

$('.main-bottle-container div:nth-child(2)').on('click', function(){
  window.location = "http://easycarewater.com/product/beautec";
});

$('.main-bottle-container div:nth-child(3)').on('click', function(){
  window.location = "http://easycarewater.com/product/pooltec";
});

$('.main-bottle-container div:nth-child(4)').on('click', function(){
  window.location = "http://easycarewater.com/product/pooltec-winter";
});

$('.main-bottle-container div:nth-child(5)').on('click', function(){
  window.location = "http://easycarewater.com/product/scaletect";
});

$('.main-bottle-container div:nth-child(6)').on('click', function(){
  window.location = "http://easycarewater.com/product/startup-tec";
});

$('.main-bottle-container div:nth-child(7)').on('click', function(){
  window.location = "http://easycarewater.com/product/fountec";
});

$('.main-bottle-container div:nth-child(8)').on('click', function(){
  window.location = "http://easycarewater.com/product/protec";
});
