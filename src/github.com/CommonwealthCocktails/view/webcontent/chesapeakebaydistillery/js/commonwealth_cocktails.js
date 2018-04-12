	var slice = [].slice;

	(function($, window) {
	  var Starrr;
	  window.Starrr = Starrr = (function() {
	    Starrr.prototype.defaults = {
	      rating: void 0,
	      max: 5,
	      readOnly: false,
	      emptyClass: 'fa fa-star-o',
	      fullClass: 'fa fa-star',
	      change: function(e, value) {}
	    };
	
	    function Starrr($el, options) {
	      this.options = $.extend({}, this.defaults, options);
	      this.$el = $el;
	      this.createStars();
	      this.syncRating();
	      if (this.options.readOnly) {
	        return;
	      }
	      this.$el.on('mouseover.starrr', 'a', (function(_this) {
	        return function(e) {
	          return _this.syncRating(_this.getStars().index(e.currentTarget) + 1);
	        };
	      })(this));
	      this.$el.on('mouseout.starrr', (function(_this) {
	        return function() {
	          return _this.syncRating();
	        };
	      })(this));
	      this.$el.on('click.starrr', 'a', (function(_this) {
	        return function(e) {
	          e.preventDefault();
	          return _this.setRating(_this.getStars().index(e.currentTarget) + 1);
	        };
	      })(this));
	      this.$el.on('starrr:change', this.options.change);
	    }
	
	    Starrr.prototype.getStars = function() {
	      return this.$el.find('a');
	    };
	
	    Starrr.prototype.createStars = function() {
	      var j, ref, results;
	      results = [];
	      for (j = 1, ref = this.options.max; 1 <= ref ? j <= ref : j >= ref; 1 <= ref ? j++ : j--) {
	        results.push(this.$el.append("<a href='#' />"));
	      }
	      return results;
	    };
	
	    Starrr.prototype.setRating = function(rating) {
	      if (this.options.rating === rating) {
	        rating = void 0;
	      }
	      this.options.rating = rating;
	      this.syncRating();
	      return this.$el.trigger('starrr:change', rating);
	    };
	
	    Starrr.prototype.getRating = function() {
	      return this.options.rating;
	    };
	
	    Starrr.prototype.syncRating = function(rating) {
	      var $stars, i, j, ref, results;
	      rating || (rating = this.options.rating);
	      $stars = this.getStars();
	      results = [];
	      for (i = j = 1, ref = this.options.max; 1 <= ref ? j <= ref : j >= ref; i = 1 <= ref ? ++j : --j) {
	        results.push($stars.eq(i - 1).removeClass(rating >= i ? this.options.emptyClass : this.options.fullClass).addClass(rating >= i ? this.options.fullClass : this.options.emptyClass));
	      }
	      return results;
	    };
	
	    return Starrr;
	
	  })();
	  return $.fn.extend({
	    starrr: function() {
	      var args, option;
	      option = arguments[0], args = 2 <= arguments.length ? slice.call(arguments, 1) : [];
	      return this.each(function() {
	        var data;
	        data = $(this).data('starrr');
	        if (!data) {
	          $(this).data('starrr', (data = new Starrr($(this), option)));
	        }
	        if (typeof option === 'string') {
	          return data[option].apply(data, args);
	        }
	      });
	    }
	  });
	})(window.jQuery, window);








	function search_item_include(include, exclude, item, hidden) {
	  document.getElementById(exclude).className = "fa fa-times btn btn-default";
	  document.getElementById(hidden).value = "";
	  document.getElementById(item).className = "list-group-item";
	  if (document.getElementById(include).className.search("active") !== -1) {
	    document.getElementById(include).className = "fa fa-check btn btn-default";
	  }
	  else {
	    document.getElementById(item).className = "list-group-item active";
	    document.getElementById(include).className = "fa fa-check btn btn-default active";
	    document.getElementById(hidden).value = "include";
	  }
	  document.getElementById(include).blur()
	}

	function search_item_exclude(include, exclude, item, hidden) {
	  document.getElementById(include).className = "fa fa-check btn btn-default";
	  document.getElementById(hidden).value = "";
	  document.getElementById(item).className = "list-group-item";
	  if (document.getElementById(exclude).className.search("active") !== -1) {
	    document.getElementById(exclude).className = "fa fa-times btn btn-default";
	  }
	  else {
	    document.getElementById(item).className = "list-group-item active";
	    document.getElementById(exclude).className = "fa fa-times btn btn-default active";
	    document.getElementById(hidden).value = "exclude";
	  }
	  document.getElementById(exclude).blur()
	}
	