document.addEventListener('DOMContentLoaded', function() {
  var nav = document.getElementsByClassName('navMaterial')
  var children = nav[0].children
  for (var i = 0; i < children.length; i++)  {
    t = children[i]
    t.addEventListener('click', function(e){
      window.location = e.target.childNodes[(e.target.childNodes.length > 1 ? 1 : 0)].getAttribute('href')
    })
  }
}, true);
