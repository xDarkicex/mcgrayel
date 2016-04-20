
$('document').ready(function(){
	$('<img src="assets/images/logo.png">').appendTo(".logo").hide().delay(1000).fadeIn(1000);
	$('<img src="assets/images/wavefade.png">').appendTo(".wave").hide().delay(100).fadeIn(900);
	$('<div class="cn-child"><p>EasyCare<span class="small-h1-R-top">®</span> Products U.S.A.</p><p>P.O. Box 12362 &#0149; 5361 S. Villa Ave. Fresno, California, 93725 U.S.A.</p><p>Tel. Direct: <a href="tel:5592997660">559-299-7660</a> Toll-Free Tel: <a href="tel:8002897660">800-289-7660</a> Fax: 800-440-7660</p></div>').appendTo('.company-name-container').hide().delay(1000).fadeIn(1000);




})

jQuery(function($) {
	$.fn.addClassTemp = function( options ){
		// Options For addCLassTemp

		var settings = $.extend({
			Delay		: null,
			Class 		: null,
			complete	: null
		}, options);
		return this.each( function(){
			// $(this).delay(settings.Delay);
			if (settings.Class){
				$(this).addClass(settings.Class);
			}

			if ( $.isFunction( settings.complete ) ) {
				settings.complete.call(this);
			}
		}).delay(settings.Delay).queue(function(next){
			$(this).removeClass(settings.Class);
			next();
		});

	}


}(jQuery));

$("#content-container").click(function(){
    $("#content-container").toggleClass("out");
});
