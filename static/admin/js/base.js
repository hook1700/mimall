$(function(){
	app.init()
})

var app={
	init:function(){
		this.slideToggle();
		this.resizeIframe();
		this.confirmDelete();
	},
	slideToggle:function(){
		$('.aside h4').click(function(){

			$(this).siblings('ul').slideToggle();
		})
	},
	resizeIframe:function(){
		$("#rightMain").height($(window).height()-80)
	},
	// 删除提示
	confirmDelete:function(){
		$(".delete").click(function(){
			var flag=confirm("您确定要删除吗?")
			return flag
		})
	}
}
