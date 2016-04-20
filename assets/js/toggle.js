(function() {

  "use strict";

  var toggles = document.querySelectorAll(".menu-wrapper");

  for (var i = toggles.length - 1; i >= 0; i--) {
    var toggle = toggles[i];
    toggleHandler(toggle);
  };

  function toggleHandler(toggle) {
    toggle.addEventListener( "click", function(e) {

      (this.classList.contains("active") === true) ? this.classList.remove("active") : this.classList.add("active");
    });
  }

})();


(function() {


  var toggles = document.querySelectorAll(".arrow-img");

  for (var i = toggles.length - 1; i >= 0; i--) {
    var toggle = toggles[i];
    toggleHandler(toggle);
  };

  function toggleHandler(toggle) {
    toggle.addEventListener( "click", function(e) {

      (this.classList.contains("active") === true) ? this.classList.remove("active") : this.classList.add("active");
    });
  }

})();