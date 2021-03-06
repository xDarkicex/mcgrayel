window.onload = function() {
   window.addEventListener("load", function() {
     document.getElementById("form").addEventListener("submit", function(event) {
       event.preventDefault(); // Prevent form submission and contact with server
       event.stopPropagation();
     }, false);
   }, false);
  ///////////////////////////////////////////////
  // Handing adding and removing for key points
  ///////////////////////////////////////////////
  // click function for adding a point to key points list
  document.querySelector('.addPoint').onclick = function(e) {
    e.preventDefault();
    var wrapper = document.querySelector(".keyPoints");
    var label = document.createElement('label');
    var input = document.createElement('input');
    input.setAttribute("type", "text");
    input.setAttribute("name", "points");
    label.appendChild(input);
    wrapper.appendChild(label);
  }

  // Click function for removing Point and preventing first point from being removed
  document.querySelector('.removePoint').onclick = function(e) {
    e.preventDefault();
    var wrapper = document.querySelector(".keyPoints");
    if (wrapper.childElementCount > 1){
      wrapper.removeChild(wrapper.querySelector('label:last-child'));
    }
  }

  ///////////////////////////////////////////////
  // Handing adding and removing for features
  ///////////////////////////////////////////////

  // click function for adding features
  document.querySelector('.addFeature').onclick = function(e) {
    e.preventDefault();
    var wrapper = document.querySelector('.features');
    var label = document.createElement('label');
    var input = document.createElement('input');
    input.setAttribute("type", "text");
    input.setAttribute("name", "features");
    label.appendChild(input);
    wrapper.appendChild(label);
  }
  // click function for removing and preventing features
  document.querySelector('.removeFeature').onclick = function(e) {
    e.preventDefault();
    var wrapper = document.querySelector('.features');
    if (wrapper.childElementCount > 1) {
      wrapper.removeChild(wrapper.querySelector('label:last-child'));
    }
  }

  document.querySelector('.addSellingPoint').onclick = function(e) {
    e.preventDefault();
    var wrapper = document.querySelector('.sellingPoints');
    var label = document.createElement('label');
    var input = document.createElement('input');
    input.setAttribute("type", "text");
    input.setAttribute("name", "sellingPoints");
    label.appendChild(input);
    wrapper.appendChild(label);
  }

  document.querySelector('.removeSellingPoint').onclick = function(e) {
    e.preventDefault();
    var wrapper = document.querySelector('.sellingPoints');
    if (wrapper.childElementCount > 1) {
      wrapper.removeChild(wrapper.querySelector('label:last-child'));
    }
  }

  document.querySelector(".addBenefit").onclick = function(e) {
    e.preventDefault();
    var wrapper = document.querySelector('.benefits');
    var label = document.createElement('label');
    var p = document.createElement('p');
    var input = document.createElement('input');
    var remove = document.createElement('button')
    remove.className = "remove"
    remove.innerText = "X"
    remove.addEventListener("click", function(e){
      e.preventDefault();
      e.target.parentNode.remove()
    })
    input.setAttribute("type", "text")
    input.setAttribute("name", "benefits");
    p.appendChild(input);
    p.appendChild(remove);
    label.appendChild(p);
    wrapper.appendChild(label);
  }

  document.querySelector('.removeBenefit').onclick = function(e) {
    e.preventDefault();
    var wrapper = document.querySelector('.benefits');
    if (wrapper.childElementCount > 1){
      wrapper.removeChild(wrapper.querySelector('label:last-child'));
    }
  }
  document.querySelector('.remove').onclick = function(e) {
    e.preventDefault();
    e.target.parentNode.remove()
  }
}
